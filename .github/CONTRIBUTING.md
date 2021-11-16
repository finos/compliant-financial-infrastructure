# Contributing to Compliant Financial Infrastructure
:+1: First off, thanks for taking the time to contribute! :+1:

## Contributor License Agreement (CLA)
A CLA is a document that specifies how a project is allowed to use your
contribution; they are commonly used in many open source projects.

**_All_ contributions to _all_ projects hosted by [FINOS](https://www.finos.org/)
must be made with a
[Foundation CLA](https://finosfoundation.atlassian.net/wiki/spaces/FINOS/pages/83034172/Contribute)
in place, and there are [additional legal requirements](https://finosfoundation.atlassian.net/wiki/spaces/FINOS/pages/75530375/Legal+Requirements)
that must also be met.**

Commits and pull requests to FINOS repositories will only be accepted from those contributors with an active, executed Individual Contributor License Agreement (ICLA) with FINOS OR who are covered under an existing and active Corporate Contribution License Agreement (CCLA) executed with FINOS. Commits from individuals not covered under an ICLA or CCLA will be flagged and blocked by the FINOS Clabot tool. Please note that some CCLAs require individuals/employees to be explicitly named on the CCLA.

As a result, PRs submitted to the Cloud Services Certification project cannot be accepted until you have a CLA in place with the Foundation.

Need an ICLA? Unsure if you are covered under an existing CCLA? Email [help@finos.org](mailto:help@finos.org?subject=CLA)

## Contributing Issues

### Prerequisites

* [ ] Have you [searched for duplicates](https://github.com/{program name}/{project name}/issues?utf8=%E2%9C%93&q=)?  A simple search for exception error messages or a summary of the unexpected behaviour should suffice.
* [ ] Are you running the latest version?
* [ ] Are you sure this is a bug or missing capability?

### Raising an Issue
* Create your issue [here](https://github.com/{program name}/{project name}/issues/new).
* New issues contain two templates in the description: bug report and enhancement request. Please pick the most appropriate for your issue, **then delete the other**.
  * Please also tag the new issue with either "Bug" or "Enhancement".
* Please use [Markdown formatting](https://help.github.com/categories/writing-on-github/)
liberally to assist in readability.
  * [Code fences](https://help.github.com/articles/creating-and-highlighting-code-blocks/) for exception stack traces and log entries, for example, massively improve readability.

## Contributing Pull Requests (Code & Docs)
To make review of PRs easier, please:

 * Please make sure your PRs will merge cleanly - PRs that don't are unlikely to be accepted.
 * For code contributions, follow the existing code layout.
 * For documentation contributions, follow the general structure, language, and tone of the [existing docs](https://github.com/{program name}/{project name}/wiki).
 * Keep commits small and cohesive - if you have multiple contributions, please submit them as independent commits (and ideally as independent PRs too).
 * Reference issue #s if your PR has anything to do with an issue (even if it doesn't address it).
 * Minimise non-functional changes (e.g. whitespace shenanigans).
 * Ensure all new files include a header comment block containing the [Apache License v2.0 and your copyright information](http://www.apache.org/licenses/LICENSE-2.0#apply).
 * If necessary (e.g. due to 3rd party dependency licensing requirements), update the [NOTICE file](https://github.com/{program name}/{project name}/blob/master/NOTICE) with any new attribution or other notices

### Goals

When making contributions, it's important to remember the project goals. As you have likely read elsewhere by now, the goals of this project are threefold, and each is reflected within the codebase.

1. Service Accelerator Templates (SAT)
    * Details the recommended infrastructure. Organized by cloud provider & topic/resource

1. Infrastructure As Code (IaC)
    * Provides a functional example of how the infrastructure recommended in a SAT may be implemented
    * Terraform is currently the only approved format for IaC contributions

1. IaC Verification & Validation
    * Provides pre-deployment verification that any IaC contributions are compatible with the associated SAT
    * Provides post-deployment validation that any IaC contributions successfully provide the recommended infrastructure outlined in the associated SAT

## Preparing your contributions

### Structure

Now, with the goals freshly in mind... below is the expected structure that all contributions should adhere to.

1. Cloud Service Providers
    * Content should be sorted based on the CSP that it addresses
    * An individual directory for CSP resources should live at the top level of the Compliant Financial Infrastructure repo
    * The directory should contain any high-level items such as `.md` files that pertain to all resources for that CSP
    * A subdirectory should exist that contains all IaC resources and their corresponding SATs
        * Only Terraform is currently approved for contributions, so that should be the only subdirectory here

1. SATs & Terraform Configs
    * Two types of content should exist within this subdirectory: resources and reusable modules
    * Each resource that has a SAT should have it's own directory, named after that resource
    * Terraform configs should be created after SATs, and live in the same subdirectory
    * Terraform configs should import modules that are defined in the modules directory

1. Terraform Modules
    * The central logic of all IaC should be reusable and live in a module
    * Modules should all live in their own subdirectory, with each module in another subdirectory named after the resource it pertains to
    * Modules should be usable by any config that needs it (agnostic and configurable)
    * Ideally, a separate module should exist for each type of resource that needs to be created
  
#### Example CSP Directory Structure

```sh
# finos/cloud-services-certification
├── aws
│   └── ...
└── azure
    ├── aks
    │   └── aks-kubernetes.md
    └── postgresql
        ├── ServiceApprovalAcceleratorPostgreSQL.md
        ├── media
        │   └── net-security.png
        └── terraform
            ├── README.md
            ├── main.tf
            ├── modules
            │   └── postgres
            │       ├── network.tf
            │       ├── outputs.tf
            │       ├── postgres.tf
            │       └── variables.tf
            ├── modules.tf
            ├── outputs.tf
            ├── terraform.tfvars
            └── variables.tf

```

### Workflow

Below is the recommended workflow for contributing to Compliant Financial Infrastructure. Unless otherwise stated, these are recommendations.

Each step in this workflow should be spread across separate PRs
to allow for peer review and contribution prior to advancing to the next step.

1. Establish an approved SAT prior to any other work
1. Create your IaC based on the SAT, test it in your own environment as much as possible
1. Add robust documentation for your IaC
    * Documentation is REQUIRED before CI contributions will be approved
    * Docs should include necessary perms for the service account that will execute the IaC
1. Create pre-deployment tests for the IaC (conftest, OPA)
1. Create a CI pipeline to run approved tests automatically on future changes to this IaC
1. Create a CI pipeline to deploy, validate (Probr), and destroy IaC
