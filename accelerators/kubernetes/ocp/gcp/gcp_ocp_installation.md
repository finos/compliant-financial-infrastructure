# OpenShift Compliant Financial Infrastructure - Google Cloud Platform

This folder provides documentation and working code to install and customise OpenShift(OCP) on Google Cloud Platform(GCP) to meet the policies set out in the [Red Hat OpenShift Compliant Financial Infrastructure Service Accelerator document](accelerators/kubernetes/ocp/sat_rh_ocp.adoc). 

There are four steps to complete installation and customisation:

1. [Cloud provider setup and Cluster Installation](accelerators/kubernetes/ocp/gcp/cluster_installation)
2. [Identity provider configuration](accelerators/kubernetes/ocp/gcp/htpasswd-identity-provider)
3. [Updating the self signed certificates for the API Server and Router](accelerators/kubernetes/ocp/gcp/02_replace_api_router_certs)
4. [Day Two Customisation](accelerators/kubernetes/ocp/gcp/day2_customisation)