### PostgreSQL server variables

#Resource group name for potgresql server
variable "postgres_resource_group_name" {
  type        = string
  description = "Name of resource group to hold postgres"
}

#Resource group tags for potgresql server
variable "postgres_resource_group_tags" {
  type        = map(any)
  description = "Tags to be added to the postgres resource group"
}

#Resource group location for potgresql server
variable "postgres_resource_group_location" {
  type        = string
  description = "The location of the resource group which will hold postgres"
}

#Name for postgres server
variable "postgres_name" {
  type        = string
  description = "Name for postgres server"
}

#Location of postgres server
variable "postgres_location" {
  type        = string
  description = "Location of the resource group"
}

#Version for postgres server
variable "postgres_version" {
  type        = string
  description = "Version of postgres"
}

variable "postgres_tags" {
  type = map(any)
}

#Database storage in megabytes
variable "storagemb" {
  type        = string
  description = "Amount of storage for the database, in mb."
  default     = "5120"
}

#Number of days to retain backups
variable "backup_retention_days" {
  type    = number
  default = 7
}

#Enable geo redundancy for Azure postgres
variable "geo_redundent_enabled" {
  type        = bool
  description = "Geo redundent postgres enablement, currently supports only: true"
  default     = true
}

#Enable auto grow for Azure postgres
variable "auto_grow_enabled" {
  type        = bool
  default     = true
  description = "Auto grow postgres enablement, currently supports only: true"
}

#Postgres sku name (GP_Gen5_2)
variable "sku_name" {
  type        = string
  description = "Sku name for postgres, currently supports only: General and Memory optimizes tiers"
  default     = "GP_Gen5_2"
}

#Name for database in postgres
variable "database_name" {
  type        = string
  description = "Name for database in postgres"
}

#Login for the database
variable "database_login" {
  type        = string
  description = "Login for server and database admin"
}

#Password for the database
variable "database_password" {
  type        = string
  description = "Password for the database in postgres"
}

### Network variables

#Target vnet name
variable "vnet_name" {
  type        = string
  description = "Name of the virtual network that the private endpoint can use. "
}

#Target resource group for vnet
variable "vnet_resource_group_name" {
  type        = string
  description = "Name of the resource group for the virtual network"
}

variable "subnet_name" {
  type        = string
  description = "Name of the subnet that the private endpoint can use"
}

#Location for private endpoint
variable "private_endpoint_location" {
  type        = string
  description = "Location of private endpoint"
}

#Private endpoint name
variable "private_endpoint_name" {
  type        = string
  description = "Name for the postgres private endpoint"
}

#Private service connection name
variable "private_service_connection_name" {
  type        = string
  description = "Name for the postgres private service connection"
}