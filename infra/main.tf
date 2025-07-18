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
    name        = "main-sg"
    description = "Main security group for web app. Allows HTTP connections"
    ingress_rules = [{
      from_port   = 80
      to_port     = 80
      ip_protocol = "tcp"
      cidr_block  = "0.0.0.0/0"
    }]

    egress_rules = [{
      cidr_ipv4   = "0.0.0.0/0"
      ip_protocol = "-1"
    }]
  }
}

locals {
  security_group = [module.network.security_groups.id]
}

resource "aws_lb" "app" {
  name               = "${var.project_name}-lb"
  internal           = false
  load_balancer_type = "application"
  subnets            = module.network.public_subnets
  security_groups    = local.security_group
}

resource "aws_lb_target_group" "app" {
  name        = "${var.project_name}-tg"
  port        = var.container_settings.port
  protocol    = "HTTP"
  vpc_id      = module.network.vpc_id
  target_type = "ip"

  health_check {
    path                = "/"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.app.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.app.arn
  }

}

module "application" {
  source       = "./modules/application"
  region       = var.aws_region
  account_id   = var.aws_account_id
  project_name = local.resource_suffix
  vpc_id       = module.network.vpc_id

  lb_settings = {
    target_group = aws_lb_target_group.app
    listener     = aws_lb_listener.http
  }

  container_settings = {
    image = var.container_settings.image,
    port  = var.container_settings.port,
    secrets = [{
      name      = "VITE_API_URL",
      valueFrom = aws_ssm_parameter.app_base_url.arn
    }]
  }

  desired_count   = 2
  public_subnets  = module.network.public_subnets
  security_groups = local.security_group
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
    path_pattern  = "/api/*"

    container_settings = {
      image = var.service.container_settings.image
      port  = data.aws_ssm_parameter.port.value
      secrets = [
        {
          name      = "PORT"
          valueFrom = data.aws_ssm_parameter.port.arn
        },
        {
          name      = "SEED_API"
          valueFrom = data.aws_ssm_parameter.seed_api.arn
        },
        {
          name      = "CORS_ORIGIN"
          valueFrom = aws_ssm_parameter.app_base_url.arn
        },
        {
          name      = "API_URL"
          valueFrom = data.aws_ssm_parameter.api_url.arn
        },
        {
          name      = "API_AUTH_TOKEN",
          valueFrom = data.aws_ssm_parameter.api_auth_token.arn
        },
        {
          name      = "DB_HOST",
          valueFrom = data.aws_ssm_parameter.db_host.arn
        },
        {
          name      = "DB_PORT",
          valueFrom = data.aws_ssm_parameter.db_port.arn
        },
        {
          name      = "DB_USER",
          valueFrom = data.aws_ssm_parameter.db_user.arn
        },
        {
          name      = "DB_PASSWORD",
          valueFrom = data.aws_ssm_parameter.db_password.arn
        },
        {
          name      = "DB_SSL",
          valueFrom = data.aws_ssm_parameter.db_ssl.arn
        },
        {
          name      = "DB_NAME",
          valueFrom = aws_ssm_parameter.db_name.arn
        }
      ]
    }

    public_subnets  = module.network.public_subnets
    security_groups = local.security_group
  }]
}
