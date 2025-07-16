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
