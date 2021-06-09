variable "name" {
  type        = string
  description = "Name of subnet."
}

variable "description" {
  type        = string
  description = "Description usage of subnet"
  default     = ""
}

variable "ip_cidr_range" {
  type        = string
  description = "IP range in CIDR format of the subnet"
}

variable "network" {
  type        = string
  description = "Name of self-link to the VPC this subnet will be linked to. Defaults to 'default' network"
}

variable "region" {
  type        = string
  description = "Region in which subnet will be created. Defaults to the region in the terraform provider"
  default     = ""
}

variable "create_secondary_ranges" {
  default     = false
  description = "Enable secondary ip ranges to be used with 'secondary_ranges' variable"
}

/*
* Parameters authorized: 
* - range_name    
* - ip_cidr_range 
*/
variable "secondary_ranges" {
  type        = list(any)
  default     = []
  description = "Create up to 5 alternative CIDR range to represent this subnetwork"
}