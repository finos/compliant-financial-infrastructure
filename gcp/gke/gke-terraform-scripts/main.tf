provider "google" {
  credentials = file(var.cred_url)
  project     = var.project_id
  region      = var.region
}

// VPC and firewall creation 

resource "google_compute_network" "gke-vpc" {
  name                    = var.gke-vpc.name
  auto_create_subnetworks = "false"
}

resource "google_compute_firewall" "allow-iap" {
  name    = "allow-iap"
  network = google_compute_network.gke-vpc.name


  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["35.235.240.0/20"]
  target_tags   = ["allow-iap"]
}


//Subnets creation for gke-vpc

module "admin-subnet" {
  source        = "./modules/subnetwork"
  name          = var.gke-vpc.subnets[0].name
  description   = var.gke-vpc.subnets[0].description
  ip_cidr_range = var.gke-vpc.subnets[0].ip_cidr_range
  region        = var.gke-vpc.subnets[0].region
  network       = google_compute_network.gke-vpc.self_link
}

module "k8s-subnet" {
  source        = "./modules/subnetwork"
  name          = var.gke-vpc.subnets[1].name
  description   = var.gke-vpc.subnets[1].description
  ip_cidr_range = var.gke-vpc.subnets[1].ip_cidr_range
  region        = var.gke-vpc.subnets[1].region
  network       = google_compute_network.gke-vpc.self_link
}

//bastion host for gke-vpc

module "bastion-host" {
  source                  = "./modules/Bastion_Host"
  project_id              = var.project_id
  bastion_subnet_name     = var.gke-vpc.subnets[0].name
  internal_ip_address     = var.bastion_host.internal_ip_address
  vm_name                 = var.bastion_host.vm_name
  machine_type            = var.bastion_host.machine_type
  zone                    = var.bastion_host.zone
  machine_image           = var.bastion_host.machine_image
  tags                    = var.bastion_host.tags
  static-internal-ip-name = "bastion-ip"
  region                  = var.bastion_host.region
}

//GKE Cluster 

module "gke-cluster" {
  source                  = "./modules/kubernetes"
  project_id              = var.project_id
  service_account_name    = var.gke-cluster.service_account_name
  region                  = var.gke-cluster.region
  cluster_name            = var.gke-cluster.cluster_name
  vpc_name                = google_compute_network.gke-vpc.name
  subnet_name             = var.gke-vpc.subnets[1].name
  bastion_ip              = "${var.bastion_host.internal_ip_address}/32"
  master_cidr             = var.gke-cluster.master_cidr
  cluster_ipv4_cidr_block = var.gke-cluster.cluster_ipv4_cidr_block
  machine_type            = var.gke-cluster.machine_type
  encryption_key_name     = var.encryption_key_name
}


