module "database" {
  source        = "./modules/database"
  database_name = var.cockroachdb_database
  cluster_id    = var.cockroachdb_cluter_id
}

module "network" {
  source       = "./modules/network"
  project_name = local.resource_suffix
  region       = var.aws_region
  cidr_block   = "10.0.0.0/20"
  subnets = [{
    private_cidr_block = "10.0.10.0/24",
    public_cidr_block  = "10.0.11.0/24",
    availability_zone  = "a"
    }, {
    private_cidr_block = "10.0.12.0/24",
    public_cidr_block  = "10.0.13.0/24",
    availability_zone  = "b"
  }]

  security_group = {
    ingress_rules = [{
      from_port   = 80
      to_port     = 80
      ip_protocol = "tcp"
    }]

    egress_rules = [{
      cidr_ipv4   = "0.0.0.0/0"
      ip_protocol = "-1"
    }]
  }
}


module "application" {
  source       = "./modules/application"
  project_name = local.resource_suffix
  vpc_id       = module.network.vpc_id
  container_settings = {
    image = var.container_settings.image,
    port  = var.container_settings.port
  }

  desired_count   = 2
  public_subnets  = module.network.public_subnets
  security_groups = module.network.security_groups
  services = [{
    name     = "api"
    port     = 80
    protocol = "HTTP"

    desired_count = 2

    health_check = {
      path     = "/health"
      protocol = "HTTP"
    }

    rule_priority = 10
    path_pattern  = "api*"

    container_settings = {
      image = var.services.container_settings.image
      port  = var.services.container_settings.port
    }

    public_subnets  = module.network.public_subnets
    security_groups = module.network.security_groups
  }]
}
