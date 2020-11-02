# Service Approval Accelerator for Azure Kubernetes Service (AKS)

## Identity & access management

### Authentication

Kubernetes itself does not manage users and authentication--this is intended to be done by the service provider [1].
This can be done by AKS, but it is recommended to use the integration with Azure Active Directory [2].
Using AD lets you manage users in a central place that can be used across many different Azure services.

1.  https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/
2.  https://docs.microsoft.com/en-us/azure/aks/concepts-identity

### Authorization

It is considered best practice to manage user authorization with role-based access control (RBAC); see below.

"Managed identities" can be used for authorization between pods and other Azure services, such as Storage and databases.
Pod identities are defined and assigned permissions.
Specified pods then authenticate as an identity, receiving an access token which can be used by the application to authorize with permitted services.
Hence, your application code doesn't need to include credentials to access a service; each pod authenticates with its own identity, so you can audit and review access.

1.  https://docs.microsoft.com/en-us/azure/aks/developer-best-practices-pod-security

### RBAC

RBAC is split into two domains.

The first is at the Azure service level.
User accounts in Active Directory can be assigned roles which restrict access to a certain set of services or resources in the subscription.
User access to AKS can be controlled to a fine level of detail, but there are some built-in roles [2]:

*   Azure Kubernetes Service Cluster Admin Role
*   Azure Kubernetes Service Cluster User Role
*   Azure Kubernetes Service Contributor Role

Secondly, there is the Kubernetes RBAC system.
This permits fine-grained control of the Kubernetes resources themselves via Role and ClusterRole definitions.
These are then linked with users via integration with Active Directory by RoleBindings and ClusterRoleBindings.

There is a feature in Preview which integrates the Kubernetes RBAC authorization with Azure RBAC [1].
This means you can define Kubernetes roles in the same system you define roles for Azure resources, and AKS will use Azure to authorize Kubernetes access requests.

1.  https://docs.microsoft.com/en-us/azure/aks/concepts-identity
2.  https://docs.microsoft.com/en-us/azure/role-based-access-control/built-in-roles#containers
3.  https://kubernetes.io/docs/reference/access-authn-authz/rbac/

### Privileged Access Management

As discussed, identities should be managed via Azure Active Directory.
Azure provides guidance and AD functionality for managing privileged identities, such as the Privileged Access Management notification service.

1.  https://docs.microsoft.com/en-us/azure/active-directory/users-groups-roles/directory-admin-roles-secure

## Encryption & secure data management

### Encryption of data at rest

AKS allows connecting pods to persistent storage via volumes.
Volumes are backed by the Azure Storage product, which encrypts all data as standard.
The encryption, decryption, and key management processes are transparent to users.
Customers can also choose to manage their own keys using Azure Key Vault.
There are a variety of methods for authenticating with the Storage platform.

1.  https://docs.microsoft.com/en-us/azure/storage/common/storage-introduction
2.  https://docs.microsoft.com/en-gb/azure/aks/concepts-storage

### Encryption of data in transit

In-transit encryption is generally more an application layer concern and not considered by the lower-layer routing and network services that Kubernetes offers.
However, it does support "ingress controllers" which can process data at the application level, including decryption of TLS for routing to pods.
This allows centralised management of TLS at the Kubernetes networking level, even for pod applications which do not support it.

1.  https://docs.microsoft.com/en-gb/azure/aks/concepts-network

### Certificate and key management

Ingress controllers can be configured with user-managed certificates (stored as AKS secrets) [1] or Let's Encrypt (via a Kubernetes Issuer resource) [2].

Kubernetes has a secrets system which manages sensitive application information like certificates and keys and can provide them securely to containers [3].
This is managed by AKS.

Many Azure services support configuration of customer-managed keys; see the BYOK section below.

1.  https://docs.microsoft.com/en-gb/azure/aks/ingress-own-tls
2.  https://docs.microsoft.com/en-gb/azure/aks/ingress-tls
3.  https://kubernetes.io/docs/concepts/configuration/secret/

### BYOK/HYOK Management

Azure products in general support two encryption models [1]:

*   Service-managed keys: Provides a combination of control and convenience with low overhead.

*   Customer-managed keys: Gives you control over the keys, including Bring Your Own Keys (BYOK) support, or allows you to generate new ones.

Customer-managed keys are managed via the Azure Key Vault service.
Azure Storage supports both of these models.

Additionally, the AKS OS disk can be encrypted with a customer-managed key via an Azure DiskEncryptionSet [3].

HYOK is only available in the Azure Information Protection service, which is irrelevant to Kubernetes [2].

1.  https://docs.microsoft.com/en-gb/azure/security/fundamentals/encryption-overview
2.  https://docs.microsoft.com/en-us/azure/information-protection/configure-adrms-restrictions
3.  https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys

