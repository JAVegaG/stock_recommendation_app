resource "aws_vpc" "main_vpc" {

  cidr_block       = var.cidr_block
  instance_tenancy = "default"

  enable_dns_hostnames = true

  tags = {
    Name = "${var.project_name}-vpc"
  }
}

module "subnets" {
  for_each = { for idx, subnet in var.subnets : idx => subnet }
  source   = "./module/vpc_subnets"

  project_name = var.project_name
  vpc_id       = aws_vpc.main_vpc.id
  region       = var.region
  subnet = {
    private_cidr_block = each.value.private_cidr_block
    public_cidr_block  = each.value.public_cidr_block
    availability_zone  = each.value.availability_zone
  }
}
