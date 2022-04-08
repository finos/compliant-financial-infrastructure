# OpenSHift Compliant Financial Infrastructure

## Setting up an [identity provider] (https://docs.openshift.com/container-platform/4.10/authentication/understanding-identity-provider.html)

OpenShift supports a number oif Identify providers, including LDAP, OpenID Connect, GitHub, Google and others. So that the KubeAdmin user can be deleted we must first create additional administrator accounts for the cluster. The following provides a working example using an [HTPasswd] (https://docs.openshift.com/container-platform/4.10/authentication/identity_providers/configuring-htpasswd-identity-provider.html) provider.

The following steps provide guidance on how this can be done.

1. Create an HTPasswd file with the first administrator or user acounts that you require

'htpasswd -c -B -b </path/to/htpasswd> <user_name> <password>'

2. Add any additional administrator or user accounts

'htpasswd -c -B -b </path/to/users.htpasswd> <user_name> <password>'

An example htpasswd file can be found [here] (htpasswd), this file contains the folloing accounts:
finos-admin-1, finos-admin-2, finos-admin-3 
finos-user-1, finos-user-2, finos-user-3 

All using the password 'F1n0s_R3dH4t_123'

3. Create a generic secret using the htpasswd file.
'oc create secret generic htpass-secret --from-file=htpasswd=htpasswd -n openshift-config

4. Update the OAUTH config to use the HTPasswd identity provider. 
- First step is to get current copy of the config and save as a yaml file.
'oc get oauth -o yaml -n openshift-config > oauth.yaml'
- Using an editor of your choice update the file, a updated sample oauth.yaml file can be found [here] (oauth.yaml). 
- Replace the existing OAUTH configuration with the update oauth.yaml file using the follwing command 'oc replace -f oauth.yaml'.
- The oauth pods should the restart using the new configuration, you can check they are updated using the following command
'oc get pods -n openshift-authentication'
- The last step is to give the newly created admin account *cluster admin* priveledges, this can be done using the following command
'oc adm policy add-cluster-role-to-user cluster-admin finos-admin-1'
We have also provided a [script] (add_cluster_admin_role.sh) to automate this.

Once the above steps have been completed you can now login into the OpenShift cluster using your new credentials. 

## The next step is to complete the day 2 customisations needed to meet the policied in the service accelerator. 
