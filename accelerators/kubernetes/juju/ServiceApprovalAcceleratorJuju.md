Date: 21-Feb-2022

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
    <td class="tg-7btt" rowspan="4">Identity &amp; Access Management</td>
    <td class="tg-fymr">Authentication</td>
    <td class="tg-f8tv">By default Juju supports username and password authentication, however it can be integrated with external identity providers that support the SAML, OpenID Connect and LDAP protocols through Candid identity broker.
All actions described in the sections below can only be performed by authenticated users.

Adding credentials
In order for Juju to be able to create and manipulate resources on the underlying cloud, it needs to be provided with a credential that has an appropriate set of permissions associated with it. Depending on the cloud substrate, the credential will take the form of a username/password pair, a secret key, or a certificate.

Upon adding a new credential to Juju, users are asked to provide an alias name which is subsequently used to refer to the stored credential. Credentials may be added to the local client, the currently active controller (if any), or both.

Credentials become active as soon as they are related to a Juju model. Typically, this is facilitated via the bootstrap and add-model commands, as both trigger the creation of a new model.

Juju supports three methods for adding credentials:

- Manually providing credentials via an interactive session with the command line client.
- Auto-detecting credentials by scanning environment variables and/or “rc” files (only supported for certain providers).
- Importing credentials from a user-provided, YAML-formatted file.

The Juju client stores any added credentials into `$HOME/.local/share/juju/credentials.yaml`.

The table below shows the authentication type available for each cloud type. It does not include the interactive type as it does not apply in the context of adding a cloud manually.

| `cloud type` | `authentication type`|
|--|--|
| azure| `service-principal-secret`|
| cloudsigma| `userpass`|
| ec2| `access-key`|
| gce| `jsonfile,oauth2`|
| lxd| n/a, `certificate` (`v.2.5`) |
| maas| `oauth1`|
| manual| n/a|
| oci| `httpsig`|
| openstack| `access-key,userpass`|
| oracle| `userpass`|
| rackspace| `userpass`|
| vsphere| `userpass`|
</td>
    <td class="tg-0pky">https://juju.is/docs/olm/adding-clouds
    https://juju.is/docs/olm/credentials</td>
  </tr>
  <tr>
    <td class="tg-fymr">Authorization</td>
    <td class="tg-f8tv">Juju has an internal user framework that allows for the sharing of controllers and models. To achieve this, a Juju user can be created, disabled, and have rights granted and revoked. Users remote to the system that created a controller can use their own Juju client to log in to the controller and manage the environment based on the rights conferred. Multiple users can be accommodated by the same Juju client.

Various categories of users can be defined based on the permissions they have been granted. In turn, these permissions lead to certain abilities.
</td>
    <td class="tg-0pky">https://juju.is/docs/olm/working-with-multiple-users
    https://juju.is/docs/olm/user-types-and-abilities</td>
  </tr>
  <tr>
    <td class="tg-fymr">RBAC</td>
    <td class="tg-f8tv">Juju implements a multi tiered role based access control model. Access rights can be provided at the model and controller level.

**Controller access**

Controller access rights are governed by the following roles:
- login: standard access level, enabling a user to log in to a controller.
- add-model: login permission with the additional possibility of adding and removing models.
- superuser: controller administrator role.

**Model access**

Model access rights are governed by the following roles which can be granted by controller superusers:
- read: A user can view the state of a model (e.g. models, machines, and status)
- write: In addition to ‘read’ abilities, a user can modify/configure models (e.g. model-config and model-defaults).
- admin: model ownership role that allows model upgrades SSH connection to machines.</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Privileged Access Management</td>
    <td class="tg-f8tv">Secrets belonging to personal and technical accounts are protected by Juju and access to privilege entitlements is governed by the RBAC model described above, however privileged access management capabilities are outside of the application remit.
SSH connection to machines is allowed only to users who have model admin rights.
</span></td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="4">Encryption &amp; Secure Data Management</td>
    <td class="tg-fymr">Encryption in Transit</td>
    <td class="tg-f8tv">All Juju connections are TLS encrypted between all agents.</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Encryption at Rest</td>
    <td class="tg-f8tv">The Juju controller can be bootstrapped in a machine with full disk encryption.</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">Certificate and Key Management</td>
    <td class="tg-f8tv">When a controller is either created or registered a passphraseless SSH keypair will be generated and placed under ~/.local/share/juju/ssh The public key juju_id_rsa.pub, as well as a possibly existing ~/.ssh/id_rsa.pub, will be placed within any newly-created model.

This means that a model creator will always be able to connect to any machine within that model without having to add keys since the creator is also granted ‘admin’ model access by default
</span></td>
    <td class="tg-0pky">https://juju.is/docs/olm/accessing-individual-machines-with-ssh</td>
  </tr>
  <tr>
    <td class="tg-fymr">BYOK/HYOK Management</td>
    <td class="tg-f8tv">When adding a cloud, Juju supports the import of a wide series of credentials, including secret keys and certificates. This can be done either by using an interactive command line method or by supplying a YAML file.
