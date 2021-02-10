# Detailed Security Configuration

Overview

This section is meant to provide an opinionated approach towards the
implementation security controls by domain. Although the approaches may
not be a fit for all use cases or complete for production use, they are
meant to guide the reader towards current best practices and design
considerations to achieve security control objectives.

**Controls and Architectures**

The information provided in the table below is compiled from product documentation,
blog posts, white papers and other resources and emphasize recommended configurations
that ensure Red Hat OpenShift Security and Platform Integrity.


<table>
<tbody>
<tr class="odd">
<td><strong>Security Domain</strong></td>
<td><strong>Control &amp; Architectural Suggestions</strong></td>
<td><strong>References</strong></td>
</tr>
<tr class="even">
<td><strong>Authentication & Authorization</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle authentication?</td>
<td><p>
There are several Related to Authentication and Authorization -
Users - The primary entity that interacts with the API Server, Assign permissions to the user by adding roles to the user or groups the user belongs to
Identity - Keeps a record of successful auth attempts from a specific user and identity provider.  All data concerning the source of the authentication is stored on the provider.
Service Account - Applications can directly communicate with the API, Service accounts are used in place of sharing user accouts.
Groups - Groups contain a set of specific users, users can belong to multiple groups. Permissions can be assigned to multiple users via groups.

Users are allowed interaction with OpenShift via the Authentication and Authorization Layers, A user makes a request to the API, and the authentication layer authenticates the user. The authorization layer uses RBAC to determin privileges.

There are two authentication methods: 1) OAuth Access Tokens and 2) X.509 Client Certificates

The Authentication Operator runs an OAuth Service which provide the access tokens to users when they attempt to authenticate to the API. The OAuth Server uses an identity provider to validate the request. OpenShift creates the identity and user resources after a successful login.

Configurable Identity providers include: HTPAsswd. Keystone, LDAP, GitHub or GitHub Enterprise, OpenID Connect, Google, GitLab and Basic Authentication.

By default the OpenShift Container Platform creates a cluster administrator kubeadmin after installation which should be removed once authentication is configured.

https://docs.openshift.com/container-platform/4.6/authentication/remove-kubeadmin.html

$ oc delete secrets kubeadmin -n kube-system
 <br><br>

</p>
<td><ol type="1">
<li><p>Understanding authentication: <a href="https://docs.openshift.com/container-platform/4.6/authentication/understanding-authentication.html">https://docs.openshift.com/container-platform/4.6/authentication/understanding-authentication.html</a></p></li>
<li><p>Orgs Management and Team Onboarding in OpenShift: <a href="https://www.openshift.com/blog/orgs-management-and-team-onboarding-in-openshift-a-fully-automated-approach">https://www.openshift.com/blog/orgs-management-and-team-onboarding-in-openshift-a-fully-automated-approach</a></p></li>
<li><p>Removing the kubeadmin user: <a href="https://docs.openshift.com/container-platform/4.6/authentication/remove-kubeadmin.html">https://docs.openshift.com/container-platform/4.6/authentication/remove-kubeadmin.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>How does the service you are evaluating handle Authorization?</td>
<td><ul>
<br>
After authentication, the OpenShift API request is passed along (with the asserted User info) to the Kubernetes authorization layer (after a visit to the
Audit layer). This layer is responsible for ensuring that the user has been granted permissions, by policy, to perform the requested action against the
requested resource. Although the Kubernetes authorization layer is pluggable, OpenShift does not allow customization here, and only uses the Role-Based Access Control (RBAC) authorization type

Authorization is handeled by rules, roles and bindings.
Rules - Set of permitted verbs on a set of objects. For example, whether a user or service account can create pods.
Roles - Collections of rules. You can associate, or bind, users and groups to multiple roles.
Bindings - Associations between users and/or groups with a role.

