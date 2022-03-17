# Detailed Security Configuration

Overview

This section is meant to provide an opinionated approach towards the
implementation security controls by domain. Although the approaches may
not be a fit for all use cases or complete for production use, they are
meant to guide the reader towards current best practices and design
considerations to achieve security control objectives.

**Controls and Architectures**

This table maps Security Domain to the corresponding controls and
architectural best practices as documented in GCP’s public documentation, 
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
By default, GKE encrypts customer content stored at rest, including Secrets. GKE handles and manages this default encryption for the user without any additional actions. Persistent disks in GKE are already encrypted at the hardware layer by default, however, there is also the option of adding additional encryption where you, the user, can manage the encryption keys. <br><br>
To encrypt persistent disks in GKE, you must use the GCP Persistent Disk CSI plugin, which lets you protect disks in GKE with a key that you manage in Cloud KMS—by creating a StorageClass referencing a key. This encryption key is used to encrypt the disks created with that StorageClass. If your organization is required to manage its own key material, the CSI plugin provides the same functionality available in traditional CMEK for persistent disks in GKE.<br> For this, you will need to create a Cloud KMS key to use for encryption, then you can create a StorageClass on Kubernetes that specifies the Cloud KMS key KMS_KEY_ID to use to encrypt the disk. <br><br>
Eg of storage class manifest:<br> <br>
apiVersion: storage.k8s.io/v1beta1 <br>
kind: StorageClass<br>
metadata:<br>
&nbsp; name: csi-gce-pd <br>
provisioner: pd.csi.storage.gke.io <br>
parameters: <br>
&nbsp;  type: pd-standard <br>
&nbsp;  disk-encryption-kms-key: KMS_KEY_ID <br> <br>

Application-layer Secrets Encryption provides an additional layer of security for sensitive data, such as Secrets, stored in etcd. Using this functionality, you can use a key managed with Cloud KMS to encrypt data at the application layer. This protects against attackers who gain access to an offline copy of etcd.<br><br>

To use Application-layer Secrets Encryption, you must first create a Cloud KMS key and give the GKE service account access to the key. The key must be in the same location as the cluster to decrease latency and to prevent cases where resources depend on services spread across multiple failure domains. Then, you can enable the feature on a new or existing cluster by specifying the key you would like to use.<br><br>

Note: Default GKE cluster has encrypted persistent disks and encryption layers for secrets stored in etcd/cluster <br><br>

Migrating from a workload without Customer Managed Encryption Keys (CMEK) disks to an environment with CMEK disks will require the creation of new storage classes and Persistent Volume Claims. 
</p>
<td><ol type="1">
<li><p>Default encryption at rest on GCP: <a href="https://cloud.google.com/security/encryption-at-rest/default-encryption">https://cloud.google.com/security/encryption-at-rest/default-encryption</a></p></li>
<li><p>Using your own keys to protect data on GKE:<a href="https://cloud.google.com/blog/products/containers-kubernetes/exploring-container-security-use-your-own-keys-to-protect-your-data-on-gke">https://cloud.google.com/blog/products/containers-kubernetes/exploring-container-security-use-your-own-keys-to-protect-your-data-on-gke</a></p></li>
<li><p>Application Layer Secrets Encryption: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/encrypting-secrets">https://cloud.google.com/kubernetes-engine/docs/how-to/encrypting-secrets</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Encryption of data in-transit</td>
<td><ul>
<p>The GKE Ingress Load Balancer is the ideal SSL termination point to ensure that data in transit is encrypted.
<br>
There are three ways to provide SSL certificates to an HTTPS load balancer on GKE:<br><br>

**1) Google-managed certificates** <br>
Google-managed SSL certificates are provisioned, deployed, renewed, and managed for your domains. Managed certificates do not support wildcard domains.<br><br>

**2) Self-managed certificates shared with Google Cloud**<br>
You can provision your own SSL certificate and create a certificate resource in your Google Cloud project. You can then list the certificate resource in an annotation on an Ingress to create an HTTP(S) load balancer that uses the certificate<br><br>

**3) Self-managed certificates as Secret resources**<br>
You can provision your own SSL certificate and create a Secret to hold it. You can then refer to the Secret in an Ingress specification to create an HTTP(S) load balancer that uses the certificate<br><br>.

