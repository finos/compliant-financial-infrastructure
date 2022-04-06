# OpenShift Compliant Financial Infrastructure Introduction

OpenShift is a open hybrid cloud enterprise Kubernetes platform that can be installed on a number of cloud providers including AWS, Azure and GCP. 
This section provides an opinionated approach, docuemntation and working code to implement a set of security policies laid out in the [OpenShift Security Configuration (Service Accelerator)](accelerators/kubernetes/ocp/sat_rh_ocp.adoc).
Additional information can be found in [OpenShift Extended Configuration Patterns](accelerators/kubernetes/ocp/expanded-sec-details.adoc) which is an addendum to the Service Accelerator.

Initial focus will be on implementing the Service Accelerator policies on OpenShift 4.10 running on [Google Cloud Platform](accelerators/kubernetes/ocp/gcp), in the future this will be expanded to cover other cloud providers that OpenShift supports. 

For each cloud provider docuemntation and working code is proveded to:

1. Configure the cloud environment ready for OpenShift installation.
2. Install an OpenShift cluster with any day 1 configuration changes to ensure complaince to the service accelerator.
3. Implement HTPasswd as an Identify Provider.
3. Day two configuration for remaining policies and in addition monitoring of policies and where possible automatic remediation of non compliance.