</p>
<td><ol type="1">
<li><p>Using RBAC to define and apply permissions: <a href="https://docs.openshift.com/container-platform/4.6/authentication/using-rbac.html">https://docs.openshift.com/container-platform/4.6/authentication/using-rbac.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle RBAC?</td>
<td><p>
Authorization in OpenShift is managed using role-based access control
(RBAC). OpenShift takes a deny-by-default approach to RBAC. There are two types of roles within OpenShift that determine which actions RBAC can authorize, Cluster and Local.

Cluster Role - give Users or Groups the ability to manage the OpenShift Cluster
Local Role - Users or Groups that are managing objects and attributes at the project level

Default Roles available in OpenShift:
admin - Can Manage All project Resources
basic-user - read access to the project
cluster-admin - Users with this role have access to the cluster resources. These users have full control of the cluster.
cluster-statue - this role grants the ability to get status information
edit - create, edit, change and delete common application resources from the project
self-provisioner - this role allows the creation of new projects (cluster role not a project level role)
view - Users with this role can view project resources.
</p></td>
<td><ol type="1">
<li><p>Using RBAC to define and apply permissions: <a href="https://docs.openshift.com/container-platform/4.5/authentication/using-rbac.html">https://docs.openshift.com/container-platform/4.5/authentication/using-rbac.html</a></p></li>
<li><p>How to customize OpenShift RBAC permissions: <a href="https://developers.redhat.com/blog/2017/12/04/customize-openshift-rbac-permissions/">https://developers.redhat.com/blog/2017/12/04/customize-openshift-rbac-permissions/</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle Privileged Access Management?</td>
<td><p>
OpenShift can use Security Context Constraints to control permissions for pods. These permissions include actions that a pod, a collection of containers, can perform and what resources it can access. You can use SCCs to define a set of conditions that a pod must run with in order to be accepted into the system.

SCCs allow an administrator to control:

* Whether a pod can run privileged containers.
* The capabilities that a container can request.
* The use of host directories as volumes.
* The SELinux context of the container.
* The container user ID.
* The use of host namespaces and networking.
* The allocation of an FSGroup that owns the pod’s volumes.
* The configuration of allowable supplemental groups.
* Whether a container requires the use of a read only root file system.
* The usage of volume types.
* The configuration of allowable seccomp profiles.
</p></td>
<td><ol type="1">
<li><p>Managing security context constraints: <a href="https://docs.openshift.com/container-platform/4.6/authentication/managing-security-context-constraints.html">https://docs.openshift.com/container-platform/4.6/authentication/managing-security-context-constraints.html</a></p></li>
<li><p>Managing SCCs in OpenShift: <a href="https://www.openshift.com/blog/managing-sccs-in-openshift">https://www.openshift.com/blog/managing-sccs-in-openshift</a></p></li>
<li><p>Introduction to Security Contexts and SCCs: <a href="https://www.openshift.com/blog/introduction-to-security-contexts-and-sccs">https://www.openshift.com/blog/introduction-to-security-contexts-and-sccs</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle Privileged Access Management?</td>
<td><p>
Accessing Nodes & Pods directly
</p></td>
<td><ol type="1">
<li><p></a></p></li>
<li><p></a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>Security Monitoring & Alerting</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">

<td>OpenShift Security Approach</td>
<td>
 <p>The security tooling provided and inherent in the platform encourages the utilization of security as a fluid methodology strengthening each layer of the platform and each stage of the application delivery lifecycle.

</ul>
</p>
</td>
<td><ol type="1">
<li><p>A layered approach to container and Kubernetes security: <a href="https://www.redhat.com/en/resources/layered-approach-security-detail">https://www.redhat.com/en/resources/layered-approach-security-detail</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td>Does the service you are evaluating offer Auditing Capabilities</td>
<td><p>
In OpenShift Container Platform, auditing occurs at both a host operating system context and at an OpenShift API context.