The 3rd option is generally the most preferred and the most secure option if you incorporate Application Layer Secrets Encryption as mentioned in the previous sub-category.<br><br>

Note: The above information is in the context of utilizing GKE LB which supports multiple TLS certificates. 3rd party/OpenSource Ingress tools such as Nginx and Istio are not covered here
</p>
<td><ol type="1">
<li><p>SSL Certs on GCP: <a href="https://cloud.google.com/load-balancing/docs/ssl-certificates">https://cloud.google.com/load-balancing/docs/ssl-certificates</a></p></li>
<li><p>Using Multiple SSL Certs with GKE Ingress <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/ingress-multi-ssl">https://cloud.google.com/kubernetes-engine/docs/how-to/ingress-multi-ssl</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Encryption Key Management</td>
<td><p>
If you want to control and manage the encryption on GCP yourself, you can use key encryption keys. Key encryption keys do not directly encrypt your data but are used to encrypt the data encryption keys that encrypt your data. <br><br>

You have 3 options for encryption in GCP:
<ol>
<li>Use Cloud Key Management Service to create and manage key encryption keys.</li> 
<li>Create and manage your own key encryption keys, also known as customer-supplied encryption keys (CSEK). </li>
<li>Use Cloud HSM: Cloud HSM is a cloud-hosted Hardware Security Module (HSM) service that allows you to host encryption keys and perform cryptographic operations in a cluster of FIPS 140-2 Level 3 certified HSMs. Google manages the HSM cluster for you, so you don't need to worry about clustering, scaling, or patching. Because Cloud HSM uses Cloud KMS as its front end, you can leverage all the conveniences and features that Cloud KMS provides. Note: Cloud HSM is only available in a few regions at the moment </li>
</ol>
<br><br>

Key Rotation should occur as often as organization security guidelines dictate.
<br><br>
Key rotations can be automated like so: <br><br>

gcloud kms keys update key-name \ <br>
 &nbsp; &nbsp; &nbsp;    --location location \ <br>
 &nbsp; &nbsp; &nbsp;    --keyring key-ring-name \ <br>
 &nbsp; &nbsp; &nbsp;    --rotation-period rotation-period \ <br>
 &nbsp; &nbsp; &nbsp;    --next-rotation-time next-rotation-time \ <br>
</p></td>
<td><ol type="1">
<li><p>Protecting Resources with KMS keys: <a href="https://cloud.google.com/compute/docs/disks/customer-managed-encryption">https://cloud.google.com/compute/docs/disks/customer-managed-encryption</a></p></li>
<li><p>Using KMS: <a href="https://cloud.google.com/kms/docs/quickstart">https://cloud.google.com/kms/docs/quickstart</a></p></li>
<li><p>Configuring automatic Key Rotation<a href="https://cloud.google.com/kms/docs/rotating-keys#kms-enable-key-version-cli">https://cloud.google.com/kms/docs/rotating-keys#kms-enable-key-version-cli</a></p></li>
<li><p>Using Cloud HSM<a href="https://cloud.google.com/kms/docs/hsm">https://cloud.google.com/kms/docs/hsm</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>Infrastructure</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Isolation of infrastructure and logical components</td>
<td>
 <p> <b>Layers of isolation in Kubernetes: </b> <br><br>

Kubernetes has several nested layers, each of which provides some level of isolation and security. 
<ul>
 <li><b>Container (not specific to Kubernetes):</b> A container provides basic management of resources, but does not isolate identity or the network, and can suffer from a noisy neighbor on the node for resources that are not isolated by cgroups. It provides some security isolation, but only provides a single layer, compared to our desired double layer. </li>

<li> <b>Pod:</b> A pod is a collection of containers. A pod isolates a few more resources than a container, including the network. It does so with micro-segmentation using Kubernetes Network Policy, which dictates which pods can speak to one another. At the moment, a pod does not have a unique identity, but the Kubernetes community has made proposals to provide this. A pod still suffers from noisy neighbors on the same host. </li>

