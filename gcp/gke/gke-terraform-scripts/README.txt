README

This terraform script deploys the following resources (The script deploys resources in the 'europe-west2' a.k.a London DC region):

1) Custom VPC with 2 subnets (1 subnet for Bastion Host and other administrative machines, the other subnet for hosting the k8s nodes).  

2) Firewall rule to allow IAP (Identity Aware Proxy) for securely logging in to Bastion Host VM

3) Bastion Host for communicating with Kube API server

4) Security Hardened Private GKE Cluster with minimal permissions and privileges. This cluster will only be accessible via the Bastion host which is whitelisted to use it. (

Notes:

-> This script does not include the creation of service accounts nor roles
-> This script does not include the creation of Key Rings and Encryption Keys
-> This script does not include the creation and configuration of NAT Gateway, it is recommended to setup CloudNAT and configure it for the VPC where your cluster is hosted, else the private K8s nodes will not be able to access the internet.
-> Any access from pods/jobs to Google Cloud Services that are not part of the K8s nodes service account permissions need to be granted granular permissions via Workload Identity

Pre-requisites:

1) service account for terraform with the following roles:

-> Editor 

2) service account for bastion host with the following roles:

-> Monitoring Viewer
-> Monitoring Metric Writer
-> Logs Writer
-> Storage Object Viewer
-> Kubernetes Engine Developer

3) service account for k8s nodes with the following roles:

-> Monitoring Viewer
-> Monitoring Metric Writer
-> Logs Writer
-> Storage Object Viewer

4) encryption key in cloud KMS for encrypting ETCD 

Edits required before terraform apply:

1) root/variables.tf:

-> Line 4: Add path to credentials file
-> Line 16: Add project ID
-> Line 122: Add service acount ID of k8s-nodes
-> Line 123 (optional): Change node pool machine type
-> line 130: Add encryption key name for ETCD

2) modules/kubernetes/main.tf

-> Line 31 (optional): Change GKE master version
-> Line 57: Add key ID where encryption key for ETCD is contained (the full resource ID must be mentioned for this, not just the key name)
-> Set maintenance window and maintenance exclusions based on your time zone (I have intentionally left this out consdiering that different organizations have different peak traffic hours and varying time zones)

3) modules/Bastion_Host/main.tf

-> Line 30: Add service account ID of Bastion Host VM



