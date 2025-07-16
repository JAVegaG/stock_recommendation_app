variable "aws_region" {
  default = "us-east-1"
  type    = string
}

variable "aws_profile" {
  default = "default"
  type    = string
}

variable "aws_account_id" {
  type = string
}

variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "cockroachdb_cluter_id" {
  description = "CockroachDB cluster id"
  type        = string
}

variable "cockroachdb_database" {
  description = "CockroachDB database name"
  type        = string
}

variable "container_settings" {
  description = "General container settings"
  type = object({
    image = string
    port  = number
  })
}

variable "services" {
  description = "List of services to be deployed"
  type = list(object({
    container_settings = object({
      image = string
      port  = number
    })
  }))
}
