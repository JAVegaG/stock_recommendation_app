variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "region" {
  description = "Region for the vpc"
  type        = string
}

variable "cidr_block" {
  description = "CIDR block for vpc. e.g 10.0.0.0/20"
  type        = string
}

variable "subnets" {
  description = "List of subnets (private/public) with AZs"
  type = list(object({
    private_cidr_block = string
    public_cidr_block  = string
    availability_zone  = string
  }))
}

variable "security_group" {
  description = "Security group configuration"
  type = object({
    name        = string
    description = string
    ingress_rules = list(object({
      from_port   = number
      to_port     = number
      ip_protocol = string
    }))
    egress_rules = list(object({
      cidr_ipv4   = string
      from_port   = number
      to_port     = number
      ip_protocol = string
    }))
  })
}
