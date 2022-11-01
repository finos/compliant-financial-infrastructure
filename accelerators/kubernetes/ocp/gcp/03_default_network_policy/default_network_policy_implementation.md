# OpenShift Compliant Financial Infrastructure

A CIS policy requirement is for non control plane projects/namespaces to have network policies to isolate traffic in the cluster network.

OCP and the OVNKubernetes CNI support network policies. To implement network policies as a default we will use a [default project template](https://docs.openshift.com/container-platform/4.10/networking/network_policy/default-network-policy.html) to define these network policies.


## Creating a default network policy for projects


1. As a cluster admin user create a default project template:

```bash
oc adm create-bootstrap-project-template -o yaml > template.yaml
```

2. Create the default project template in the openshift-config namespace

```bash
oc create -f template.yaml -n openshift-config
```

3. So that this template is used we must edit the project config resiource to use template created in the last step

```bash
oc edit project.config.openshift.io/cluster
```
add the following:

```yaml
spec:
  projectRequestTemplate:
    name: project-request
```
see following example

```yaml
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: config.openshift.io/v1
kind: Project
metadata:
  annotations:
    include.release.openshift.io/ibm-cloud-managed: "true"
    include.release.openshift.io/self-managed-high-availability: "true"
    include.release.openshift.io/single-node-developer: "true"
    release.openshift.io/create-only: "true"
  creationTimestamp: "2022-07-12T11:55:34Z"
  generation: 2
  name: cluster
  ownerReferences:
  - apiVersion: config.openshift.io/v1
    kind: ClusterVersion
    name: version
    uid: 1ca64482-ba72-4021-8ae2-cfc1140449a8
  resourceVersion: "333899"
  uid: a33cf655-92f5-4623-aa61-06f10be2de5c
spec:
  projectRequestTemplate:
    name: project-request
```


4. The following will setup a network policy which provides [multitenant isolation](https://docs.openshift.com/container-platform/4.10/networking/network_policy/multitenant-network-policy.html). This comprises of three network polices:

    - allow-from-openshift-ingress
    - allow-from-openshift-monitoring
    - allow-same-namespace

To add these policies to the default project template we will edit the default project config resource:

```bash
oc edit template project-request -n openshift-config
```

Copy the network policies yaml from this [sample](/accelerators/kubernetes/ocp/gcp/03_default_network_policy/multi_tenant_isolation_netpol.yaml) 

Sample can be seen below:

```yaml
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: template.openshift.io/v1
kind: Template
metadata:
  creationTimestamp: "2022-07-13T08:46:57Z"
  name: project-request
  namespace: openshift-config
  resourceVersion: "337531"
  uid: 975b5d14-d836-4cc8-8e72-447e005a0760
objects:
- apiVersion: networking.k8s.io/v1
  kind: NetworkPolicy
  metadata:
    name: allow-from-openshift-ingress
  spec:
    ingress:
    - from:
      - namespaceSelector:
          matchLabels:
            policy-group.network.openshift.io/ingress: ""
    podSelector: {}
    policyTypes:
    - Ingress
- apiVersion: networking.k8s.io/v1
  kind: NetworkPolicy
  metadata:
    name: allow-from-openshift-monitoring
  spec:
    ingress:
    - from:
      - namespaceSelector:
          matchLabels:
            network.openshift.io/policy-group: monitoring
    podSelector: {}
    policyTypes:
    - Ingress
- apiVersion: networking.k8s.io/v1
  kind: NetworkPolicy
  metadata:
    name: allow-same-namespace
  spec:
    ingress:
    - from:
      - podSelector: {}
    podSelector: null
- apiVersion: project.openshift.io/v1
  kind: Project
  metadata:
    annotations:
      openshift.io/description: ${PROJECT_DESCRIPTION}
      openshift.io/display-name: ${PROJECT_DISPLAYNAME}
      openshift.io/requester: ${PROJECT_REQUESTING_USER}
    creationTimestamp: null
    name: ${PROJECT_NAME}
  spec: {}
  status: {}
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    creationTimestamp: null
    name: admin
    namespace: ${PROJECT_NAME}
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: admin
  subjects:
  - apiGroup: rbac.authorization.k8s.io
    kind: User
    name: ${PROJECT_ADMIN_USER}
parameters:
- name: PROJECT_NAME
- name: PROJECT_DISPLAYNAME
- name: PROJECT_DESCRIPTION
- name: PROJECT_ADMIN_USER
- name: PROJECT_REQUESTING_USER
```

5. To test that network policies are being added to new, non control plane, projects / namespaces by default create a new project and check that the multi-tenant network polices have been created: 

```bash
oc new-project test-network-policy
```

```bash
oc get networkpolicy
```

The following policies should be seen:

```console
NAME                              POD-SELECTOR   AGE
allow-from-openshift-ingress      <none>         16s
allow-from-openshift-monitoring   <none>         16s
allow-same-namespace              <none>         16s
```

The next [step](/accelerators/kubernetes/ocp/gcp/04_replace_api_router_certs/replace_api_router_certs.md) will replace the self signed certificates for the API Server and Router. 
