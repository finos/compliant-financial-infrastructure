# OpenShift Compliant Financial Infrastructure Introduction

OpenShift (OCP) is an open hybrid cloud enterprise Kubernetes platform that can be installed on a number of cloud providers including Amazon Web Services (AWS), Azure and Google Cloud Platform (GCP).

This section provides an opinionated approach, documentation and working code to implement a set of security policies laid out in the [OpenShift Security Configuration (Service Accelerator)](accelerators/kubernetes/ocp/sat_rh_ocp.adoc).
Additional information on these policies can be found in [OpenShift Extended Configuration Patterns](accelerators/kubernetes/ocp/expanded-sec-details.adoc) which is an addendum to the Service Accelerator.

The initial focus of the project team is to implement the Service Accelerator policies on OpenShift 4.10 running on [Google Cloud Platform](accelerators/kubernetes/ocp/gcp), in the future this will be expanded to futher automation and include other cloud providers that OCP supports. 

For each cloud provider documentation and working code will be provided to:

1. Cloud provider setup and Cluster Installation
2. Identity provider configuration
3. Setup default network policies
4. Updating the self signed certificates for the API Server and Router
5. Implement OCP Compliance Operator
6. Manual Remediation of CIS Controls

