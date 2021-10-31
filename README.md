[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
[<img src="https://img.shields.io/badge/slack-@finos/compliant%20financial%20infrastructure-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

<img src="https://github.com/finos/branding/blob/master/project-logos/active-project-logos/Cloud%20Service%20Certification%20Logo/Horizontal/2020_CloudServicesCertification_Horizontal.png?raw=true" width="400">

### Important Notice
The Compliant Financial Infrastructure default branch has been renamed. 

`master` is now named `main`

If you have a local clone, you can update it by running:

```
git branch -m master main
git fetch origin
git branch -u origin/main main
git remote set-head origin -a
```

For more complex branch changes, please read the following issue comment https://github.com/finos/compliant-financial-infrastructure/issues/119#issuecomment-843231539 

# Compliant Financial Infrastructure
Compliant Financial Infrastructure accelerates the development, deployment and adoption of services provided for AWS, Azure and Google in a way that meets existing regulatory and internal security controls.

## Business Problem and Opportunity
Cloud services controls and tests are used to demonstrate adherence with regulatory and internal compliance requirements mandated for financial institutions when using cloud services. The majority of cloud security incidents are due to misconfiguration; services are not secure by default, configuration is often complex, nuanced and difficult to validate. To some degree or another all financial institutions are re-inventing the wheel – institutions have similar control frameworks and each is trying to secure and stand up the same providers and services within the same regulatory frameworks.

Having robust controls and tests developed and in place removes a barrier to faster adoption of cloud services such as those provided by Amazon/AWS, Microsoft/Azure and Google/GCP, among others. Addressing this barrier will benefit both financial services IT departments, many of whom are looking to move more quickly to the cloud, and the providers themselves, who wish to sell more cloud services into financial institutions.

Controls for cloud service compliance afford banks no particular strategic or competitive advantage while also representing a task something all banks who look to deploy more applications onto the cloud needs to do, and as such are conducive to being developed together as part of the "public commons". The focused project and collaboration with other banks will increase the amount of controls produced and, it's expected, help increase the rate of adoption of cloud services.

## Approach and Proposed Solution

This FINOS project produces multiple artifacts (together forming one or multiple accelerators) that provide functional code that implements regulatory compliant configurations of cloud services with BDD tests to validate efficacy. The group review the artifacts for an accelerator and then gather feedback on process and content before iterating on additional services. A key part of the project's approach is to set quality standards across artifacts; members of all tiers can contribute to the project and ensure a common high level of quality is delivered and in less time. The group will also work with cloud service providers to produce more industry specific content and solutions.

## Project Kanban
Find the [Compliant Financial Infrastructure Project Kanban](https://github.com/orgs/finos/projects/1) in the parent FINOS organisation on GitHub.

# Compliant Financial Infrastructure - Agile Workflow

The Agile Workflow for Compliant Financial Infrastructure falls into three main work streams which are overseen by Project Maintainers and fulfilled by the project team and wider FINOS community.

- [Agile Delivery of Prioritised Work Items](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#agile)
- [Community Contributions and Pull Requests](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#community)
- [Compliant Financial Infrastructure Asynchronous Pull Request and Code Reviews](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#reviews)

# Contributing

## Forking, Feature Branches and Pull Requests

1. Fork it (<https://github.com/finos/compliant-financial-infrastructure/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Service Approval Accelerator

The [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) describes each service contributed to Compliant Financial Infrastructure alongside test cases and infrastructure as code.

A single [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) document should be contributed with every service contributed to Compliant Financial Infrastructure. 

_See AWS Redshift example below_.

## Contributed Cloud Service by Example, AWS Redshift

The [AWS RedShift Service Definition](https://github.com/finos/compliant-financial-infrastructure/tree/master/aws/redshift) has been created to demonstrate through example the assets required with each service contribution to Compliant Financial Infrastructure.

* [Redshift Test Cases](aws/redshift/RedshiftTestCases.md) : 
  * A document containing test cases from the point of view of AWS Redshift. 
* [Redshift Service Approval Accelerator](aws/redshift/ServiceApprovalAcceleratorRedshift.md) : 
  * A document containing the Service Approval Accelerator from the point of view of AWS Redshift.
* [The Redshift Service Definition](aws/redshift/redshift_template_public.yml) : 
  * A YAML file containing the description of the AWS Redshift service as code.

## Compliant Financial Infrastructure Mailing List
Compliant Financial Infrastructure email communications are conducted through the fdx-cloud-service-certification@finos.org mailing list. Email fdx-cloud-service-certification@finos.org with questions or suggestions related to Compliant Financial Infrastructure.

Subscribe to the Compliant Financial Infrastructure mailing list by sending an email to fdx-cloud-service-certification+subscribe@finos.org.

## Join the Compliant Financial Infrastructure Slack Channel
Join Compliant Financial Infrastructure on the FINOS Slack by signing up at https://finos-lf.slack.com/. The Compliant Financial Infrastructure channel on Slack is found directly at https://finos-lf.slack.com/messages/compliant-financial-infrastructure/.

[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

Reach out to help@finos.org for any issues when joining Compliant Financial Infrastructure on the FINOS Slack.

## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
