<table>
<thead>
<tr class="header">
<th><h2 id="section-10"></h2></th>
<th></th>
<th><p><strong>Service Status</strong></p>
<table>
<tbody>
<tr class="odd">
<td>Pass</td>
<td>Review</td>
<td>Remediate</td>
</tr>
</tbody>
</table></th>
<th><strong>Vendor Status</strong></th>
<th><strong>Vendor Notes</strong></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><h2 id="section">1</h2></td>
<td>External Certification</td>
<td>DynamoDB is covered as part of the SOC 1, 2 &amp; 3 reports in the following regions: US East,<br />
US West, Canada, EU (Ireland, Frankfurt, London) Asia Pacific (Singapore, Sydney, Tokyo, Seoul, Mumbai), and South America (São Paulo)</td>
<td>All reports can be downloaded from AWS Artifact.</td>
<td><a href="https://aws.amazon.com/compliance/services-in-scope/">https://aws.amazon.com/compliance/services-in-scope/</a></td>
</tr>
<tr class="even">
<td><h2 id="section-1">2</h2></td>
<td>Authentication</td>
<td>IAM is leveraged to handle authentication.</td>
<td>Bank is likely to use IAM roles which will be mapped to on premise Active Directory Groups</td>
<td><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/authentication-and-access-control.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/authentication-and-access-control.html</a></td>
</tr>
<tr class="odd">
<td><h2 id="section-2">3</h2></td>
<td>Authorization</td>
<td><p>IAM is leveraged to handle authorization.</p>
<p>DynamoDB offers very granular policy enforcement, providing that ability to restrict access not only at the <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_dynamodb_specific-table.html">table</a>, but also the <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_dynamodb_columns.html">column</a> and <a href="https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_dynamodb_rows.html">row</a> level via IAM policies.</p>
<p>User permissions can be scoped down to tables and items.</p>
<p>Resource based policies are not supported for DynamoDB.</p></td>
<td></td>
<td><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/authentication-and-access-control.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/authentication-and-access-control.html</a></td>
</tr>
<tr class="even">
<td><h2 id="section-3">4</h2></td>
<td>Encryption At Rest</td>
<td><p>DynamoDB supports encryption-at-rest encryption for the following regions</p>
<ul>
<li><p>US East (Ohio)</p></li>
<li><p>US East (N. Virginia)</p></li>
<li><p>US West (Oregon)</p></li>
<li><p>US West (N. California)</p></li>
<li><p>Canada (Central)</p></li>
<li><p>South America (São Paulo)</p></li>
<li><p>EU (Ireland)</p></li>
<li><p>EU (London)</p></li>
<li><p>EU (Frankfurt)</p></li>
<li><p>Asia Pacific (Mumbai)</p></li>
<li><p>Asia Pacific (Tokyo)</p></li>
<li><p>Asia Pacific (Seoul)</p></li>
<li><p>Asia Pacific (Singapore)</p></li>
<li><p>Asia Pacific (Sydney)</p></li>
</ul>
<p>CMKs are not supported by the SSE utilised by DynamoDB; you can only make use of service default keys. These keys are created by DynamoDB for every region where it is used if they do not already exist.</p>
<p>You can also use client side encryption, using KMS, to encrypt data before it is stored in DynamoDB.</p></td>
<td></td>
<td><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html</a></td>
</tr>
<tr class="odd">
<td><h2 id="section-4">5</h2></td>
<td>Encryption In Transit</td>
<td><p>DynamoDB provides a dedicated “Encryption Client” which developers and operators can make use of to encrypt data. This can transparently encrypt data client side.</p>
<p>DynamoDB supports the use of TLS backed endpoints.</p>
<p>DynamoDB does not have the notion of resource based policies; restrictions to the transport type (e.g. HTTPS only) can be attached to user policies to prevent them from using plain HTTP connections.</p></td>
<td><p>It is not possible to enforce or monitor clients using the encryption client.</p>
<p>It is possible to enforce and monitor clients using encrypted service endpoints.</p>
<p>The following policy statement, when applied to a policy and attached to a principle, will prevent the use of DynamoDB over HTTP</p>
<p>{</p>
<p>"Action": "dynamodb:*",</p>
<p>"Effect": "Deny",</p>
<p>"Resource": "arn:aws:dynamodb:*:*:*/*",</p>
<p>"Condition": {</p>
<p>"Bool": {</p>
<p>"aws:SecureTransport": "false"</p>
<p>}</p>
<p>}</p>
<p>}</p></td>
<td><a href="https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/client-server-side.html">https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/client-server-side.html</a></td>
</tr>
<tr class="even">
<td><h2 id="section-5">6</h2></td>
<td>Underlying OS</td>
<td>DynamoDB is a managed service. The underlying technologies are not configurable by Customer.</td>
<td>DynamoDB is in-scope for SOC and PCI.</td>
<td>https://aws.amazon.com/compliance/services-in-scope/</td>
</tr>
<tr class="odd">
<td><h2 id="section-6">7</h2></td>
<td>CSP access</td>
<td>Access to the DynamoDB data store can be controlled via IAM policy and encryption using KMS.</td>
<td>The AWS Production network is segregated from the Amazon Corporate network and requires a separate set of credentials for logical access. The Amazon Corporate network relies on user IDs, passwords, and Kerberos, while the AWS Production network requires SSH public-key authentication through a bastion host.<br />
AWS developers and administrators on the Amazon Corporate network who need to access AWS cloud components must explicitly request access through the AWS access management system. All requests are reviewed and approved by the appropriate owner or manager [1]<br />
<br />
Details of these controls and how they are validated by an independent 3rd party are available via AWS Artifact (SOC 2 - Security, Availability &amp; Confidentiality Report) [2]</td>
<td>AWS Security Whitepaper <a href="https://d1.awsstatic.com/whitepapers/Security/AWS_Security_Whitepaper.pdf">https://d1.awsstatic.com/whitepapers/Security/AWS_Security_Whitepaper.pdf</a><br />
<br />
AWS SOC FAQs: <a href="https://aws.amazon.com/compliance/soc-faqs/">https://aws.amazon.com/compliance/soc-faqs/</a></td>
</tr>
<tr class="even">
<td><h2 id="section-7">8</h2></td>
<td>Endpoint Localization</td>
<td>You can use VPC endpoints to keep traffic between your Amazon VPC and DynamoDB from leaving the Amazon network.</td>
<td></td>
<td>https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/vpc-endpoints-dynamodb.html</td>
</tr>
<tr class="odd">
<td><h2 id="section-8">9</h2></td>
<td>Storage/Process Isolation</td>
<td>The DynamoDB platform is serverless and makes use of a shared-tenancy host model - there are no physical hosts to set up or manage.</td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td><h2 id="section-9">10</h2></td>
<td>Supporting Services</td>
<td>KMS, VPC, CloudTrail, CloudWatch, IAM</td>
<td>DynamoDB is a managed service and other services can be used to augment or provide greater security/compliance control. They are not required for base functionality.</td>
<td></td>
</tr>
</tbody>
</table>

