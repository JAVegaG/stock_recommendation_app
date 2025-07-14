variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "vpc_id" {
  description = "main vpc id"
  type        = string
}

variable "region" {

}

variable "subnet" {
  description = "Subnet basic settings"
  type = object({
    private_cidr_block = string
    public_cidr_block  = string
    availability_zone  = string
  })
}
