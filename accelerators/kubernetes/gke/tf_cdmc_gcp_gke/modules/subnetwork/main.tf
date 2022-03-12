#REFER THE https://github.com/idwall/terraform-google-network-subnet
/******************************************
  Create subnetwork without secondary_range
 *****************************************/
resource "google_compute_subnetwork" "basic" {

  name                     = var.name
  description              = var.description
  region                   = var.region
  ip_cidr_range            = var.ip_cidr_range
  network                  = var.network
  private_ip_google_access = true
}

output "bastion_subnet_name" {
  value = google_compute_subnetwork.basic.name
}

