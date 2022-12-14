# OpenShift Compliant Financial Infrastructure - Google Cloud

This folder provides documentation and working code to install and customise OpenShift(OCP) on Google Cloud(GC) to meet the policies set out in the [Red Hat OpenShift Compliant Financial Infrastructure Service Accelerator document](../ServiceApprovalAccelerator_OCP.md). 

There are six steps to complete cluster installation and customisation:

1. [Cloud provider setup and Cluster Installation](01_cluster_installation/cluster_installation.md)
2. [Identity provider configuration](02_htpasswd_identity_provider/htpasswd_implementation.md)
3. [Setup default network policies](03_default_network_policy/default_network_policy_implementation.md)
4. [Updating the self signed certificates for the API Server and Router](04_replace_api_router_certs/replace_api_router_certs.md)
5. [Implement OCP Compliance Operator](05_implement_ocp_compliance_operator/implement_ocp_compliance_operator.md)
6. [Manual Remediation of CIS Controls](06_remediation_of_manual_CIS_controls/Remediation_of_manual_CIS_controls.md)


