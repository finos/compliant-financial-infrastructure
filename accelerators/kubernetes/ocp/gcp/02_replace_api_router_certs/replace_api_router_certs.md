# OpenShift Compliant Financial Infrastructure

OpenShift uses certificates to encrypt the communication with the Web Console as well as applications exposed as Routes. Without any further customisation the install process will create self-signed certificates. While these work they usually trigger severe security warnings about unknown certificates in Web Browsers and should be replaced with approved certificates for both the API and router endpoints. The following steps uses [acme.sh](https://github.com/acmesh-official/acme.sh) and [Let's Encrypt](https://letsencrypt.org/) this example is for PoC purposes and should a firms approved certificates should be used for a production deployment. 

## Day 2 Replace API and Router Certificates

1. Install [acme.sh](https://github.com/acmesh-official/acme.sh) a number of installation options are provided in [here](https://github.com/acmesh-official/acme.sh#1-how-to-install)

2. By default acme.sh uses [ZeroSSL.com](https://github.com/acmesh-official/acme.sh/wiki/ZeroSSL.com-CA) we will use [Let's Encrypt](https://letsencrypt.org/), to set that as the default use the following command

`acme.sh  --set-default-ca  --server letsencrypt`

3. Identify the FQDN for both the API and Router endpoints (you need to be logged onto OCP) and assign them as variables

API endpoint:

`oc whoami --show-server | cut -f 2 -d ':' | cut -f 3 -d '/' | sed 's/-api././'`

*api.finos1.ahamgcp.com*

Router endpoint:

`oc get ingresscontroller default -n openshift-ingress-operator -o jsonpath='{.status.domain}`

*apps.finos1.ahamgcp.com*


4. acme.sh support a large number of DNS providers including Google Cloud DNS, this [link](https://github.com/acmesh-official/acme.sh/wiki/dnsapi#49-use-google-cloud-dns-api-to-automatically-issue-cert) provides more information. To issue a certificate use the following steps:

    - authenticate with GCP using 'gcloud init' go through the dialog to initiase the connection 
    - in the .acme.sh use following command to issue the certificates
      
      ```bash
      acme.sh --issue -d api.finos1.ahamgcp.com --dns dns_gcloud
      acme.sh --issue -d *.apps.finos1.ahamgcp.com --dns dns_gcloud
      ```

    - move the certificates from the acme.sh path to a known working directory, for this example we are using *home*/certificates/api and *home*/certificates/router

      ```bash
      acme.sh --install-cert -d api.finos1.ahamgcp.com --cert-file /Users/*home*/certificates/api/cert.pem --key-file /Users/*home*/certificates/api/key.pem --fullchain-file /Users/*home*/certificates/api/fullchain.pem --ca-file /Users/*home*/certificates/api/ca.cer

      acme.sh --install-cert -d *.apps.finos1.ahamgcp.com --cert-file /Users/*home*/certificates/router/cert.pem --key-file /Users/*home*/certificates/router/key.pem --fullchain-file /Users/*home*/certificates/router/fullchain.pem --ca-file /Users/*home*/certificates/router/ca.cer
      ```

      check that the certificates exist in the target directories.

5. To replace the API endpoint certificate use following command, more details can be founfd in the [OCP documentation]{https://docs.openshift.com/container-platform/4.10/security/certificates/api-server.html}

    - Create a secret using the API endpoint certifcate chain and private key created in the previous step

    `oc create secret tls api-certs --cert=fullchain.pem --key=key.pem -n openshift-config`

    - Patch the apiserver to use the new certificate
    
    ```oc patch apiserver cluster --type=merge -p '{"spec":{"servingCerts": {"namedCertificates": [{"names": ["api.finos1.ahamgcp.com"], "servingCertificate": {"name": "api-certs"}}]}}}'```

    - To check the update has been made to the apiserver review the servingCerts using the following command

    `oc get apiserver cluster -o yaml`
    
    *servingCerts:*
    ```yaml
    namedCertificates:
    - names:
      - api.finos1.ahamgcp.com
      servingCertificate:
        name: api-certs
     ```

    - The changes will be rolled out which will take a few minutes to check progress use the following command

      `oc get clusteroperators kube-apiserver`

      The following output shows that the change has been made, do not proceed until progressing is false


``console
NAME              VERSION    AVAILABLE    PROGRESSING   DEGRADED   SINCE    MESSAGE 
kube-apiserver    4.10.3     True         True          False      98m     NodeInstallerProgressing: 3 nodes are at revision 9; 0 nodes have achieved new revision 10 
kube-apiserver    4.10.3     True         True          False      98m     NodeInstallerProgressing: 3 nodes are at revision 9; 0 nodes have achieved new revision 10 
kube-apiserver    4.10.3     True         True          False      98m     NodeInstallerProgressing: 3 nodes are at revision 9; 0 nodes have achieved new revision 10 
kube-apiserver    4.10.3     True         True          False      98m     NodeInstallerProgressing: 3 nodes are at revision 9; 0 nodes have achieved new revision 11 
kube-apiserver    4.10.3     True         True          False      102m    NodeInstallerProgressing: 2 nodes are at revision 9; 1 nodes are at revision 11 
kube-apiserver    4.10.3     True         False         False      122m 
```


Once completed you will need to log back into the API server. To confirm that the new certificate is being used the following curl command can be used:

`curl -v https://api.finos1.ahamgcp.com:6443`

6. To replace the default Router endpoint certificate use following command, more details can be founfd in the [OCP documentation]{https://docs.openshift.com/container-platform/4.10/security/certificates/replacing-default-ingress-certificate.html}

 - Create a secret using the Router endpoint certifcate chain and private key created in the previous step

    `oc create secret tls router-certs --cert=fullchain.pem --key=key.pem -n openshift-ingress`

  - Patch the apiserver to use the new certificate
    
    `oc patch ingresscontroller.operator default --type=merge -p '{"spec":{"defaultCertificate": {"name": "router-certs"}}}' -n openshift-ingress-operator`

It will take a few minutes for the update to replicate, to check login to the OCP Console and check that the connection is secure and using the new certificate. 

The next step is to complete the day 2 customisation.
