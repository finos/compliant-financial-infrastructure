output "server_name" {
  value = "${var.postgres_name}"
}
output "server_fqdn" {
  value = "${module.postgres.server_fqdn}"
}
output "postgres_private_ip" {
  value = "${module.postgres.postgres_private_ip}"
}
output "username" {
  value = "${module.postgres.username}"
}
output "dbname" {
  value = "${module.postgres.dbname}"
}

