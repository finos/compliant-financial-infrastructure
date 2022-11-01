# OpenShift Compliant Financial Infrastructure

To complete the setup of the cluster to meet the policy requirements laid out in the [OpenShift Security Configuration (Service Accelerator)](/accelerators/kubernetes/ocp/sat_rh_ocp.adoc) we will use the [OpenShift Compliance Operator](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-release-notes.html). The Compliance Operator lets OpenShift Container Platform administrators describe the required compliance state of a cluster and it provides them with an overview of gaps and ways to remediate them. The compliance operator supports a number of [compliance standards](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-supported-profiles.html) for example NIST and CIS. 

To meet the policies as laid out in the service accelerator the OpenShift CIS Benchamrk profile will be used.

The compliance operator is a Customer Resource Definition (CRD) and includes a number of kubernetes objects, more details on these object can be found in [here](https://github.com/openshift/compliance-operator/blob/master/doc/crds.md). 

The Compliance Operator, or any other operator, can be installed on OCP using the Administrator UI or the CLI. Instructions on both methods can be found [here](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-installation.html). 

Below we will decribe both the installation and configuration of the CIS profile using the command line. 

## Day 2 Customisation 

### Compliance Operator Installation and verification

1. Create the compliance operator namespace

```shell
oc create -f 00-namespace-object.yaml
```

2. Go into the openshift-compliance project

```shell
oc project openshift-compliance
```

3. Create the compliance operator OperatorGroup object

```shell
oc create -f 01-operator-group-object.yaml
```

4. Create the compliance operator subscription object, this will initiate the operator installation 

```shell
oc create -f 02-subscription-object.yaml
```

5. Verify that the installation was sucessful and that the operator is running and look at the compliance operator deployment. 

```shell
oc get csv -n openshift-compliance -w
```

```console
NAME                          DISPLAY               VERSION   REPLACES   PHASE
compliance-operator.v0.1.52   Compliance Operator   0.1.52               
compliance-operator.v0.1.52   Compliance Operator   0.1.52               
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Pending
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Pending
compliance-operator.v0.1.52   Compliance Operator   0.1.52               InstallReady
compliance-operator.v0.1.52   Compliance Operator   0.1.52               InstallReady
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Installing
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Installing
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Succeeded
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Failed
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Pending
compliance-operator.v0.1.52   Compliance Operator   0.1.52               InstallReady
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Installing
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Installing
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Installing
compliance-operator.v0.1.52   Compliance Operator   0.1.52               Succeeded
```

```shell
oc get deploy -n openshift-compliance
```

```console
NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
compliance-operator              1/1     1            1           2m12s
ocp4-openshift-compliance-pp     0/1     1            0           45s
rhcos4-openshift-compliance-pp   0/1     1            0           45s
```

6. The compliance operator ships with a number of complaince profiles, these can be seen using the following command

```shell
oc get profiles.compliance
```

```console
NAME                 AGE
ocp4-cis             63s
ocp4-cis-node        63s
ocp4-e8              63s
ocp4-high            63s
ocp4-high-node       63s
ocp4-moderate        63s
ocp4-moderate-node   63s
ocp4-nerc-cip        62s
ocp4-nerc-cip-node   62s
ocp4-pci-dss         62s
ocp4-pci-dss-node    62s
rhcos4-e8            58s
rhcos4-high          57s
rhcos4-moderate      57s
rhcos4-nerc-cip      57s
```

7. The next step is to set up the compliance scanning using the CIS profile, the default profiles that is being used will run scans on a daily base and, where possible. auto-remiate any complaince issues.

```shell
oc create -f 03-scansetting-cis-ocp4-default.yaml
```

8. The creation of the scan object will trigger the initial compliance scan to take please against the CIS compliance profile, to check that the compliance scan is running use the following command: 

```shell
oc get compliancesuite -w
```

```console
NAME                  PHASE       RESULT
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
cis-compliance-ocp4   DONE          NON-COMPLIANT
```

9. Once the scan is complete the following command can be used to check the complaince scan results, when setting up the compliance scan object (step 7), we set the complaince operator to automatical remidate compliance failures. 

```shell
oc get ccr
```

An extract of the oc get ccr command can be seen below.


```console
NAME                                                                           STATUS   SEVERITY
ocp4-cis-accounts-restrict-service-account-tokens                              MANUAL   medium
ocp4-cis-accounts-unique-service-account                                       MANUAL   medium
ocp4-cis-api-server-admission-control-plugin-alwaysadmit                       PASS     medium
ocp4-cis-api-server-admission-control-plugin-alwayspullimages                  PASS     high
ocp4-cis-api-server-admission-control-plugin-namespacelifecycle                PASS     medium
ocp4-cis-api-server-admission-control-plugin-noderestriction                   PASS     medium
ocp4-cis-api-server-admission-control-plugin-scc                               PASS     medium
ocp4-cis-api-server-admission-control-plugin-securitycontextdeny               PASS     medium
ocp4-cis-api-server-admission-control-plugin-serviceaccount                    PASS     medium
ocp4-cis-api-server-anonymous-auth                                             PASS     medium
ocp4-cis-api-server-api-priority-gate-enabled                                  PASS     medium
ocp4-cis-api-server-audit-log-maxbackup                                        PASS     low
ocp4-cis-api-server-audit-log-maxsize                                          PASS     medium
ocp4-cis-api-server-audit-log-path                                             PASS     high
ocp4-cis-api-server-auth-mode-no-aa                                            PASS     medium
ocp4-cis-api-server-auth-mode-node                                             PASS     medium
ocp4-cis-api-server-auth-mode-rbac                                             PASS     medium
ocp4-cis-api-server-basic-auth                                                 PASS     medium
ocp4-cis-api-server-bind-address                                               PASS     low
ocp4-cis-api-server-client-ca                                                  PASS     medium
ocp4-cis-api-server-encryption-provider-cipher                                 FAIL     medium
ocp4-cis-api-server-encryption-provider-config                                 FAIL     medium
```

Initially a number of the compliance checks will FAIL, many of them will be remediated automatically by the compliance operator. This remediation process can take sometime (hours). These remediation a actioned by the OCP cluster operators, using the following command the progress of these cluster operators can be seen.

'''shell
oc get co
```
Below is an example of the output of this command.

```console
NAME                                       VERSION   AVAILABLE   PROGRESSING   DEGRADED   SINCE   MESSAGE
authentication                             4.10.3    True        True          False      76m     APIServerDeploymentProgressing: deployment/apiserver.openshift-oauth-apiserver: 1/3 pods have been updated to the latest generation
baremetal                                  4.10.3    True        False         False      3h51m   
cloud-controller-manager                   4.10.3    True        False         False      3h56m   
cloud-credential                           4.10.3    True        False         False      3h57m   
cluster-autoscaler                         4.10.3    True        False         False      3h51m   
config-operator                            4.10.3    True        False         False      3h53m   
console                                    4.10.3    True        False         False      76m     
csi-snapshot-controller                    4.10.3    True        False         False      3h52m   
dns                                        4.10.3    True        False         False      3h51m   
etcd                                       4.10.3    True        False         False      3h50m   
image-registry                             4.10.3    True        False         False      3h43m   
ingress                                    4.10.3    True        False         False      3h42m   
insights                                   4.10.3    True        False         False      3h33m   
kube-apiserver                             4.10.3    True        True          False      3h38m   NodeInstallerProgressing: 2 nodes are at revision 9; 1 nodes are at revision 10
kube-controller-manager                    4.10.3    True        False         False      3h49m   
kube-scheduler                             4.10.3    True        False         False      3h49m   
kube-storage-version-migrator              4.10.3    True        False         False      56s     
machine-api                                4.10.3    True        False         False      3h48m   
machine-approver                           4.10.3    True        False         False      3h51m   
machine-config                             4.10.3    True        False         False      3h51m   
marketplace                                4.10.3    True        False         False      3h51m   
monitoring                                 4.10.3    True        False         False      3h31m   
network                                    4.10.3    True        False         False      3h53m   
node-tuning                                4.10.3    True        False         False      3h51m   
openshift-apiserver                        4.10.3    True        True          False      3h38m   APIServerDeploymentProgressing: deployment/apiserver.openshift-apiserver: 1/3 pods have been updated to the latest generation
openshift-controller-manager               4.10.3    True        False         False      3h51m   
openshift-samples                          4.10.3    True        False         False      3h49m   
operator-lifecycle-manager                 4.10.3    True        False         False      3h52m   
operator-lifecycle-manager-catalog         4.10.3    True        False         False      3h52m   
operator-lifecycle-manager-packageserver   4.10.3    True        False         False      29s     
service-ca                                 4.10.3    True        False         False      3h53m   
storage                                    4.10.3    True        True          False      3h52m   GCPPDCSIDriverOperatorCRProgressing: GCPPDDriverControllerServiceControllerProgressing: Waiting for Deployment to deploy pods
```

### The OCP CIS Benchmark

The OCP CIS Benchmark consists of three sets of policies:

- *ocp-cis* for Cluster polies
- *ocp-node-master* for control node policies
- *ocp-node-worker* for compute node policies

The following commmands can be used to do compliance checks for each of the above policies. The examples below look for compliance checks that have *FAILED*, to check for policies that have passed use *PASS* or policies that require manual evaluation use *MANUAL*.

```bash
oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis,compliance.openshift.io/check-status=FAIL

oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis-node-master,compliance.openshift.io/check-status=FAIL

oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis-node-worker,compliance.openshift.io/check-status=FAIL
```

Once any remediation have be made the following commands can be used to trigger a compliance rescan. A [script](rescan.sh) to trigger a rescan of all scans has been provided. 

```bash
oc annotate compliancescans/ocp4-cis compliance.openshift.io/rescan=

oc annotate compliancescans/ocp4-cis-node-master compliance.openshift.io/rescan=

oc annotate compliancescans/ocp4-cis-node-worker compliance.openshift.io/rescan=
```

The compliance operator will auto-remediate all CIS policies with the exception of those policies that the CIS define as requiring manual rediation. In the next section we will address these [manual remdiations](/accelerators/kubernetes/ocp/gcp/06_remediation_of_manual_CIS_controls/Remediation_of_manual_CIS_controls.md).
