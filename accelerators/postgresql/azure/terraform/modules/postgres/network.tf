# Get vnet
data "azurerm_virtual_network" "vnet" {
  name                = var.vnet_name
  resource_group_name = var.vnet_resource_group_name
}

# Get subnet
data "azurerm_subnet" "subnet" {
  name                 = var.subnet_name
  virtual_network_name = var.vnet_name
  resource_group_name  = var.vnet_resource_group_name
}

# Create private endpoint
resource "azurerm_private_endpoint" "postgres-endpoint" {
  name                = var.private_endpoint_name
  location            = var.private_endpoint_location
  resource_group_name = azurerm_resource_group.postgres-rg.name
  subnet_id           = data.azurerm_subnet.subnet.id

  private_service_connection {
    name                           = var.private_service_connection_name
    private_connection_resource_id = azurerm_postgresql_server.postgres-server.id
    is_manual_connection           = false
    subresource_names              = ["postgresqlServer"]
  }

  depends_on = [azurerm_postgresql_server.postgres-server]
}