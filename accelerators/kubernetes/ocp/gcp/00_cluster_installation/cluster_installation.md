# OpenShift Compliant Financial Infrastructure

## GCP Project setup and Cluster Installation 

The OpenShift (OCP) Installer supports two installation methods, Installer Provisioned Infrastructure(IPI) and User Provisioned Infrastructure(UPI). IPI is an opinionated automated installation, this is the installation menthod that will be used. UPI gives users more flexibility to install OCP on pre-provisioned infrastructure, for example an on-premises installation where a firm's IT standards and policies prevent the use of an opinionated and automated installation. More details on OpenShift Installation can be found [here](https://docs.openshift.com/container-platform/4.10/installing/index.html).

It is possile for OCP to be installed into a disconnected or air-gapped environment or be configured to have no public endpoints. To meet the current service accelerator requirements this is not required, the following instructions will implement a cluster that is internet connected and has public end-points. 

The following provides an overview of the steps needed to install OCP on GCP, to meet the requirements of the service accelerator customisation needs to be made both at install time and as a day two change. To make the changes at install time we will use the [Installing a Cluster on GCP with Customisations] (https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html) installation method.

Following are the high level steps to complete OCP installation, including the service accelerator polices to implement [FIPS cryptography](https://docs.openshift.com/container-platform/4.10/installing/installing-fips.html) and [OVNKubernetes Container Network Interface (CNI) plugin](https://docs.openshift.com/container-platform/4.10/networking/ovn_kubernetes_network_provider/about-ovn-kubernetes.html#about-ovn-kubernetes). It is planned that in the future code to automate these steps will be provided. 

1. Setup the [GCP project](https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-account.html) ready for the OCP installation. This includes:
- Creating the GCP project Folder
- Enabling the API's that the OCP Installer requires
- Create a DNS public zone
- Increasing GCP quotas (if needed)
- Creating a GCP service account for the OCP Installer and give it required permissions
- [download installer and service account key](https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html) onto local machine or a bastion where the installer is run and the installer state files can be stored.   

3. Create the [installation configuration file](https://docs.openshift.com/container-platform/4.10/installing/installing_gcp/installing-gcp-customizations.html#installation-initializing_installing-gcp-customizations)
- Create an OCP installation directory
- Run the OCP installer to create the *install-config* file that will be used to implement install time customisations. Here is an example command: './openshift-install create cluster-config --dir=/Users/adrianhammond/ocp_clusters/finosocp'
- Edit the install-config YAML file to implement FIPS and OVNKubernets, here is an example [config-yaml](sample-install-config.yaml)
- Run the OCP installer to create the cluster. Here is an example command: './openshift-install create cluster --dir=/Users/adrianhammond/ocp_clusters/finosocp --log-level debug'

Once the cluster installation completes, which normally takes 30-40 minutes, the installer output includes the OCP console url, kubeadmin userid and password and KUBECONFIG.

To confirm that FIPS and OVNKubernetes changes have been implemeted log onto the cluster and run the following commands.

1. ## NEED TO ADD A CHECK FOR FIPS
2. Check OVNKubernetes by using the follwing command 'oc describe network.config/cluster'. Check the output for 'Network Type:  OVNKubernetes'


## The next step of the installation process is to set up an identity provider. An example using [HTPASSWD] (htpassed-identity-provider) has been provided, OCP supports a number of identity providers, more detail can be found [here](https://docs.openshift.com/container-platform/4.10/authentication/understanding-identity-provider.html)