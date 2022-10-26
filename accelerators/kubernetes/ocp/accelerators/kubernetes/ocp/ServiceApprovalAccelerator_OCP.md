# Service Approval Accelerator (SAA) Template

_**Template Notes:**_

*DRAFT WORK IN PROGRESS*

This is the Service Approval Accelerator (SAA) for Red Hat OpenShift (OCP).  

The documentation provided to implement OCP so that it is compliant to this SAA will be achieved through both day one and day two configuration, the majority of day two changes being made through the use of the [Red Hat OpenShift Compliance Operator](https://docs.openshift.com/container-platform/4.11/security/compliance_operator/compliance-operator-understanding.html)

## 1. Baseline

_**Template Notes:**_

The policies included in the following document(s) are required for compliance,
unless otherwise noted in section 2: [Amendments to the Baseline](#2-amendments-to-the-baseline)

| Standard | Name | Link |
|---|---|---|
| CIS      | CIS RedHat OpenShift Container Platform Benchmark v1.2.0 | [downloads.cisecurity.org](https://downloads.cisecurity.org/#/) |


## 2. Amendments to the Baseline

_**Template Notes:**_



- _If anything in the above baseline is to be excluded or handled differently, provide details in this section_
_following the format demonstrated in section 3: [Extensions to the Baseline](#3-extensions-to-the-baseline)._
- _If no amendments are necessary, simply provide the following text:_

All guidelines provided by the aforementioned baseline are to be followed without modification or exception.

## 3. Extensions to the Baseline

_**Template Notes:**_

- _Use this section to provide a summary of any additional measures that must be taken in addition to the baseline._
- _Formatting for extensions and amendments should imitate the CIS benchmark format as closely as possible._
- _The example below is actually covered by a different CIS control, 5.5.1,_
_but I'm adding a custom version here simply for demonstration of formatting._

### 1. Ensure deployment from an unauthorized container registry is denied

**Profile Applicability:**
- Level 2 - Master Node

**Description:**

- Only accept known, authorised container registries for pod deployments.

**Rationale:**

- Containers deployed using unauthorized registries may introduce insecure images that are not approved for firm use.
To mitigate this risk, only trusted registries with approved images should be allowed.

**Impact:**

- You reject any deployment that requests an unauthorized registry.

**Audit:**

- Attempt to deploy a pod using an image from `docker.io`.

**Remediation:**

- Follow the Kubernetes documentation and setup image provenance, or restrict access to unauthorized registries on the network.

**References:**
  1. https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#imagepolicywebhook
  1. https://stackoverflow.com/questions/54463125/how-to-reject-docker-registries-in-kubernetes
CIS Compliance Operators FAILS and MANUAL




NAME                                                STATUS   SEVERITY
ocp4-cis-accounts-restrict-service-account-tokens   MANUAL   medium
ocp4-cis-accounts-unique-service-account            MANUAL   medium
ocp4-cis-api-server-oauth-https-serving-cert        MANUAL   medium
ocp4-cis-api-server-openshift-https-serving-cert    MANUAL   medium
ocp4-cis-general-apply-scc                          MANUAL   medium
ocp4-cis-general-configure-imagepolicywebhook       MANUAL   medium
ocp4-cis-general-default-namespace-use              MANUAL   medium
ocp4-cis-general-default-seccomp-profile            MANUAL   medium
ocp4-cis-general-namespaces-in-use                  MANUAL   medium
ocp4-cis-rbac-limit-cluster-admin                   MANUAL   medium
ocp4-cis-rbac-limit-secrets-access                  MANUAL   medium
ocp4-cis-rbac-pod-creation-access                   MANUAL   medium
ocp4-cis-rbac-wildcard-use                          MANUAL   medium
ocp4-cis-scc-drop-container-capabilities            MANUAL   medium
ocp4-cis-scc-limit-ipc-namespace                    MANUAL   medium
ocp4-cis-scc-limit-net-raw-capability               MANUAL   medium
ocp4-cis-scc-limit-network-namespace                MANUAL   medium
ocp4-cis-scc-limit-privilege-escalation             MANUAL   medium
ocp4-cis-scc-limit-privileged-containers            MANUAL   medium
ocp4-cis-scc-limit-process-id-namespace             MANUAL   medium
ocp4-cis-scc-limit-root-containers                  MANUAL   medium
ocp4-cis-secrets-consider-external-storage          MANUAL   medium
ocp4-cis-secrets-no-environment-variables           MANUAL   medium

ocp4-cis-audit-log-forwarding-enabled            FAIL     medium
ocp4-cis-configure-network-policies-namespaces   FAIL     high