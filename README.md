[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
[<img src="https://img.shields.io/badge/slack-@finos/compliant%20financial%20infrastructure-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

<img src="https://github.com/finos/branding/blob/master/project-logos/active-project-logos/Compliant%20Financial%20Infrastructure%20Logo/Horizontal/2021_CFI_Logo_Horizontal.png?raw=true" width="450">


# Compliant Financial Infrastructure
Compliant Financial Infrastructure (CFI) accelerates the development, deployment and adoption of services provided for infrastructure in a way that meets existing regulatory and internal security controls.

## What's the value?

### What does this project provide to financial institutions?

Our community collaboratively creates accelerators to aid in the adoption of specific infrastructure based on input and experience from professionals across a variety of roles and firms.

A firm may take benefit from one or more of the following:

- Opinionated compliance documentation provided by our *service approval accelerators*
- Vetted *infrastructure as code* that is ready to import to your internal registry
- CI/CD-friendly *post-deployment validation tests* to ensure your deployed resources are compliant

### Why should I contribute to this project?

Every active community member serves to accelerate the frequency that we are able to make resources available! Participation from security and compliance professionals helps to ensure that opinions made by our community are aligned to every firm's needs. Contribution from infrastructure engineers helps the community create more and better infrastructure as code. Collaboration from programmers and testers helps the community more quickly offer validation tests for more services.

Anything you are interested in contributing to the project will make a huge impact! Contributions will be publicly displayed on the project's GitHub history and often highlighted within FINOS and CFI community sessions.

### How does this project add to FINOS?

This question is answered fully in the [project charter](docs/project-charter.md).

## What is the meaning of "compliant" in CFI?

Our community has created an authoritative source for _defining, creating, and validating_ infrastructure for financial services that is compliant with regulatory standards.
As detailed in [the following section](#services--workflows), our community focuses on these three key offerings.

Our compliance definition comes from the extension of a baseline into a comprehensive standard that is suited to accelerate adoption in highly regulated environments.

### Defining Compliant Infrastructure

It must be noted that every institution is made up of highly skilled professionals who collaborate to define what _compliance_ means within their firm.
As such, _compliance_ may mean something different from one institution to the next. The goal of CFI is _not_ to create a single solution 
that all firms must adhere to, instead our goal is to streamline adoption and free up security teams to focus on non-redundant activities.

This effort is only made possible through the efforts of our Financial Institution members to provide insight and guidance in the development of Service Approval Accelerators.

A template for the Service Approval Accelerator can be found [here](templates/ServiceApprovalAcceleratorTemplate.md).

### Creating Compliant Infrastructure

Infrastructure as Code (IaC) is a key component of every modern firm's infrastructure efforts. 

Many times, however, our community has seen the task of IaC development fall on the shoulders of engineers and developers who are not fully familiar with the required technologies.
This results in added time, stress, and risks as teams become familiar with the necessary technologies.

In other cases, we see nearly identical IaC efforts being undertaken by multiple teams within a single business unit, resulting in unncessesary costs and delays.

And most importantly, we have seen teams struggling to follow best-practices
(even when their most senior people are leading the charge) because these technologies are constantly changing.

By developing _registry-ready resources_ in an open source context, specialists from all over can come together to create a solid foundation for IaC efforts.

In most cases, our community uses Terraform for our IaC. As noted [below](#services--workflows), these are developed in a way that can be
individually reviewed by your security teams and imported to a firm registry for extention or consumption.

### Validating Compliant Infrastructure

A key element to compliance is _validating the result of an activity_. This is a pillar to the CFI efforts.

In addition to technology-specific automated tests and scans in the code repository, our community believes that it is essential to provide
the tools necessary to ensure that infrastructure is compliant after it has been deployed into an environment.

In some cases, full compliance may only be possible (or preferable) when deployed behind a firewall or with other similar contexts not reflected by the IaC code.
In these cases, only post-deployment validation will provide confidence that a deployment is actually prod-ready.

To organize and facilitate these tests, our community uses the [Probr](https://github.com/probr) test harness to create tests for each service.
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

CFI addresses multiple services, across multiple providers.
These services range from not-yet-implemented, to fully featured,
and automatically tested for compliance with the CDMC framework.

This feature matrix is intended to show all current and planned services along with the current status of each,
to help contributors looking for a task know where best to spend their effort.

|                    |         AWS                                                                             |         GCP                                                     |        Azure                                                 |    OpenShift                                               |
| ------------------:|:---------------------------------------------------------------------------------------:|:--------------------------------------------------------------: |:------------------------------------------------------------:|:----------------------------------------------------------:|
| Kubernetes Cluster | ![](https://byob.yarr.is/finos/compliant-financial-infrastructure/eks-terraform)        | <img src="docs/_images/bronze.png"  alt="bronze" height="20"/>  | <img src="docs/_images/paper.png" alt="paper" height="20"/>  |<img src="docs/_images/paper.png" alt="paper" height="20"/> |
|  Postgres Database |                                                                                         |                                                                 | <img src="docs/_images/bronze.png" alt="bronze" height="20"/>|                                                            |
|           DynamoDB |  <img src="docs/_images/paper.png" alt="paper" height="20"/>                            |                                                                 |                                                              |                                                            |
|           RedShift |  <img src="docs/_images/paper.png" alt="paper" height="20"/>                            |                                                                 |                                                              |                                                            |
|                SQS |                                                                                         |                                                                 |                                                              |                                                            |


### Key

- <img src="docs/_images/diamond.png" alt="diamond" height="20"/> - Post-deployment validation tests exist for every aspect of the SAA,
and all tests pass in the CI pipeline testing phase.
- <img src="docs/_images/gold.png" alt="gold" height="20"/> - Post-deployment validation tests are run in CI as part of the testing phase,
and some tests pass.
- <img src="docs/_images/silver.png" alt="silver" height="20"/> - The service is automatically spun up and destroyed via CI before being merged to `main`.
- <img src="docs/_images/bronze.png"  alt="bronze" height="20"/> - IaC has been produced that is able to create and destroy the service,
where said service meets the SAA specification.
- <img src="docs/_images/paper.png" alt="paper" height="20"/> - A complete SAA document for this service has been merged to `main`.

## Join the Community!

For more information about how to engage with the rest of the community and contribute to the project, view the documentation and links [here](docs/CONTRIBUTING.md).

## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)

## Suggestions and Change Requests

Please feel free to engage the community via [Slack](https://finos-lf.slack.com/messages/compliant-financial-infrastructure), [GitHub Discussions](https://github.com/finos/compliant-financial-infrastructure/discussions), or [GitHub Issues](https://github.com/finos/compliant-financial-infrastructure/issues).

## Security Concerns

If you have any security concerns related to this project, please [create an issue on this repository](https://github.com/finos/compliant-financial-infrastructure/issues/new/choose) _or_ create an issue on the repository associated with your concern.