Auditing of the host operating system consists of the standard auditing capabilities provided by the auditd service in Red Hat Enterprise Linux
(RHEL) and Red Hat CoreOS (RHCOS). Audit is enabled by default in Red Hat Enterprise Linux CoreOS (RHCOS); however, the audit subsystem is running in a default configuration and without any audit rules. The auditd configuration ( /etc/audit/auditd.conf ) file should be modified as necessary to meet common organizational audit requirements such as retention and fault tolerance. Additionally, audit rules must be configured to record events.

Auditing at the OpenShift context consists of recording the HTTP requests made to the OpenShift API. The OpenShift API consists of two
components : the Kubernetes API server and the OpenShift API server. Both of these components provide an audit log, each recording the events that
have affected the system by individual users, administrators, or other components of the system. OpenShift API audit is enabled by default and is produced by both the kube-apiserver and openshift-apiserver components. The audit configuration of each is defined by a combination of default settings and corresponding custom resources named KubeAPIServer and OpenShiftAPIServer, respectively. For more information, consult the Kubernetes Auditing documentation https://kubernetes.io/docs/tasks/debug-application-cluster/audit/.
</p></td>
<td><ol type="1">
<li><p>Viewing audit logs: <a href="https://docs.openshift.com/container-platform/4.6/security/audit-log-view.html#audit-log-view">https://docs.openshift.com/container-platform/4.6/security/audit-log-view.html#audit-log-view</a></p></li>
<li><p>Configuring the audit log policy: <a href="https://docs.openshift.com/container-platform/4.6/security/audit-log-policy-config.html">https://docs.openshift.com/container-platform/4.6/security/audit-log-policy-config.html</a></p></li>
<li><p>Auditing the OS: <a href="https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/security_hardening/auditing-the-system_security-hardening">https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/security_hardening/auditing-the-system_security-hardening</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Does the service you are evaluating enforce Compliance Capabilities</td>
<td><p>
The Compliance Operator lets OpenShift Container Platform administrators describe the desired compliance state of a cluster and provides them with an overview of gaps and ways to remediate them. The Compliance Operator assesses compliance of both the Kubernetes API resources of OpenShift Container Platform, as well as the nodes running the cluster. The Compliance Operator uses OpenSCAP, a NIST-certified tool, to scan and enforce security policies provided by the content. Currently the following profiles are available for Compliance:

$ oc get -n <namespace> profiles.compliance
NAME              
ocp4-cis         
ocp4-cis-node     
ocp4-e8          
ocp4-moderate     
ocp4-ncp         
rhcos4-e8         
rhcos4-moderate   
rhcos4-ncp        
</p></td>
<td><ol type="1">
<li><p>Understanding the Compliance Operator: <a href="https://docs.openshift.com/container-platform/4.6/security/compliance_operator/compliance-operator-understanding.html">https://docs.openshift.com/container-platform/4.6/security/compliance_operator/compliance-operator-understanding.html</a></p></li>
<li><p>How does Compliance Operator work for OpenShift?:<a href="https://www.openshift.com/blog/how-does-compliance-operator-work-for-openshift-part-1">https://www.openshift.com/blog/how-does-compliance-operator-work-for-openshift-part-1</a></p></li>
<li><p>RHEL CoreOS Compliance Scanning in OpenShift 4:<a href="https://www.openshift.com/blog/rhel-coreos-compliance-scanning-in-openshift-4">https://www.openshift.com/blog/rhel-coreos-compliance-scanning-in-openshift-4</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Does the service you are evaluating enforce File Integrity & Intrusion Detection</td>
<td><p>
The File Integrity Operator is an OpenShift Container Platform Operator that continually runs file integrity checks on the cluster nodes. It deploys a daemon set that initializes and runs privileged advanced intrusion detection environment (AIDE) containers on each node, providing a status object with a log of files that are modified during the initial run of the daemon set pods.
</p></td>
<td><ol type="1">
<li><p>Understanding the File Integrity Operator: <a href="https://docs.openshift.com/container-platform/4.6/security/file_integrity_operator/file-integrity-operator-understanding.html">https://docs.openshift.com/container-platform/4.6/security/file_integrity_operator/file-integrity-operator-understanding.html</a></p></li>
<li><p>Configuring the Custom File Integrity Operator: <a href="https://docs.openshift.com/container-platform/4.6/security/file_integrity_operator/file-integrity-operator-configuring.html">https://docs.openshift.com/container-platform/4.6/security/file_integrity_operator/file-integrity-operator-configuring.html</a></p></li>
<li><p>How to install and use the File Integrity Operator in Red Hat OpenShift Container Platform 4.6: <a href="https://access.redhat.com/solutions/5751261">https://access.redhat.com/solutions/5751261</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Does the service you are evaluating offer Alerting</td>
<td><p>
In OpenShift Container Platform 4.6, the Alerting UI enables you to manage alerts, silences, and alerting rules.

