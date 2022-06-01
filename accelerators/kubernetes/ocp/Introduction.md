# OpenShift Compliant Financial Infrastructure Introduction

OpenShift (OCP) is an open hybrid cloud enterprise Kubernetes platform that can be installed on a number of cloud providers including Amazon Web Services (AWS), Azure and Google Cloud Platform (GCP).

This section provides an opinionated approach, documentation and working code to implement a set of security policies laid out in the [OpenShift Security Configuration (Service Accelerator)](accelerators/kubernetes/ocp/sat_rh_ocp.adoc).
Additional information on these policies can be found in [OpenShift Extended Configuration Patterns](accelerators/kubernetes/ocp/expanded-sec-details.adoc) which is an addendum to the Service Accelerator.

The initial focus of the project team is to implement the Service Accelerator policies on OpenShift 4.10 running on [Google Cloud Platform](accelerators/kubernetes/ocp/gcp), in the future this will be expanded to futher automation and include other cloud providers that OCP supports. 

For each cloud provider documentation and working code will be provided to:

1. Configure the cloud environment ready for an OCP installation.
2. Install OCP with any day 1 configuration changes to ensure compliance to the service accelerator.
3. Implement HTPasswd as an Identify Provider and remove KubeAdmin access.
4. Implement custom certifcate for the API server and Ingress Router
5. Complete day two configuration for remaining policies and implement the [OpenShift Compliance Operator](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-understanding.html#understanding-compliance-operator) which asseses compliance of the clusters kubernetes API and the nodes running the cluster. 

