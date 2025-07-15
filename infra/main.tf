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
}


module "application" {
  source       = "./modules/application"
  project_name = local.resource_suffix
  vpc_id       = module.network.vpc_id
  container_settings = {
    image = "",
    port  = 8080
  }
  desired_count  = 2
  public_subnets = toset([for s in module.network.public_subnets : s.value])
}