Alerting rules. Alerting rules contain a set of conditions that outline a particular state within a cluster. Alerts are triggered when those conditions are true. An alerting rule can be assigned a severity that defines how the alerts are routed.

Alerts. An alert is fired when the conditions defined in an alerting rule are true. Alerts provide a notification that a set of circumstances are apparent within an OpenShift Container Platform cluster.

Silences. A silence can be applied to an alert to prevent notifications from being sent when the conditions for an alert are true. You can mute an alert after the initial notification, while you work on resolving the underlying issue.
</p></td>
<td><ol type="1">
<li><p>Configuring alert notifications: <a href="https://docs.openshift.com/container-platform/4.6/post_installation_configuration/configuring-alert-notifications.html">https://docs.openshift.com/container-platform/4.6/post_installation_configuration/configuring-alert-notifications.html</a></p></li>
<li><p>Managing alerts: <a href="https://docs.openshift.com/container-platform/4.6/monitoring/managing-alerts.html">https://docs.openshift.com/container-platform/4.6/monitoring/managing-alerts.html</a></p></li>
<li><p>Understanding cluster logging alerts: <a href="https://docs.openshift.com/container-platform/4.6/logging/troubleshooting/cluster-logging-alerts.html">https://docs.openshift.com/container-platform/4.6/logging/troubleshooting/cluster-logging-alerts.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Does the Platform provide monitoring for this service?</td>
<td><p>
OpenShift Container Platform includes a pre-configured, pre-installed, and self-updating monitoring stack that provides monitoring for core platform components. OpenShift Container Platform delivers monitoring best practices out of the box. A set of alerts are included by default that immediately notify cluster administrators about issues with a cluster. Default dashboards in the OpenShift Container Platform web console include visual representations of cluster metrics to help you to quickly understand the state of your cluster.

After installing OpenShift Container Platform 4.6, cluster administrators can optionally enable monitoring for user-defined projects. By using this feature, cluster administrators, developers, and other users can specify how services and pods are monitored in their own projects. You can then query metrics, review dashboards, and manage alerting rules and silences for your own projects in the OpenShift Container Platform web console.
</p></td>
<td><ol type="1">
<li><p>Understanding the monitoring stack: <a href="https://docs.openshift.com/container-platform/4.6/monitoring/understanding-the-monitoring-stack.html">https://docs.openshift.com/container-platform/4.6/monitoring/understanding-the-monitoring-stack.html</a></p></li>
<li><p>Configuring the monitoring stack: <a href="https://docs.openshift.com/container-platform/4.6/monitoring/configuring-the-monitoring-stack.html">https://docs.openshift.com/container-platform/4.6/monitoring/configuring-the-monitoring-stack.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>Is there Alert & Incident Management capabilities?</td>
<td><p>
Alerting is built into the platform.  Alerts can be managed via rules, queried upon, and surfaced on a visual dashboard.  Alerts can send notices to external systems
</p></td>
<td><ol type="1">
<li><p>Managing alerts: <a href="https://docs.openshift.com/container-platform/4.6/monitoring/managing-alerts.html">https://docs.openshift.com/container-platform/4.6/monitoring/managing-alerts.html</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>Data Resilience (back-up/replication)</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>Are data backups and replication capabilities provided if needed?</td>
<td>Replication is an underlying feature of etcd.  The ability to snapshot the etcd data is available via the CLI.  For real-time synchonous backups of the data store, the storage provider of the storage used to persist the etcd would provide data replication capabilities.

