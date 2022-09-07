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
For Terraform Enterprise users, _submodules_ are also visible in [the documentation](https://www.terraform.io/enterprise/registry/using#viewing-nested-modules-and-examples) which are nested modules in subdirectories of configuration repositories.  These can also be structured in open source use but should be considered only if deployables are complex and subdivision is warranted because intra-module references will have to be included to take this structure into account.

## Best Practices
This section details bullets which draw attention to good practices when handling Terraform deployments.

- Read the documentation linked in the references before looking to develop one's own modules.
- Do not put resources in root modules, only put them into child modules.
  - Doing so drives duplication of effort which already blights financial services.
  - Root modules should consist of module calls to child modules only.  Doing this means that all resources are arranged in child modules, which means that others can use the same child modules obviating the need for them to write their own (thus duplicating effort).
  - This may mean that there will be some child modules which contain single resources.  This is acceptable effect of removing duplication of effort.  It also means you will thus be able to apply a standard across your organisation.
-

## References
While this section comprises linked resources, it is strongly recommended that the reader consume all pages linked.

- [Terraform Modules](https://www.terraform.io/language/modules)
- [Module Creation - Recommended Pattern](https://learn.hashicorp.com/tutorials/terraform/pattern-module-creation?in=terraform/recommended-patterns)
