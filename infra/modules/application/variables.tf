variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "container_settings" {
  description = "Container settings (image, port)"
  type = object({
    image = string
    port  = number
  })
}

variable "desired_count" {
  description = "Desired units count"
  type        = number
}

variable "vpc_id" {
  description = "main vpc id"
  type        = string
}

variable "public_subnets" {
  description = "Subnets associated with the task or service"
  type        = set(string)
}

variable "security_groups" {
  description = "Security groups associated with the task or service"
  type        = set(string)
  default     = null
}