<td><ol type="1">
<li><p>Backing up etcd: <a href="https://docs.openshift.com/container-platform/4.6/backup_and_restore/backing-up-etcd.html">https://docs.openshift.com/container-platform/4.6/backup_and_restore/backing-up-etcd.html</a></p></li>
</ol></td>
</tr>

<tr class="even">
<td><strong>Compute High Availability</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating provide high availability to suit your requirements?</td>
<td>The standard OpenShift Architecture consists of 3 control plane nodes or masters and at least two worker nodes providing localized high availability in the event a master node or worker node is lost. For multi-site High Availability there are two ways to achieve this:
1) Install complete clusters across Availability Zones or Sites and have applications deployed to multiple clusters with load balancing between the two separate clusters.
2) Have an OpenShift Cluster span 3 sites with a master node in each site and workers distributed. This type of cluster is extremely sensitive to network and other external variables and extreme consideration and testing should be applied to the architecture and deployment.
<td><ol type="1">
<li><p>OpenShift Container Platform architecture: <a href="https://docs.openshift.com/container-platform/4.6/architecture/architecture.html">https://docs.openshift.com/container-platform/4.6/architecture/architecture.html</a></p></li>
<li><p>Stretch and multi-site clusters Capabilities and Support: <a href="https://access.redhat.com/articles/3220991#policies">https://access.redhat.com/articles/3220991#policies</a></p></li>
</ol></td>
</tr>
<tr class="even">
<td><strong>Cluster Versioning</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>How does this service ensure it is using the latest stable/secure version of the underlying software?</td>
<td>OpenShift provides over the air updates to both the underlying RHCOS nodes as well as the cluster itself. The entire platform is treated as one composable platform. Updates are packed in containers and can be set to apply automatically from selected channels and releases.Cluster Upgrades are managed via the Cluster Version Operator, the Machine Config Operator and some individual Operators. Updates and Patches are managed by the Cluster Administrator. Updates to nodes are done in a rolling fashion ensuring zero cluster downtime for applications designed according to cloud native principals. ALL platform Operators ensure that any drift from unsupported configuration changes are reset to the baseline configuration.

RHCOS is tightly coupled with the platform in order to consistently apply OS Updates, the Machine Config Operator can apply upgrades automatically in a coordinated fashion, minimizing cluster impact. Updates are released with the OpenShift Cluster update payload ensuring OS releases are in sync with Cluster releases.

OpenShift Updates can bee applied via the Web Console or CLI -

From the Web Console the user will be notified if the update is available, from there they can simply click update.
From the command line, there are a few steps:
1) Check if the cluster is available -   oc get clusterversion
2) Check if an update is available - oc adm upgrade
3)  Apply an update to the latest release - oc adm upgrade --to-latest=true
<td><ol type="1">
<li><p>Updating a cluster within a minor version by using the CLI
: <a href="https://docs.openshift.com/container-platform/4.6/updating/updating-cluster-cli.html">https://docs.openshift.com/container-platform/4.6/updating/updating-cluster-cli.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating manage the underlying operating system the service is built on?</td>
<td>As mentioned in the latest version category the operating system  (Red Hat CoreOS Operating System) and OpenShift kubernetes orchestration platform are tightly coupled together to ensure consistency and interoperability between fast moving components. Since RHCOS is only intended to be used by Red Hat OpenShift, it's installable via the OpenShift Installer Provisioned Infrastructure (IPI) or User Provisioned Infrastructure (UPI) where the user is responsible to downloading the image and generating the ingnition control scripts.

