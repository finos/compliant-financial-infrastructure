# OpenShift Compliant Financial Infrastructure

To complete the setup of the cluster to meet the policy requirements laid out in the [OpenShift Security Configuration (Service Accelerator)](accelerators/kubernetes/ocp/sat_rh_ocp.adoc) we will use the [OpenShift Compliance Operator](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-release-notes.html). The Compliance Operator lets OpenShift Container Platform administrators describe the required compliance state of a cluster and provides them with an overview of gaps and ways to remediate them. The compliance operator support a number of [compliance standards](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-supported-profiles.html for example NIST and CIS. 

To meet the policies as laid out in the service accelerator the OpenShift CIS Benchamrk profile will be used.

The compliance operator is a Customer Resource Definition (CRD) and includes a number of kubernetes objects, more details on these object can be found in [here](https://github.com/openshift/compliance-operator/blob/master/doc/crds.md). 

The Compliance Operator, or any other operator, can be installed on OCP using the Administrator UI or the CLI. Instructions on both methods can be found [here](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-installation.html). 

Below we will decribe both the installation and configuration of the FINOS profile using the command line. 

## Day 2 Customisation 

### Compliance Operator Installation and verification

1. Create the compliance operator namespace

`oc create -f 00-namespace-object.yaml`

2. Go into the openshift-compliance project

`oc project openshift-compliance`

3. Create the compliance operator OperatorGroup object

`oc create -f 01-operator-group-object.yaml`

4. Create the compliance operator subscription object, this will initiate the operator installation 

`oc create -f 02-subscription-object.yaml`

5. Verify that the installation was sucessful and that the operator is running

`oc get csv -n openshift-compliance -w`

*NAME                          DISPLAY               VERSION   REPLACES   PHASE
compliance-operator.v0.1.49   Compliance Operator   0.1.49               Pending

NAME                          DISPLAY               VERSION   REPLACES   PHASE
compliance-operator.v0.1.49   Compliance Operator   0.1.49               Installing

NAME                          DISPLAY               VERSION   REPLACES   PHASE
compliance-operator.v0.1.49   Compliance Operator   0.1.49               Succeeded*

`oc get deploy -n openshift-compliance`

*NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
compliance-operator   1/1     1            1           66s*


6. The compliance operator ships with a number of complaince profiles, these can be seen using the following command

`oc get profiles.compliance`

*NAME                 AGE
ocp4-cis             31s
ocp4-cis-node        32s
ocp4-e8              31s
ocp4-moderate        31s
ocp4-moderate-node   31s
ocp4-nerc-cip        31s
ocp4-nerc-cip-node   31s
ocp4-pci-dss         31s
ocp4-pci-dss-node    31s
rhcos4-e8            24s
rhcos4-moderate      23s
rhcos4-nerc-cip      23s*


7. The next step is to set up the compliance scanning, the default profiles that is being used will run scans on a daily base and, where possible. auto-remiate any complaince issues.

`oc create -f 03-scansettingbinding-cis-default.yaml`

8. The creation of the scan object will trigger the initial complaince scan to take please, to check that the compliance scan is running use the following command: 

`oc get compliancesuite -w`

*NAME                  PHASE       RESULT
cis-compliance-ocp4   LAUNCHING   NOT-AVAILABLE
cis-compliance-ocp4   LAUNCHING   NOT-AVAILABLE
cis-compliance-ocp4   LAUNCHING   NOT-AVAILABLE
cis-compliance-ocp4   RUNNING     NOT-AVAILABLE
cis-compliance-ocp4   RUNNING     NOT-AVAILABLE
cis-compliance-ocp4   RUNNING     NOT-AVAILABLE
cis-compliance-ocp4   AGGREGATING   NOT-AVAILABLE
cis-compliance-ocp4   AGGREGATING   NOT-AVAILABLE
cis-compliance-ocp4   AGGREGATING   NOT-AVAILABLE
cis-compliance-ocp4   DONE          NON-COMPLIANT
cis-compliance-ocp4   DONE          NON-COMPLIANT*

9. Once the scan is complete the following command can be used to check the complaince scan results so that an compliance scan failures can be remediated

`oc get ccr`

Initially a number of the compliance checks will FAIL, many of them will be remediated automatically by the compliance operator. This remediation process can take sometime (hours). 

The OCP CIS Benchmark consists of three sets of policies:

- *ocp-cis* for Cluster polies
- *ocp-node-master* for control node policies
- *ocp-node-worker* for compute node policies

The following commmands can be used to do compliance checks for each of the above policies. The examples below look for compliance checks that have *FAILED*, to check for policies that have passed use *PASS* or policies that require manual evaluation use *MANUAL*.

`oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis,compliance.openshift.io/check-status=FAIL`

`oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis-node-master,compliance.openshift.io/check-status=FAIL`

`oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis-node-worker,compliance.openshift.io/check-status=FAIL`

Once any remediation have be made the following commands can be used to trigger a compliance rescan. A [script](rescan.sh) to trigger a rescan of all scans has been provided. 

`oc annotate compliancescans/ocp4-cis compliance.openshift.io/rescan=`

`oc annotate compliancescans/ocp4-cis-node-master compliance.openshift.io/rescan=`

`oc annotate compliancescans/ocp4-cis-node-worker compliance.openshift.io/rescan=`

The compliance operator will auto-remiate all CIS policies with the exception of those policies that the CIS define as requirement manual rediation. In the next section we will address these [manual remdiations](accelerators/kubernetes/ocp/gcp/04_remediation_of_manual_CIS_controls).
