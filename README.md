[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
[<img src="https://img.shields.io/badge/slack-@finos/compliant%20financial%20infrastructure-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

<img src="https://github.com/finos/branding/blob/master/project-logos/active-project-logos/Compliant%20Financial%20Infrastructure%20Logo/Horizontal/2021_CFI_Logo_Horizontal.png?raw=true" width="450">

# Compliant Financial Infrastructure
Compliant Financial Infrastructure (CFI) accelerates the development, deployment and adoption of services provided for _AWS_, _Azure_, _Google Cloud_, and _IBM Cloud_ in a way that meets existing regulatory and internal security controls.

## Business Problem and Opportunity
Cloud services controls and tests are used to demonstrate adherence with regulatory and internal compliance requirements mandated for financial institutions when using cloud services. The majority of cloud security incidents are due to misconfiguration; services are not secure by default, configuration is often complex, nuanced and difficult to validate. To some degree or another all financial institutions are re-inventing the wheel – institutions have similar control frameworks and each is trying to secure and stand up the same providers and services within the same regulatory frameworks.

Having robust controls and tests developed and in place removes a barrier to faster adoption of cloud services such as those provided by Amazon/AWS, Microsoft/Azure and Google/GCP, among others. Addressing this barrier will benefit both financial services IT departments, many of whom are looking to move more quickly to the cloud, and the providers themselves, who wish to sell more cloud services into financial institutions.

Controls for cloud service compliance afford banks no particular strategic or competitive advantage while also representing a task something all banks who look to deploy more applications onto the cloud needs to do, and as such are conducive to being developed together as part of the "public commons". The focused project and collaboration with other banks will increase the amount of controls produced and, it's expected, help increase the rate of adoption of cloud services.

## What is the meaning of "compliant" in CFI?

Our community intends to create an authoritative source for _defining, creating, and validating_ compliant infrastructure for financial services.
As detailed in [the following section](#services--workflows), our community focuses on these three key offerings.

Our compliance definition comes from the an extension of a baseline into a
comprehensive standard that is suited to accelerate adoption in highly regulated environements.

### Defining Compliant Infrastructure

It must be noted that every institution is made up of highly skilled professionals who collaborate to define what _compliance_ means within their firm.
As such, _compliance_ may mean something different from one institution to the next. The goal of CFI is _not_ to create a single solution that all firms must adhere to,
instead our goal is to streamline adoption and free up security teams to focus on less redundant activities.

This effort is only made possible through the efforts of our Financial Institution members to provide insight and guidance in the development of Service Approval Accelerators.

### Creating Compliant Infrastructure

Infrastructure as Code (IaC) is a key component of every modern bank's infrastructure efforts. 

Many times, however, our community has seen the task of IaC development fall on the shoulders of engineers and developers who are not fully familiar with the required technologies.
This results in added time, stress, and risks as teams become familiar with the necessary technologies.

In other cases, we see nearly identical IaC efforts being undertaken by multiple teams within a single business unit, resulting in uncessesary costs and delays.

And most importantly, we have seen teams struggling to follow best-practices
(even when their most senior people are leading the charge) because these technologies are constantly changing.

By developing _registry-ready resources_ in an open source context, specialists from all over can come together to create a solid foundation for IaC efforts.

In most cases, our community uses Terraform for our IaC. As noted [below](#services--workflows), these are developed in a way that can be
individually reviewed by your security teams and imported to a firm registry for extention or consumption.

### Validating Compliant Infrastructure

A key element to compliance is _validating the result of an activity_. This is a pillar to the CFI efforts.

In addition to technology-specific automated tests and scans in the code repository, our community believes that it is essential to provide
the tools necessary to ensure that infrastructure is compliant after it has been deployed into an environment.

In some cases, full compliance may only be possible (or preferrable) when deployed behind a firewall or with other similar contexts not reflected by the IaC code.
In these cases, only post-deployment validation will provide confidence that a deployment is actually prod-ready.

To organize and facilitate these tests, our community uses the (Probr)[http://github.com/probr] test harness to create tests for each service.
These tests are run in our validation pipelines after a deployment is complete,
and can also be run independently in any environment regardless of how the service was deployed.

## Services & Workflows

As noted [above](#what-is-the-meaning-of-compliant-in-cfi), our community produces multiple artifacts for each CFI service.
Where necessary, we have segmented the aforementioned artifacts across multiple repositories
to ensure that they are ready for quick and easy adoption by users in highly regulated environments.

1. Detailed documentation in the form of a _Service Approval Accelerator_ (SAA)
    - All SAAs live within this repository
1. Infrastructure as Code (IaC) that meets the specifications described in the SAA
    - This IaC may include multiple services if necessary to properly ensure compliance of the specified service
    - IaC registry-ready repos should be linked as git submodules within this repository
1. Post-Deployment Validation Tests to ensure that IaC is compliant
    - Test packs for each service should live in separate repos so they can be used independently of the provided IaC
1. An automated pipeline to execute the post-deployment tests and apply accurate badges to the service
    - CI pipelines should exist in this repo to tie together all other elements

## Feature Matrix

CFI provides IaC for multiple services in multiple environments. 
These services range from not-yet-implemented, to fully featured,
and automatically tested for compliance with the CDMC framework.

This feature matrix is intended to show all current and planned services along with the current status of each,
to help contributors looking for a task know where best to spend their effort.

|                    |         AWS         |         GCP         |        Azure        |    OpenShift     |
| ------------------:|:-------------------:|:-------------------:|:-------------------:|:----------------:|
| Kubernetes Cluster |      :trophy:       |       :cookie:      |  :page_facing_up:   | :page_facing_up: |
|  Postgres Database |                     |                     |      :cookie:       |                  |
|           DynamoDB |  :page_facing_up:   |                     |                     |                  |
|           RedShift |  :page_facing_up:   |                     |                     |                  |
|                SQS |  :page_facing_up:   |                     |                     |                  |


### Key

- :gem: diamond - Post-deployment validation tests exist for every aspect of the SAA,
and all tests pass in the CI pipeline testing phase.
- :trophy: gold - Post-deployment validation tests are run in CI as part of the testing phase,
and some tests pass.
- :cd: silver - The service is automatically spun up and destroyed via CI before being merged to `main`.
- :cookie: bronze - IaC has been produced that is able to create and destroy the service,
where said service meets the SAA specification.
- :page_facing_up: paper - A complete SAA document for this service has been merged to `main`.

## Project Kanban
Find the [Compliant Financial Infrastructure Project Kanban](https://github.com/orgs/finos/projects/1) in the parent FINOS organisation on GitHub.

### Compliant Financial Infrastructure - Agile Workflow

The Agile Workflow for Compliant Financial Infrastructure falls into three main work streams which are overseen by Project Maintainers and fulfilled by the project team and wider FINOS community.

- [Agile Delivery of Prioritised Work Items](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#agile)
- [Community Contributions and Pull Requests](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#community)
- [Compliant Financial Infrastructure Asynchronous Pull Request and Code Reviews](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#reviews)

### Contributing - Forking, Feature Branches and Pull Requests

1. Fork it (<https://github.com/finos/compliant-financial-infrastructure/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Service Approval Accelerator

The [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) (SAA) describes each service contributed to Compliant Financial Infrastructure alongside test cases and infrastructure as code.

A single SAA should be contributed with every service contributed to Compliant Financial Infrastructure. 

You may review existing services for examples of existing SAAs.

## Compliant Financial Infrastructure Mailing List
Compliant Financial Infrastructure email communications are conducted through the compliant-financial-infrastructure@finos.org mailing list. Email compliant-financial-infrastructure@finos.org with questions or suggestions related to Compliant Financial Infrastructure.

Subscribe to the Compliant Financial Infrastructure mailing list by sending an email to compliant-financial-infrastructure+subscribe@finos.org.

## Join the Compliant Financial Infrastructure Slack Channel
Join Compliant Financial Infrastructure on the FINOS Slack by signing up at https://finos-lf.slack.com/. The Compliant Financial Infrastructure channel on Slack is found directly at https://finos-lf.slack.com/messages/compliant-financial-infrastructure/.

[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

Reach out to help@finos.org for any issues when joining Compliant Financial Infrastructure on the FINOS Slack.

## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
