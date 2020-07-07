---
name: 🛠 New Cloud Service Request
about: To create a request for a new set of cloud service controls

---

## New Cloud Service

A new cloud service must follow the standard completion template. The 'definition of done' of a cloud service request is defined as follows:

- [ ] Cloud Platform Target

The directory in which the Cloud Service will reside (/<cloud_platform>/<service_name>).
- [ ] Service Name

The name of the service.
- [ ] Service Approval Accelerator Definition

The written definition of a compliant service, including security, regulatory and quality standards for that service.
- [ ] Define Control Framework with Test Case definition

Test cases to be created to prove the integrity of the given service. These tests will form the service controls.
- [ ] Deployment Scripts Created

Terraform or YAML scripts to provision resources in keeping with the control framework.
- [ ] Gherkin Scripts Created

BDD feature scripts that map the defined test cases onto the control framework.
