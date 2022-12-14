# OpenShift Compliant Financial Infrastructure Introduction

OpenShift (OCP) is an open hybrid cloud enterprise Kubernetes platform that can be installed on a number of cloud providers including Amazon Web Services (AWS), Azure and Google Cloud Platform (GCP).

This section provides an opinionated approach, documentation and working code to implement CFI security policies laid out in the [OpenShift Security Configuration (Service Accelerator)](ServiceApprovalAccelerator_OCP.md).


The initial focus of the project team is to implement the Service Accelerator policies on OpenShift 4.11 running on [Google Cloud Platform](./gcp/), in the future this will be expanded to futher automation and include other cloud providers that OCP supports. 

For each cloud provider documentation and working code will be provided to:

1. Cloud provider setup and Cluster Installation
2. Identity provider configuration
3. Setup default network policies
4. Updating the self signed certificates for the API Server and Router
5. Implement OCP Compliance Operator
6. Manual Remediation of CIS Controls

[] Todo: Need to add a section here that covers requirement to have a secure and compliant bastion to install OCP from, maybe as Juo to write??
