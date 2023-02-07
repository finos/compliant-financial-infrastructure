[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6557/badge)](https://bestpractices.coreinfrastructure.org/projects/6557)
[<img src="https://img.shields.io/badge/slack-@finos/compliant%20financial%20infrastructure-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cfi/)

<img src="https://github.com/finos/branding/blob/master/project-logos/active-project-logos/Compliant%20Financial%20Infrastructure%20Logo/Horizontal/2021_CFI_Logo_Horizontal.png?raw=true" width="450">


# Compliant Financial Infrastructure

Compliant Financial Infrastructure (CFI) is a project that exists to accelerate the development, deployment and adoption of services provided for infrastructure in a way that meets common regulatory and internal security controls.

Through our three working groups, we provide:

- Opinionated compliance documentation provided by our *service approval accelerators*
- Vetted *infrastructure as code* that is ready to import to your internal registry
- CI/CD-friendly *runtime validation tests* to ensure your deployed resources are compliant

## Policy Working Group [<img src="https://img.shields.io/badge/slack-@finos/cfi%20policy-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cfi-policy-wg)

This WG exists to define and document best practice and process for implementing compliant infrastructure, while streamlining the process for contributions from financial institutions in a frictionless manner.

_Compliance_ may mean something different from one institution to the next. The goal of CFI is _not_ to create a single solution that all firms must adhere to, instead our goal is to streamline adoption and free up security teams to focus on non-redundant activities.

Detailed documentation in the form of _Service Approval Accelerators_ (SAAs) live within this main CFI repository.

### High level objectives

1. Maintain a knowledge base of up-to-date compliance requirements from member financial institutions (Inputs)
1. Document how to achieve compliance for different infrastructure resources from a financial perspective (Outputs)

### Approach

- Document opinionated configurations, mitigations, and decisions to accelerate compliance for infrastructure services in SAAs.
- Ensure all SAAs are informed by industry-wide experience/feedback
- Ensure CFI communication methods (both inputs and outputs) are streamlined to best serve our community and users

A template Service Approval Accelerator is maintained [here](templates/ServiceApprovalAcceleratorTemplate.md).


### Contributions

- Work for this WG is tracked in [GitHub issues on the main CFI repository](https://github.com/finos/compliant-financial-infrastructure/issues).
- Approved and active work is visualized on the [Policy WG project board](https://github.com/orgs/finos/projects/50).

## Reproducible Infrastructure Working Group [<img src="https://img.shields.io/badge/slack-@finos/cfi%20reproducible%20infrastructure-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cfi-reproducible-infrastructure-wg)

This WG exists to develop, maintain, and document easily consumable infrastructure as code (IaC) which can be used as a base for deploying systems in highly-regulated environments.

Detailed documentation regarding the process for developing and delivering IaC can be found [here](https://github.com/finos/compliant-financial-infrastructure/blob/docs/wg-readme/docs/terraform-module-best-practices.md).

### High level objectives

1. Create and maintain IaC to deploy services that meet policies as defined by the Policy Working Group

### Approach

- Review Service Accelerators and work with the Policy Working Group to agree on each approach to codify policies
- Build and maintain the IaC to meet requirements set out in the SAA
  - Where this is not possible then any policy gaps will be documented

### Contributions

- Work for this WG that does not yet have a dedicated repo is tracked in [GitHub issues on the main CFI repository](https://github.com/finos/compliant-financial-infrastructure/issues).
- Work for IaC that has already begun will be tracked on the respective repo.
- Approved and active work is visualized on the [Reproducible Infrastructure WG project board](https://github.com/orgs/finos/projects/45).

## Runtime Validation Working Group [<img src="https://img.shields.io/badge/slack-@finos/cfi%20runtime%20validation-green.svg?logo=slack">](https://finos-lf.slack.com/messages/cfi-runtime-validation-wg)

This WG exists to maintain a suite of tools that may be used to validate that deployed infrastructure is compliant with the documentation provided by the Policy Working Group, and provide actionable information for users who are working toward compliance.

Detailed documentation regarding the process for developing and delivering runtime validation test packs can be found [here](?).

### High level objectives

1. Maintain tests matching each SAA to validate the compliance of any deployed resource
1. Maintain test harness to streamline approach across all services

### Approach

- Execute tests that match the accelerators provided by the Policy WG (no more, no less)
- Ensure harnes is easily configurable & can be used for diverse validation purposes
- Maintain smooth logging functionality for validation and development purposes
- Ensure common human-readable output format for all test packs

### Contributions

- Work for this WG that does not yet have a dedicated repo is tracked in [GitHub issues on the main CFI repository](https://github.com/finos/compliant-financial-infrastructure/issues).
- Work on test packs that has already begun will be tracked on the respective repo.
- Approved and active work is visualized on the [Reproducible Infrastructure WG project board](https://github.com/orgs/finos/projects/51).


## Join the Community!

For more information about how to engage with the rest of the community and contribute to the project, view the documentation and links [here](docs/CONTRIBUTING.md).

Please feel free to request changes via [GitHub Issues](https://github.com/finos/compliant-financial-infrastructure/issues).

Everyone is encouraged to join our public community meetings found on the [FINOS community calendar](https://www.finos.org/finos-community-calendar), and join us on [Slack](https://finos-lf.slack.com/messages/cfi).


## License

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)

## Security Concerns

If you have any security concerns related to this project, please [create an issue on this repository](https://github.com/finos/compliant-financial-infrastructure/issues/new/choose) _or_ create an issue on the repository associated with your concern.
