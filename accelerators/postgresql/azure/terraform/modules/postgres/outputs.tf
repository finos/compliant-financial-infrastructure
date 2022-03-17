output "server_fqdn" {
  description = "The fully qualified domain name (FQDN) of the PostgreSQL server"
  value       = azurerm_postgresql_server.postgres-server.fqdn
}

output "postgres_private_ip" {
  value = azurerm_private_endpoint.postgres-endpoint.private_service_connection[0].private_ip_address
}

output "username" {
  value = azurerm_postgresql_server.postgres-server.administrator_login
}

output "dbname" {
  value = azurerm_postgresql_database.postgres-db.name
}