<li> <b>Node:</b> This is a machine, either physical or virtual. A node includes a collection of pods, and has a superset of the privileges of those pods. A node leverages a hypervisor or hardware for isolation, including for its resources. Modern Kubernetes nodes run with distinct identities, and are authorized only to access the resources required by pods that are scheduled to the node. There can still be attacks at this level, such as convincing the scheduler to assign sensitive workloads to the node. You can use firewall rules to restrict network traffic to the node. </li>

<li> <b>Cluster</b> A cluster is a collection of nodes and a control plane. This is a management layer for your containers. Clusters offer stronger network isolation with per-cluster DNS. </li>

<li> <b>Project:</b> A GCP project is a collection of resources, including Kubernetes Engine clusters. A project provides all of the above, plus some additional controls that are GCP-specific, like project-level IAM for Kubernetes Engine and org policies. Resource names, and other resource metadata, are visible up to this layer.</li></ul>
<br>

Isolating applications on GKE can be achieved with the following scenarios: <br><br>
<ul>
 <li> <b>Seperating applications by namespace</b>: This is the simplest option for isolating different services on kubernetes, however, namespaces are only meant to incorporate logical isolation and doesn't provide actual physical isolation between services/pods. <br><br>Eg: login microservice resources will be hosted in the login namespace, and similarly shipping microservice resources will be contained in the shipping namespace</li>
 <li> <b>Seperating applications by node pool</b>: This method can be used to achieve physical isolation if used appropriately in tandem with Node Selectors and Affinities. This option also provides flexibility in using different node types for different type of workloads. <br><br> Eg: You can have ML workloads running on a Nodepool with GPUs and a datastore Nodepool which contains SSD Disks.
 <li> <b>Seperating applications by cluster</b>: This option achieves the maximum level of physical and logical isolation for kubernetes clusters, also providing all the benefits of the second option, but with a relatively heavier management and overhead and will often involve the inclusion of implementing network structures that allow for communication between clusters if needed. <br><br>Eg: You may want to run a seperate cluster on GKE for hosting your CI/CD workloads, however this will require a VPN connection if the CI/CD cluster is expected to deploy changes to a separate private cluster.
</ul>
</p>
</td>
<td><ol type="1">
<li><p>Security Isolation on GKE <a href="https://cloud.google.com/blog/products/gcp/exploring-container-security-isolation-at-different-layers-of-the-kubernetes-stack">https://cloud.google.com/blog/products/gcp/exploring-container-security-isolation-at-different-layers-of-the-kubernetes-stack</a></p></li>
<li><p>Network Policy for inter-pod communication: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy">https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Network Isolation</td>
<td><p>
There are three different flavors of network modes for clusters that can deliver network level protection: <br><br>
<ol>
<li><b>Public endpoint access disabled, master authorized networks enabled (recommended):</b> This is the most secure option as it prevents all internet access to both masters and nodes, except for certain whitelisted IPs and Google Service IP ranges.</li>
 <li><b> Public endpoint access enabled, master authorized networks enabled:</b> This option provides restricted access to the master from source IP addresses that you define. This is a good choice if you don't have existing VPN infrastructure or have remote users or branch offices that connect over the public internet instead of the corporate VPN and Cloud Interconnect or Cloud VPN.</li>
 <li><b>Public endpoint access enabled, master authorized networks disabled (least secure): </b>This is the default and allows anyone on the internet to make network connections to the control plane.</li><br>
 
The most secure option is utilizing private GKE cluster with no public endpoints (on either master or nodes), and whitelisting Kube API server access in the master to only select machines.<br>This best practive ensures the control plane is only reachable by: <br>
<ul>
<li>The whitelisted CIDRs in master authorized networks (bastion host or other admin machines). </li>
<li> Nodes within your cluster's VPC.</li>
<li> Google's internal production jobs that manage your master.</li>
</ul>
<br><br>
This corresponds to the following gcloud flags at cluster creation time:
<ul>
 <li> --enable-ip-alias</li>
 <li> --enable-private-nodes</li>
 <li> --enable-master-authorized-networks</li>
</ul>
</p></td>
<td><ol type="1">
<li><p>Restricting Network Access <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#restrict_network_access_to_the_control_plane_and_nodes">https://cloud.google.com/kubernetes-engine/docs/how-to/hardening-your-cluster#restrict_network_access_to_the_control_plane_and_nodes</a></p></li>
<li><p>Creating a Private GKE Cluster: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters">https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters</a></p></li>
 <li><p>Master Authorized Networks: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/authorized-networks">https://cloud.google.com/kubernetes-engine/docs/how-to/authorized-networks</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>GKE Master Network</td>
