#Create Resource group for the postgres server
resource "azurerm_resource_group" "postgres-rg" {
  name     = var.postgres_resource_group_name
  location = var.postgres_resource_group_location
  tags     = var.postgres_resource_group_tags
}

#Create a Postgres server
resource "azurerm_postgresql_server" "postgres-server" {
  name                             = var.postgres_name
  resource_group_name              = azurerm_resource_group.postgres-rg.name
  location                         = var.postgres_location
  version                          = var.postgres_version
  administrator_login              = var.database_login
  administrator_login_password     = var.database_password

  sku_name                         = var.sku_name

  storage_mb                       = var.storagemb
  backup_retention_days            = var.backup_retention_days
  geo_redundant_backup_enabled     = var.geo_redundent_enabled
  auto_grow_enabled                = var.auto_grow_enabled
  public_network_access_enabled    = false
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"

  tags                             = var.postgres_tags

}

#Add a database to the server
resource "azurerm_postgresql_database" "postgres-db" {
  name                = var.database_name
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
}

#Configure postgresql
resource "azurerm_postgresql_configuration" "log_connections" {
  name                = "log_connections"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "log_disconnections" {
  name                = "log_disconnections"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "log_duration" {
  name                = "log_duration"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "log_checkpoints" {
  name                = "log_checkpoints"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "connection_throttling" {
  name                = "connection_throttling"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "backslash_quote" {
  name                = "backslash_quote"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "SAFE_ENCODING"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "client_min_messages" {
  name                = "client_min_messages"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "NOTICE"
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "log_retention_days" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = 7
}
#Configure postgresql
resource "azurerm_postgresql_configuration" "logging_collector" {
  name                = "logging_collector"
  resource_group_name = azurerm_resource_group.postgres-rg.name
  server_name         = azurerm_postgresql_server.postgres-server.name
  value               = "on"
}