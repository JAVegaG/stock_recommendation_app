resource "aws_vpc" "main_vpc" {

  cidr_block       = var.cidr_block
  instance_tenancy = "default"

  enable_dns_hostnames = true

  tags = {
    Name = "${var.project_name}-vpc"
  }
}

# --- Internet Gateway ---

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main_vpc.id

  tags = {
    Name = "${var.project_name}-igw"
  }
}

# --- Route Tables ---

resource "aws_route_table" "public_subnet_routes" {
  vpc_id = aws_vpc.main_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "${var.project_name}-public-route-table"
  }
}

# --- security group ---

resource "aws_security_group" "sg" {
  name        = var.security_group.name
  description = var.security_group.description
  vpc_id      = aws_vpc.main_vpc.id

  tags = {
    Name = var.security_group.name
  }
}

resource "aws_vpc_security_group_ingress_rule" "sg_ingress_rule" {
  for_each          = { for idx, ingress_rule in var.security_group.ingress_rules : idx => ingress_rule }
  security_group_id = aws_security_group.sg.id
  cidr_ipv4         = aws_vpc.main_vpc.cidr_block
  from_port         = each.value.from_port
  ip_protocol       = each.value.ip_protocol
  to_port           = each.value.to_port
}

resource "aws_vpc_security_group_egress_rule" "sg_egress_rule" {
  for_each          = { for idx, egress_rule in var.security_group.egress_rules : idx => egress_rule }
  security_group_id = aws_security_group.sg.id
  cidr_ipv4         = each.value.cidr_block
  from_port         = each.value.from_port
  ip_protocol       = each.value.ip_protocol
  to_port           = each.value.to_port
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

  public_subnet = {
    route_table_id = aws_route_table.public_subnet_routes.id
  }

  depends_on = [aws_internet_gateway.igw]

}
