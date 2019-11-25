[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://finosfoundation.atlassian.net/wiki/display/FINOS/Incubating)

Find the [Cloud Service Certification Project Kanban](https://github.com/orgs/finos/projects/1) in the parent FINOS organisation on GitHub.

# Cloud Service Certification

Enable banks to use services provided by AWS, Azure and Google in a way that meets existing regulatory and internal security controls.

# Project Structure
The project is expected to grow to cover the different CSP (Cloud Service Providers) and in turn services for each of the providers. The project structure maps this hierarchy of CSPs -> Services.

Each service folder will contain markdown files (the service control description), test cases documentation and ultimately infrastructure-as-code (e.g. cloud formation or Terraform configuration). Where applicable, pre-existing source documents (e.g. docx) are stored within "archive" folders under the indidivual service folders.

## Contributing

1. Fork it (<https://github.com/finos-fdx/cloud-service-certification/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

# Tips

## Converting DocX to Markdown
After several tests, the best results to convert input .docx (Microsoft Word Documents) was achieved by using the Pandoc cmdline tool using the gfm (Github flavored Markdown) as output format. 

For example, after [installing Pandoc](https://pandoc.org/installing.html#), using [templates/archive/ServiceApprovalAccelerator_template_draftrelease.docx](templates/archive/ServiceApprovalAccelerator_template_draftrelease.docx) as input, we obtain  as output [templates/ServiceApprovalAcceleratorTemplate.md](templates/ServiceApprovalAcceleratorTemplate.md) with the following command: 

`pandoc -s ServiceApprovalAccelerator_template_draftrelease.docx -t gfm -o ServiceApprovalAcceleratorTemplate.md`

See [this thread](https://stackoverflow.com/questions/16383237/how-can-doc-docx-files-be-converted-to-markdown-or-structured-text) for reference. 

## License

Copyright 2019 JPMC

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
