#Call the module from the modules folder
module "postgres" {
  source = "./modules/postgres"

  postgres_resource_group_name     = var.postgres_resource_group_name
  postgres_resource_group_tags     = var.postgres_resource_group_tags
  postgres_resource_group_location = var.postgres_resource_group_location
  postgres_name                    = var.postgres_name
  postgres_location                = var.postgres_location
  postgres_version                 = var.postgres_version
  postgres_tags                    = var.postgres_tags
  storagemb                        = var.storagemb
  backup_retention_days            = var.backup_retention_days
  geo_redundent_enabled            = var.geo_redundent_enabled
  auto_grow_enabled                = var.auto_grow_enabled
  sku_name                         = var.sku_name
  database_name                    = var.database_name
  database_login                   = var.database_login
  database_password                = var.database_password
  vnet_name                        = var.vnet_name
  vnet_resource_group_name         = var.vnet_resource_group_name
  subnet_name                      = var.subnet_name
  private_endpoint_name            = var.private_endpoint_name
  private_endpoint_location        = var.private_endpoint_location
  private_service_connection_name  = var.private_service_connection_name

}
