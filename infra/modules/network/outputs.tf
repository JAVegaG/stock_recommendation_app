output "vpc_id" {
  value = aws_vpc.main_vpc.id
}

output "private_subnets" {
  value = [for subnet in module.subnets : subnet.private_subnets.id]
}

output "public_subnets" {
  value = [for subnet in module.subnets : subnet.public_subnets.id]
}

output "security_groups" {
  value = [for sg in aws_security_group.sg : sg.id]
}
