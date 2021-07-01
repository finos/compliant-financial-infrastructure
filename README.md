[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cloud-service-certification/)

<img src="https://github.com/finos/branding/blob/master/project-logos/active-project-logos/Cloud%20Service%20Certification%20Logo/Horizontal/2020_CloudServicesCertification_Horizontal.png?raw=true" width="400">

# Cloud Service Certification
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

1. Fork it (<https://github.com/finos/cloud-service-certification/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Service Approval Accelerator

The [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) describes each service contributed to Cloud Service Certification alongside test cases and infrastructure as code.

A single [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) document should be contributed with every service contributed to Cloud Service Certifcation. 

_See AWS Redshift example below_.

## Contributed Cloud Service by Example, AWS Redshift

The [AWS RedShift Service Definition](https://github.com/finos/cloud-service-certification/tree/master/aws/redshift) has been created to demonstrate through example the assets required with each service contribution to Cloud Service Certification.

* [Redshift Test Cases](aws/redshift/RedshiftTestCases.md) : 
  * A document containing test cases from the point of view of AWS Redshift. 
* [Redshift Service Approval Accelerator](aws/redshift/ServiceApprovalAcceleratorRedshift.md) : 
  * A document containing the Service Approval Accelerator from the point of view of AWS Redshift.
* [The Redshift Service Definition](aws/redshift/redshift_template_public.yml) : 
  * A YAML file containing the description of the AWS Redshift service as code.

## CSC Mailing List
Cloud Service Certification email communications are conducted through the fdx-cloud-service-certification@finos.org mailing list. Email fdx-cloud-service-certification@finos.org with questions or suggestions related to Cloud Service Certification.

Subscribe to the Cloud Service Certification mailing list by sending an email to fdx-cloud-service-certification+subscribe@finos.org.

## Join the CSC Slack Channel
Join Cloud Service Certification on the FINOS Slack by signing up at https://finos-lf.slack.com/. The Cloud Service Certification channel on Slack is found directly at https://finos-lf.slack.com/messages/cloud-service-certification/.

[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cloud-service-certification/)

Reach out to help@finos.org for any issues when joining CSC on the FINOS Slack.

## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
