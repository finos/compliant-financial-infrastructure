---
title: 'AWS Service Accelerator: RedShift'
---

Date: 2019-04-08

#######  

####### 

####### Detailed Security Configuration

Overview

This section is meant to provide an opinionated approach towards the
implementation security controls by domain. Although the approaches may
not be a fit for all use casees or complete for production use, they are
meant to guide the audience towards current best practices and design
considerations to achieve security control objectives.

**Controls and Architectures**

This table maps Security Domain to the corresponding controls and
architectural best practices as documented in AWS' public documentation,
white papers, and blog posts.

+-----------------------+-----------------------+-----------------------+
| **Security Domain**   | **Control &           | **References**        |
|                       | Architectural         |                       |
|                       | Suggestions**         |                       |
+-----------------------+-----------------------+-----------------------+
| **Encryption**        |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Encryption of data    | AWS RedShift supports | 1.  <https://docs.aws |
| at-rest               | KMS and HSM to        | .amazon.com/redshift/ |
|                       | provide key material  | latest/mgmt/working-w |
|                       | management and        | ith-db-encryption.htm |
|                       | encryption services.  | l>                    |
|                       | Encryption at rest of |                       |
|                       | RedShift encrypts the | 2.  [Migrating an     |
|                       | data blocks and       |     Unencrypted       |
|                       | system metadata of    |     Cluster to an     |
|                       | the cluster and its   |     Encrypted         |
|                       | snapshots.\[1\]       |     Cluster](https:// |
|                       |                       | docs.aws.amazon.com/r |
|                       | **Note:** By default, | edshift/latest/mgmt/m |
|                       | Amazon Redshift       | igrating-to-an-encryp |
|                       | selects the account   | ted-cluster.html)     |
|                       | service default key   |                       |
|                       | as the master key.    | 3.  <https://docs.aws |
|                       | The default key is an | .amazon.com/cli/lates |
|                       | AWS-managed key that  | t/reference/redshift/ |
|                       | is created for your   | describe-clusters.htm |
|                       | AWS account to use in | l>                    |
|                       | Amazon Redshift. Some |                       |
|                       | customer security     |                       |
|                       | controls prevent the  |                       |
|                       | use of default        |                       |
|                       | service KMS keys for  |                       |
|                       | sensitive workloads.  |                       |
|                       | Users should          |                       |
|                       | pre-create a customer |                       |
|                       | managed CMK for       |                       |
|                       | RedShift usage.       |                       |
|                       |                       |                       |
|                       | **Implementation      |                       |
|                       | Note:** Encryption is |                       |
|                       | an optional,          |                       |
|                       | immutable setting of  |                       |
|                       | a cluster. If you     |                       |
|                       | want encryption, you  |                       |
|                       | enable it during the  |                       |
|                       | cluster launch        |                       |
|                       | process. To go from   |                       |
|                       | an unencrypted        |                       |
|                       | cluster to an         |                       |
|                       | encrypted cluster or  |                       |
|                       | the other way around, |                       |
|                       | unload your data from |                       |
|                       | the existing cluster  |                       |
|                       | and reload it in a    |                       |
|                       | new cluster with the  |                       |
|                       | chosen encryption     |                       |
|                       | setting. \[2\]        |                       |
|                       |                       |                       |
|                       | For S3 encryption     |                       |
|                       | details see S3        |                       |
|                       | Accelerator.          |                       |
|                       |                       |                       |
|                       | For KMS details see   |                       |
|                       | KMS Accelerator       |                       |
+-----------------------+-----------------------+-----------------------+
| Encryption of data    | -   To support SSL    | 1.  How to encrypt    |
| in-transit            |     connections,      |     end to end:       |
|                       |     Amazon Redshift   |     <https://aws.amaz |
|                       |     creates and       | on.com/blogs/big-data |
|                       |     installs an [AWS  | /encrypt-your-amazon- |
|                       |     Certificate       | redshift-loads-with-a |
|                       |     Manager           | mazon-s3-and-aws-kms/ |
|                       |     (ACM)](https://aw | >                     |
|                       | s.amazon.com/certific |                       |
|                       | ate-manager/) issued  | 2.  To make client    |
|                       |     SSL certificate   |     side encryption   |
|                       |     on each cluster.  |     work follow this  |
|                       |     The set of        |     pattern           |
|                       |     Certificate       |     <https://docs.aws |
|                       |     Authorities that  | .amazon.com/AmazonS3/ |
|                       |     you must trust in | latest/dev/UsingClien |
|                       |     order to properly | tSideEncryption.html> |
|                       |     support SSL       |                       |
|                       |     connections can   | 3.  https://docs.aws. |
|                       |     be found          | amazon.com/acm/latest |
|                       |     at <https://s3.am | /userguide/import-cer |
|                       | azonaws.com/redshift- | tificate.html         |
|                       | downloads/redshift-ca |                       |
|                       | -bundle.crt>.         |                       |
|                       |                       |                       |
|                       | > Note: Customers can |                       |
|                       | > import certificates |                       |
|                       | > into AWS ACM to use |                       |
|                       | > custom certs and    |                       |
|                       | > still take          |                       |
|                       | > advantage of the    |                       |
|                       | > integration ACM has |                       |
|                       | > with Redshift.\[3\] |                       |
|                       |                       |                       |
|                       | -   RedShift          |                       |
|                       |     endpoints are     |                       |
|                       |     available over    |                       |
|                       |     HTTPS at a        |                       |
|                       |     selection of      |                       |
|                       |     regions.          |                       |
|                       |                       |                       |
|                       | Best practice:        |                       |
|                       |                       |                       |
|                       | -   Set               |                       |
|                       |     the **"require\_S |                       |
|                       | SL"** parameter       |                       |
|                       |     to **"true"** in  |                       |
|                       |     the parameter     |                       |
|                       |     group that is     |                       |
|                       |     associated with   |                       |
|                       |     the cluster.      |                       |
|                       |                       |                       |
|                       | -   For workloads     |                       |
|                       |     that require      |                       |
|                       |     FIPS-140-2 SSL    |                       |
|                       |     compliance an     |                       |
|                       |     additional step   |                       |
|                       |     is required to    |                       |
|                       |     set parameter     |                       |
|                       |     **"use\_fips\_ssl |                       |
|                       | "**                   |                       |
|                       |     to **"true"**     |                       |
+-----------------------+-----------------------+-----------------------+
| Encryption Key        | -   Key management is | 1.  <https://docs.aws |
| Management            |     handled by KMS or | .amazon.com/redshift/ |
|                       |     HSM.              | latest/mgmt/working-w |
|                       |                       | ith-db-encryption.htm |
|                       | -   For encryption    | l>                    |
|                       |     functions to      |                       |
|                       |     occur in Redshift | 2.  <https://docs.aws |
|                       |     the DEK/CEK are   | .amazon.com/kms/lates |
|                       |     stored on disk in | t/developerguide/conc |
|                       |     an encrypted      | epts.htm>             |
|                       |     state. They       |                       |
|                       |     persist on disk   | 3.  <https://docs.aws |
|                       |     after cluster     | .amazon.com/cli/lates |
|                       |     reboots and then  | t/reference/redshift/ |
|                       |     require a request | rotate-encryption-key |
|                       |     to KMS to use the | .html>                |
|                       |     CMK to decrypt    |                       |
|                       |     the CEK to be     |                       |
|                       |     used again in     |                       |
|                       |     memory.           |                       |
|                       |                       |                       |
|                       | -   Key Rotation can  |                       |
|                       |     occur as often as |                       |
|                       |     data requirement  |                       |
|                       |     define.\[1\]      |                       |
|                       |                       |                       |
|                       |     Example:          |                       |
|                       |     Commandline key   |                       |
|                       |     rotation\[3\]     |                       |
|                       |                       |                       |
|                       | rotate-encryption-key |                       |
|                       |                       |                       |
|                       | \--cluster-identifier |                       |
|                       | \<value\>             |                       |
|                       |                       |                       |
|                       | \[\--cli-input-json   |                       |
|                       | \<value\>\]           |                       |
|                       |                       |                       |
|                       | \[\--generate-cli-ske |                       |
|                       | leton                 |                       |
|                       | \<value\>\]           |                       |
|                       |                       |                       |
|                       | **Note:** Snapshots   |                       |
|                       | stored in S3 will     |                       |
|                       | need to be decrypted  |                       |
|                       | prior to key rotation |                       |
|                       | and then re-encrypted |                       |
|                       | using the new DEK.    |                       |
|                       | This is a process     |                       |
|                       | that should be tested |                       |
|                       | prior to production   |                       |
|                       | use.                  |                       |
+-----------------------+-----------------------+-----------------------+
| **Infrastructure**    |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Isolation of physical | **N/A**: RedShift is  |                       |
| hosts                 | a fully-managed       |                       |
|                       | service and for       |                       |
|                       | cluster nodes the     |                       |
|                       | isolation of hosts is |                       |
|                       | not currently         |                       |
|                       | possible via          |                       |
|                       | dedicated host        |                       |
|                       | ec2-resources.        |                       |
|                       | Reserved instances    |                       |
|                       | can be purchased to   |                       |
|                       | ensure availability   |                       |
|                       | of ec2 resource       |                       |
|                       | types.                |                       |
+-----------------------+-----------------------+-----------------------+
| Network Isolation     | When an Amazon        | 1.  <https://docs.aws |
|                       | Redshift cluster is   | .amazon.com/redshift/ |
|                       | provisioned, it is    | latest/mgmt/working-w |
|                       | locked down by        | ith-security-groups.h |
|                       | default so nobody has | tml>                  |
|                       | access to it except   |                       |
|                       | to IAM entities with  | 2.  <https://docs.aws |
|                       | Console access from   | .amazon.com/redshift/ |
|                       | within the            | latest/mgmt/working-w |
|                       | provisioned network   | ith-security-groups.h |
|                       | and with the default  | tml>                  |
|                       | credentials. Amazon   |                       |
|                       | Redshift provides a   | 3.  <https://docs.aws |
|                       | cluster security      | .amazon.com/redshift/ |
|                       | group called default, | latest/mgmt/copy-unlo |
|                       | which is created      | ad-iam-role.html>     |
|                       | automatically when    |                       |
|                       | you launch your first | 4.  <https://docs.aws |
|                       | cluster. Initially,   | .amazon.com/redshift/ |
|                       | this cluster security | latest/mgmt/enhanced- |
|                       | group is empty. You   | vpc-enabling-cluster. |
|                       | can add inbound       | html>                 |
|                       | access rules to the   |                       |
|                       | default cluster       | 5.  <https://docs.aws |
|                       | security group and    | .amazon.com/redshift/ |
|                       | then associate it     | latest/gsg/rs-gsg-pre |
|                       | with your Amazon      | req.html>             |
|                       | Redshift cluster. To  |                       |
|                       | grant other users     |                       |
|                       | inbound access to an  |                       |
|                       | Amazon Redshift       |                       |
|                       | cluster, you          |                       |
|                       | associate the cluster |                       |
|                       | with a security       |                       |
|                       | group. To grant       |                       |
|                       | access use an         |                       |
|                       | existing Amazon VPC   |                       |
|                       | security group or     |                       |
|                       | define a new one and  |                       |
|                       | then associate it     |                       |
|                       | with a cluster. For   |                       |
|                       | more information on   |                       |
|                       | managing a cluster on |                       |
|                       | the EC2-VPC platform, |                       |
|                       | see [Managing         |                       |
|                       | Clusters in an Amazon |                       |
|                       | Virtual Private Cloud |                       |
|                       | (VPC)](https://docs.a |                       |
|                       | ws.amazon.com/redshif |                       |
|                       | t/latest/mgmt/managin |                       |
|                       | g-clusters-vpc.html). |                       |
|                       |                       |                       |
|                       | Amazon RedShift       |                       |
|                       | relies on EC2         |                       |
|                       | security groups to    |                       |
|                       | provide               |                       |
|                       | infrastructure        |                       |
|                       | security, and thus    |                       |
|                       | initial protection    |                       |
|                       | from unauthorized     |                       |
|                       | traffic connecting to |                       |
|                       | the cluster. \[1\]    |                       |
|                       |                       |                       |
|                       | Best Practice         |                       |
|                       |                       |                       |
|                       | -   SecurityGroups    |                       |
|                       |     should follow a   |                       |
|                       |     naming convention |                       |
|                       |     for the entire    |                       |
|                       |     account           |                       |
|                       |                       |                       |
|                       | -   The cluster       |                       |
|                       |     leader node is    |                       |
|                       |     the only EC2      |                       |
|                       |     instance that is  |                       |
|                       |     allowed to        |                       |
|                       |     communicate with  |                       |
|                       |     the cluster nodes |                       |
|                       |     in the AWS        |                       |
|                       |     Service Account.  |                       |
|                       |     Ensure to enable  |                       |
|                       |     VPC FlowLogs on   |                       |
|                       |     the leader node   |                       |
|                       |     ENI and capture   |                       |
|                       |     logs from the OS  |                       |
|                       |     to ensure only    |                       |
|                       |     authorized access |                       |
|                       |     and activity has  |                       |
|                       |     occurred.         |                       |
|                       |                       |                       |
|                       | -   For the leader    |                       |
|                       |     node it is        |                       |
|                       |     recommended the   |                       |
|                       |     SecurityGroup     |                       |
|                       |     have no outbound  |                       |
|                       |     entries to        |                       |
|                       |     prevent egress of |                       |
|                       |     data if the node  |                       |
|                       |     is compromised.   |                       |
|                       |                       |                       |
|                       | -   Attempt to        |                       |
|                       |     reference other   |                       |
|                       |     SecurityGroups    |                       |
|                       |     instead of using  |                       |
|                       |     IP                |                       |
|                       |                       |                       |
|                       | -   Enable            |                       |
|                       |     EnhancedVPCRoutin |                       |
|                       | g                     |                       |
|                       |     \[4\]             |                       |
|                       |                       |                       |
|                       | -   Enable            |                       |
|                       |     VPCEndpoint with  |                       |
|                       |     S3                |                       |
|                       |                       |                       |
|                       | -   For most use      |                       |
|                       |     cases the cluster |                       |
|                       |     should not be     |                       |
|                       |     publically        |                       |
|                       |     accessible.       |                       |
|                       |                       |                       |
|                       | -   Configure default |                       |
|                       |     port to           |                       |
|                       |     authorized port   |                       |
|                       |     for SecurityGroup |                       |
|                       |     usage. The        |                       |
|                       |     default port is   |                       |
|                       |     5439 and cannot   |                       |
|                       |     be changed after  |                       |
|                       |     cluster is        |                       |
|                       |     built.\[5\]       |                       |
|                       |                       |                       |
|                       | See S3 Accelerator    |                       |
|                       | for controls around   |                       |
|                       | S3 and data           |                       |
|                       | isolation.            |                       |
+-----------------------+-----------------------+-----------------------+
| AWS Network           | -   A special use     | 1.  <https://docs.aws |
|                       |     case exists for   | .amazon.com/redshift/ |
|                       |     RedShift network  | latest/mgmt/working-w |
|                       |     isolation and     | ith-db-encryption.htm |
|                       |     must be noted but | l>                    |
|                       |     requires no       |                       |
|                       |     action. When      |                       |
|                       |     database          |                       |
|                       |     encryption is     |                       |
|                       |     enabled with KMS, |                       |
|                       |     KMS exports a     |                       |
|                       |     CEK/DEK that is   |                       |
|                       |     stored on a       |                       |
|                       |     separate network  |                       |
|                       |     from the cluster. |                       |
|                       |     This network is   |                       |
|                       |     part of the       |                       |
|                       |     managed service   |                       |
|                       |     of RedShift and   |                       |
|                       |     is not customer   |                       |
|                       |     configurable or   |                       |
|                       |     monitored.        |                       |
|                       |                       |                       |
|                       | <!-- -->              |                       |
|                       |                       |                       |
|                       | -   Another note to   |                       |
|                       |     mention is that   |                       |
|                       |     Redshift clusters |                       |
|                       |     exist in another  |                       |
|                       |     AWS account       |                       |
|                       |     managed by AWS.   |                       |
|                       |     This is important |                       |
|                       |     to be aware of    |                       |
|                       |     for monitoring so |                       |
|                       |     it is clear what  |                       |
|                       |     account traffic   |                       |
|                       |     is going towards  |                       |
|                       |     and coming from   |                       |
|                       |     is actually       |                       |
|                       |     authorized vs     |                       |
|                       |     rogue.            |                       |
+-----------------------+-----------------------+-----------------------+
| **IAM**               |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Admin Accounts        | -   A superuser       | 1.  <https://docs.aws |
|                       |     account must be   | .amazon.com/redshift/ |
|                       |     created to        | latest/dg/r_Privilege |
|                       |     perform password  | s.html>               |
|                       |     disable/enable    |                       |
|                       |     function (the     | 2.  <https://docs.aws |
|                       |     default user      | .amazon.com/redshift/ |
|                       |     created when a    | latest/dg/r_CREATE_US |
|                       |     cluster is        | ER.html>              |
|                       |     launched is       |                       |
|                       |     called            |                       |
|                       |     **masteruser**)   |                       |
|                       |                       |                       |
|                       | -   The IAM entity    |                       |
|                       |     that creates the  |                       |
|                       |     cluster is the    |                       |
|                       |     default owner and |                       |
|                       |     first superuser.  |                       |
|                       |                       |                       |
|                       | -   A database        |                       |
|                       |     superuser         |                       |
|                       |     bypasses all      |                       |
|                       |     permission        |                       |
|                       |     checks. Be very   |                       |
|                       |     careful when      |                       |
|                       |     using a superuser |                       |
|                       |     role. It is       |                       |
|                       |     recommended that  |                       |
|                       |     you do most of    |                       |
|                       |     your work as a    |                       |
|                       |     role that is not  |                       |
|                       |     a superuser.      |                       |
|                       |     Superusers retain |                       |
|                       |     all privileges    |                       |
|                       |     regardless of     |                       |
|                       |     GRANT and REVOKE  |                       |
|                       |     commands.         |                       |
|                       |                       |                       |
|                       | -   When you launch a |                       |
|                       |     new cluster using |                       |
|                       |     the AWS           |                       |
|                       |     Management        |                       |
|                       |     Console, AWS CLI, |                       |
|                       |     or Amazon         |                       |
|                       |     Redshift API, you |                       |
|                       |     must supply a     |                       |
|                       |     clear text        |                       |
|                       |     password for the  |                       |
|                       |     master database   |                       |
|                       |     user.             |                       |
|                       |                       |                       |
|                       | > To protect the      |                       |
|                       | > password make sure  |                       |
|                       | > to encrypt the      |                       |
|                       | > plaintext password  |                       |
|                       | > using the following |                       |
|                       | > command (assuming   |                       |
|                       | > CMK used to encrypt |                       |
|                       | > cluster) "aws kms   |                       |
|                       | > encrypt \--key-id   |                       |
|                       | > \<kms\_key\_id\>    |                       |
|                       | > \--plaintext        |                       |
|                       | > \<password\>"       |                       |
+-----------------------+-----------------------+-----------------------+
| Role Based Access     | -   RedShift uses IAM | 1.  [[Managing        |
| Control               |     principles to     |     federation in AWS |
|                       |     assign rights to  |     IAM]{.underline}] |
|                       |     actions. When     | (https://aws.amazon.c |
|                       |     different roles   | om/iam/details/manage |
|                       |     are created and   | -federation/)         |
|                       |     mapped from       |                       |
|                       |     customer domain   | 2.  <https://docs.aws |
|                       |     groups to AWS IAM | .amazon.com/redshift/ |
|                       |     roles consider    | latest/mgmt/iam-redsh |
|                       |     some best         | ift-user-mgmt.html>   |
|                       |     practices:        |                       |
|                       |                       | 3.  <https://docs.aws |
|                       |     Limit potential   | .amazon.com/redshift/ |
|                       |     over privilege by | latest/mgmt/redshift- |
|                       |     using             | iam-access-control-id |
|                       |     redshift:RequestT | entity-based.html>    |
|                       | ag                    |                       |
|                       |     Condition key to  | 4.  <https://docs.aws |
|                       |     limit any action  | .amazon.com/redshift/ |
|                       |     to a specific     | latest/mgmt/redshift- |
|                       |     deployment or     | policy-resources.reso |
|                       |     environment:      | urce-permissions.html |
|                       |                       | >                     |
|                       | {                     |                       |
|                       |                       | 5.  <https://docs.aws |
|                       | \"Version\":          | .amazon.com/redshift/ |
|                       | \"2012-10-17\",       | latest/mgmt/using-ser |
|                       |                       | vice-linked-roles.htm |
|                       | \"Statement\": {      | l>                    |
|                       |                       |                       |
|                       | \"Sid\":\"AllowCreate | 6.  <https://docs.aws |
|                       | ProductionCluster\",  | .amazon.com/IAM/lates |
|                       |                       | t/UserGuide/list_amaz |
|                       | \"Effect\":           | onredshift.html>      |
|                       | \"Allow\",            |                       |
|                       |                       |                       |
|                       | \"Action\":           |                       |
|                       | \"redshift:CreateClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"Resource\": \"\*\"  |                       |
|                       |                       |                       |
|                       | \"Condition\":{\"Stri |                       |
|                       | ngEquals\":{\"redshif |                       |
|                       | t:RequestTag/usage\": |                       |
|                       | \"production\"}       |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | -   It should go      |                       |
|                       |     without saying    |                       |
|                       |     but policies      |                       |
|                       |     should not        |                       |
|                       |     include "\*"      |                       |
|                       |     without having a  |                       |
|                       |     following deny    |                       |
|                       |     statement and/or  |                       |
|                       |     condition         |                       |
|                       |     statements        |                       |
|                       |                       |                       |
|                       |     For example:      |                       |
|                       |                       |                       |
|                       |     A condition       |                       |
|                       |     statement to      |                       |
|                       |     restrict access   |                       |
|                       |     by                |                       |
|                       |     redshift:Resource |                       |
|                       | Tag                   |                       |
|                       |     Condition key     |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Version\":          |                       |
|                       | \"2012-10-17\",       |                       |
|                       |                       |                       |
|                       | \"Statement\": {      |                       |
|                       |                       |                       |
|                       | \"Sid\":\"AllowModify |                       |
|                       | TestCluster\",        |                       |
|                       |                       |                       |
|                       | \"Effect\":           |                       |
|                       | \"Allow\",            |                       |
|                       |                       |                       |
|                       | \"Action\":           |                       |
|                       | \"redshift:ModifyClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"Resource\":         |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | cluster:\*\"          |                       |
|                       |                       |                       |
|                       | \"Condition\":{\"Stri |                       |
|                       | ngEquals\":{\"redshif |                       |
|                       | t:ResourceTag/environ |                       |
|                       | ment\":\"test\"}      |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | For example: A deny   |                       |
|                       | policy to limit       |                       |
|                       | actions to a specific |                       |
|                       | Redshift Cluster      |                       |
|                       | environment           |                       |
|                       | "production\*".       |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Version\":          |                       |
|                       | \"2012-10-17\",       |                       |
|                       |                       |                       |
|                       | \"Statement\": \[     |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Sid\":\"AllowCluste |                       |
|                       | rManagement\",        |                       |
|                       |                       |                       |
|                       | \"Action\": \[        |                       |
|                       |                       |                       |
|                       | \"redshift:CreateClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"redshift:DeleteClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"redshift:ModifyClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"redshift:RebootClus |                       |
|                       | ter\"                 |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Resource\": \[      |                       |
|                       |                       |                       |
|                       | \"\*\"                |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Effect\": \"Allow\" |                       |
|                       |                       |                       |
|                       | },                    |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Sid\":\"DenyDeleteM |                       |
|                       | odifyProtected\",     |                       |
|                       |                       |                       |
|                       | \"Action\": \[        |                       |
|                       |                       |                       |
|                       | \"redshift:DeleteClus |                       |
|                       | ter\",                |                       |
|                       |                       |                       |
|                       | \"redshift:ModifyClus |                       |
|                       | ter\"                 |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Resource\": \[      |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | cluster:production\*\ |                       |
|                       | "                     |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Effect\": \"Deny\"  |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | \]                    |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | -   Ensure you        |                       |
|                       |     separate roles    |                       |
|                       |     that perform      |                       |
|                       |     snapshot/restore  |                       |
|                       |     from regular      |                       |
|                       |     users. Remember   |                       |
|                       |     to test this as   |                       |
|                       |     IAM assumes deny  |                       |
|                       |     but if a broad    |                       |
|                       |     access role is    |                       |
|                       |     assigned and a    |                       |
|                       |     regular user can  |                       |
|                       |     assume that role  |                       |
|                       |     then they user    |                       |
|                       |     can perform the   |                       |
|                       |     API action. The   |                       |
|                       |     only way to       |                       |
|                       |     prevent this is   |                       |
|                       |     to have a         |                       |
|                       |     condition         |                       |
|                       |     statement         |                       |
|                       |     limiting actions  |                       |
|                       |     to a specific     |                       |
|                       |     role or an        |                       |
|                       |     explicit deny.    |                       |
|                       |                       |                       |
|                       | -   Become very       |                       |
|                       |     familiar with all |                       |
|                       |     API actions for   |                       |
|                       |     RedShift \[6\]    |                       |
|                       |                       |                       |
|                       | -   To properly       |                       |
|                       |     manage a          |                       |
|                       |     Service-Linked    |                       |
|                       |     role be sure to   |                       |
|                       |     separate the      |                       |
|                       |     \"create\" and    |                       |
|                       |     \"delete\"        |                       |
|                       |     actions to unique |                       |
|                       |     IAM entity so     |                       |
|                       |     access to manage  |                       |
|                       |     data and manage   |                       |
|                       |     the cluster are   |                       |
|                       |     separated.        |                       |
|                       |                       |                       |
|                       |     For example:      |                       |
|                       |     Allow IAM         |                       |
|                       |     identity to       |                       |
|                       |     delete            |                       |
|                       |     Service-Linked    |                       |
|                       |     Role              |                       |
|                       |                       |                       |
|                       |     {                 |                       |
|                       |                       |                       |
|                       |     \"Effect\":       |                       |
|                       |     \"Allow\",        |                       |
|                       |                       |                       |
|                       |     \"Action\": \[    |                       |
|                       |                       |                       |
|                       |     \"iam:DeleteServi |                       |
|                       | ceLinkedRole\",       |                       |
|                       |                       |                       |
|                       |     \"iam:GetServiceL |                       |
|                       | inkedRoleDeletionStat |                       |
|                       | us\"                  |                       |
|                       |                       |                       |
|                       |     \],               |                       |
|                       |                       |                       |
|                       |     \"Resource\":     |                       |
|                       |     \"arn:aws:iam::\< |                       |
|                       | AWS-account-ID\>:role |                       |
|                       | /aws-service-role/red |                       |
|                       | shift.amazonaws.com/A |                       |
|                       | WSServiceRoleForRedsh |                       |
|                       | ift\",                |                       |
|                       |                       |                       |
|                       |     \"Condition\":    |                       |
|                       |     {\"StringLike\":  |                       |
|                       |     {\"iam:AWSService |                       |
|                       | Name\":               |                       |
|                       |     \"redshift.amazon |                       |
|                       | aws.com\"}}           |                       |
|                       |                       |                       |
|                       |     }                 |                       |
+-----------------------+-----------------------+-----------------------+
| Authorization between | -   Redshift will     | 1.  <https://docs.aws |
| AWS services          |     need to access    | .amazon.com/redshift/ |
|                       |     other AWS         | latest/mgmt/authorizi |
|                       |     services (S3,     | ng-redshift-service.h |
|                       |     Glue, Athena,     | tml>                  |
|                       |     KMS, etc\..., to  |                       |
|                       |     perform functions | 2.  <https://docs.aws |
|                       |     in an automated   | .amazon.com/redshift/ |
|                       |     way. To setup a   | latest/mgmt/using-ser |
|                       |     role for a        | vice-linked-roles.htm |
|                       |     specific service  | l>                    |
|                       |     function use      |                       |
|                       |     reference \[1\].  |                       |
|                       |                       |                       |
|                       | -   It is recommended |                       |
|                       |     to run a Service  |                       |
|                       |     Linked Role for   |                       |
|                       |     RedShift to limit |                       |
|                       |     service specific  |                       |
|                       |     actions to only   |                       |
|                       |     the RDS service   |                       |
|                       |     endpoint \[2\]    |                       |
|                       |                       |                       |
|                       | -   You don\'t need   |                       |
|                       |     to manually       |                       |
|                       |     create an         |                       |
|                       |     AWSServiceRoleFor |                       |
|                       | Redshift              |                       |
|                       |     service-linked    |                       |
|                       |     role. Amazon      |                       |
|                       |     Redshift creates  |                       |
|                       |     the               |                       |
|                       |     service-linked    |                       |
|                       |     role for you. If  |                       |
|                       |     the               |                       |
|                       |     AWSServiceRoleFor |                       |
|                       | Redshift              |                       |
|                       |     service-linked    |                       |
|                       |     role has been     |                       |
|                       |     deleted from your |                       |
|                       |     account, Amazon   |                       |
|                       |     Redshift creates  |                       |
|                       |     the role when you |                       |
|                       |     launch a new      |                       |
|                       |     Amazon Redshift   |                       |
|                       |     cluster.          |                       |
|                       |                       |                       |
|                       | > .                   |                       |
+-----------------------+-----------------------+-----------------------+
| Authentication to AWS | RedShift supports     | 1.  <https://docs.aws |
| platform              | local and IAM based   | .amazon.com/redshift/ |
|                       | authentication. To    | latest/mgmt/options-f |
|                       | align to best         | or-providing-iam-cred |
|                       | practice Redshift     | entials.html>         |
|                       | local users should    |                       |
|                       | have passwords        | 2.  <https://docs.aws |
|                       | disabled which forces | .amazon.com/redshift/ |
|                       | authentication based  | latest/mgmt/generatin |
|                       | on IAM.               | g-user-credentials.ht |
|                       |                       | ml>                   |
|                       | > To generate         |                       |
|                       | > credentials for an  |                       |
|                       | > IAM user Granting   |                       |
|                       | > permission to       |                       |
|                       | > GetClusterCredentia |                       |
|                       | ls API                |                       |
|                       | > action should be    |                       |
|                       | > limited to          |                       |
|                       | > authorized IAM      |                       |
|                       | > entities and        |                       |
|                       | > limited to only     |                       |
|                       | > specific cluster,   |                       |
|                       | > database,           |                       |
|                       | > usernames, and      |                       |
|                       | > group names.        |                       |
|                       | >                     |                       |
|                       | > aws redshift        |                       |
|                       | > get-cluster-credent |                       |
|                       | ials                  |                       |
|                       | > \--cluster-identifi |                       |
|                       | er                    |                       |
|                       | > examplecluster      |                       |
|                       | > \--db-user          |                       |
|                       | > temp\_creds\_user   |                       |
|                       | > \--db-name          |                       |
|                       | > exampledb           |                       |
|                       | > \--duration-seconds |                       |
|                       | > 3600                |                       |
|                       |                       |                       |
|                       | -   Db credentials    |                       |
|                       |     should make use   |                       |
|                       |     of IAM            |                       |
|                       |                       |                       |
|                       | -   This means db     |                       |
|                       |     users password    |                       |
|                       |     must be disabled  |                       |
|                       |     which forces      |                       |
|                       |     login credentials |                       |
|                       |     based on temp IAM |                       |
|                       |     credentials       |                       |
+-----------------------+-----------------------+-----------------------+
| Authorization (AWS    | Amazon Redshift       | 1.  <https://docs.aws |
| IAM) of corporate     | requires IAM          | .amazon.com/redshift/ |
| users via Active      | credentials that AWS  | latest/mgmt/options-f |
| Directory for access  | can use to            | or-providing-iam-cred |
| to RedShift           | authenticate your     | entials.html>         |
| resources.            | requests. Those       |                       |
|                       | credentials must have | 2.  [Overview of      |
|                       | permissions to access |     Managing Access   |
|                       | AWS resources, such   |     Permissions to    |
|                       | as an Amazon Redshift |     Your Amazon       |
|                       | cluster.              |     Redshift          |
|                       |                       |     Resources](https: |
|                       | AWS Identity and      | //docs.aws.amazon.com |
|                       | Access Management     | /redshift/latest/mgmt |
|                       | (IAM) enable          | /redshift-iam-access- |
|                       | organizations with    | control-overview.html |
|                       | multiple employees to | )                     |
|                       | create and manage     |                       |
|                       | multiple users,       | 3.  [Using            |
|                       | groups and roles      |     Identity-Based    |
|                       | under a single AWS    |     Policies (IAM     |
|                       | account. With IAM     |     Policies) for     |
|                       | policies, companies   |     Amazon            |
|                       | can grant IAM         |     Redshift](https:/ |
|                       | users/groups/roles    | /docs.aws.amazon.com/ |
|                       | fine-grained control  | redshift/latest/mgmt/ |
|                       | to their Amazon       | redshift-iam-access-c |
|                       | RedShift data while   | ontrol-identity-based |
|                       | also retaining full   | .html)                |
|                       | control over          |                       |
|                       | everything the users  | 4.  <https://docs.aws |
|                       | do.                   | .amazon.com/redshift/ |
|                       |                       | latest/mgmt/generatin |
|                       | Most customers have   | g-user-credentials.ht |
|                       | setup federation with | ml>                   |
|                       | AWS accounts.         |                       |
|                       | Therefore, the only   | 5.  <https://docs.aws |
|                       | decision to make is   | .amazon.com/redshift/ |
|                       | what API actions are  | latest/mgmt/redshift- |
|                       | needed for roles,     | iam-access-control-id |
|                       | groups, or users      | entity-based.html>    |
|                       | within IAM \[2\].     |                       |
|                       |                       |                       |
|                       | To combine these      |                       |
|                       | concepts to control   |                       |
|                       | access to RedShift    |                       |
|                       | resources, a user     |                       |
|                       | would:                |                       |
|                       |                       |                       |
|                       | -   Determine how     |                       |
|                       |     authentication    |                       |
|                       |     will occur with   |                       |
|                       |     Redshift\[2\]     |                       |
|                       |                       |                       |
|                       | -   Create the        |                       |
|                       |     appropriate IAM   |                       |
|                       |     roles \[5\]       |                       |
|                       |                       |                       |
|                       |     For example, a    |                       |
|                       |     policy to allow   |                       |
|                       |     IAM users to      |                       |
|                       |     request           |                       |
|                       |     credentials for   |                       |
|                       |     temporary access. |                       |
|                       |     The following     |                       |
|                       |     policy enables    |                       |
|                       |     the GetCredential |                       |
|                       | s, CreateCluserUser,  |                       |
|                       |     and JoinGroup act |                       |
|                       | ions.                 |                       |
|                       |     The policy uses   |                       |
|                       |     condition keys to |                       |
|                       |     allow             |                       |
|                       |     the GetClusterCre |                       |
|                       | dentials and CreateCl |                       |
|                       | usterUser actions     |                       |
|                       |     only when the AWS |                       |
|                       |     user ID           |                       |
|                       |     matches \"AIDIODR |                       |
|                       | 4TAW7CSEXAMPLE:\${red |                       |
|                       | shift:DbUser}\@yourdo |                       |
|                       | main.com\".           |                       |
|                       |     IAM access is     |                       |
|                       |     requested for     |                       |
|                       |     the \"testdb\" da |                       |
|                       | tabase                |                       |
|                       |     only. The policy  |                       |
|                       |     also allows users |                       |
|                       |     to join a group   |                       |
|                       |     named \"common\_g |                       |
|                       | roup\".               |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Version\":          |                       |
|                       | \"2012-10-17\",       |                       |
|                       |                       |                       |
|                       | \"Statement\": \[     |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Sid\":              |                       |
|                       | \"GetClusterCredsStat |                       |
|                       | ement\",              |                       |
|                       |                       |                       |
|                       | \"Effect\":           |                       |
|                       | \"Allow\",            |                       |
|                       |                       |                       |
|                       | \"Action\": \[        |                       |
|                       |                       |                       |
|                       | \"redshift:GetCluster |                       |
|                       | Credentials\"         |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Resource\": \[      |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | dbuser:examplecluster |                       |
|                       | /\${redshift:DbUser}\ |                       |
|                       | ",                    |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | dbname:examplecluster |                       |
|                       | /testdb\",            |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | dbgroup:examplecluste |                       |
|                       | r/common\_group\"     |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Condition\": {      |                       |
|                       |                       |                       |
|                       | \"StringEquals\": {   |                       |
|                       |                       |                       |
|                       | \"aws:userid\":\"AIDI |                       |
|                       | ODR4TAW7CSEXAMPLE:\${ |                       |
|                       | redshift:DbUser}\@you |                       |
|                       | rdomain.com\"         |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | },                    |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Sid\":              |                       |
|                       | \"CreateClusterUserSt |                       |
|                       | atement\",            |                       |
|                       |                       |                       |
|                       | \"Effect\":           |                       |
|                       | \"Allow\",            |                       |
|                       |                       |                       |
|                       | \"Action\": \[        |                       |
|                       |                       |                       |
|                       | \"redshift:CreateClus |                       |
|                       | terUser\"             |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Resource\": \[      |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | dbuser:examplecluster |                       |
|                       | /\${redshift:DbUser}\ |                       |
|                       | "                     |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Condition\": {      |                       |
|                       |                       |                       |
|                       | \"StringEquals\": {   |                       |
|                       |                       |                       |
|                       | \"aws:userid\":\"AIDI |                       |
|                       | ODR4TAW7CSEXAMPLE:\${ |                       |
|                       | redshift:DbUser}\@you |                       |
|                       | rdomain.com\"         |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | },                    |                       |
|                       |                       |                       |
|                       | {                     |                       |
|                       |                       |                       |
|                       | \"Sid\":              |                       |
|                       | \"RedshiftJoinGroupSt |                       |
|                       | atement\",            |                       |
|                       |                       |                       |
|                       | \"Effect\":           |                       |
|                       | \"Allow\",            |                       |
|                       |                       |                       |
|                       | \"Action\": \[        |                       |
|                       |                       |                       |
|                       | \"redshift:JoinGroup\ |                       |
|                       | "                     |                       |
|                       |                       |                       |
|                       | \],                   |                       |
|                       |                       |                       |
|                       | \"Resource\": \[      |                       |
|                       |                       |                       |
|                       | \"arn:aws:redshift:us |                       |
|                       | -west-2:123456789012: |                       |
|                       | dbgroup:examplecluste |                       |
|                       | r/common\_group\"     |                       |
|                       |                       |                       |
|                       | \]                    |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | \]                    |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | }                     |                       |
|                       |                       |                       |
|                       | -   Map the           |                       |
|                       |     appropriate AD    |                       |
|                       |     groups to those   |                       |
|                       |     roles             |                       |
|                       |                       |                       |
|                       | -   Determine if IAM  |                       |
|                       |     users are allowed |                       |
|                       |     to create user    |                       |
|                       |     credentials       |                       |
|                       |     within RedShift   |                       |
|                       |     \[4\]             |                       |
+-----------------------+-----------------------+-----------------------+
| Tagging               | Amazon Redshift       | 1.  <https://docs.aws |
|                       | supports tagging to   | .amazon.com/redshift/ |
|                       | provide metadata      | latest/mgmt/amazon-re |
|                       | about resources at a  | dshift-tagging.html>  |
|                       | glance, IAM           |                       |
|                       | enforcement, and to   |                       |
|                       | categorize your       |                       |
|                       | billing reports based |                       |
|                       | on cost allocation.   |                       |
|                       |                       |                       |
|                       | Best Practice:        |                       |
|                       |                       |                       |
|                       | Since tags can be     |                       |
|                       | used to enforce IAM   |                       |
|                       | entity abilities it   |                       |
|                       | is important to set   |                       |
|                       | appropriate controls  |                       |
|                       | on changes to tags    |                       |
|                       | with API actions like |                       |
|                       | (DeleteTags,          |                       |
|                       | CreateTags)           |                       |
|                       |                       |                       |
|                       | **Note: tags are not  |                       |
|                       | retained if you copy  |                       |
|                       | a snapshot to another |                       |
|                       | region, so you must   |                       |
|                       | recreate the tags in  |                       |
|                       | the new region**      |                       |
+-----------------------+-----------------------+-----------------------+
| **Logging &           |                       |                       |
| Monitoring**          |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Logging activity      | Audit logging is not  | 1.  <https://docs.aws |
| within RedShift       | enabled by default in | .amazon.com/redshift/ |
|                       | Amazon Redshift. When | latest/mgmt/db-auditi |
|                       | you enable logging on | ng.html>              |
|                       | your cluster, Amazon  |                       |
|                       | Redshift creates and  | 2.  <https://docs.aws |
|                       | uploads logs to       | .amazon.com/redshift/ |
|                       | Amazon S3 that        | latest/mgmt/working-w |
|                       | capture data from the | ith-parameter-groups. |
|                       | creation of the       | html>                 |
|                       | cluster to the        |                       |
|                       | present time.         |                       |
|                       |                       |                       |
|                       | To meet logging       |                       |
|                       | requirements make     |                       |
|                       | sure to enable audit  |                       |
|                       | logging:              |                       |
|                       |                       |                       |
|                       | -   Connection log    |                       |
|                       |                       |                       |
|                       | -   User log          |                       |
|                       |                       |                       |
|                       | -   User activity log |                       |
|                       |     (requires         |                       |
|                       |     additional step   |                       |
|                       |     after enable of   |                       |
|                       |     audit logging)    |                       |
|                       |                       |                       |
|                       | <!-- -->              |                       |
|                       |                       |                       |
|                       | -   To enable this    |                       |
|                       |     you must enable   |                       |
|                       |     the **"enable\_us |                       |
|                       | er\_activity\_logging |                       |
|                       | "** database          |                       |
|                       |     parameter\[2\]    |                       |
|                       |                       |                       |
|                       |     For example:      |                       |
|                       |                       |                       |
|                       |     aws redshift      |                       |
|                       |     modify-cluster-pa |                       |
|                       | rameter-group         |                       |
|                       |                       |                       |
|                       |     \--parameter-grou |                       |
|                       | p-name                |                       |
|                       |     myclusterparamete |                       |
|                       | rgroup                |                       |
|                       |                       |                       |
|                       |     \--parameters     |                       |
|                       |     ParameterName=sta |                       |
|                       | tement\_timeout,Param |                       |
|                       | eterValue=20000       |                       |
|                       |     ParameterName=ena |                       |
|                       | ble\_user\_activity\_ |                       |
|                       | logging,ParameterValu |                       |
|                       | e=true                |                       |
|                       |                       |                       |
|                       | As a best practice,   |                       |
|                       | after a configuration |                       |
|                       | of RedShift is found  |                       |
|                       | to be functional and  |                       |
|                       | meet requirements     |                       |
|                       | make sure to commit   |                       |
|                       | all settings into a   |                       |
|                       | [parameter            |                       |
|                       | group](https://docs.a |                       |
|                       | ws.amazon.com/redshif |                       |
|                       | t/latest/mgmt/working |                       |
|                       | -with-parameter-group |                       |
|                       | s.html)               |                       |
|                       | so all databases      |                       |
|                       | within a cluster are  |                       |
|                       | configured the same   |                       |
|                       | and each new cluster  |                       |
|                       | can be configured the |                       |
|                       | same. (a final        |                       |
|                       | deployed cluster      |                       |
|                       | should not have       |                       |
|                       | parameter group =     |                       |
|                       | default.redshift-1.0  |                       |
|                       | because this will not |                       |
|                       | enable logging or     |                       |
|                       | other settings        |                       |
|                       | specific to customer  |                       |
|                       | requirements.)        |                       |
+-----------------------+-----------------------+-----------------------+
| Logging API actions   | Cloudtrail will be    | 1.  <https://docs.aws |
|                       | enabled in every      | .amazon.com/redshift/ |
|                       | account as part of a  | latest/mgmt/db-auditi |
|                       | default account       | ng.html#rs-db-auditin |
|                       | build.                | g-cloud-trail>        |
+-----------------------+-----------------------+-----------------------+
| Alerting and Incident | You can use the       | 1.  <https://github.c |
| Management            | following automated   | om/awslabs/amazon-red |
|                       | monitoring tools to   | shift-monitoring>     |
|                       | watch RedShift and    |                       |
|                       | report when something | 2.  <https://docs.aws |
|                       | is wrong:             | .amazon.com/redshift/ |
|                       |                       | latest/mgmt/metrics-l |
|                       | -   **Amazon          | isting.html>          |
|                       |     CloudWatch        |                       |
|                       |     Alarms** -- Watch | 3.  <https://docs.aws |
|                       |     a single metric   | .amazon.com/redshift/ |
|                       |     over a time       | latest/mgmt/working-w |
|                       |     period that you   | ith-event-notificatio |
|                       |     specify, and      | ns.html>              |
|                       |     perform one or    |                       |
|                       |     more actions      |                       |
|                       |     based on the      |                       |
|                       |     value of the      |                       |
|                       |     metric relative   |                       |
|                       |     to a given        |                       |
|                       |     threshold over a  |                       |
|                       |     number of time    |                       |
|                       |     periods. The      |                       |
|                       |     action is a       |                       |
|                       |     notification sent |                       |
|                       |     to an Amazon      |                       |
|                       |     Simple            |                       |
|                       |     Notification      |                       |
|                       |     Service (Amazon   |                       |
|                       |     SNS) topic or     |                       |
|                       |     Auto Scaling      |                       |
|                       |     policy.           |                       |
|                       |     CloudWatch alarms |                       |
|                       |     do not invoke     |                       |
|                       |     actions simply    |                       |
|                       |     because they are  |                       |
|                       |     in a particular   |                       |
|                       |     state; the state  |                       |
|                       |     must have changed |                       |
|                       |     and been          |                       |
|                       |     maintained for a  |                       |
|                       |     specified number  |                       |
|                       |     of periods. For   |                       |
|                       |     more information, |                       |
|                       |     see <https://docs |                       |
|                       | .aws.amazon.com/Amazo |                       |
|                       | nCloudWatch/latest/mo |                       |
|                       | nitoring/rs-metricsco |                       |
|                       | llected.html>.        |                       |
|                       |                       |                       |
|                       | -   **AWS CloudTrail  |                       |
|                       |     Log               |                       |
|                       |     Monitoring** --   |                       |
|                       |     Share log files   |                       |
|                       |     between accounts, |                       |
|                       |     monitor           |                       |
|                       |     CloudTrail log    |                       |
|                       |     files in real     |                       |
|                       |     time by sending   |                       |
|                       |     them to           |                       |
|                       |     CloudWatch Logs,  |                       |
|                       |     write log         |                       |
|                       |     processing        |                       |
|                       |     applications in   |                       |
|                       |     Java, and         |                       |
|                       |     validate that     |                       |
|                       |     your log files    |                       |
|                       |     have not changed  |                       |
|                       |     after delivery by |                       |
|                       |     CloudTrail. For   |                       |
|                       |     more information, |                       |
|                       |     see [Logging      |                       |
|                       |     Amazon RedShift   |                       |
|                       |     API Calls By      |                       |
|                       |     Using AWS         |                       |
|                       |     CloudTrail](https |                       |
|                       | ://docs.aws.amazon.co |                       |
|                       | m/redshift/latest/mgm |                       |
|                       | t/db-auditing.html#rs |                       |
|                       | -db-auditing-cloud-tr |                       |
|                       | ail).                 |                       |
|                       |                       |                       |
|                       | **Implementation      |                       |
|                       | Note:** To access     |                       |
|                       | logs and data for     |                       |
|                       | monitoring the data   |                       |
|                       | must be decrypted. To |                       |
|                       | decrypt logs/data a   |                       |
|                       | customer managed CMK  |                       |
|                       | must be defined. Use  |                       |
|                       | the same CMK created  |                       |
|                       | to encrypt the        |                       |
|                       | cluster and create a  |                       |
|                       | new policy to grant   |                       |
|                       | access only to API    |                       |
|                       | actions necessary for |                       |
|                       | tables and actions    |                       |
|                       | that are authorized.  |                       |
|                       | For example: Use      |                       |
|                       | guides in reference   |                       |
|                       | \[1\] where           |                       |
|                       | cloudformation        |                       |
|                       | templates already     |                       |
|                       | exist and can be used |                       |
|                       | to provide a          |                       |
|                       | prescriptive approach |                       |
|                       | to collecting and     |                       |
|                       | monitoring logs.      |                       |
|                       |                       |                       |
|                       | Make note of the      |                       |
|                       | minimum requirements  |                       |
|                       | for access to the     |                       |
|                       | Redshift user that is |                       |
|                       | required. Be cautious |                       |
|                       | not to enable more    |                       |
|                       | than the necessary    |                       |
|                       | "**grant select on    |                       |
|                       | all tables in schema  |                       |
|                       | pg\_catalog to        |                       |
|                       | tamreporting**"       |                       |
|                       | entitlement.          |                       |
|                       |                       |                       |
|                       | **Note**              |                       |
|                       |                       |                       |
|                       | Audit logging to      |                       |
|                       | Amazon S3 is an       |                       |
|                       | optional, manual      |                       |
|                       | process. When you     |                       |
|                       | enable logging on     |                       |
|                       | your cluster, you are |                       |
|                       | enabling logging to   |                       |
|                       | Amazon S3 only.       |                       |
|                       | Logging to system     |                       |
|                       | tables is not         |                       |
|                       | optional and happens  |                       |
|                       | automatically for the |                       |
|                       | cluster. For more     |                       |
|                       | information about     |                       |
|                       | logging to system     |                       |
|                       | tables, see [System   |                       |
|                       | Tables                |                       |
|                       | Reference](http://doc |                       |
|                       | s.aws.amazon.com/reds |                       |
|                       | hift/latest/dg/cm_cha |                       |
|                       | p_system-tables.html) |                       |
|                       |  in                   |                       |
|                       | the Amazon Redshift   |                       |
|                       | Database Developer    |                       |
|                       | Guide.                |                       |
+-----------------------+-----------------------+-----------------------+
| **Patch/Updates**     |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Update/Patch for      | When Redshift is      | 1.  <https://aws.amaz |
| RedShift              | configured a          | on.com/premiumsupport |
|                       | maintenance window    | /knowledge-center/not |
|                       | must be defined to    | ification-maintenance |
|                       | allow for updates to  | -rds-redshift/>       |
|                       | be installed. This    |                       |
|                       | window should align   |                       |
|                       | with expected change  |                       |
|                       | management times and  |                       |
|                       | a complete            |                       |
|                       | understanding of      |                       |
|                       | outages, if any, will |                       |
|                       | occur during this     |                       |
|                       | window.               |                       |
|                       |                       |                       |
|                       | Additional            |                       |
|                       | considerations        |                       |
|                       | include:              |                       |
|                       |                       |                       |
|                       | -   Refer to          |                       |
|                       |     <https://docs.aws |                       |
|                       | .amazon.com/cli/lates |                       |
|                       | t/index.html>         |                       |
|                       |     to understand     |                       |
|                       |     what API action   |                       |
|                       |     will require a    |                       |
|                       |     reboot of the     |                       |
|                       |     cluster or will   |                       |
|                       |     not.              |                       |
|                       |                       |                       |
|                       | -   Understand        |                       |
|                       |     automatic version |                       |
|                       |     updates may       |                       |
|                       |     change expected   |                       |
|                       |     behavior and a    |                       |
|                       |     setting is        |                       |
|                       |     offered to        |                       |
|                       |     prevent cluster   |                       |
|                       |     version changes   |                       |
|                       |     without approval. |                       |
+-----------------------+-----------------------+-----------------------+
| **Availability**      |                       |                       |
+-----------------------+-----------------------+-----------------------+
| Backup and Restore    | -   Amazon RedShift   | 1.  <https://docs.aws |
|                       |     makes use of      | .amazon.com/redshift/ |
|                       |     Snapshots to      | latest/mgmt/working-w |
|                       |     provide customers | ith-snapshots.html>   |
|                       |     with a way of     |                       |
|                       |     recovering to an  | 2.  <https://docs.aws |
|                       |     RPO. Snapshots    | .amazon.com/redshift/ |
|                       |     are stored in     | latest/mgmt/working-w |
|                       |     Amazon S3,        | ith-db-encryption.htm |
|                       |     managed by AWS;   | l>                    |
|                       |     Snapshots are     |                       |
|                       |     transferred to S3 |                       |
|                       |     over SSL, and     |                       |
|                       |     where the data in |                       |
|                       |     the database is   |                       |
|                       |     already encrypted |                       |
|                       |     in the cluster,   |                       |
|                       |     it remains        |                       |
|                       |     encrypted in the  |                       |
|                       |     snapshot too.     |                       |
|                       |     \[1\]             |                       |
|                       |                       |                       |
|                       | -   If you enable     |                       |
|                       |     copying of        |                       |
|                       |     snapshots from an |                       |
|                       |     encrypted cluster |                       |
|                       |     and use AWS KMS   |                       |
|                       |     for your master   |                       |
|                       |     key, you cannot   |                       |
|                       |     rename your       |                       |
|                       |     cluster because   |                       |
|                       |     the cluster name  |                       |
|                       |     is part of the    |                       |
|                       |     encryption        |                       |
|                       |     context. If you   |                       |
|                       |     must rename your  |                       |
|                       |     cluster, you can  |                       |
|                       |     disable copying   |                       |
|                       |     of snapshots in   |                       |
|                       |     the source        |                       |
|                       |     region, rename    |                       |
|                       |     the cluster, and  |                       |
|                       |     then configure    |                       |
|                       |     and enable        |                       |
|                       |     copying of        |                       |
|                       |     snapshots again.  |                       |
|                       |     \[2\]             |                       |
|                       |                       |                       |
|                       | -   Important Note:   |                       |
|                       |     If you rotate a   |                       |
|                       |     DEK/CEK that is   |                       |
|                       |     used to encrypt a |                       |
|                       |     cluster all data  |                       |
|                       |     will be encrypted |                       |
|                       |     with the new key  |                       |
|                       |     except for        |                       |
|                       |     snapshots stored  |                       |
|                       |     in S3. A process  |                       |
|                       |     should be         |                       |
|                       |     developed to      |                       |
|                       |     ensure snapshots  |                       |
|                       |     are encrypted     |                       |
|                       |     with the new key  |                       |
|                       |     to ensure         |                       |
|                       |     recovery point    |                       |
|                       |     objectives (RPO)  |                       |
|                       |     are met.          |                       |
+-----------------------+-----------------------+-----------------------+
| Limits                | Understanding the     | 1.  <https://docs.aws |
|                       | limitations of the    | .amazon.com/redshift/ |
|                       | service can help      | latest/mgmt/amazon-re |
|                       | prevent unintentional | dshift-limits.html>   |
|                       | outages or ability to |                       |
|                       | meet requirements.    |                       |
|                       |                       |                       |
|                       | Items like:           |                       |
|                       |                       |                       |
|                       | -   \# of tables by   |                       |
|                       |     ec2 instance type |                       |
|                       |                       |                       |
|                       | -   Spectrum limits   |                       |
|                       |                       |                       |
|                       | -   Quotas            |                       |
|                       |                       |                       |
|                       | -   IAM roles allowed |                       |
|                       |                       |                       |
|                       | -   Naming            |                       |
|                       |     constraints       |                       |
+-----------------------+-----------------------+-----------------------+
