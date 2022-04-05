# OpenShift Compliant Financial Infrastructure

## Cluster Installation Instructions

An OpenShift Installer supports two installtion methods, Installer Provisioned Infrastructure (IPI) and User Provisioned Infrastructure (UPI). IPI is an opinionated automated installation, which we will be using, and UPI which gives use more flexibility to install OpenShift on pre-provisioned infrastructire. More details on OpenShift Installation can be found [here] (https://docs.openshift.com/container-platform/4.10/installing/index.html).

It is possile for an OpenShift cluster to be installed into a disconnected or air-gapped enviormnt with no public endpoints. To meet these current service accelerator policies this is not required, the cluster being installed will be internet connected and have public end-points. 

The following provides an overview of the steps needed to install OpenShift on GCP. To meet the requirements of the service accelerator customisation need to be made both at install time and as a day two change. To make the changes at install time we will use the [Installing a Cluster on GCP with Customisations] (https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html) method.

Following are the high level steps to complete cluster installation, including the service acceprator polices to implement FIPS and OVNKubernetes. The contributors to this project plan to provide automations for the following steps in the future.

1. Setup the [GCP project] (https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-account.html) ready for installation of the OpenShiift Cluster. This includes:
- Creating the GCP project Folder
- Enabling the API's that the OpenShift Installer requires
- Create DNS public zone
- Increasing GCP quoatas (if needed)
- Create a GCPP service account for the OpenShift Installer and give it required permissions

2. Complete installation setup, [download installer and service account key] (https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html) 

3. Create the [installation configuration file] (https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html#installation-initializing_installing-gcp-customizations)
- Create OCP installation directory
- Run OpenShift installer to create *install-config* file that will be used to implement install time customisations
for example: './openshift-install create cluster-config --dir=/Users/adrianhammond/ocp_clusters/finosocp'
- Edit the install-config YAML file to implement FIPS and OVNKubernets, here is an example [config-yaml] (sample-install-config.yaml)
- Run OpenShift installer to create the cluster
for example: './openshift-install create cluster --dir=/Users/adrianhammond/ocp_clusters/finosocp --log-level debug'

Once the cluster installation completes, normally takes 30-40 minutes, the installer output includes the OCP consoe url, kubeadmin password and KUBECONFIG

To confirm that FIPS and OVNKubernetes changes have been implemeted log onto the cluster and run the following commands.

1. Check OVNKubernetes by using the follwing command 'oc describe network.config/cluster'. Check the output for 'Network Type:  OVNKubernetes'
2. ## NEED TO ADD A CHECK FOR FIPS


## Next step is to set up an identity provider, as an example we are using [HTPASSWD] (htpassed-identity-provider)