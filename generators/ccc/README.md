# Documentation for Using the CCC to CFI Generator

This generator is designed for converting CCC YAML formated control definitions into a wireframe CFI validator.

## Overview

The Go program generates Nuclei templates and Go code from a specified YAML component definition file. The generated files are organized by service and control identifiers, facilitating automated security testing and code generation.

### Prerequisites

1. Go: Ensure you have Go installed.
2. Nuclei: Install Nuclei, as the program relies on it for signing the templates.

### Command-Line Arguments

The program expects the following command-line arguments:

1. YAML Component Definition File: The path to the YAML file containing the component definitions.
1. CSP Name: The abbreviated name of the Cloud Service Provider (CSP). This should be less than 8 characters.
1. Service Name: The abbreviated name of the service. This should also be less than 8 characters.
1. Version: The version of the service, which should be shorter than 8 characters.

## Usage

### Run the Program

A Makefile command has been provided at the root level of the CFI directory. This command will execute the generator and move the generated files to their correct location in the repo.

```sh
make ccc-wireframe
```

To run the generator directly from this directory, use the following command:

```sh
go run main.go <component_def_file.yaml> <csp_name> <service_name> <version>
```

Example
```sh
go run main.go definitions.yaml AWS S3 v1.0
```

### Output

The program generates several outputs:

Go Files:

- main.go: Contains the main function wireframe.
- security.go: Contains security function wireframes.
- go.mod: Contains the Go module definition.
- Nuclei Templates: For each control in the component definition, a Nuclei template YAML file is created in the security directory within the specified output path.
- Nuclei Profile: A security-profile.yaml file is created, listing all the generated template files.

### Directory Structure

The output is organized into a directory structure based on the CSP and service names. If any of the directories do not exist, they will be created.

_The example below is using GCP's GCS service retrieved from a [development URL](https://raw.githubusercontent.com/eddie-knight/cfi-nuclei-demo/main/yaml/CCC.OS.yaml) for this example. The component definition has 5 controls defined in the CCC component definition._

```sh
GCP
└── GCS2
    ├── security-profile.yaml
    ├── security-templates
    │   ├── GCS2_CCC_OS_C1.yaml
    │   ├── GCS2_CCC_OS_C2.yaml
    │   ├── GCS2_CCC_OS_C3.yaml
    │   ├── GCS2_CCC_OS_C4.yaml
    │   └── GCS2_CCC_OS_C5.yaml
    └── src
        ├── go.mod
        ├── main.go
        └── security.go
```

### Error Handling

The program checks for several conditions and outputs appropriate error messages, such as:

- Missing arguments or incorrect argument lengths.
- Failure to read the YAML file or URL.
- Issues in creating output directories or files.

### Notes

- Ensure you are running this from the top level, of the repository, as it will create all subdirectories relative to the execution location.
- The generators templates directory contains main.txt, security.txt, and go.mod template files.
- The generated Nuclei templates use Zsh for executing security tests. Adjust the Engine field in the final output if using a different shell.
