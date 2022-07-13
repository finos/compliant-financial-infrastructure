# OpenShift Compliant Financial Infrastructure - Google Cloud Platform

This folder provides documentation and working code to install and customise OpenShift(OCP) on Google Cloud Platform(GCP) to meet the policies set out in the [Red Hat OpenShift Compliant Financial Infrastructure Service Accelerator document](accelerators/kubernetes/ocp/sat_rh_ocp.adoc). 

There are six steps to complete cluster installation and customisation:

1. [Cloud provider setup and Cluster Installation](/accelerators/kubernetes/ocp/gcp/01_cluster_installation/cluster_installation.md)
2. [Identity provider configuration](/accelerators/kubernetes/ocp/gcp/02_htpasswd_identity_provider/htpasswd_implementation.md)
3. [Setup default network policies](/accelerators/kubernetes/ocp/gcp/03_default_network_policy/default_network_policy_implementation.md)
4. [Updating the self signed certificates for the API Server and Router](/accelerators/kubernetes/ocp/gcp/04_replace_api_router_certs/replace_api_router_certs.md)
5. [Implement OCP Compliance Operator](/accelerators/kubernetes/ocp/gcp/05_implement_ocp_compliance_operator/implement_ocp_compliance_operator.md)
6. [Manual Remiadation of CIS Controls](/accelerators/kubernetes/ocp/gcp/06_remediation_of_manual_CIS_controls/Remediation_of_manual_CIS_controls.md)


