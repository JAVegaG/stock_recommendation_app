variable "project_name" {
  description = "Name of the project used for naming resources"
  type        = string
}

variable "name" {
  description = "Name of the specific service or path-based route (e.g., svc-a)"
  type        = string
}

variable "port" {
  description = "Port on which the target group receives traffic"
  type        = number
  default     = 80
}

variable "protocol" {
  description = "Protocol used by the target group (e.g., HTTP)"
  type        = string
  default     = "HTTP"
}

variable "vpc_id" {
  description = "The VPC ID where the target group resides"
  type        = string
}

variable "health_check" {
  description = "Health check configuration for the target group"
  type = object({
    path     = string
    protocol = string
  })
}

variable "lb_listener" {
  description = "Listener resource (must contain at least the ARN)"
  type = object({
    arn = string
  })
}

variable "rule_priority" {
  description = "Unique priority number for the listener rule"
  type        = number
}

variable "path_pattern" {
  description = "Path pattern to route requests (e.g., /svcA*)"
  type        = string
}

variable "desired_count" {
  description = "Desired units count"
  type        = number
}

variable "public_subnets" {
  description = "Subnets associated with the task or service"
  type        = set(string)
}

variable "security_groups" {
  description = "Security groups associated with the task or service"
  type        = set(string)
}

variable "container_settings" {
  description = "Container settings (image, port)"
  type = object({
    image = string
    port  = number
  })
}
