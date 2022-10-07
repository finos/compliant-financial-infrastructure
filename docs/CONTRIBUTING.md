# Contributing

## Project Kanban
Find the [Compliant Financial Infrastructure Project Kanban](https://github.com/orgs/finos/projects/1) in the parent FINOS organisation on GitHub.

### Compliant Financial Infrastructure - Agile Workflow

The Agile Workflow for Compliant Financial Infrastructure falls into three main work streams which are overseen by Project Maintainers and fulfilled by the project team and wider FINOS community.

- [Agile Delivery of Prioritised Work Items](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#agile)
- [Community Contributions and Pull Requests](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#community)
- [Compliant Financial Infrastructure Asynchronous Pull Request and Code Reviews](https://github.com/finos/compliant-financial-infrastructure/tree/main/docs/agile-workflow#reviews)

## Forking, Feature Branches and Pull Requests

1. Fork it (<https://github.com/finos/compliant-financial-infrastructure/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Feature Requests & Bug Reports

If you'd like to report a bug or request a new feature, create an issue on the associated repository. If you're not sure, feel free to just make an issue on [this repo](https://github.com/finos/compliant-financial-infrastructure/issues). If you'd like to discuss the matter with a maintainer or other contributors first, see below for information about Slack and community meetings.

## Service Approval Accelerators

The [Service Approval Accelerator](templates/ServiceApprovalAcceleratorTemplate.md) (SAA) describes each service contributed to Compliant Financial Infrastructure alongside test cases and infrastructure as code.

A single SAA should be contributed with every service contributed to Compliant Financial Infrastructure. 

You may review existing services for examples of existing SAAs.

## Infrastructure as Code

All IaC should live in external repositories that are independently tested and ready for users to import into their own module registries. Initial contributions do not need to be 100% compliant, but a badge or other documentation should be included to demonstrate the module's level of maturity.

The [child module template repository](https://github.com/finos/cfi-terraform-template-child-module) is designed to streamline the creation of compliant Terraform child modules. This template repo comes with CI tests that will be run automatically when a pull request is made to the respective repo. You may replicate these tests locally by reviewing the CI to see how the tests are installed and executed.

## Post-Deployment Validation Tests

CI/CD validation test packs should be created independently of IaC. These test packs should be ready to execute against any running resources, and should validate that the resources *can do what they're supposed to do* while simultaneously *not being able to do anything they shouldn't be able to do*.

Post-deployment validation tests do not replace pre-deployment tests, such as *semgrep* or *tfsec*.

Currently the [Probr](https://github.com/probr/) toolbox is being used as a harness to harmonize the inputs, outputs, and logs for CFI post-deployment validation tests. A new service pack (plugin) must be created for each resource that we would like to validate using Probr.

## Community Channels

### Join the CFI Mailing List
Compliant Financial Infrastructure email communications are conducted through the compliant-financial-infrastructure@finos.org mailing list. Email compliant-financial-infrastructure@finos.org with questions or suggestions related to Compliant Financial Infrastructure.

Subscribe to the Compliant Financial Infrastructure mailing list by sending an email to compliant-financial-infrastructure+subscribe@finos.org.

### Join the CFI Slack Channel
Join Compliant Financial Infrastructure on the FINOS Slack by signing up at https://finos-lf.slack.com/. The Compliant Financial Infrastructure channel on Slack is found directly at https://finos-lf.slack.com/messages/compliant-financial-infrastructure/.

[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/compliant-financial-infrastructure/)

Reach out to help@finos.org for any issues when joining Compliant Financial Infrastructure on the FINOS Slack.


### Join the Community Collaboration Meetings

The project meets every other Friday at 10am ET / 3pm UK.

#### Webex:
  - https://finos.webex.com/finos/j.php?MTID=m4b1d85127d1f3ca179545d6bd3291975

#### Dial-in
  - **US** +1-415-655-0003 US Toll
  - **UK** +44-20319-88141 UK Toll
  - **Access code:** 127 846 2278
