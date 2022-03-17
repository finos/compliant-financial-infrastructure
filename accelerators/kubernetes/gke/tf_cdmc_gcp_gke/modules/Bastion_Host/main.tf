resource "google_compute_address" "bastion_internal_ip" {
  name         = var.static-internal-ip-name
  address_type = "INTERNAL"
  address      = var.internal_ip_address
  subnetwork   = var.bastion_subnet_name
  region       = var.region
}

resource "google_compute_instance" "default" {
  name         = var.vm_name
  machine_type = var.machine_type
  zone         = var.zone
  boot_disk {
    initialize_params {
      image = var.machine_image
    }
  }

  network_interface {
    subnetwork         = var.bastion_subnet_name
    subnetwork_project = var.project_id
    network_ip         = google_compute_address.bastion_internal_ip.address
  }

  metadata = {
    startup-script = "curl -LO \"https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl\" && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl"
  }

  service_account {
    email  = "<bastion-service-account-id>"
    scopes = ["logging-write", "monitoring-write", "storage-ro", "service-management", "service-control", "cloud-platform"]
  }

  tags = var.tags

}
