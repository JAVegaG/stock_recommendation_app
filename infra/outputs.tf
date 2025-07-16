output "lb_dns_name" {
  description = "DNS name of the load balancer"
  value       = module.application.lb_dns_name
}
