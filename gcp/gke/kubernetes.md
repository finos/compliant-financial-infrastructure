Date: 2019-04-08

# Detailed Security Configuration

Overview

This section is meant to provide an opinionated approach towards the
implementation security controls by domain. Although the approaches may
not be a fit for all use casees or complete for production use, they are
meant to guide the audience towards current best practices and design
considerations to achieve security control objectives.

**Controls and Architectures**

This table maps Security Domain to the corresponding controls and
architectural best practices as documented in GCP public documentation,
white papers, and blog posts.

<table>
<tbody>
<tr class="odd">
<td><strong>Security Domain</strong></td>
<td><strong>Control &amp; Architectural Suggestions</strong></td>
<td><strong>References</strong></td>
</tr>
<tr class="even">
<td><strong>Encryption</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Encryption of data at-rest</td>
<td><p>
By default, GKE encrypts customer content stored at rest, including Secrets. GKE handles and manages this default encryption for the user without any additional actions. Persistent disks in GKE are already encrypted at the hardware layer by default, however, there is also the option of adding additional encryption where you, the user, can manage the encryption keys.
To encrypt persistent disks in GKE, you must use the GCP Persistent Disk CSI plugin, which lets you protect disks in GKE with a key that you manage in Cloud KMS—by creating a StorageClass referencing a key. This encryption key is used to encrypt the disks created with that StorageClass. If your organization is required to manage its own key material, the CSI plugin provides the same functionality available in traditional CMEK for persistent disks in GKE. For this, you will need to create a Cloud KMS key to use for encryption, then you can create a StorageClass on Kubernetes that specifies the Cloud KMS key KMS_KEY_ID to use to encrypt the disk. <br>
Eg of storage class manifest:
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: csi-gce-pd
provisioner: pd.csi.storage.gke.io
parameters:
  type: pd-standard
  disk-encryption-kms-key: KMS_KEY_ID

Application-layer Secrets Encryption provides an additional layer of security for sensitive data, such as Secrets, stored in etcd. Using this functionality, you can use a key managed with Cloud KMS to encrypt data at the application layer. This protects against attackers who gain access to an offline copy of etcd.

To use Application-layer Secrets Encryption, you must first create a Cloud KMS key and give the GKE service account access to the key. The key must be in the same location as the cluster to decrease latency and to prevent cases where resources depend on services spread across multiple failure domains. Then, you can enable the feature on a new or existing cluster by specifying the key you would like to use.



Note: Default GKE cluster has encrypted persistent disks and encryption layers for secrets stored in etcd/cluster.