## Summary

This document provides a compiled list of security controls and best
practices for DynamoDB to assist security teams in assessing DynamoDB’s
fitness for use.

Amazon DynamoDB is a fully managed NoSQL database service that provides
fast and predictable performance with seamless scalability. DynamoDB
lets you offload the administrative burdens of operating and scaling a
distributed database, so that you don't have to worry about hardware
provisioning, setup and configuration, replication, software patching,
or cluster scaling. Also, DynamoDB offers encryption at rest, which
eliminates the operational burden and complexity involved in protecting
sensitive data. For more information, see [Amazon DynamoDB Encryption at
Rest](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html).

This table maps common security concerns to the corresponding controls
and architectural best practices as documented in AWS’ public
documentation, white papers, and blog posts. For additional AWS
security, compliance, penetration testing, Security Bulletins,
resources, etc, please visit [AWS Cloud
Security](https://aws.amazon.com/security/) website.

<table>
<thead>
<tr class="header">
<th>Security Need</th>
<th>AWS Control &amp; Architectural Options</th>
<th>References</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>Network Isolation</td>
<td><p>As a managed service, DynamoDB is protected by the AWS global network security procedures as described in the “Network Security” section of the <a href="https://d0.awsstatic.com/whitepapers/Security/AWS_Security_Whitepaper.pdf">AWS Security Whitepaper</a> and in more detail in <a href="https://aws.amazon.com/compliance/resources/">compliance reports</a> and 3<sup>rd</sup> party audit findings available to AWS customers.</p>
<p>While these APIs are callable from any network location, DynamoDB supports <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/api-permissions-reference.html">resource-level</a> and <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/specifying-conditions.html">item-level</a> access policies that can be applied to identities that have DynamoDB access.</p></td>
<td><p><a href="https://d0.awsstatic.com/whitepapers/Security/AWS_Security_Whitepaper.pdf">AWS Security Whitepaper</a></p>
<p><a href="https://aws.amazon.com/compliance/resources/">AWS compliance reports</a></p></td>
</tr>
<tr class="even">
<td>Inter-network traffic privacy (VPC to DynamoDB)</td>
<td><p>Amazon DynamoDB supports VPC endpoints. This means that DynamoDB can be provisioned as service within your AWS VPC. This means that resources within your VPC (e.g. EC2 instances) do not need to use publically available endpoints to reach DynamoDB.</p>
<p>By default when you create a VPC endpoint for DynamoDB, a default endpoint policy is attached to the endpoint. The default policy allows access by any user or service within the VPC, using credentials from any AWS account, to any DynamoDB resource. You can restrict this and lock down access to a subset of tables and API actions. An example of a locked down policy is shown below:</p>
<p>{</p>
<p>"Statement": [</p>
<p>{</p>
<p>"Sid": "AccessToSpecificTable",</p>
<p>"Principal": "*",</p>
<p>"Action": [</p>
<p>"dynamodb:Batch*",</p>
<p>"dynamodb:Delete*",</p>
<p>"dynamodb:DescribeTable",</p>
<p>"dynamodb:GetItem",</p>
<p>"dynamodb:PutItem",</p>
<p>"dynamodb:Update*"</p>
<p>],</p>
<p>"Effect": "Allow",</p>
<p>"Resource": "arn:aws:dynamodb:<em>us-east-1</em>:<em>123456789012</em>:table/BobbyDropTables"</p>
<p>}</p>
<p>]</p>
<p>}</p>
<p>There are additional restrictions that apply if you choose to use VPC endpoints for access to DynamoDB</p>
<ul>
<li><p>DynamoDB does not support resource-based policies (for example, on tables). Access to DynamoDB is controlled though the endpoint policy and IAM policies for individual IAM users and roles.</p></li>
<li><p>You cannot access Amazon DynamoDB Streams through a VPC endpoint.</p></li>
<li><p>Endpoints currently do not support cross-region requests—ensure that you create your endpoint in the same region as your DynamoDB tables.</p></li>
<li><p>If you use AWS CloudTrail to log DynamoDB operations, the log files contain the private IP address of the EC2 instance in the VPC and the endpoint ID for any actions performed through the endpoint.</p></li>
<li><p>The source IPv4 addresses from instances in your affected subnets change from public IPv4 addresses to the private IPv4 addresses from your VPC. An endpoint switches network routes, and disconnects open TCP connections. Your tasks are interrupted during the changeover, and any previous connections using public IPv4 addresses are not resumed. We recommend that you do not have any critical tasks running when you create or modify an endpoint; or that you test to ensure that your software can automatically reconnect to DynamoDB after the connection break.</p></li>
</ul></td>
<td></td>
</tr>
<tr class="odd">
<td>Encryption of data at-rest</td>
<td><p>You can transparently encrypt data stored in Amazon DynamoDB in 15 regions (at time of publication), allowing you to protect your data. This feature makes use of KMS to provide a consistent experience. CMKs are not supported by this feature, however. If upon first use the service default key does not exist, one will be created for you. The following are encrypted:</p>
<ul>
<li><p>DynamoDB base tables</p></li>
<li><p>Local secondary indexes</p></li>
<li><p>Global secondary indexes</p></li>
</ul>
<p>You can choose to use client side encryption to encrypt data before it is sent to DynamoDB.</p></td>
<td><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html">Amazon DynamoDB Encryption at Rest</a></td>
</tr>
</tbody>
</table>

<table>
<thead>
<tr class="header">
<th>Encryption of data in-transit</th>
<th><p>When service clients use HTTPS to send data, Amazon DynamoDB encrypts data in-flight between clients which protects against someone eavesdropping on data items being transferred.</p>
<p>You may also make use of the DynamoDB Data Encryption Client to encrypt you payload before it’s transmitted to the service endpoint.</p></th>
<th><a href="https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/client-server-side.html">https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/client-server-side.html</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>Encryption Key Management</td>
<td>Where enabled, Amazon DynamoDB supports Server-Side Encryption using AWS Key Management Service (AWS KMS) to encrypt DynamoDB data items using the service default key. The key material for this key is managed by AWS; users of the service do not have access to it.</td>
<td><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/encryption.howitworks.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/encryption.howitworks.html</a></td>
</tr>
<tr class="even">
<td>Isolation of physical hosts</td>
<td>Amazon DynamoDB platform is a serverless platform that runs on shared tenancy, like other serverless platforms.</td>
<td></td>
</tr>
<tr class="odd">
<td>Restricting administrative access to certain individuals within the company</td>
<td><p>Administrative Access</p>
<ul>
<li><p>Use AWS Identity &amp; Access Management (IAM) to create administrative users with appropriate permissions to Amazon DynamoDB</p></li>
<li><p>Alternatively, you can federate users from an on-premise (i.e. Microsoft Active Directory, or LDAP) to AWS IAM Roles</p></li>
</ul>
<p>DynamoDB does not support resource based policies, administrative access is determined using IAM permissions only.</p></td>
<td><p><a href="http://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html">IAM Best Practices</a></p>
<p><a href="https://aws.amazon.com/identity/federation/">Identify Federation in the AWS Cloud</a></p>
<p><a href="https://aws.amazon.com/blogs/security/enabling-federation-to-aws-using-windows-active-directory-adfs-and-saml-2-0/">Enabling Federation to AWS Using Windows Active Directory, ADFS, and SAML 2.0</a></p></td>
</tr>
<tr class="even">
<td>Authentication and authorization of corporate users via Active Directory for access to Amazon DynamoDB</td>
<td><p>You can federate users from an on-premise (i.e. Microsoft Active Directory, or LDAP) to AWS IAM Roles.</p>
<p>It is possible to restrict access at a very granular level to DynamoDB column and row level data, by using IAM policies: <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/specifying-conditions.html">Specifying Conditions</a>.</p></td>
<td><p><a href="https://aws.amazon.com/identity/federation/">Identify Federation in the AWS Cloud</a></p>
<p><a href="https://aws.amazon.com/blogs/security/enabling-federation-to-aws-using-windows-active-directory-adfs-and-saml-2-0/">Enabling Federation to AWS Using Windows Active Directory, ADFS, and SAML 2.0</a></p></td>
</tr>
<tr class="odd">
<td>Auditing of all Interactions with Amazon DynamoDB</td>
<td><p>Amazon DynamoDB is integrated with AWS CloudTrail, which captures API calls made by or on behalf of Streams and delivers the log files to the Amazon S3 bucket that you specify. The API calls can be made indirectly by using the console or directly by using the DynamoDB API. Using the information collected by CloudTrail, you can determine what request was made to DynamoDB, the source IP address from which the request was made, who made the request, when it was made, and so on.</p>
<p>CloudTrail currently logs the following API commands:</p>
<ul>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateBackup.html">CreateBackup</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateGlobalTable.html">CreateGlobalTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html">CreateTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteBackup.html">DeleteBackup</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteTable.html">DeleteTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeBackup.html">DescribeBackup</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeContinuousBackups.html">DescribeContinuousBackups</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeGlobalTable.html">DescribeGlobalTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeLimits.html">DescribeLimits</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTable.html">DescribeTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTimeToLive.html">DescribeTimeToLive</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListBackups.html">ListBackups</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListTables.html">ListTables</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListTagsOfResource.html">ListTagsOfResource</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ListGlobalTables.html">ListGlobalTables</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_RestoreTableFromBackup.html">RestoreTableFromBackup</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TagResource.html">TagResource</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UntagResource.html">UntagResource</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateGlobalTable.html">UpdateGlobalTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html">UpdateTable</a></p></li>
<li><p><a href="http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTimeToLive.html">UpdateTimeToLive</a></p></li>
<li><p>DescribeReservedCapacity</p></li>
<li><p>DescribeReservedCapacityOfferings</p></li>
<li><p>PurchaseReservedCapacityOfferings</p></li>
</ul></td>
<td><p><a href="http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-working-with-log-files.html">Working with CloudTrail Log Files</a></p>
<p><a href="http://docs.aws.amazon.com/awscloudtrail/latest/userguide/">AWS CloudTrail User Guide</a></p>
<p><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/logging-using-cloudtrail.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/logging-using-cloudtrail.html</a></p></td>
</tr>
<tr class="even">
<td>Compliance Framework reference material for Amazon DynamoDB</td>
<td>For compliance documents related to AWS please visit <a href="https://aws.amazon.com/artifact/">AWS Artifact</a>.</td>
<td><a href="https://aws.amazon.com/compliance/services-in-scope/">https://aws.amazon.com/compliance/services-in-scope/</a></td>
</tr>
<tr class="odd">
<td>Alerting and Incident Management</td>
<td><p>Any low-level API calls made to DynamoDB actions are tracked in log files. DynamoDB records are written together with other AWS service records in a log file. CloudTrail determines when to create and write to a new file based on a time period and file size.</p>
<p>Full details on the actions that are logged to CloudTrail by DynamoDB can be found here: <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/logging-using-cloudtrail.html">Logging DynamoDB Operations using CloudTrail</a></p>
<p>If DDB tables are provisioned on behalf of users of tables, i.e. administrators of tables are different from the users, the following CloudTrail events should be monitored:</p>
<p>- DeleteTable</p>
<p>- DeleteBackup</p>
<p>- CreateTable</p>
<p>Users, Roles, and Groups with policies that have DynamoDB access restricted to a subset of tables &amp; attributes should be monitored for changes that make them more permissive than before.</p></td>
<td><p><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/logging-using-cloudtrail.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/logging-using-cloudtrail.html</a></p>
<p><a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/PointInTimeRecovery_Howitworks.html">https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/PointInTimeRecovery_Howitworks.html</a></p></td>
</tr>
<tr class="even">
<td>Cross-region replication/key management.</td>
<td>DynamoDB supports Global Tables (multi-region/multi-master) in 6 regions (as of May 2018), which encrypts data at rest using the service default key for the region. The following regions are supported for GlobalTables - US East (N.Virginia &amp; Ohio), US West (Oregon), Asia Pacific (Singapore), EU (Frankfurt), EU(Ireland)</td>
<td></td>
</tr>
<tr class="odd">
<td>Update/Patch/Hardening for Amazon DynamoDB</td>
<td>DynamoDB is a managed service and there are no components that can be patched or updated by the customer.</td>
<td></td>
</tr>
</tbody>
</table>