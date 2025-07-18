output "lb_dns_name" {
  description = "DNS name of the load balancer"
  value       = aws_lb.app.dns_name
}

output "lb_protocol" {
  description = "protocol of the load balancer"
  value       = aws_lb_listener.http.protocol
}

output "lb_port" {
  description = "port of the load balancer"
  value       = aws_lb_listener.http.port
}
