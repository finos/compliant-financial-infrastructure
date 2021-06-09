variable "cred_url" {
  type        = string
  description = "Your service account full URL"
  default     = "<path to json key credentials of service account that tf uses>"
}

variable "service_account_name" {
  type        = string
  description = "The service account name"
  default     = "<ID of K8s nodes service account>"
}

variable "project_id" {
  type        = string
  description = "Your project id in GCP"
  default     = "<project-id>"
}


variable "region" {
  type        = string
  description = "The region of the project resources in GCP"
  default     = "europe-west2"
}

variable "zone" {
  type        = string
  description = "The zone of the project resources in GCP"
  default     = "a"
}

variable "bastion_subnet_name" {
  type        = string
  description = "name of the subnet to deploy bastion host on"
  default     = "admin-subnet"
}


variable "gke-vpc" {
  type = object({
    name = string
    subnets = list(object({
      name          = string
      description   = string
      ip_cidr_range = string
      region        = string
    }))
  })
  description = "The name of the production VPC"
  default = {
    name = "gke-vpc"
    subnets = [
      {
        name          = "admin-subnet"
        description   = "Subnet for bastion host and other administrative VMs"
        ip_cidr_range = "10.0.1.0/24"
        region        = "europe-west2"
      },
      {
        name          = "k8s-nodes-subnet"
        description   = "Subnet for GKE nodes "
        ip_cidr_range = "10.0.2.0/24"
        region        = "europe-west2"
      }
    ]
  }
}


variable "bastion_host" {
  type = object({
    internal_ip_address = string
    vm_name             = string
    machine_type        = string
    zone                = string
    machine_image       = string
    tags                = list(string)
    bastion_subnet_name = string
    region              = string
  })
  description = "The Bastion host config for production"
  default = {
    internal_ip_address = "10.0.1.2"
    vm_name             = "prod-bastion-host"
    machine_type        = "n1-standard-1"
    zone                = "europe-west2-a"
    machine_image       = "ubuntu-1604-lts"
    tags                = ["allow-iap"]
    bastion_subnet_name = "admin-subnet"
    region              = "europe-west2"
  }
}

variable "master_cidr" {
  type        = string
  description = "CIDR block address of GKE master."
  default     = "172.16.0.0/28"
}

variable "cluster_name" {
  type        = string
  description = "Cluster name for the GCP Cluster."
  default     = "gke-cluster"
}
//GKE CLUSTERS..

variable "gke-cluster" {
  type = object({
    region                  = string
    cluster_name            = string
    master_cidr             = string
    cluster_ipv4_cidr_block = string
    service_account_name    = string
    machine_type            = string
  })
  description = "The GKE app cluster for production"
  default = {
    region                  = "europe-west2"
    cluster_name            = "gke-cluster"
    master_cidr             = "172.168.10.0/28"
    cluster_ipv4_cidr_block = "10.1.0.0/16"
    service_account_name    = "<k8s-nodes-service-account-id>"
    machine_type            = "e2-standard-4"
  }
}

variable "encryption_key_name" {
  type        = string
  description = "Name of the encryption key for ETCD"
  default     = "<encryption-key-id>"
}
