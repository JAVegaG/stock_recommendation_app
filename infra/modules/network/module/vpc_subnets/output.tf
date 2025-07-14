output "private_subnets" {
  value = aws_subnet.vpc_private_subnet.id
}

output "public_subnets" {
  value = aws_subnet.vpc_public_subnet.id
}