The RHCOS Operating System is designed to be a single purpose container Operating System only supported in the capacity of OpenShift Container Platform usage which allows it to be more targeted and controlled than general purpose Operating Systems. It's based on Red Hat Enterprise Linux and inherits all the of the security and hardware certifications of RHEL in addition to the secure and stable OS Lifecycle. Reiterating the single use, the OS lacks non-critical components generally found in multi-purpose Operating Systems, which greatly reduce the attack surface. The state of the Operating System is stored within the OpenShift Container Platform ensuring controlled immutability, allowing the nodes to be scaled in either direction. The container Runtime is CRI-O which is designed to be specifically used with Kubernetes implementing only the features needed, again minimizing the attack surface. The last two capabilities to highloght are the way in which software is updated using rpm-ostree and the way RHCOS is configured using the Machine Config Operator. RPM-OSTREE features transactional upgrades. Updates are delivered by way of a container as part of the OpenShift upgrade process and extracted to disk. From there the bootloader is modified to boot into the updated version. There is the ability to rollback as neccessary. The Machine Config Operator handles the OS upgrades directed using rpm-ostree as well as maintaining and applying node configurations. This state is maintained accross all cluster nodes.

Please see the associated links for RHCOS Configuration, Hardening, Compliance Scanning and Installation.
<td><ol type="1">
<li><p>Creating Red Hat Enterprise Linux CoreOS (RHCOS) machines: <a href="https://docs.openshift.com/container-platform/4.6/installing/installing_bare_metal/installing-bare-metal.html#creating-machines-bare-metal">https://docs.openshift.com/container-platform/4.6/installing/installing_bare_metal/installing-bare-metal.html#creating-machines-bare-metal</a></p></li>
<li><p>RHEL CoreOS Compliance Scanning in OpenShift 4: <a href="https://www.openshift.com/blog/rhel-coreos-compliance-scanning-in-openshift-4">https://www.openshift.com/blog/rhel-coreos-compliance-scanning-in-openshift-4</a></p></li>
<li><p>Hardening RHCOS: <a href="https://docs.openshift.com/container-platform/4.6/security/container_security/security-hardening.html">https://docs.openshift.com/container-platform/4.6/security/container_security/security-hardening.html</a></p></li>
</ol></td>
</tr>
<td><strong>Certificate and Key Management</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle Certificate and Key Management?</td>
<td>All certificates for internal traffic is managed by OpenShift and rotated automatically. Egress (Proxy) traffic CA is configurable. Ingres traffic is configurable
<td><ol type="1">
<li><p>User-provided certificates for the API server: <a href="https://docs.openshift.com/container-platform/4.6/security/certificate_types_descriptions/user-provided-certificates-for-api-server.html">https://docs.openshift.com/container-platform/4.6/security/certificate_types_descriptions/user-provided-certificates-for-api-server.html</a></p></li>
</ol></td>
</tr>
<td><strong>Encryption</strong></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle Encryption at Rest?</td>
<td>Encrypting etcd data - To be updated
<td><ol type="1">
<li><p>Encrypting ETCD: <a href="https://docs.openshift.com/container-platform/4.6/security/encrypting-etcd.html">https://docs.openshift.com/container-platform/4.6/security/encrypting-etcd.html</a></p></li>
</ol></td>
</tr>
<tr class="odd">
<td>How does the service you are evaluating handle Encryption in Transit?</td>
<td>OpenShift 4.7 adds IPSec support for the OVN-Kubernetes CNI - To be updated
<td><ol type="1">
<li>To Be Added</li>
</ol></td>
</tr>
<td><strong>Network Policy and Security</strong></td>
<td></td>
<td></td>
</tr>

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
