# OpenShift Compliant Financial Infrastructure

To complete the setup of the cluster to meet the policy requirements laid out in the [OpenShift Security Configuration (Service Accelerator)](accelerators/kubernetes/ocp/sat_rh_ocp.adoc) we will use the [OpenShift Compliance Operator](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-release-notes.html). The Compliance Operator lets OpenShift Container Platform administrators describe the required compliance state of a cluster and provides them with an overview of gaps and ways to remediate them. The compliance operator support a number of [compliance standards](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-supported-profiles.html for example NIST and CIS. In support of the FINOS project Red Hat is working on a FINOS profile for the compliance operator, an upstream version of this profile is being used for this project. 

The compliance operator is a Customer Resource Definition (CRD) and includes a number of kubernetes objects, more details on these object can be found in [here](https://github.com/openshift/compliance-operator/blob/master/doc/crds.md). 

The Compliance Operator, or any other operator, can be installed on OCP using the Administrator UI or the CLI. Instructions on both methods can be found [here](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-installation.html). 

Below we will decribe both the installation and configuration of the FINOS profile using the command line. 

## Day 2 Customisation 

### Compliance Operator Installation and verification

1. Create the compliance operator namespace

`oc create -f 00-namespace-object.yaml`

2. Go into the openshift-compliance project

`oc project openshift-compliance`

2. Create the compliance operator OperatorGroup object

`oc create -f 01-operator-group-object.yaml`

3. Create the compliance operator subscription object, this will initiate the operator installation 

`oc create -f 02-subscription-object.yaml`

4. Verify that the installation was sucessful and that the operator is running

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

### Add the FINOS upstream profile

The compliance operator ships with a number of complaince profiles, these can be seen using the following command

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

Red Hat is investigation adding a Finos profile to the compliance operator, until this work is completed we will install an upstream version of the Finos profile. 

1. The first step is to create the compliance bundles

`oc create -f 03-finos-ocp4-profilebundle.yaml`

`oc create -f 04-finos-rhcos4-profilebundle.yaml`

To check that the Finos profiles are present use the following command

`oc get profiles.compliance`

*NAME                       AGE
ocp4-cis                   90s
ocp4-cis-node              91s
ocp4-e8                    90s
ocp4-finos-cis             1s
ocp4-finos-cis-node        1s
ocp4-finos-e8              1s
ocp4-finos-finos           1s
ocp4-finos-finos-node      1s
ocp4-finos-high            0s
ocp4-finos-high-node       0s
ocp4-finos-moderate        0s
ocp4-finos-moderate-node   0s
ocp4-moderate              90s
ocp4-moderate-node         90s
ocp4-nerc-cip              90s
ocp4-nerc-cip-node         90s
ocp4-pci-dss               90s
ocp4-pci-dss-node          90s
rhcos4-e8                  83s
rhcos4-moderate            82s
rhcos4-nerc-cip            82s*

2. The next step is to set up the compliance scanning, the profile that is being used will auto-remiate any complaince issues where possible

`oc create -f 05-scansettingbinding-finos-default.yaml`

`oc create -f 06-scansettingbinding-finos-default.yaml`

3. We can now check that the compliance scan is running 

`oc get compliancesuite -w`

*NAME                      PHASE     RESULT
finos-compliance-ocp4     DONE      NON-COMPLIANT
finos-compliance-rhcos4   RUNNING   NOT-AVAILABLE
finos-compliance-rhcos4   LAUNCHING   NOT-AVAILABLE
finos-compliance-rhcos4   RUNNING     NOT-AVAILABLE
finos-compliance-rhcos4   AGGREGATING   NOT-AVAILABLE
finos-compliance-rhcos4   DONE          COMPLIANT*

4. Once the scan is complete the following command can be used to check the complaince scan results so that an compliance scan failures can be remediated

`oc get ccr`

To just list the compliance checks that failed the following command can be used, in this example we are looking at the *rhcos-finos-finos-master' scan is being reviewed

`oc get compliancecheckresults -lcompliance.openshift.io/scan-name=rhcos4-finos-finos-master,compliance.openshift.io/check-status=FAIL`

Once any remediation have be made the following command can be used to rescan a scan, in this example we are again looking at the *rhcos4-finos-finos-master* scan. A [script](rescan.sh) to trigger a rescan of all scans has been provided. 

`oc annotate compliancescans/ocp4-finos-finos-node-master compliance.openshift.io/rescan=`



