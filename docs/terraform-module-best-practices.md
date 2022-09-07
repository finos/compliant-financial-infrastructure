# Terraform Child Module Best Practices
This document guides the reader in respect of contributions to the FINOS Terraform module repositories.
Where this document contradicts HashiCorp documentation for the same, the HashiCorp document takes precedence in all cases.
This document will evolve over time, as the feature set supplied evolves correspondingly.
The best practices listed here are devised with scale in mind, typical of financial services organisations, most of which already have national or international presence, or have such as an element of the mission statement.  However, it is recommended that all practitioners apply these everywhere.

## Preamble
A distinction should be made between terminology as there are functional and architectural implications as expounded below.
A _child module_ (the default meaning if documentation refers to a _module_) is defined [here](https://www.terraform.io/language/modules#child-modules).
Child modules are called in a root module configuration.
A _root module_ is the top level Terraform configuration on which a `terraform plan` would be run.  It is defined [here](https://www.terraform.io/language/modules#the-root-module).
A _terraform configuration_ generally refers to a whole configuration processed by Terraform during a run, but colloquially may be meant to refer to the root module.
For Terraform Enterprise users, _submodules_ are also visible in [the documentation](https://www.terraform.io/enterprise/registry/using#viewing-nested-modules-and-examples) which are nested modules in subdirectories of configuration repositories.  These can also be structured in open source.

## Best Practices
This section details bullets which draw attention to good practices when handling Terraform deployments. Items referred in linked references are reproduced here to drive home specific best practices.

- Read the documentation, which form the main best practices:
  - Read the [Terraform module documentation](https://www.terraform.io/language/modules) before attempting to develop one's own child module.
  - Further read the [Module Creation - Recommended Pattern](https://learn.hashicorp.com/tutorials/terraform/pattern-module-creation?in=terraform/recommended-patterns).
- Do not put resources in root modules, only put them into child modules.
  - Doing so drives duplication of effort which already blights financial services.
  - Root modules should consist of module calls to child modules only.  Doing this means that all resources are arranged in child modules, which means that others can use the same child modules obviating the need for them to write their own (thus duplicating effort).
  - This may mean that there will be some child modules which contain single resources.  This is acceptable effect of removing duplication of effort.  It also means you will thus be able to apply a standard across your organisation.
- Child modules should be versioned in their own repositories, and versioned and tagged.
  - This sets up the child module as a software component in its own right.  An agile backlog may also be appropriate.
  - Open source users will thus be able to refer to specific versions of the module using git source strings.
  - Enterprise users need to tag the repository for it to be recognised by the Terraform Enterprise registry.
  - Do not use a monorepo.
- Imagine others' use of your child modules.
  - Keep your child module as generic as possible, in order for it to be reused as easily as possible.
  - Provide the bulk of the attributes for contained resource as input variables to the child module - others will be able to configure the child module from their root modules without having to contribute changes to the repository in order to get what they want.
  - Output the bulk of the attributes and argument values exported by instantiation of resources in your child module via the outputs.tf file.  Together with the previous practice, this means that the child module will provide the maximal usable dynamic data items back to the calling root module so that they can be used without modification.
  - These practices will save you future time reviewing PRs to your module repository.
- In certain circumstances, modules may return large numbers of outputs, and these may need to be returned to an orchestrator by the calling root module.  In this case, definition of a naming convention for your outputs would be required to more easily machine returned information (such as VPC ID) which would then be used down-pipe by the orchestrator for the purposes of application component deployment on infrastructure Terraform has just deployed.
- Enumerate resources with switches.
  - Some resources in your child module will be deployed in every case.
  - For those resources which are only deployed in certain circumstances, control these with `count` like so:
  ```hcl
  resource "aws_lb" "nlb" {
    count = var.load_balancer_type == "nlb" ? 1 : 0                        # NLB only deployed if the child module is configured to deploy one.

    name               = "${var.friendly_name_prefix}-tfe-lb"
    load_balancer_type = "network"
    internal           = var.load_balancer_scheme == "external" ? false : true
    subnets            = var.lb_subnet_ids

    tags = merge({ "Name" = "${var.friendly_name_prefix}-tfe-lb" }, var.common_tags)
  }
  ```
- Use of submodules should be considered only if deployables are complex and subdivision is warranted because intra-module references will have to be included to take this structure into account.
  - Use of nested modules could be construed as reducing readability but may also be construed as increasing code organisation, so is a matter of taste; within HashiCorp, some teams use submodules, and some do not.
