# CFI Service Approval Accelerators

This directory is to contain the service approval accelerators (SAA)
alongside links to associated IaC and test packs.

## Organization

- All content related to a specific service should live in a
directory that is named according to the service.
- For the sake of organization, these service directories may
be organized within topical directories, such as "kubernetes."
- Example:

    ```sh
    accelerators
    ├───dynamodb # service directory
    ├───kubernetes # topical directory
    │   ├───aks # service directory
    │   ├───eks
    │   ├───gke
    │   └───ocp
    └───postgresql
        ├───aws
        ├───az
        ├───gcp
        └───rh
    ```

The following conventions should be followed for each type
of file or directory that is being created.

Naming conventions are designed to provide a simplified workflow
when users are navigating directories and repositories.

### Service Directories

This directory should house the SAA file and submodules for the IaC
& test pack. Additional ad-hoc resources are acceptable if
strictly necessary.

#### Naming Conventions for Service Directories

- All lowercase
- Underscore delimiters
- Use acronyms to shorten the name when possible
- Do not use abbreviations (such as "k8s" for Kubernetes)
- Avoid multi-word names, but use underscores for spaces if necessary
- Example: `gke/`

### Service Approval Accelerators

This file is a guide to accelerate the creation of IaC and test packs.
It is preferred to use markdown format ahead of other options.

#### Naming Conventions for Service Approval Accelerators

- All lowercase
- Underscore delimiters
- First value is `saa_`
- Second value is the cloud provider
- Third and final value is the service name, matching the directory name
- Format: `saa_<cloud-provider>_<service-name>.md`
- Example: `saa_gcp_gke.md`

### IaC Submodules or Subdirectories

The infrastructure as code to provision a particular resource should
live outside this repo when possible (some content may still be
pending relocation).
If the code has not yet been extracted, it should be in a subdirectory.
Otherwise, it should be linked as a git submodule.

Be sure to follow the technology-specific registry guidelines in the event
that they conflict with these naming guidelines.

#### Naming Conventions for IaC

- All lowercase
- Underscore delimiters
- First value is an abbreviation of the IaC technology name
- Second value is an abbreviation of the control standard
- Third and final value is the service name, matching the directory name
- Format:
`<technology-abbreviation>_<control-standard>_<cloud-provider>_<service-name>`
- Example: `tf_cdmc_gcp_gke`
