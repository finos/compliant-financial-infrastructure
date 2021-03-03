resource "google_container_cluster" "gke-cluster" {
  project    = var.project_id
  name       = var.cluster_name
  location   = var.region
  subnetwork = var.subnet_name

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1
  enable_shielded_nodes    = true

  ip_allocation_policy {
    cluster_ipv4_cidr_block = var.cluster_ipv4_cidr_block
  }

  master_auth {
    client_certificate_config {
      issue_client_certificate = false
    }
  }

  master_authorized_networks_config {
    cidr_blocks {
      cidr_block   = var.bastion_ip
      display_name = "Bastion Host IP"
    }
  }

  min_master_version = "1.16.13-gke.401"

  network = var.vpc_name

  network_policy {
    enabled = true
  }

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = true
    master_ipv4_cidr_block  = var.master_cidr
  }



  vertical_pod_autoscaling {
    enabled = true
  }

  workload_identity_config {
    identity_namespace = "${var.project_id}.svc.id.goog"
  }

  database_encryption {
    state    = "ENCRYPTED"
    key_name = var.encryption_key_name
  }
}

resource "google_container_node_pool" "node-pool-1" {
  name     = "app-node-pool"
  location = google_container_cluster.gke-cluster.location
  cluster  = google_container_cluster.gke-cluster.name
  autoscaling {
    min_node_count = 1
    max_node_count = 3
  }

  initial_node_count = 1

  management {
    auto_repair  = true
    auto_upgrade = true
  }

  upgrade_settings {
    max_surge       = 1
    max_unavailable = 0
  }

  node_config {
    preemptible     = false
    machine_type    = var.machine_type
    disk_size_gb    = 20
    image_type      = "COS_CONTAINERD"
    service_account = var.service_account_name

    shielded_instance_config {
      enable_secure_boot          = true
      enable_integrity_monitoring = true
    }

    metadata = {
      disable-legacy-endpoints = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
