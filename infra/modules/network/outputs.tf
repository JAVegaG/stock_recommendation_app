output "vpc_id" {
  value = aws_vpc.main_vpc.id
}

output "private_subnets" {
  value = module.subnets.private_subnets
}

output "public_subnets" {
  value = module.subnets.public_subnets
}