Once credentials are added Juju provides a way to list, modify and delete them. All the activities can only be performed by a user with administrative privileges.</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="2">Network Security </td>
    <td class="tg-fymr">Endpoint Localisation</td>
    <td class="tg-f8tv">Endpoint localisation is referred to by where a machine is located from a cloud, host, machine, host  and port perspectives. It is tied to the "Link-Layer Device Address" entry in the linklayerdevices collection. Juju is also aware of public cloud specific location identifiers like availability zones and VPC.

Juju discovers an endpoint network connectivity via the following mechanisms:
The machine agent collects and stores link-layer device information and IP addresses sourced directly from machines/VMs.
The instance-poller discovers network configuration for machines/VMs by asking the providers via their API.
The Kubernetes provider discovers container/pod/service addresses via the Kubernetes API.

The discovered data resides in the following Mongo DB collections:
linklayerdevices and ip.addresses for configuration discovered by the machine agent, and where it can be related to those devices, the instance-poller.
machines for IP addresses discovered by the instance-poller that may not be specifically associated with known link-layer devices.
cloudservices and cloudcontainers for IP addresses gleaned from Kubernetes.

Charm endpoints are bound to network spaces. Charms deployed to machines/VMs can query Juju about the specifics of their connectivity, optionally within a relation context to negotiate workload connectivity.

In general terms, this is done as follows:
The link-layer devices are queried for the machine where the charm is deployed.
Those device IP addresses are matched to their subnets.
The subnets are matched to their spaces.
Based on the bound space for an endpoint, the link-layer device(s) and IP address information are resolved, including subnets for ingress and egress.</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-fymr">IP Firewall Rules</td>
    <td class="tg-f8tv">Firewall rules control ingress to a well known service within a Juju model. A rule consists of the service name and a whitelist of allowed ingress subnets.

Modes available include:
- instance: Requests the use of an individual firewall per instance.
- global: Uses a single firewall for all instances (access for a network port is enabled to one instance if any instance requires that port).
- none: Requests that no firewalling should be performed inside the model, which is useful for clouds without support for either global or per instance security groups.</td>
    <td class="tg-0pky">https://juju.is/docs/olm/configure-a-model#heading--firewall-mode
    https://discourse.charmhub.io/t/command-set-firewall-rule/1811</td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="2">Logging &amp; Monitoring</td> 
    <td class="tg-fymr">Logging</td>
    <td class="tg-f8tv">There are three logging resources available to the Juju operator:

**Model logs**
Model logs can be considered as Juju’s “regular logs” and are intended to be inspected with the debug-log command. This method provides logs on a per-model basis and is therefore more convenient than reading individual logs on multiple (Juju) machines directly on the file system. The latter can nonetheless be done in exceptional circumstances and some explanation is provided here.

**Remote logging**

On a per-model basis log messages can optionally be forwarded to a remote syslog server over a secure TLS connection.

**Audit logging**

Juju audit logging provides a chronological account of all events by capturing invoked user commands. These logs reside on the controller involved in the transmission of commands affecting the Juju client, Juju machines, and the controller itself.

The audit log filename is /var/log/juju/audit.log and contains records which are either:

- a Conversation, a collection of API methods associated with a single top-level CLI command
- a Request , a single API method
- a ResponseErrors, errors resulting from an API method
- Information can be filtered out of the audit log to prevent its file(s) from growing without bounds and making it difficult to read.</td>
    <td class="tg-0pky">https://juju.is/docs/olm/juju-logs</td>
  </tr>
  <tr>
    <td class="tg-fymr">Service Monitoring & Alerting</td>
    <td class="tg-f8tv">The Canonical Observability Stack (COS) is the observability and monitoring stack for Juju. COS is made of the following Juju charmed operators:
    - Prometheus charmed operator
    - Alertmanager charmed operator
    - Loki charmed operator
    - Grafana charmed operator

The charmed operators that make up COS are available as the pre-configured, batteries-included COS Lite bundle deployable with `juju deploy cos-lite --channel=edge --trust`.</td>
    <td class="tg-0pky">https://charmhub.io/cos-lite
https://charmhub.io/grafana-k8s
https://charmhub.io/loki-k8s
https://charmhub.io/prometheus-k8s
https://charmhub.io/alertmanager-k8s
https://ubuntu.com/observability</td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="2">Resilience &amp; Recovery</td>
    <td class="tg-fymr">Data Resilience (back-up/replication)</td>
    <td class="tg-f8tv">Scaling Applications
