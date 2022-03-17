variable "project_id" {
  type        = string
  description = "Project ID"
}

variable "service_account_name" {
  type        = string
  description = "name of k8s nodes' service account"
}

variable "region" {
  type        = string
  description = "Regional location of the GKE Cluster."
}

variable "cluster_name" {
  type        = string
  description = "Cluster name for the GCP Cluster."
}
variable "vpc_name" {
  type        = string
  description = "Name of the VPC to deploy GKE cluster on."
}
variable "subnet_name" {
  type        = string
  description = "Name of the subnet to deploy GKE cluster on."
}

variable "machine_type" {
  type        = string
  description = "Defines node configuration"
}

variable "bastion_ip" {
  type        = string
  description = "Internal IP address of the Bastion Host VM used to connect to GKE cluster"
}

variable "master_cidr" {
  type        = string
  description = "CIDR range of GKE Master"
}

variable "cluster_ipv4_cidr_block" {
  type        = string
  description = "CIDR range of pods"
}

variable "encryption_key_name" {
  type        = string
  description = "Name of the encryption key for ETCD"
}

