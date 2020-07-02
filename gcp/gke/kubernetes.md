Detailed Security Configuration
Overview

This section is meant to provide an opinionated approach towards the implementation security controls by domain. Although the approaches may not be a fit for all use casees or complete for production use, they are meant to guide the audience towards current best practices and design considerations to achieve security control objectives.

Controls and Architectures

This table maps Security Domain to the corresponding controls and architectural best practices as documented in AWS’ public documentation, white papers, and blog posts.

Security Domain	Control & Architectural Suggestions	References
Encryption		
Encryption of data at-rest	
AWS RedShift supports KMS and HSM to provide key material management and encryption services. Encryption at rest of RedShift encrypts the data blocks and system metadata of the cluster and its snapshots.[1]

Note: By default, Amazon Redshift selects the account service default key as the master key. The default key is an AWS-managed key that is created for your AWS account to use in Amazon Redshift. Some customer security controls prevent the use of default service KMS keys for sensitive workloads. Users should pre-create a customer managed CMK for RedShift usage.

Implementation Note: Encryption is an optional, immutable setting of a cluster. If you want encryption, you enable it during the cluster launch process. To go from an unencrypted cluster to an encrypted cluster or the other way around, unload your data from the existing cluster and reload it in a new cluster with the chosen encryption setting. [2]

For S3 encryption details see S3 Accelerator.

For KMS details see KMS Accelerator

https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html

Migrating an Unencrypted Cluster to an Encrypted Cluster

https://docs.aws.amazon.com/cli/latest/reference/redshift/describe-clusters.html

Encryption of data in-transit	
To support SSL connections, Amazon Redshift creates and installs an AWS Certificate Manager (ACM) issued SSL certificate on each cluster. The set of Certificate Authorities that you must trust in order to properly support SSL connections can be found at https://s3.amazonaws.com/redshift-downloads/redshift-ca-bundle.crt.

Note: Customers can import certificates into AWS ACM to use custom certs and still take advantage of the integration ACM has with Redshift.[3]

RedShift endpoints are available over HTTPS at a selection of regions.

Best practice:

Set the “require_SSL” parameter to “true” in the parameter group that is associated with the cluster.

For workloads that require FIPS-140-2 SSL compliance an additional step is required to set parameter “use_fips_ssl” to “true”

How to encrypt end to end: https://aws.amazon.com/blogs/big-data/encrypt-your-amazon-redshift-loads-with-amazon-s3-and-aws-kms/

To make client side encryption work follow this pattern https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html

https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html

Encryption Key Management	
Key management is handled by KMS or HSM.

For encryption functions to occur in Redshift the DEK/CEK are stored on disk in an encrypted state. They persist on disk after cluster reboots and then require a request to KMS to use the CMK to decrypt the CEK to be used again in memory.

Key Rotation can occur as often as data requirement define.[1]

Example: Commandline key rotation[3]

rotate-encryption-key --cluster-identifier  [--cli-input-json ] [--generate-cli-skeleton ]
Note: Snapshots stored in S3 will need to be decrypted prior to key rotation and then re-encrypted using the new DEK. This is a process that should be tested prior to production use.

https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html

https://docs.aws.amazon.com/kms/latest/developerguide/concepts.htm

https://docs.aws.amazon.com/cli/latest/reference/redshift/rotate-encryption-key.html

Infrastructure		
Isolation of physical hosts	N/A: RedShift is a fully-managed service and for cluster nodes the isolation of hosts is not currently possible via dedicated host ec2-resources. Reserved instances can be purchased to ensure availability of ec2 resource types.	
Network Isolation	
When an Amazon Redshift cluster is provisioned, it is locked down by default so nobody has access to it except to IAM entities with Console access from within the provisioned network and with the default credentials. Amazon Redshift provides a cluster security group called default, which is created automatically when you launch your first cluster. Initially, this cluster security group is empty. You can add inbound access rules to the default cluster security group and then associate it with your Amazon Redshift cluster. To grant other users inbound access to an Amazon Redshift cluster, you associate the cluster with a security group. To grant access use an existing Amazon VPC security group or define a new one and then associate it with a cluster. For more information on managing a cluster on the EC2-VPC platform, see Managing Clusters in an Amazon Virtual Private Cloud (VPC).

Amazon RedShift relies on EC2 security groups to provide infrastructure security, and thus initial protection from unauthorized traffic connecting to the cluster. [1]

Best Practice

SecurityGroups should follow a naming convention for the entire account

The cluster leader node is the only EC2 instance that is allowed to communicate with the cluster nodes in the AWS Service Account. Ensure to enable VPC FlowLogs on the leader node ENI and capture logs from the OS to ensure only authorized access and activity has occurred.

For the leader node it is recommended the SecurityGroup have no outbound entries to prevent egress of data if the node is compromised.

Attempt to reference other SecurityGroups instead of using IP

Enable EnhancedVPCRouting [4]

Enable VPCEndpoint with S3

For most use cases the cluster should not be publically accessible.

Configure default port to authorized port for SecurityGroup usage. The default port is 5439 and cannot be changed after cluster is built.[5]

See S3 Accelerator for controls around S3 and data isolation.

https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html

https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html

https://docs.aws.amazon.com/redshift/latest/mgmt/copy-unload-iam-role.html

https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-enabling-cluster.html

https://docs.aws.amazon.com/redshift/latest/gsg/rs-gsg-prereq.html

AWS Network	
A special use case exists for RedShift network isolation and must be noted but requires no action. When database encryption is enabled with KMS, KMS exports a CEK/DEK that is stored on a separate network from the cluster. This network is part of the managed service of RedShift and is not customer configurable or monitored.

Another note to mention is that Redshift clusters exist in another AWS account managed by AWS. This is important to be aware of for monitoring so it is clear what account traffic is going towards and coming from is actually authorized vs rogue.

https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html

IAM		
