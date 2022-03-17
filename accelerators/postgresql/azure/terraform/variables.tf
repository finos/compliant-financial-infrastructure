### PostgreSQL server variables

#Resource group name for potgresql server
variable "postgres_resource_group_name" {
  type = string
}
#Resource group tags for potgresql server
variable "postgres_resource_group_tags" {
  type = map(any)
}
#Resource group location for potgresql server
variable "postgres_resource_group_location" {
  type = string
}
#Name for postgres server
variable "postgres_name" {
  type = string
}
#Location of postgres server
variable "postgres_location" {
  type = string
}
#Version for postgres server
variable "postgres_version" {
  type = string
}
variable "postgres_tags" {
  type = map(any)
}
#Database storage in megabytes
variable "storagemb" {
  type = string
}
#Number of days to retain backups
variable "backup_retention_days" {
  type = number
}
#Enable geo redundency for Azure postgres
variable "geo_redundent_enabled" {
  type = bool
}
#Enable auto grow for Azure postgres
variable "auto_grow_enabled" {
  type = bool
}
#Postres sku name (GP_Gen5_2)
variable "sku_name" {
  type = string
}
#Name for database in postgres
variable "database_name" {
  type = string
}
#Login for the database
variable "database_login" {
  type = string
}
#Password for the database
variable "database_password" {
  type = string
}

### Network variables

#Target vnet name
variable "vnet_name" {
  type = string
}
#Target resource group for vnet
variable "vnet_resource_group_name" {
  type = string
}
variable "subnet_name" {
  type = string
}
#Location for privete endpoint
variable "private_endpoint_location" {
  type = string
}
#Private endpoint name
variable "private_endpoint_name" {
  type = string
}
#Private service connection name
variable "private_service_connection_name" {
  type = string
}