Migrating from a workload without Customer Managed Encryption Keys (CMEK) disks to an environment with CMEK disks will require the creation of new storage classes and Persistent Volume Claims. 
</p>
<td><ol type="1">
<li><p><a href="https://cloud.google.com/security/encryption-at-rest/default-encryption">Default encryption at rest on GCP: </a></p></li>
<li><p><a href="https://cloud.google.com/blog/products/containers-kubernetes/exploring-container-security-use-your-own-keys-to-protect-your-data-on-gke">Using your own keys to protect data on GKE:</a></p></li>
<li><p><a href="https://cloud.google.com/kubernetes-engine/docs/how-to/encrypting-secrets"Application Layer Secrets Encryption: </a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Encryption of data in-transit</td>
<td><ul>
<li><p>To support SSL connections, Amazon Redshift creates and installs an <a href="https://aws.amazon.com/certificate-manager/">AWS Certificate Manager (ACM)</a> issued SSL certificate on each cluster. The set of Certificate Authorities that you must trust in order to properly support SSL connections can be found at <a href="https://s3.amazonaws.com/redshift-downloads/redshift-ca-bundle.crt">https://s3.amazonaws.com/redshift-downloads/redshift-ca-bundle.crt</a>.</p></li>
</ul>
<blockquote>
<p>Note: Customers can import certificates into AWS ACM to use custom certs and still take advantage of the integration ACM has with Redshift.[3]</p>
</blockquote>
<ul>
<li><p>RedShift endpoints are available over HTTPS at a selection of regions.</p></li>
</ul>
<p>Best practice:</p>
<ul>
<li><p>Set the <strong>“require_SSL”</strong> parameter to <strong>“true”</strong> in the parameter group that is associated with the cluster.</p></li>
<li><p>For workloads that require FIPS-140-2 SSL compliance an additional step is required to set parameter <strong>“use_fips_ssl”</strong> to <strong>“true”</strong></p></li>
</ul></td>
<td><ol type="1">
<li><p>How to encrypt end to end: <a href="https://aws.amazon.com/blogs/big-data/encrypt-your-amazon-redshift-loads-with-amazon-s3-and-aws-kms/">https://aws.amazon.com/blogs/big-data/encrypt-your-amazon-redshift-loads-with-amazon-s3-and-aws-kms/</a></p></li>
<li><p>To make client side encryption work follow this pattern <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html">https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html</a></p></li>
<li><p>https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html</p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Encryption Key Management</td>
<td><ul>
<li><p>Key management is handled by KMS or HSM.</p></li>
<li><p>For encryption functions to occur in Redshift the DEK/CEK are stored on disk in an encrypted state. They persist on disk after cluster reboots and then require a request to KMS to use the CMK to decrypt the CEK to be used again in memory.</p></li>
<li><p>Key Rotation can occur as often as data requirement define.[1]</p>
<p>Example: Commandline key rotation[3]</p></li>
</ul>
<pre>rotate-encryption-key --cluster-identifier <value> [--cli-input-json <value>] [--generate-cli-skeleton <value>]</pre>
<p><strong>Note:</strong> Snapshots stored in S3 will need to be decrypted prior to key rotation and then re-encrypted using the new DEK. This is a process that should be tested prior to production use.</p></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.htm">https://docs.aws.amazon.com/kms/latest/developerguide/concepts.htm</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/cli/latest/reference/redshift/rotate-encryption-key.html">https://docs.aws.amazon.com/cli/latest/reference/redshift/rotate-encryption-key.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>Infrastructure</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Isolation of physical hosts</td>
<td><strong>N/A</strong>: RedShift is a fully-managed service and for cluster nodes the isolation of hosts is not currently possible via dedicated host ec2-resources. Reserved instances can be purchased to ensure availability of ec2 resource types.</td>
<td></td>
</tr>
<tr class="even">
<td>Network Isolation</td>
<td><p>When an Amazon Redshift cluster is provisioned, it is locked down by default so nobody has access to it except to IAM entities with Console access from within the provisioned network and with the default credentials. Amazon Redshift provides a cluster security group called default, which is created automatically when you launch your first cluster. Initially, this cluster security group is empty. You can add inbound access rules to the default cluster security group and then associate it with your Amazon Redshift cluster. To grant other users inbound access to an Amazon Redshift cluster, you associate the cluster with a security group. To grant access use an existing Amazon VPC security group or define a new one and then associate it with a cluster. For more information on managing a cluster on the EC2-VPC platform, see <a href="https://docs.aws.amazon.com/redshift/latest/mgmt/managing-clusters-vpc.html">Managing Clusters in an Amazon Virtual Private Cloud (VPC)</a>.</p>
<p>Amazon RedShift relies on EC2 security groups to provide infrastructure security, and thus initial protection from unauthorized traffic connecting to the cluster. [1]</p>
<p>Best Practice</p>
<ul>
<li><p>SecurityGroups should follow a naming convention for the entire account</p></li>
<li><p>The cluster leader node is the only EC2 instance that is allowed to communicate with the cluster nodes in the AWS Service Account. Ensure to enable VPC FlowLogs on the leader node ENI and capture logs from the OS to ensure only authorized access and activity has occurred.</p></li>
<li><p>For the leader node it is recommended the SecurityGroup have no outbound entries to prevent egress of data if the node is compromised.</p></li>
<li><p>Attempt to reference other SecurityGroups instead of using IP</p></li>
<li><p>Enable EnhancedVPCRouting [4]</p></li>
<li><p>Enable VPCEndpoint with S3</p></li>
<li><p>For most use cases the cluster should not be publically accessible.</p></li>
<li><p>Configure default port to authorized port for SecurityGroup usage. The default port is 5439 and cannot be changed after cluster is built.[5]</p></li>
</ul>
<p>See S3 Accelerator for controls around S3 and data isolation.</p></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/copy-unload-iam-role.html">https://docs.aws.amazon.com/redshift/latest/mgmt/copy-unload-iam-role.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-enabling-cluster.html">https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-enabling-cluster.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/gsg/rs-gsg-prereq.html">https://docs.aws.amazon.com/redshift/latest/gsg/rs-gsg-prereq.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>AWS Network</td>
<td><ul>
<li><p>A special use case exists for RedShift network isolation and must be noted but requires no action. When database encryption is enabled with KMS, KMS exports a CEK/DEK that is stored on a separate network from the cluster. This network is part of the managed service of RedShift and is not customer configurable or monitored.</p></li>
</ul>
<ul>
<li><p>Another note to mention is that Redshift clusters exist in another AWS account managed by AWS. This is important to be aware of for monitoring so it is clear what account traffic is going towards and coming from is actually authorized vs rogue.</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>IAM</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Admin Accounts</td>
<td><ul>
<li><p>A superuser account must be created to perform password disable/enable function (the default user created when a cluster is launched is called <strong>masteruser</strong>)</p></li>
<li><p>The IAM entity that creates the cluster is the default owner and first superuser.</p></li>
<li><p>A database superuser bypasses all permission checks. Be very careful when using a superuser role. It is recommended that you do most of your work as a role that is not a superuser. Superusers retain all privileges regardless of GRANT and REVOKE commands.</p></li>
<li><p>When you launch a new cluster using the AWS Management Console, AWS CLI, or Amazon Redshift API, you must supply a clear text password for the master database user.</p></li>
</ul>
<blockquote>
<p>To protect the password make sure to encrypt the plaintext password using the following command (assuming CMK used to encrypt cluster) “aws kms encrypt --key-id &lt;kms_key_id&gt; --plaintext &lt;password&gt;”</p>
</blockquote></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/dg/r_Privileges.html">https://docs.aws.amazon.com/redshift/latest/dg/r_Privileges.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_USER.html">https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_USER.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Role Based Access Control</td>
<td><ul>
<li><p>RedShift uses IAM principles to assign rights to actions. When different roles are created and mapped from customer domain groups to AWS IAM roles consider some best practices:</p>
<p>Limit potential over privilege by using redshift:RequestTag Condition key to limit any action to a specific deployment or environment:</p></li>
</ul>
<pre>
{
  "Version": "2012-10-17",
  "Statement": {
    "Sid":"AllowCreateProductionCluster",
    "Effect": "Allow",
    "Action": "redshift:CreateCluster",
    "Resource": "*"
    "Condition":{"StringEquals":{"redshift:RequestTag/usage":"production"}
  }
}
</pre>
<ul>
<li><p>It should go without saying but policies should not include “*” without having a following deny statement and/or condition statements</p>
<p>For example:</p>
<p>A condition statement to restrict access by redshift:ResourceTag Condition key</p></li>
</ul>
<pre>
{
  "Version": "2012-10-17",
  "Statement": {
    "Sid":"AllowModifyTestCluster",
    "Effect": "Allow",
    "Action": "redshift:ModifyCluster",
    "Resource": "arn:aws:redshift:us-west-2:123456789012:cluster:*"
    "Condition":{"StringEquals":{"redshift:ResourceTag/environment":"test"}
  }
}          
</pre>
<p>For example: A deny policy to limit actions to a specific Redshift Cluster environment “production*”.</p>
<pre>
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid":"AllowClusterManagement",
      "Action": [
        "redshift:CreateCluster",
        "redshift:DeleteCluster",
        "redshift:ModifyCluster",
        "redshift:RebootCluster"
      ],
      "Resource": [
        "*"
      ],
      "Effect": "Allow"
    },
    {
      "Sid":"DenyDeleteModifyProtected",
      "Action": [
        "redshift:DeleteCluster",
        "redshift:ModifyCluster"
      ],
      "Resource": [
        "arn:aws:redshift:us-west-2:123456789012:cluster:production*"
      ],
      "Effect": "Deny"
    }
  ]
}
</pre>
<ul>
<li><p>Ensure you separate roles that perform snapshot/restore from regular users. Remember to test this as IAM assumes deny but if a broad access role is assigned and a regular user can assume that role then they user can perform the API action. The only way to prevent this is to have a condition statement limiting actions to a specific role or an explicit deny.</p></li>
<li><p>Become very familiar with all API actions for RedShift [6]</p></li>
<li><p>To properly manage a Service-Linked role be sure to separate the "create" and "delete" actions to unique IAM entity so access to manage data and manage the cluster are separated.</p>
<p>For example: Allow IAM identity to delete Service-Linked Role</p>
<pre>
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid":"AllowClusterManagement",
      "Action": [
        "redshift:CreateCluster",
        "redshift:DeleteCluster",
        "redshift:ModifyCluster",
        "redshift:RebootCluster"
      ],
      "Resource": [
        "*"
      ],
      "Effect": "Allow"
    },
    {
      "Sid":"DenyDeleteModifyProtected",
      "Action": [
        "redshift:DeleteCluster",
        "redshift:ModifyCluster"
      ],
      "Resource": [
        "arn:aws:redshift:us-west-2:123456789012:cluster:production*"
      ],
      "Effect": "Deny"
    }
  ]
}
</pre>
</li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://aws.amazon.com/iam/details/manage-federation/"><span class="underline">Managing federation in AWS IAM</span></a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/iam-redshift-user-mgmt.html">https://docs.aws.amazon.com/redshift/latest/mgmt/iam-redshift-user-mgmt.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html">https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-policy-resources.resource-permissions.html">https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-policy-resources.resource-permissions.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/using-service-linked-roles.html">https://docs.aws.amazon.com/redshift/latest/mgmt/using-service-linked-roles.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonredshift.html">https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonredshift.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Authorization between AWS services</td>
<td><ul>
<li><p>Redshift will need to access other AWS services (S3, Glue, Athena, KMS, etc..., to perform functions in an automated way. To setup a role for a specific service function use reference [1].</p></li>
<li><p>It is recommended to run a Service Linked Role for RedShift to limit service specific actions to only the RDS service endpoint [2]</p></li>
<li><p>You don't need to manually create an AWSServiceRoleForRedshift service-linked role. Amazon Redshift creates the service-linked role for you. If the AWSServiceRoleForRedshift service-linked role has been deleted from your account, Amazon Redshift creates the role when you launch a new Amazon Redshift cluster.</p></li>
</ul>
<blockquote>
<p>.</p>
</blockquote></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/authorizing-redshift-service.html">https://docs.aws.amazon.com/redshift/latest/mgmt/authorizing-redshift-service.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/using-service-linked-roles.html">https://docs.aws.amazon.com/redshift/latest/mgmt/using-service-linked-roles.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Authentication to AWS platform</td>
<td><p>RedShift supports local and IAM based authentication. To align to best practice Redshift local users should have passwords disabled which forces authentication based on IAM.</p>
<blockquote>
<p>To generate credentials for an IAM user Granting permission to GetClusterCredentials API action should be limited to authorized IAM entities and limited to only specific cluster, database, usernames, and group names.</p>
<p>aws redshift get-cluster-credentials --cluster-identifier examplecluster --db-user temp_creds_user --db-name exampledb --duration-seconds 3600</p>
</blockquote>
<ul>
<li><p>Db credentials should make use of IAM</p></li>
<li><p>This means db users password must be disabled which forces login credentials based on temp IAM credentials</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/options-for-providing-iam-credentials.html">https://docs.aws.amazon.com/redshift/latest/mgmt/options-for-providing-iam-credentials.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html">https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Authorization (AWS IAM) of corporate users via Active Directory for access to RedShift resources.</td>
<td><p>Amazon Redshift requires IAM credentials that AWS can use to authenticate your requests. Those credentials must have permissions to access AWS resources, such as an Amazon Redshift cluster.</p>
<p>AWS Identity and Access Management (IAM) enable organizations with multiple employees to create and manage multiple users, groups and roles under a single AWS account. With IAM policies, companies can grant IAM users/groups/roles fine-grained control to their Amazon RedShift data while also retaining full control over everything the users do.</p>
<p>Most customers have setup federation with AWS accounts. Therefore, the only decision to make is what API actions are needed for roles, groups, or users within IAM [2].</p>
<p>To combine these concepts to control access to RedShift resources, a user would:</p>
<ul>
<li><p>Determine how authentication will occur with Redshift[2]</p></li>
<li><p>Create the appropriate IAM roles [5]</p>
<p>For example, a policy to allow IAM users to request credentials for temporary access. The following policy enables the GetCredentials, CreateCluserUser, and JoinGroup actions. The policy uses condition keys to allow the GetClusterCredentials and CreateClusterUser actions only when the AWS user ID matches "AIDIODR4TAW7CSEXAMPLE:${redshift:DbUser}@yourdomain.com". IAM access is requested for the "testdb" database only. The policy also allows users to join a group named "common_group".</p></li>
</ul>
<p>{</p>
<p>"Version": "2012-10-17",</p>
<p>"Statement": [</p>
<p>{</p>
<p>"Sid": "GetClusterCredsStatement",</p>
<p>"Effect": "Allow",</p>
<p>"Action": [</p>
<p>"redshift:GetClusterCredentials"</p>
<p>],</p>
<p>"Resource": [</p>
<p>"arn:aws:redshift:us-west-2:123456789012:dbuser:examplecluster/${redshift:DbUser}",</p>
<p>"arn:aws:redshift:us-west-2:123456789012:dbname:examplecluster/testdb",</p>
<p>"arn:aws:redshift:us-west-2:123456789012:dbgroup:examplecluster/common_group"</p>
<p>],</p>
<p>"Condition": {</p>
<p>"StringEquals": {</p>
<p>"aws:userid":"AIDIODR4TAW7CSEXAMPLE:${redshift:DbUser}@yourdomain.com"</p>
<p>}</p>
<p>}</p>
<p>},</p>
<p>{</p>
<p>"Sid": "CreateClusterUserStatement",</p>
<p>"Effect": "Allow",</p>
<p>"Action": [</p>
<p>"redshift:CreateClusterUser"</p>
<p>],</p>
<p>"Resource": [</p>
<p>"arn:aws:redshift:us-west-2:123456789012:dbuser:examplecluster/${redshift:DbUser}"</p>
<p>],</p>
<p>"Condition": {</p>
<p>"StringEquals": {</p>
<p>"aws:userid":"AIDIODR4TAW7CSEXAMPLE:${redshift:DbUser}@yourdomain.com"</p>
<p>}</p>
<p>}</p>
<p>},</p>
<p>{</p>
<p>"Sid": "RedshiftJoinGroupStatement",</p>
<p>"Effect": "Allow",</p>
<p>"Action": [</p>
<p>"redshift:JoinGroup"</p>
<p>],</p>
<p>"Resource": [</p>
<p>"arn:aws:redshift:us-west-2:123456789012:dbgroup:examplecluster/common_group"</p>
<p>]</p>
<p>}</p>
<p>]</p>
<p>}</p>
<p>}</p>
<p>}</p>
<ul>
<li><p>Map the appropriate AD groups to those roles</p></li>
<li><p>Determine if IAM users are allowed to create user credentials within RedShift [4]</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/options-for-providing-iam-credentials.html">https://docs.aws.amazon.com/redshift/latest/mgmt/options-for-providing-iam-credentials.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-overview.html">Overview of Managing Access Permissions to Your Amazon Redshift Resources</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html">Using Identity-Based Policies (IAM Policies) for Amazon Redshift</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html">https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html">https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Tagging</td>
<td><p>Amazon Redshift supports tagging to provide metadata about resources at a glance, IAM enforcement, and to categorize your billing reports based on cost allocation.</p>
<p>Best Practice:</p>
<p>Since tags can be used to enforce IAM entity abilities it is important to set appropriate controls on changes to tags with API actions like (DeleteTags, CreateTags)</p>
<p><strong>Note: tags are not retained if you copy a snapshot to another region, so you must recreate the tags in the new region</strong></p></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-tagging.html">https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-tagging.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td><strong>Logging &amp; Monitoring</strong></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td>Logging activity within RedShift</td>
<td><p>Audit logging is not enabled by default in Amazon Redshift. When you enable logging on your cluster, Amazon Redshift creates and uploads logs to Amazon S3 that capture data from the creation of the cluster to the present time.</p>
<p>To meet logging requirements make sure to enable audit logging:</p>
<ul>
<li><p>Connection log</p></li>
<li><p>User log</p></li>
<li><p>User activity log (requires additional step after enable of audit logging)</p></li>
</ul>
<ul>
<li><p>To enable this you must enable the <strong>“enable_user_activity_logging”</strong> database parameter[2]</p>
<p>For example:</p>
<p>aws redshift modify-cluster-parameter-group</p>
<p>--parameter-group-name myclusterparametergroup</p>
<p>--parameters ParameterName=statement_timeout,ParameterValue=20000 ParameterName=enable_user_activity_logging,ParameterValue=true</p></li>
</ul>
<p>As a best practice, after a configuration of RedShift is found to be functional and meet requirements make sure to commit all settings into a <a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html">parameter group</a> so all databases within a cluster are configured the same and each new cluster can be configured the same. (a final deployed cluster should not have parameter group = default.redshift-1.0 because this will not enable logging or other settings specific to customer requirements.)</p></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html">https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Logging API actions</td>
<td>Cloudtrail will be enabled in every account as part of a default account build.</td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#rs-db-auditing-cloud-trail">https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#rs-db-auditing-cloud-trail</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Alerting and Incident Management</td>
<td><p>You can use the following automated monitoring tools to watch RedShift and report when something is wrong:</p>
<ul>
<li><p><strong>Amazon CloudWatch Alarms</strong> – Watch a single metric over a time period that you specify, and perform one or more actions based on the value of the metric relative to a given threshold over a number of time periods. The action is a notification sent to an Amazon Simple Notification Service (Amazon SNS) topic or Auto Scaling policy. CloudWatch alarms do not invoke actions simply because they are in a particular state; the state must have changed and been maintained for a specified number of periods. For more information, see <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/rs-metricscollected.html">https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/rs-metricscollected.html</a>.</p></li>
<li><p><strong>AWS CloudTrail Log Monitoring</strong> – Share log files between accounts, monitor CloudTrail log files in real time by sending them to CloudWatch Logs, write log processing applications in Java, and validate that your log files have not changed after delivery by CloudTrail. For more information, see <a href="https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#rs-db-auditing-cloud-trail">Logging Amazon RedShift API Calls By Using AWS CloudTrail</a>.</p></li>
</ul>
<p><strong>Implementation Note:</strong> To access logs and data for monitoring the data must be decrypted. To decrypt logs/data a customer managed CMK must be defined. Use the same CMK created to encrypt the cluster and create a new policy to grant access only to API actions necessary for tables and actions that are authorized. For example: Use guides in reference [1] where cloudformation templates already exist and can be used to provide a prescriptive approach to collecting and monitoring logs.</p>
<p>Make note of the minimum requirements for access to the Redshift user that is required. Be cautious not to enable more than the necessary “<strong>grant select on all tables in schema pg_catalog to tamreporting</strong>” entitlement.</p>
<p><strong>Note</strong></p>
<p>Audit logging to Amazon S3 is an optional, manual process. When you enable logging on your cluster, you are enabling logging to Amazon S3 only. Logging to system tables is not optional and happens automatically for the cluster. For more information about logging to system tables, see <a href="http://docs.aws.amazon.com/redshift/latest/dg/cm_chap_system-tables.html">System Tables Reference</a> in the Amazon Redshift Database Developer Guide.</p></td>
<td><ol type="1">
<li><p><a href="https://github.com/awslabs/amazon-redshift-monitoring">https://github.com/awslabs/amazon-redshift-monitoring</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/metrics-listing.html">https://docs.aws.amazon.com/redshift/latest/mgmt/metrics-listing.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td><strong>Patch/Updates</strong></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td>Update/Patch for RedShift</td>
<td><p>When Redshift is configured a maintenance window must be defined to allow for updates to be installed. This window should align with expected change management times and a complete understanding of outages, if any, will occur during this window.</p>
<p>Additional considerations include:</p>
<ul>
<li><p>Refer to <a href="https://docs.aws.amazon.com/cli/latest/index.html">https://docs.aws.amazon.com/cli/latest/index.html</a> to understand what API action will require a reboot of the cluster or will not.</p></li>
<li><p>Understand automatic version updates may change expected behavior and a setting is offered to prevent cluster version changes without approval.</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://aws.amazon.com/premiumsupport/knowledge-center/notification-maintenance-rds-redshift/">https://aws.amazon.com/premiumsupport/knowledge-center/notification-maintenance-rds-redshift/</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td><strong>Availability</strong></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td>Backup and Restore</td>
<td><ul>
<li><p>Amazon RedShift makes use of Snapshots to provide customers with a way of recovering to an RPO. Snapshots are stored in Amazon S3, managed by AWS; Snapshots are transferred to S3 over SSL, and where the data in the database is already encrypted in the cluster, it remains encrypted in the snapshot too. [1]</p></li>
<li><p>If you enable copying of snapshots from an encrypted cluster and use AWS KMS for your master key, you cannot rename your cluster because the cluster name is part of the encryption context. If you must rename your cluster, you can disable copying of snapshots in the source region, rename the cluster, and then configure and enable copying of snapshots again. [2]</p></li>
<li><p>Important Note: If you rotate a DEK/CEK that is used to encrypt a cluster all data will be encrypted with the new key except for snapshots stored in S3. A process should be developed to ensure snapshots are encrypted with the new key to ensure recovery point objectives (RPO) are met.</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html</a></p></li>
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html">https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Limits</td>
<td><p>Understanding the limitations of the service can help prevent unintentional outages or ability to meet requirements.</p>
<p>Items like:</p>
<ul>
<li><p># of tables by ec2 instance type</p></li>
<li><p>Spectrum limits</p></li>
<li><p>Quotas</p></li>
<li><p>IAM roles allowed</p></li>
<li><p>Naming constraints</p></li>
</ul></td>
<td><ol type="1">
<li><p><a href="https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html">https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html</a></p></li>
</ol></td>
</tr>
</tbody>
</table>
