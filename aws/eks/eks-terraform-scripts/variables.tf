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