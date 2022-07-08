# OpenShift Compliant Financial Infrastructure

## In this section we will remediate the CIS Manual policies.

In the previous section we implemented the OCP Compliance Operator and
the CIS Benchmark for the OCP cluster, control and compute nodes. The
Compliance Operator was able to auto-remediate all failing policies and
at this point we should only have policies that require manual intervention.

There are two kinds of checks that require a manual intervention:
 - checks that would result in a FAIL, but do not have a corresponding `ComplianceRemediation`
   object. These are rules that require setting up something that the operator can't know
   beforehand. An example is setting up cluster forwarding
 - checks that would result in a MANUAL result. These need to be evaluated
   on a case-by-case basis.

Let's first illustrate how can we get more information on the checks and the rules
that those check cover.

### Getting more information about compliance rules

In general, it is often useful to get more information about a particular check
or a rule - why does the check exist, why could it fail etc.

After applying all the remediations, the following command:
```shell
oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis,compliance.openshift.io/check-status=FAIL
```

Should return two rules that are still failing:
```shell
NAME                                                      STATUS   SEVERITY
ocp4-cis-audit-log-forwarding-enabled            FAIL     medium
ocp4-cis-configure-network-policies-namespaces   FAIL     high
```

Note that if you started from a default installation, there would have been
more failing checks, e.g. one that tests that kubeadmin had been removed
or that an IDP is configured, but by following this guide, those had been
remediated already.

Let's illustrate that on the first check that enables log forwarding.
We'll examine the check first:
```shell
oc describe ccr ocp4-cis-audit-log-forwarding-enabled
```

Several attributes are useful in the describe output:
 - instructions: This is a shell command you can run to evaluate the rule manually
 - description: Why is the rule included, or why is the check important
 - annotations.compliance.openshift.io/rule: Which rule.compliance object does this check represent

Let's find more information from the rule object:
```shell
$ oc get rules | grep audit-log-forwarding-enabled
ocp4-audit-log-forwarding-enabled                                                            117m
$ oc describe rules ocp4-audit-log-forwarding-enabled
```

The `rule` object contains some of the same data as the `ComplianceCheckResult`
object, but its description also contains more information, including links
to OpenShift documentation that tells us how to enable cluster log forwarding
which would satisfy the rule. By following the documentation, we can install
the `ClusterLoggingOperator`, create an instance of `ClusterLogging` and an
instance of `ClusterLogForwarder`, which would satisfy the check.

### Addressing the MANUAL rules

Each compliance check also has a number of MANUAL rules. These are rules that
the operator can't check automatically for one reason or another, most often
because it lacks the knowledge that pertains to the particular environment and
would allow to evaluate the result. List the manual checks with:

```shell
$ oc get compliancecheckresults -l compliance.openshift.io/scan-name=ocp4-cis,compliance.openshift.io/check-status=MANUAL
```

Let's pick one check at random:
```shell
$ oc get ccr ocp4-cis-scc-limit-privileged-containers
```

The instructions attribute of the check contains:
```shell
Inspect each SCC returned from running the following command:
$ oc get scc
Review each SCC for those that have allowPrivilegedContainer set to true.
Next, examine the outputs of the following commands:
$ oc describe roles --all-namespaces
$ oc describe clusterroles
For any role/clusterrole that reference the
securitycontextconstraints resource with the resourceNames
of the SCCs that have allowPrivilegedContainer, examine the associated
rolebindings to account for the users that are bound to the role. Review the
account to determine if allowPrivilegedContainer is truly required.
```

As you can see, this is an operational control where the organization decided
which roles or clusterRole objects are allowed to use a privileged SCC.
This check must therefore be addressed manually or using a third-party policy
engine such as OPA that would limit what roles or clusterRoles can bind to
these privileged SCCs.

To filter our manual checks that have been addressed in one way or another, we
can use a `tailoredProfile`, disabling rules that were checked already:
```yaml
apiVersion: compliance.openshift.io/v1alpha1
kind: TailoredProfile
metadata:
  name: cis-filter-checked-manual-rules
spec:
  extends: ocp4-cis
  title: OCP4 CIS with fewer manual rules
  description: This tailored profile disables manual rules that we have addressed
  disableRules:
    - name: ocp4-scc-limit-privileged-containers
      rationale: We have OPA rules that prevent new clusterRoles or roles from binding to this SCC
```

For more information on `TailoredProfile` resources, please refer to the
[compliance operator documentation](https://docs.openshift.com/container-platform/4.10/security/compliance_operator/compliance-operator-tailor.html)