<td><p>
An important thing to note is that the entire GKE control plane is managed by Google. <br><br>
Essentially, the master node is  managed in a separate Google managed project, in a separate Google managed VPC that is automatically peered with the VPC in which you deploy your cluster upon cluster creation time.<br><br>
This means that if you need to communicate to the Kube master API in a private cluster from an external source, it would require establishing a VPN tunnel and exporting the custom routes.
</p></td>
<td><ol type="1">
<li><p>Cluster Master<a href="https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#master">https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture#master</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>IAM and RBAC</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Admin Accounts</td>
<td> <p> In IAM for GKE services there are 2 types of administrative roles Kubernetes Engine Admin, and Kubernetes Engine Cluster Admin	</p> <br><br>
 <ul>
 <li><p> <b>The Kubernetes Engine Admin</b> role provides access to full management of clusters and their Kubernetes API objects. This role is kind of similar to the admin role available in kubernetes RBAC. This role is applicable at a project level and hence applies to every cluster in the project. Since this is a highly privileged role, as a best practice, it is recommended grant limited users this role and to maintain admin activity logs for audit to identify what changes were made by whom.</p></li>
  <li><p><b>The Kubernetes Engine Cluster Admin</b> role provides access to management of clusters, but does not actually provide access to the cluster. So while the Cluster Admin may make changes to the clusters in a particular project, he/she will not be able to access it or make any kube operations. The permissions involved in this role are required to create a GKE cluster, and the same permissions are found in the Kubernetes Engine Admin role mentioned above.</p></li>
</ul>

<td><ol type="1">
<li><p>GKE IAM Roles: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/iam">https://cloud.google.com/kubernetes-engine/docs/how-to/iam</a></p></li>
<li><p>Admin Activity Audit Logs: <a href="https://cloud.google.com/logging/docs/audit#admin-activity">https://cloud.google.com/logging/docs/audit#admin-activity</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Role Based Access Control</td>
<td><ul>
<li><p>The best way to grant permissions is by following the principle of least privilege. To achieve that on GKE, we have to use a combination of GCP IAM roles in tandem with granular RBAC policies. </p></li>
 
<li><p> For users that don't require cluster administrative permissions (permissions that allow infrastructural and networking changes to the cluster), it is recommended to grant those users the Kubernetes Engine Viewer role, which grants read acess to gke resources. That access canthen be extended with granular RBAC configuration that grants write and/or admin access to specific namespaces. With GKE, IAM is the foundation for authorization and RBAC is the structure built on top of that.</p></li>

<li><p> So for example if we want to provide a user named John edit access to a specific 'web' namespace where web-server pods are hosted, we would grant John the Kubernetes Engine Viewer IAM role like so: <br> <pre> gcloud projects add-iam-policy-binding PROJECT_ID --member=john@example.com --role=roles/container.viewer </pre> <br> supplemented by the following role-binding:</p></li>
</ul>

<pre>
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: john-web
  namespace: web
subjects:
- kind: User
  name: john@example.com
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: edit
  apiGroup: rbac.authorization.k8s.io
</pre>

<ul>
<li><p>Defining RBAC configuration like this works relatively well at a small scale, but can quickly become difficult to manage across a variety of namespaces or users. It can also be difficult to get a full picture of authorization in your cluster when half of the configuration is in Google Cloud IAM while the other half lives in Kubernetes RBAC. GKE allows the usage of google groups withing the kubernetes RBAC to simplify user management</p></li>
 
<li><p>One way to ease RBAC management is by using tools such as this open-source one: <a href:"https://github.com/FairwindsOps/rbac-manager">rbac-manager</a></p></li>
</ul>
</td>
<td><ol type="1">
<li><p>RBAC on GKE: <a href="https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control">https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control</a></p></li>
<li><p>RBAC manager tool: <a href="https://github.com/FairwindsOps/rbac-manager">https://github.com/FairwindsOps/rbac-manager</a></p></li>
</ol></td>
</tr>
<!--
<tr class="odd">
<td>Authorization between GCP services</td>
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
-->
