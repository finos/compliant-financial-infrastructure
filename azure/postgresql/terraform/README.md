# README

This terraform script provisions:
+ azure postresql Single Server with a database
+ private endpoint for postgresql server

## How to use

All variable are set in [terraform.tfvars](terraform.tfvars) these should be changed for each application team.

Variables that would need to change are marked as `<change-me>`.

More detail on each variable can be found in the `Inputs` section below

## How do I connect to my new Azure PostgreSQL server?
As we are using a private endpoint with the postgres server to securely connect from on-prem, the fully qualified domain name will not be reachable until a dns record is set for it. (This is the same process as for the storage account).

Meanwhile the private ip can be used to connect. The private ip is an output from the terraform module, but it can also be looked up in th portal under the private endpoint resource.

To connect you can use your favorite IDE, SQL tool or command line.

### Oracle SQL Developer with PostgreSQL driver

	Username: <username>@<postgres-server-name>
	Password: <password>
	Hostname: <postgres-private-ip>:5432/<database-name>?sslmode=require&ssl=true 
	Port: THIS IS EMPTY

## Providers

| Name | Version |
|------|---------|
| azurerm | 2.7.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| auto\_grow\_enabled | Auto grow postgres enablement, currently supports only: true | `bool` | `true` | no |
| backup\_retention\_days | Number of days to retain backups | `number` | `7` | no |
| database\_login | Login for server and database admin | `string` | n/a | yes |
| database\_name | Name for database in postgres | `string` | n/a | yes |
| database\_password | Password for the database in postgres | `string` | n/a | yes |
| geo\_redundent\_enabled | Geo redundent postgres enablement, currently supports only: true | `bool` | `true` | no |
| postgres\_location | Location of the resource group | `string` | n/a | yes |
| postgres\_name | Name for postgres server | `string` | n/a | yes |
| postgres\_resource\_group\_location | The location of the resource group which will hold postgres | `string` | n/a | yes |
| postgres\_resource\_group\_name | Name of resource group to hold postgres | `string` | n/a | yes |
| postgres\_resource\_group\_tags | Tags to be added to the postgres resource group | `map` | n/a | yes |
| postgres\_tags | n/a | `map` | n/a | yes |
| postgres\_version | Version of porstgres | `string` | n/a | yes |
| private\_endpoint\_location | Location of private endpoint | `string` | n/a | yes |
| private\_endpoint\_name | Name for the postgres private endpoint | `string` | n/a | yes |
| private\_service\_connection\_name | Name for the postgres private service connection | `string` | n/a | yes |
| sku\_name | Sku name for postgres, currently supports only: General and Memory optimizes tiers | `string` | `"GP_Gen5_2"` | no |
| storagemb | Amount of storage for the database, in mb. | `string` | `"5120"` | no |
| subnet\_name | Name of the subnet that the private endpoint can use | `string` | n/a | yes |
| vnet\_name | Name of the virtual network that the private endpoint can use. | `string` | n/a | yes |
| vnet\_resource\_group\_name | Name of the resource group for the virtual network | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| dbname | n/a |
| postgres\_private\_ip | n/a |
| server\_fqdn | The fully qualified domain name (FQDN) of the PostgreSQL server |
| username | n/a |

