variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "region" {
  type = string
}

variable "account_id" {
  type = string
}

variable "container_settings" {
  description = "Container settings (image, port)"
  type = object({
    image = string
    port  = number
    secrets = list(object({
      name      = string
      valueFrom = string
    }))
  })
}

variable "health_check" {
  description = "Health check settings"
  type = object({
    path     = string
    protocol = string
  })
  default = {
    path = "/"
  protocol = "HTTP" }
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
}


variable "services" {
  description = "Services to be deployed"
  type = list(object({
    name     = string
    port     = number
    protocol = string
    health_check = object({
      path     = string
      protocol = string
    })
    rule_priority = number
    path_pattern  = string
    container_settings = object({
      image = string
      port  = number
      secrets = list(object({
        name      = string
        valueFrom = string
      }))
    })
    desired_count   = number
    public_subnets  = set(string)
    security_groups = set(string)
  }))
}

variable "lb_settings" {
  description = "Load Balancer Settings"
  type = object({
    target_group = object({
      arn = string
    })
    listener = object({
      arn = string
    })
  })
}
