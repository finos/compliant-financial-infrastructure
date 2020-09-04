# Overview

This section is meant to provide an opinionated approach towards the implementation security controls by domain. Although the approaches may not be a fit for all use cases or complete for production use,  
they are meant to guide the audience towards current best practices and design considerations to achieve security control objectives.

Azure database for PostgreSQL is available in 2 deployment options.
- Single Server
- Hyperscale (Citus) – scale horizontally across multiple nodes

Although this document largely applies to both the offerings, some of the sections have been created with single server as the focus.

## Controls and Architectures
This table maps Security Domain to the corresponding controls and architectural best practices as documented in Azure public documentation, white papers, and blog posts.

<table class="tg">
<thead>
  <tr>
    <th class="tg-7btt">Top level</th>
    <th class="tg-fymr">Breakdown</th>
    <th class="tg-fymr">Detail</th>
    <th class="tg-fymr">Public Links</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td class="tg-amwm">External Certification</td>
    <td class="tg-8zwo"><a href="https://azure.microsoft.com/fr-fr/blog/compliance-offerings-for-azure-database-for-mysql-and-azure-database-for-postgresql/">Article on Compliance</a></td>
    <td class="tg-0lax">
        <p>Azure Database for PostgreSQL is certified compliant for:</p>
        <ul>
            <li>ISO 27001:2013</li>
            <li>ISO 27018:2014*</li>
            <li>CSA STAR Attestation</li>
            <li>CSA STAR Certification</li>
            <li>HIPAA / HITECH Act</li>
            <li>PCI DSS Level 1</li>
            <li>SOC1</li>
            <li>SOC2</li>
            <li>SOC3</li>
            <li>EU Model Clauses</li>   
            <li>UK G-cloud</li>
        </ul>
    </td>
    <td class="tg-0lax"><a href="https://azure.microsoft.com/fr-fr/blog/compliance-offerings-for-azure-database-for-mysql-and-azure-database-for-postgresql/">Article on Compliance</a></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="3">Identity &amp; Access Management</td>
    <td class="tg-fymr">Authentication</td>
    <td class="tg-f8tv">
        <p>
            It is recommended to enable AAD integration with PAzure Database for PostgreSQL.
            <br><br>
            Microsoft Azure Active Directory (Azure AD) authentication is a mechanism of connecting to Azure Database for PostgreSQL using identities defined in Azure AD.
            With Azure AD authentication, you can manage database user identities and other Microsoft services in a central location, which simplifies permission management.            
            <br><br>
            when using Azure AD for authentication, the security features that build in AAD can help to protect the login process. The AAD token can be obtained in various
            ways which could help to protect most of the password attach. 
            <br><br>
            Azure managed identity can be used by many azure resources (VM, app services, azure function, etc) to have a controled identity within AAD, which elimate the needs store and
            manage any form of password or token for deployed application. It further strenghtens the protection for authentication.
            <br><br>
            When AAD integration is not enabled, simple user/pass is used for authentication in postgres.
         </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-aad-authentication">Use Azure Active Directory for authenticating with PostgreSQL</a></td>
  </tr>
  <tr>
    <td class="tg-fymr">Authorization</td>
    <td class="tg-f8tv">
        <p>
            It is recommended to only assign AAD groups roles within postgres, and not to create indivisual users within postgres itself.            
            <br><br>
            When setting up AAD integration for Azure postgres, an AAD account (either group or user) needs to be specified as the azure ad administorator for the postgres server.             
            Only the AAD users that are specified can initially connect and add more AAD user/group into postgres and assign database roles.            
            <br><br>
            Custom database roles with least privileged access is always preferable, logically grouped permissions can be used to ease the burdon of administration
            <br><br>                        
            When a AAD group is added into the postgres server, any member of the group can login and assume the database role that is assigned to the group. It thus recommended to create and add 
            AAD groups only to postgres, and manage all users and role assignments using AAD, and eleminate the needs to create individual users within postgres.            
        </p>
    </td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Privileged Access Management</td>
    <td class="tg-f8tv">
        <p>
            Only authorised Database Administrators group members should be allowed to create or update or delete any instance of Azure database for PostgreSQL.
            Additionally, this should happen in accordance to the defined deployment process as well (e.g. via a build service, etc.).
        </p>
        <ul>
            <li>One admin group per PostgreSQL server</li>
            <li>RBAC should be used for database admins with least privileged access, custom roles per user are preferable but to ease the burden of management we
            have provided a potential RBAC model</li>
            <li>The Banks existing privileged account security system should be integrated</li>
            <li>The Banks existing systems for user account re-certification should be integrated</li>
        </ul>
        Azure AD provides PIM service to manage, control, and monitor access to the priviledge accoung/group.
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/active-directory/privileged-identity-management/pim-configure">1. Azure AD PIM</a></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="4">Encryption &amp; Secure Data Management</td>
    <td class="tg-fymr">Encryption in Transit</td>
    <td class="tg-0pky">
        <p>
        Microsoft uses the Transport Layer Security (TLS) protocol to protect data when it’s traveling between the cloud services and customers. Microsoft data-centers
        negotiate a TLS connection with client systems that connect to Azure services. TLS provides strong authentication, message privacy, and integrity (enabling
        detection of message tampering, interception, and forgery), interoperability, algorithm flexibility, and ease of deployment and use. Perfect Forward Secrecy
        (PFS) protects connections between customers’ client systems and Microsoft cloud services by unique keys. Connections also use RSA-based 2,048-bit encryption key
        lengths. This combination makes it difficult for someone to intercept and access data that is in transit. <br><br>
        All database communications from the banks application servers to Azure PostgreSQL should use TLS1.2, Azure policy should be used to enforce encryption in transit
        on all Azure PostgreSQL servers.</p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/security/fundamentals/encryption-overview#encryption-of-data-in-transit">Encryption of Data in Transit</a></td>
  </tr>
  <tr>
    <td class="tg-fymr">Encryption at Rest</td>
    <td class="tg-0pky">
        <p>
            Encryption at rest is required by an organization’s need for data governance and compliance efforts. Industry and government regulations such as HIPAA, PCI and
            FedRAMP, lay out specific safeguards regarding data protection and encryption requirements. All data that is stored in a cloud environment (Data at Rest) must be
            encrypted. The minimal encryption is using the native cloud functions to encrypt on data level, including virtual machines, containers and data storage.<br><br>
            Data encryption with customer-managed keys for Azure Database for PostgreSQL Single server enables you to bring your own key (BYOK) for data protection at rest.
            It also allows organizations to implement separation of duties in the management of keys and data. With customer-managed encryption, you are responsible for, and
            in a full control of, a key’s lifecycle, key usage permissions, and auditing of operations on keys. (Detail in BYOK/HOYK Management Section) <br><br>
            Data encryption with customer-managed keys for Azure Database for PostgreSQL Single server, is set at the server-level. For a given server, a customer-managed key,
            called the key encryption key (KEK), is used to encrypt the data encryption key (DEK) used by the service. The KEK is an asymmetric key stored in a customer-owned
            and customer-managed Azure Key Vault instance.
        </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/security/fundamentals/encryption-overview#encryption-of-data-at-rest">Encryption at rest</a></td>
  </tr>
  <tr>
    <td class="tg-fymr">Certificate and Key Management</td>
    <td class="tg-0pky">
        <p>
            Keys that are used for Data at Rest encryption need to be protected and managed, either by using Bring Your Own Key (BYOK) with a reproducible hierarchy and a
            process to protect the transfer of the key blobs, or by using Hold Your Own Key (HYOK), i.e. using HSM-protected keys on-prem. It must be ensured that all Data
            at Rest can be decrypted using keys protected by on-prem HSMs and that it is possible to revoke the access of the cloud provider to all key material.
        </p>
    </td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">BYOK/HYOK Management</td>
    <td class="tg-f8tv">
        <h4>BYOK</h4>
        <p>
            Using BYOK, the key material is generated and tested on-prem and exported to hardware-protected storage in the cloud. Keys are managed on-prem by the the banks
            internal key management solution. For transport, the key must be encrypted and authenticated. Furthermore, it must be ensured that the exported key material is
            bound to its future use-case and location. Each key has a life-time defined by criticality and algorithm, clear roll-over processes must be defined and followed.
            <br><br>
            The required processes for BYOK are:
        </p>
        <ul>
            <li>Creation / deletion of a key store</li>
            <li>Managing the access rights of the key store</li>
            <li>BYOK process: generate a key, transport it to the cloud service</li>
            <li>Import/installation of keys, ensuring intended usage</li>
            <li>Rollover a key before expiration</li>
            <li>Incident handling, e.g. in case of key compromise</li>
        </ul>
        <h4>HYOK</h4>
        <p>
            Using Hold your own keys (HYOK) is preferable, as it gives the bank more control over the keys and enables easy potential migration between platforms. In particular,
            applications developed by the bank should not use CSP-specific vaults directly. In case a native CSP service integrates with the vault, a BYOK strategy is preferable,
            e.g. for Database as a Service and Kubernetes.
            <br><br>
            HYOK are best practice future state, although this is currently not supported by Microsoft Azure
        </p>
        <h4>Key Management using Key Vault</h4>
        <p>
            Key Vault is a cloud-based, external key management system. It’s highly available and provides scalable, secure storage for RSA cryptographic keys, optionally backed by
            FIPS 140-2 Level 2 validated hardware security modules (HSMs). It doesn’t allow direct access to a stored key, but does provide services of encryption and decryption to
            authorized entities. Key Vault can generate the key, imported it, or have it transferred from an on-premises HSM device.
            <br><br>
            When you configure data encryption with a customer-managed key in Key Vault, continuous access to this key is required for the server to stay online. If the server loses
            access to the customer-managed key in Key Vault, the server begins denying all connections within 10 minutes. The server issues a corresponding error message, and changes
            the server state to Inaccessible. The only action allowed on a database in this state is deleting it.<br><br>
            After Azure Database for PostgreSQL Single server is encrypted with a customer’s managed key stored in Key Vault, any newly created copy of the server is also encrypted.
            You can make this new copy either through a local or geo-restore operation, or through read replicas. However, the copy can be changed to reflect a new customer’s managed
            key for encryption. When the customer-managed key is changed, old backups of the server start using the latest key.<br><br>
            To avoid issues while setting up customer-managed data encryption during restore or read replica creation, it’s important to follow these steps on the master and restored/replica servers:
        </p>
            <ol>
                <li>Initiate the restore or read replica creation process from the master Azure Database for PostgreSQL Single server.</li>
                <li>Keep the newly created server (restored/replica) in an inaccessible state, because it’s unique identity hasn’t yet been given permissions to Key Vault.</li>
                <li>On the restored/replica server, revalidate the customer-managed key in the data encryption settings. This ensures that the newly created server is given wrap and unwrap permissions to the key stored in Key Vault.</li>
            </ol>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-data-encryption-postgresql">Encryption with Customer managed Keys</a></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="3">Network Security </td>
    <td class="tg-fymr">Endpoint Localisation</td>
    <td class="tg-f8tv">
        <p>
            Private Link allows you to connect to various PaaS services in Azure via a private endpoint. Azure Private Link essentially brings Azure services inside your private Virtual
            Network (VNet). The PaaS resources can be accessed using the private IP address just like any other resource in the VNet.
        </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-data-access-and-security-private-link">Private Link for Azure Database for PostgreSQL-Single server</a></td>
  </tr>
  <tr>
    <td class="tg-fymr">IP Firewall Rules</td>
    <td class="tg-0pky">
        <p> PostgreSQL service utilizes postgres built-in firewall mechanism using <code>ph_hba.conf</code> file. <br><br>
            Connections to an Azure Database for PostgreSQL server are first routed through a regional gateway. The gateway has a publicly accessible IP, while the server
            IP addresses are protected. The following diagram illustrates the traffic flow:</p>
        <img src="media/net-security.png" width="700" height="300">
        <p>
            A newly created Azure Database for PostgreSQL server has a firewall that blocks all external connections. Though they reach the gateway, they are not allowed to
            connect to the server. <br><br>
            To begin using your server from another computer, you need to specify one or more server-level firewall rules to enable access to your server. Use the firewall rules
            to specify which IP address ranges from the Internet to allow. Access to the Azure portal website itself is not impacted by the firewall rules. <br><br>
            This example uses Acces to Azure Services enabled or Public peering, this will be different for private link but gives some insight to how communications are proxied
            to Azure PAAS services.<br><br>
        </p>
        <h4>Connecting from Azure</h4>
        <p>
            To allow applications from Azure to connect to your Azure Database for PostgreSQL server, Azure connections must be enabled. For example, to host an Azure Web Apps
            application, or an application that runs in an Azure VM, or to connect from an Azure Data Factory data management gateway. The resources do not need to be in the same
            Virtual Network (VNet) or Resource Group for the firewall rule to enable those connections. When an application from Azure attempts to connect to your database server,
            the firewall verifies that Azure connections are allowed. This option configures the firewall to allow all connections from Azure including connections from the
            subscriptions of other customers.<br><br>
            Allow connections from Azure should be disabled on every PostgreSQL firewall that is using an in-house application running on servers, other PAAS services connection
            into Azure PostgreSQL should be assessed individually for best practice secure configuration.
        </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-firewall-rules">Firewall Rules</a></td>
  </tr>
  <tr>
    <td class="tg-fymr">Data Exfiltration Prevention <br>&amp; Data Loss Prevention</td>
    <td class="tg-0pky">
        <p>
            Best practice is to disable all Azure service traffic to Azure Database for PostgreSQL Single server via the public endpoint by setting Allow Azure Services to OFF.<br><br>
            Ensure no public IP addresses or ranges are allowed to access the server either via NSG and PostgreSQL firewall rules, and no private IP addresses are allowed to
            connect to the Azure PostgreSQL server via virtual network service endpoints.<br><br>
            Create a Private Link to the Azure PostgreSQL server and only allow traffic to the Azure Database for PostgreSQL Single server using the Private IP address of the VM.<br><br>
            With Private Link, you can now set up network access controls like NSGs to restrict access to the private endpoint.<br><br>
            Restricting access to the Azure PostgreSQL server from only the IPV4 addresses that are found in rules applied to both the NSG attached to the subnet that the Azure PostgreSQL
            server has a private endpoint and the Azure PostgreSQL firewall reduces data exfiltration risks.<br><br>
            With Private Link, you can enable cross-premises access to the private endpoint using Express Route (ER), Then you can subsequently disable all access via public endpoints.
        </p>
    </td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="3">Logging &amp; Monitoring</td>
    <td class="tg-fymr">Security Monitoring &amp; Alerting</td>
    <td class="tg-f8tv">
        <p>
            By default, pgAudit log statements are emitted along with regular log statements by using Postgres’s standard logging facility. In Azure Database for PostgreSQL, these .log
            files can be downloaded through the Azure portal or the CLI. The maximum storage for the collection of files is 1 GB, and each file is available for a maximum of seven days
            (the default is three days). This service is a short-term storage option. <br><br>
            Alternatively, the bank can configure all logs to be emitted to Azure Monitor’s diagnostic log service. If Azure Monitor diagnostic logging is enabled, all logs will be
            automatically sent (in JSON format) to Azure Storage, Event Hubs, and/or Azure Monitor logs, depending on the configuration. <br><br>
            Azure Events that may impact security of platform to be monitored and alerted upon <br><br>
            PGAUDIT and Query store should be logging for all PostgreSQL servers.
        </p>
    </td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Service Monitoring</td>
    <td class="tg-f8tv">
        <ul>
            <li>All AzurePostgreSQL logs should be stored in a Log Analytics workspace</li>
            <li>Log analytics workspaces should have table level RBAC implemented</li>
            <li>All table level RBAC roles should have least privelaged permissions</li>
            <li>Azure monitor should be used for operational insights</li>
        </ul>
    </td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Alert &amp; Incident Management</td>
    <td class="tg-f8tv">
        <p>
            Advanced Threat Protection (ATP) provides additional layer of security that enables customers to detect and react to anamolous activities.</br>
            ATP provides alerts on multiple fronts like the database activities, vulnerabilities, anamolous database activities, various query patterns.</br>
            ATP can raise alerts on unusual activities like access from unusual location, access from unfamiliar principal, access from potentially harmful applications etc. </br>
            ATP for Azure database for PostgreSQL integrates its alerts with Azure Security Center. </br>
            The security alerts within the security center provide the list of potential alerts and any threats detected on the database. </br>
            Enable the Advanced Threat Protection to send alerts to certain email addresses to receive alerts upon detection of anamolous activities. </br>
        </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-data-access-and-security-threat-protection">Advanced Threat Protection</a></br></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="3">Resilience &amp; Recovery</td>
    <td class="tg-fymr">Data Resilience (back-up)</td>
    <td class="tg-0pky">
        <p>
            Backups are created and stored automatically in storage at either a locally redundant storage or geo-redundant storage based on the configuration set up.</br>
            Backup retention period can be anywhere between 7 and 35 days.</br>
            Point in time restore can be done in the same region as the original server. A new server is created from the original server's backups.</br>
            Geo-restore can be done if the server is configured for geo-redundant storage and can be restored to a server to a different region.</br>
            The configuration on the restored server needs to mimic the configuration on the original server (a script needs to be developed to copy the configuration)</br>
            In case of Hyperscaale, backups of each node are created and stored locally.</br>
            The cluster can be restored to a specified point in time using the backups.</br>
            Restoring the cluster creates a new cluster from the original nodes backups.</br>
        </p>
    </td>
    <td class="tg-f8tv">
     <p>
        <a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-backup">Backups</a></br>
        <a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-hyperscale-backup">Backups Hyperscale</a></br>
    </p>
    </td>
    </tr>
    <tr>
    <td class="tg-fymr">Data Resilience (Replication)</td>
    <td class="tg-0pky">
        <p>
            Azure allows to create replicas from an Azure database for PostgreSQL to a read-only server. Up to 5 replicas can be created.</br>
            Replicas are updated asynchronously using PostgresSQL asynchronous replication. The data becomes eventually consistent between the read replicas.</br>
            The read replicas can be used to improve performance of intensive workloads by isolating some read workloads to replicas and write workloads to the master.</br>
            Read replicas can be created in either the same region as the master server or in a different region. Cross region replication helps in the DR set up.</br>
            The replicas don’t inherit the firewall rules or the service endpoint of the master server and therefore the rules must be set separately for the replica.</br>
            There is no automated fail-over between the master server and the replica. The fail-over process is manual and requires the replication to be stopped. Once the
            replication is stopped, the replica becomes stand-alone server. The replica server is restarted, so it can accept writes.</br>
            The applications need to be pointed to the replica. A script that is capable of performing the above steps to stop replication and making the replica to a master
            server should be created to achieve this.</br>
            Azure replication support parameter needs to be configured for replicas as this feature is dependent on Postgres WAL (Write Ahead Log). The parameter has 3 options
            viz. 'Off, 'Replica' and 'Logical'. This needs to be set to at-least 'Replica' for logging to be configured for replica servers.</br>
        </p>
    </td>
    <td class="tg-f8tv">
     <p>
        <a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-read-replicas">Read Replicas</a></br>
        <a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-business-continuity">Business Continuity</a></br>
    </p>
    </td>
  </tr>
  <tr>
    <td class="tg-fymr">Compute High Availability</td>
    <td class="tg-0pky">
    <p>
        99.99% availability is defined in the SLA document and handled by Microsoft. <br><br>
        Additionally, high availability can be implemented by creating a read-only replica of the database in a different region to the
        one the database is deployed in. A script can be used to fail-over between the main database and the read replica. If the replication
        is broken the read-replica becomes write-able and can act as the main database. This method ensures high availability between regions,
        on top of the high availability guaranteed by Microsoft.
    </p>
    </td>
    <td class="tg-0pky"><a href="https://docs.microsoft.com/en-us/azure/postgresql/concepts-high-availability">Availability in PostgreSQL</a></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="2">Underlying OS</td>
    <td class="tg-fymr">Use of Latest Version</td>
    <td class="tg-f8tv">PostgeSQL is a fully managed service, therefore Azure handle all of the underlying OS patches and version updates.</td>
    <td class="tg-0pky"><a href="https://azure.microsoft.com/en-us/support/legal/sla/postgresql/v1_1/">SLA for PostgreSQL</a></td>
  </tr>
  <tr>
    <td class="tg-1wig">-</td>
    <td class="tg-8zwo">PostgeSQL is a fully managed service, therefore Azure handle all of the underlying OS patches and version updates.</td>
    <td class="tg-0lax"><a href="https://azure.microsoft.com/en-us/support/legal/sla/postgresql/v1_1/">SLA for PostgreSQL</a></td>
  </tr>
  <tr>
    <td class="tg-7btt">CSP Access</td>
    <td class="tg-fymr">-</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-7btt">Dependent Services</td>
    <td class="tg-fymr">-</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
  </tr>
</tbody>
</table>

### Abbreviation

<table>
    <thead>
      <tr>
        <th>Acronym</th>
        <th>Meaning</th>
        <th>Description</th>
      </tr>
    </thead>
    <tbody>
        <tr>
            <td>NSG</td>
            <td>Network Security Group</td>
            <td><a href="https://docs.microsoft.com/en-us/azure/virtual-network/security-overview">docs</a></td>
        </tr>
        <tr>
            <td>CSP</td>
            <td>Cloud Service Provider</td>
            <td>The company which provides network services, infrastructure, or business applications in the cloud</td>
        </tr>
        <tr>
            <td>SLA</td>
            <td>Service Level Agreement</td>
            <td>
                <p>A service-level agreement (SLA) defines the level of service you expect from a vendor, laying out the metrics by which service is measured,
                as well as remedies or penalties should agreed-on service levels not be achieved.</p>
            </td>
        </tr>
        <tr>
            <td>BYOK</td>
            <td>Bring Your Own Key</td>
            <td></td>
        </tr>
        <tr>
            <td>HOYK</td>
            <td>Host Your Own Key</td>
            <td></td>
        </tr>
    </tbody>
</table>