[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
# Cloud Service Certification

## Mission
Cloud Service Certification accelerates the development, deployment and adoption of services provided for AWS, Azure and Google in a way that meets existing regulatory and internal security controls.

## Business Problem and Opportunity
Cloud services controls and tests are used to demonstrate adherence with regulatory and internal compliance requirements mandated for financial institutions when using cloud services. The majority of cloud security incidents are due to misconfiguration; services are not secure by default, configuration is often complex, nuanced and difficult to validate. To some degree or another all financial institutions are re-inventing the wheel – institutions have similar control frameworks and each is trying to secure and stand up the same providers and services within the same regulatory frameworks.

Having robust controls and tests developed and in place removes a barrier to faster adoption of cloud services such as those provided by Amazon/AWS, Microsoft/Azure and Google/GCP, among others. Addressing this barrier will benefit both financial services IT departments, many of whom are looking to move more quickly to the cloud, and the providers themselves, who wish to sell more cloud services into financial institutions.

Controls for cloud service compliance afford banks no particular strategic or competitive advantage while also representing a task something all banks who look to deploy more applications onto the cloud needs to do, and as such are conducive to being developed together as part of the "public commons". The focused project and collaboration with other banks will increase the amount of controls produced and, it's expected, help increase the rate of adoption of cloud services.

## Approach and Proposed Solution

The working group will produce multiple Cloud Service Certification artifacts (together forming one or multiple accelerators) that provide functional code that implements regulatory compliant configurations of cloud services with BDD tests to validate efficacy. The group review the artifacts for an accelerator and then gather feedback on process and content before iterating on additional services. A key part of the working group's approach will be to set quality standards across artifacts; members of all tiers can contribute to the project and ensure a common high level of quality is delivered and in less time. The group will also work with cloud service providers to produce more industry specific content and solutions.

## Project Kanban
Find the [Cloud Service Certification Project Kanban](https://github.com/orgs/finos/projects/1) in the parent FINOS organisation on GitHub.

# Contributing

## Forking, Feature Branches and Pull Requests

1. Fork it (<https://github.com/finos-fdx/cloud-service-certification/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Service Approval Accelerator

The [Service Approval Accelerator](https://github.com/finos/cloud-service-certification/blob/master/templates/ServiceApprovalAcceleratorTemplate.md) describes each service contributed to Cloud Service Certification alongside test cases and infrastructure as code.

A single [Service Approval Accelerator](https://github.com/finos/cloud-service-certification/blob/master/templates/ServiceApprovalAcceleratorTemplate.md) document should be contributed with every service contributed to Cloud Service Certifcation. 

_See AWS Redshift example below_.

## Contributed Cloud Service by Example, AWS Redshift

The [AWS RedShift Service Definition](https://github.com/finos/cloud-service-certification/tree/master/aws/redshift) has been created to demonstrate through example the assets required with each service contribution to Cloud Service Certification.

* [Redshift Test Cases](https://github.com/finos/cloud-service-certification/blob/master/aws/redshift/RedshiftTestCases.md) : 
  * A document containing test cases from the point of view of AWS Redshift. 
* [Redshift Service Approval Accelerator](https://github.com/finos/cloud-service-certification/blob/master/aws/redshift/ServiceApprovalAcceleratorRedshift.md) : 
  * A document containing the Service Approval Accelerator from the point of view of AWS Redshift.
* [The Redshift Service Definition](https://github.com/finos/cloud-service-certification/blob/master/aws/redshift/redshift_template_public.yml) : 
  * A YAML file containing the description of the AWS Redshift service as code.

## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
