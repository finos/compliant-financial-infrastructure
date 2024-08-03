# Contributing

## Why should I contribute to this project?

Every active community member serves to accelerate the frequency that we are able to make resources available! Participation from security and compliance professionals helps to ensure that opinions made by our community are aligned to every firm's needs. Contribution from infrastructure engineers helps the community create more and better infrastructure as code. Collaboration from programmers and testers helps the community more quickly offer validation tests for more services.

Anything you are interested in contributing to the project will make a huge impact! Contributions will be publicly displayed on the project's GitHub history and often highlighted within CFI or recognized at FINOS community sessions.

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
3. Read our [contribution guidelines](CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

## Feature Requests & Bug Reports

If you'd like to report a bug or request a new feature, create an issue on the associated repository. If you're not sure, feel free to just make an issue on [this repo](https://github.com/finos/compliant-financial-infrastructure/issues). If you'd like to discuss the matter with a maintainer or other contributors first, see below for information about Slack and community meetings.

## Infrastructure as Code

All IaC should live in external repositories that are independently tested and ready for users to import into their own module registries. Initial contributions do not need to be 100% compliant, but a badge or other documentation should be included to demonstrate the module's level of maturity.

The [child module template repository](https://github.com/finos/cfi-terraform-template-child-module) is designed to streamline the creation of compliant Terraform child modules. This template repo comes with CI tests that will be run automatically when a pull request is made to the respective repo. You may replicate these tests locally by reviewing the CI to see how the tests are installed and executed.

## Post-Deployment Validation Tests

CI/CD validation test packs should be created independently of IaC. These test packs should be ready to execute against any running resources, and should validate that the resources *can do what they're supposed to do* while simultaneously *not being able to do anything they shouldn't be able to do*.

Post-deployment validation tests do not replace pre-deployment tests, such as *semgrep* or *tfsec*.

## Community Channels

### Join the CFI Mailing List
Compliant Financial Infrastructure email communications are conducted through the cfi@lists.finos.org mailing list. Join or reach out with questions or suggestions related to Compliant Financial Infrastructure.

Subscribe to the Compliant Financial Infrastructure mailing list by sending an email to cfi+subscribe@lists.finos.org..

### Join the CFI Slack Channel
Join Compliant Financial Infrastructure on the FINOS Slack by signing up at https://finos-lf.slack.com/. The Compliant Financial Infrastructure channel on Slack is found directly at https://finos-lf.slack.com/messages/cfi/.

[<img src="https://img.shields.io/badge/slack-@finos/cloud%20service%20certification-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cfi/)

Reach out to help@finos.org for any issues when joining Compliant Financial Infrastructure on the FINOS Slack.