## Network security

### IP Firewall Rules

Security of internal network connections within the Kubernetes architecture is ensured by the Azure service.

Application-layer filtering can be performed by ingress controller and load balancer software.

For transport layer security, the Azure Firewall service can be configured to protect AKS.

1.  https://docs.microsoft.com/en-gb/azure/firewall/protect-azure-kubernetes-service

### Data Exfiltration Prevention & Data Loss Prevention

General Kubernetes best practices should be followed for the networking of pods and management of incoming traffic to ensure storage is not accessible except by the resources that need it.

Azure offers Web Application Firewall (WAF) services for extra safeguarding of web applications.
These work at protocol-level and can detect and prevent a wide class of attacks, including injections, XSS, and HTTP violations.
The Azure Application Gateway provides a WAF and can be integrated with AKS [2].

1.  https://docs.microsoft.com/en-us/azure/aks/operator-best-practices-network
2.  https://docs.microsoft.com/en-us/azure/web-application-firewall/ag/ag-overview

## Logging & monitoring

### Security Monitoring & Alerting

The Application Gateway WAF can integrate with and report detected attacks to the Azure Monitor service.

Monitor provides a wide variety of features, including a flexible alert system, which can integrate with AKS.

Azure Security Center presents a top-level overview of the entire infrastructure and network topology (including AKS), and offers recommendations to improve security.

1.  https://docs.microsoft.com/en-us/azure/azure-monitor/platform/alerts-overview
2.  https://docs.microsoft.com/en-us/azure/security-center/security-center-introduction

### Service Monitoring

The Azure Monitor service has a wide set of features for monitoring containers and clusters.
It can measure/record

*   resource utilisation of containers, pods and clusters
*   logs from container applications
*   variations in usage over average and heaviest loads
*   load, latency, and capacity of volumes from Azure Storage

1.  https://docs.microsoft.com/en-us/azure/azure-monitor/insights/container-insights-overview
2.  https://docs.microsoft.com/en-us/azure/azure-monitor/insights/storage-insights-overview

### Alert & Incident Management

The Azure security and monitoring services can be configured to raise alerts for incidents in container applications, network firewalls, and storage services.

The Kubernetes system itself, being fully managed by Azure, does not require customer response to incidents at the cloud service level.

## Resilience & recovery

### Data Resilience (back-up/replication)

N/A

### Compute High Availability

Kubernetes itself is designed to facilitate creating highly available services.
The computing resources backing the nodes, containers, and control plane are managed by Azure.

In addition to Kubernetes pod and cluster auto-scalers, Azure Container Instances can be integrated with AKS to provide extra computing resources in the case of a burst in traffic.
Pods are created on ACI and connect to AKS via a "virtual node" over a private network.
This circumvents the provisioning of new Kubernetes nodes in AKS and allows a much faster scaling response than a typical horizontal auto-scaler.

AKS also automatically manages node repair when they fail periodic health checks.
In summary, if a node remains unhealthy for 10 consecutive minutes, the following actions are taken.

1.  Reboot the node
1.  If the reboot is unsuccessful, reimage the node
1.  If the reimage is unsuccessful, create and reimage a new node

1)  https://docs.microsoft.com/en-us/azure/aks/concepts-scale
2)  https://docs.microsoft.com/en-us/azure/aks/node-auto-repair

## Underlying OS

### Use of Latest Version

Azure fully manages the host OS for the nodes and the rest of the core Kubernetes system.

Linux nodes run an optimized Ubuntu distribution using the Moby container runtime.
Windows Server nodes run an optimized Windows Server 2019 release and also use the Moby container runtime.
When an AKS cluster is created or scaled up, the nodes are automatically deployed with the latest OS security updates and configurations.

The Azure platform automatically applies OS security patches to Linux nodes on a nightly basis.
If a Linux OS security update requires a host reboot, that reboot is not automatically performed.
You can manually reboot the Linux nodes, or a common approach is to use Kured, an open-source reboot daemon for Kubernetes.
Kured runs as a DaemonSet and monitors each node for the presence of a file indicating that a reboot is required.

For Windows Server nodes, Windows Update does not automatically run and apply the latest updates.
On a regular schedule around the Windows Update release cycle and your own validation process, you should perform an upgrade on the Windows Server node pool(s) in your AKS cluster.
This upgrade process creates nodes that run the latest Windows Server image and patches, then removes the older nodes.

To upgrade the Kubernetes software, a similar "cordon and drain" system is used.
New, upgraded nodes are added silently added to the cluster, while old existing nodes are gracefully shut down, recreating pods and resources on the new nodes.
This happens without any downtime or intervention from the customer.

1.  https://docs.microsoft.com/en-us/azure/aks/concepts-security

## CSP access

???

## Dependent services

N/A
