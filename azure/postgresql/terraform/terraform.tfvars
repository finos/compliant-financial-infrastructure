### Azure PostgreSQL variables
# Name of the resource group that will be created to house postgres
postgres_resource_group_name     = "<change-me>"

# Tags for the resource group
postgres_resource_group_tags     = { "<change-me>" = "<change-me>", "<change-me>" = "<change-me>" }

# Location of the resource group
postgres_resource_group_location = "<change-me>"

# Name of the postgres server
postgres_name                    = "<change-me>"

# Location of postgres
postgres_location                = "<change-me>"

# Version of postgres
postgres_version                 = "<change-me>"

# Tags for the postgres
postgres_tags                    = { "<change-me>" = "<change-me>", "<change-me>" = "<change-me>" }

# Storage for postgres
storagemb                        = "5120"

# Number of days for backup retention
backup_retention_days            = 7

# Geo redundent postgres enablement, currently supports only: true
geo_redundent_enabled            = true

# Auto grow postgres enablement, currently supports only: true
auto_grow_enabled                = true

# Sku name for postgres, currently supports only: General and Memory optimizes tiers
sku_name                         = "GP_Gen5_2"

# Name of the database in postgres
database_name                    = "<change-me>"

# Login for server and database admin
database_login                   = "<change-me>"

# Password for server and database admin
database_password                = "<change-me>"

# Name of the virtual network that the private endpoint can use.
vnet_name                        = "<change-me>"

# Name of the resource group for the virtual network
vnet_resource_group_name         = "<change-me>"

# Name of the subnet that the private endpoint can use
subnet_name                      = "<change-me>"

#Name for the postgres private endpoint
private_endpoint_name            = "<change-me>"

# Location of private endpoint
private_endpoint_location        = "<change-me>"

#Name for the postgres private service connection
private_service_connection_name  = "<change-me>"
