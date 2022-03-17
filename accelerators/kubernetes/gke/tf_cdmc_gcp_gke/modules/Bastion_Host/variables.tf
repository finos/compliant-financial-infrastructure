variable "internal_ip_address" {
  type        = string
  description = "interal ip address"
}


variable "bastion_subnet_name" {
  type        = string
  description = "Subnet name where Bastion Host will be deployed"
}

variable "vm_name" {
  type        = string
  description = "the VM machine name"
}

variable "machine_type" {
  type        = string
  description = "the VM machine type"
}

variable "region" {
  type        = string
  description = "the VM machine region"
}

variable "zone" {
  type        = string
  description = "the VM machine zone"
}

variable "machine_image" {
  type        = string
  description = "The machine image"
}

variable "project_id" {
  type        = string
  description = "Your project id in GCP"
}

variable "tags" {
  type        = list(string)
  description = "Your tags"
}

variable "static-internal-ip-name" {
  type        = string
  description = "Static internal ip name"
}
