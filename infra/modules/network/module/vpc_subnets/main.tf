locals {
  availability_zone = "${var.region}${var.subnet.availability_zone}" # e.g us-east-1a
}

resource "aws_subnet" "vpc_private_subnet" {
  vpc_id            = var.vpc_id
  cidr_block        = var.subnet.private_cidr_block
  availability_zone = local.availability_zone

  tags = {
    Name = "${var.project_name}-private-subnet-${local.availability_zone}"
  }
}

resource "aws_subnet" "vpc_public_subnet" {
  vpc_id            = var.vpc_id
  cidr_block        = var.subnet.public_cidr_block
  availability_zone = local.availability_zone

  tags = {
    Name = "${var.project_name}-public-subnet-${local.availability_zone}"
  }
}

# --- Elastic IP ---

resource "aws_eip" "eip" {

  tags = {
    Name = "${var.project_name}-eip-nat"
  }

}

# --- Nat Gateway ---

resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.eip.id
  subnet_id     = aws_subnet.vpc_public_subnet.id

  tags = {
    Name = "${var.project_name}-nat-gw"
  }

}

# --- Route Tables ---

resource "aws_route_table_association" "public_subnet_routes_association" {
  subnet_id      = aws_subnet.vpc_public_subnet.id
  route_table_id = var.public_subnet.route_table_id
}

resource "aws_route_table" "private_subnet_routes" {
  vpc_id = var.vpc_id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_gateway.id
  }

  tags = {
    Name = "${var.project_name}-private-route-table"
  }
}

resource "aws_route_table_association" "private_subnet_routes_association" {
  subnet_id      = aws_subnet.vpc_private_subnet.id
  route_table_id = aws_route_table.private_subnet_routes.id
}
