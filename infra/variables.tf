variable "aws_region" {
  default = "us-east-1"
  type    = string
}

variable "aws_profile" {
  default = "default"
  type    = string
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
