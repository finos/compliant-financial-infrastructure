# OpenShift Compliant Financial Infrastructure

OCP supports a number of password providers, HTPasswd is used in this implementation of the FINOS CFI service accelerator, this should be updated where a different identity provider is being used.

OCP supports these [identify providers](https://docs.openshift.com/container-platform/4.10/authentication/understanding-identity-provider.html), including LDAP, OpenID Connect, GitHub, Google and others. 

So that the KubeAdmin user can be deleted we must first create additional administrator account(s) for the cluster. The following provides a working example using an [HTPasswd](https://docs.openshift.com/container-platform/4.10/authentication/identity_providers/configuring-htpasswd-identity-provider.html) provider.

## Setting up HTPasswd 

The following steps provide guidance on how this can be done. To follow the steps below login into the cluster that was deployed in the previous step as the KubeAdmin or system:admin user. 

An example htpasswd file can be found [here](htpasswd), this file contains the following accounts:

finos-admin-1, finos-admin-2, finos-admin-3 
finos-user-1, finos-user-2, finos-user-3 

All using the password *F1n0s_R3dH4t_123* If this file is to be used then start at step 3.

1. Create an HTPasswd file with the first administrator or user acounts that you require

```shell
htpasswd -c -B -b </path/to/htpasswd> <user_name> <password>
```

2. Add any additional administrator or user accounts

```shell
htpasswd -c -B -b </path/to/users.htpasswd> <user_name> <password>
```

3. Create a generic secret using the htpasswd file.

```shell
oc create secret generic htpass-secret --from-file=htpasswd=htpasswd -n openshift-config
```

4. Update the OAUTH config to use the HTPasswd identity provider. 
- First step is to get current copy of the config and save as a yaml file.

```shell
oc get oauth -o yaml -n openshift-config > oauth.yaml
```

  - Using an editor of your choice update *oauth.yaml* to define the HTPasswd provider, a sample *spec:* definition can be found [here](sample_htpassed_provider_oauth.yaml).

    Below is an example of an updated oauth.yaml

```yaml
apiVersion: v1
items:
- apiVersion: config.openshift.io/v1
  kind: OAuth
  metadata:
    annotations:
      include.release.openshift.io/ibm-cloud-managed: "true"
      include.release.openshift.io/self-managed-high-availability: "true"
      include.release.openshift.io/single-node-developer: "true"
      release.openshift.io/create-only: "true"
    creationTimestamp: "2022-07-12T11:55:27Z"
    generation: 1
    name: cluster
    ownerReferences:
    - apiVersion: config.openshift.io/v1
      kind: ClusterVersion
      name: version
      uid: 1ca64482-ba72-4021-8ae2-cfc1140449a8
    resourceVersion: "1684"
    uid: b21d7582-8b97-464e-bff0-a7451cbc51de
  spec:
    identityProviders:
    - name: my_htpasswd_provider
      mappingMethod: claim
      type: HTPasswd
      htpasswd:
        fileData:
          name: htpass-secret
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

  - Replace the existing OAUTH configuration with the updated oauth.yaml file using the follwing command

```shell
oc replace -f oauth.yaml
```

  - The oauth pods should then restart using the new configuration, you can check they are updated using the following command

```shell
oc get pods -n openshift-authentication
```

```bash
NAME                               READY   STATUS        RESTARTS   AGE
oauth-openshift-7495648bb6-68tmb   1/1     Terminating   0          20m
oauth-openshift-7495648bb6-gqkb9   1/1     Running       0          22m
oauth-openshift-7495648bb6-l74cv   1/1     Running       0          23m
oauth-openshift-7cdbdd45ff-8xkt4   0/1     Pending       0          25s
```


  - The last step is to give the newly created administrator account *cluster admin* priveledges, this can be done using the following command

```shell
oc adm policy add-cluster-role-to-user cluster-admin finos-admin-1
```

A sample [script](add_cluster_admin_role.sh) is provided to automate this. The response that the user is not found can be ignored. 

Once the above steps have been completed it should now be possible to login into the OCP cluster using the new credentials. 

5. Login to the cluster as one of the newly created admin users

```shell
oc login -u finos-admin-1 -p F1n0s_R3dH4t_123
```

6. Check that the account has admin privedledges: this can be done by using the following command, if the cluster nodes are not displayed the account being used does not have admin priveledges so do not proceed with the next step until corrected.

```shell
oc get nodes
```

```bash
NAME                                                 STATUS   ROLES    AGE   VERSION
ocpfinos-gm7g2-master-0.c.finos-cfi.internal         Ready    master   44m   v1.23.3+e419edf
ocpfinos-gm7g2-master-1.c.finos-cfi.internal         Ready    master   45m   v1.23.3+e419edf
ocpfinos-gm7g2-master-2.c.finos-cfi.internal         Ready    master   44m   v1.23.3+e419edf
ocpfinos-gm7g2-worker-a-mwzgg.c.finos-cfi.internal   Ready    worker   27m   v1.23.3+e419edf
ocpfinos-gm7g2-worker-b-w4ljc.c.finos-cfi.internal   Ready    worker   27m   v1.23.3+e419edf
ocpfinos-gm7g2-worker-c-28msg.c.finos-cfi.internal   Ready    worker   26m   v1.23.3+e419edf
```

7. Delete the KubeAdmin user 

```shell
oc delete secrets kubeadmin -n kube-system
```


The next [step](/accelerators/kubernetes/ocp/gcp/03_replace_api_router_certs/replace_api_router_certs.md) will replace the self signed certificates for the API Server and Router. 

