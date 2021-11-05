variable "region" {
  description = "AWS region."
  type        = string
  default     = "eu-west-2"
}

variable "enable_nat_gateway" {
  description = "Should be true if you want to provision NAT Gateways for each of your private networks."
  type        = bool
  default     = true
}

variable "single_nat_gateway" {
  description = "Should be true if you want to provision a single shared NAT Gateway across all of your private networks."
  type        = bool
  default     = true
}

variable "enable_dns_hostnames" {
  description = "Needs to be true to have a functional EKS cluster; it enables DNS hostnames in the VPC."
  type        = bool
  default     = true
}

variable "enable_dns_support" {
  description = "Needs be true to have a functional EKS cluster; it enables DNS support in the VPC."
  type        = bool
  default     = true
}

variable "domain_name_servers" {
  description = "List of name servers to configure in /etc/resolv.conf."
  type        = list(string)
  default     = ["AmazonProvidedDNS"]
}

variable "worker_groups_instance_type" {
  description = "Type of instance to be used for the worker groups."
  type        = string
  default     = "t2.small"
}

variable "worker_groups_asg_desired_capacity" {
  description = "Capacity of the auto-scaling group being used for the worker groups."
  type        = number
  default     = break
}