The capability of a service to adjust its resource footprint to a level appropriate for fulfilling client demands placed upon it is known as scalability. Scaling vertically affects the resources of existing machines (memory, CPU, disk space) whereas scaling horizontally, in Juju, involves the number of application units available.

Units are not always synonymous with machines however. Multiple units can be placed onto a single machine (co-location) and still be considered horizontal scaling if sufficient resources are present on the machine.

To scale up while in a Kubernetes model the total number of desired units for the application is simply stated. Here we want a total of three units:

`juju scale-application <application_name> <number_of_units>`

Back up the controller
A backup of a controller enables one to save (and later re-establish) the configuration and state of a controller. It does not influence workload instances on the backing cloud. That is, if such an instance is terminated directly in the cloud then a controller restore cannot recreate it. The current state is held within the ‘controller’ model. Therefore, all backup commands need to operate within that model explicitly or by ensuring the current model is the ‘controller’ model.

Creating a backup
The `create-backup` command is used to create a backup. It does so by generating an archive and downloading it to the Juju client system as a ‘tar.gz’ file (a local backup). If the `--keep-copy` option is used then a copy of the archive will also remain on the controller (a remote backup). With the aid of the `--no-download` option a local backup can be prevented, but since the archive must be kept somewhere, this option implies --keep-copy.

The name of the backup is composed of the creation time (in UTC) and a unique identifier.

To create a backup of <controller_name> (and keep a copy on the controller):

``` bash
juju create-backup -m <controller_name> :controller --keep-copy
```
</td>
    <td class="tg-0pky">https://juju.is/docs/scaling-applications
    https://juju.is/docs/olm/controller-backups</td>
  </tr>
  <tr>
    <td class="tg-fymr">Compute High Availability</td>
    <td class="tg-0pky">To ensure the high availability (HA) of deployed applications, the Juju controller must itself be highly available. This necessitates the creation of additional controllers, all of which naturally reside within the ‘controller’ model. The initial controller becomes known as the master and automatic failover occurs should it lose connectivity with its cluster peers.

Controller HA is managed with the `juju enable-ha` command. It does this by ensuring that the cluster has the requisite number of controllers present. By default, this number is three but the `-n` switch can be used to change that. Therefore, this command is used to both enable HA as well as compensate for any missing controllers, as is the case if you enable HA and then remove one or more controllers.

When a controller is provisioned, API server code is installed along with a MongoDB database.

The number of controllers must be an odd number in order for a master to be “voted in” amongst its peers. A cluster with an even number of members will cause a random member to become inactive. This latter system will become a “hot standby” and automatically become active should some other member fail. Furthermore, due to limitations of the underlying database in an HA context, that number cannot exceed seven. All this means that a cluster can only have three, five, or seven active members.

Juju clients and agents talk to any of the controllers in the cluster. This means that processing at the controller (API) level is distributed. However, there is only one primary database at any given time and all controllers write to it. The “master”, therefore, actually refers to the underlying database.
</td>
    <td class="tg-0pky">https://juju.is/docs/olm/high-availability-juju-controller</td>
  </tr>
  <tr>
    <td class="tg-7btt" rowspan="2">Underlying OS</td>
    <td class="tg-fymr">Use of Latest Version</td>
    <td class="tg-f8tv"><span style="font-weight:400">Generally, charm developers set an automatic release process so when a new version of the image is released, a new charm is built.</span></td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-1wig">-</td>
    <td class="tg-8zwo">The preferred way to run workloads on Kubernetes with charms is to start the workload with Pebble. The base OS of the image should be Linux. There is no need to modify upstream container images to make use of Pebble for managing workloads. The Juju controller automatically injects Pebble into workload containers using an Init Container and Volume Mount. The entrypoint of the container is overridden so that Pebble starts first and is able to manage running services. Charms communicate with the Pebble API using a UNIX socket, which is mounted into both the charm and workload containers.

By default, the Pebble socket is at /var/lib/pebble/default/pebble.sock in the workload container, and /charm/<container>/pebble.sock in the charm container.

Most Kubernetes charms will need to define a containers map in their metadata.yaml in order to start a workload with a known OCI image:

``` bash
# ...
containers:
  myapp:
    resource: myapp-image
  redis:
    resource: redis-image

resources:
  myapp-image:
    type: oci-image
    description: OCI image for my application
  redis-image:
    type: oci-image
    description: OCI image for Redis
# ...
```

In some cases, you may wish not to specify a container map, which will result in an “operator-only” charm. These can be useful when writing “integrator charms” (sometimes known as “proxy charms”), which are used to represent some external service in the Juju model.

For each container, a resource of type oci-image must also be specified. The resource is used to inform the Juju controller how to find the correct OCI-compliant container image for your workload on Charmhub.
</td>
    <td class="tg-0lax">https://juju.is/docs/sdk/workloads#heading--kubernetes-charms</td>
  </tr>
</tbody>
</table>
