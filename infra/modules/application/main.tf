resource "aws_ecs_cluster" "this" {
  name = var.project_name
}


resource "aws_cloudwatch_log_group" "log" {
  name = "awslogs-${var.project_name}"
}

resource "aws_ecs_task_definition" "app" {
  family                   = "${var.project_name}-task"
  cpu                      = 256
  memory                   = 512
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn

  container_definitions = jsonencode([
    {
      name  = "${var.project_name}-container",
      image = var.container_settings.image,
      portMappings = [
        {
          containerPort = var.container_settings.port
        }
      ],
      essential = true,
      secrets   = var.container_settings.secrets,
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          "awslogs-create-group"  = "true",
          "awslogs-group"         = "${aws_cloudwatch_log_group.log.name}",
          "awslogs-region"        = "${var.region}",
          "awslogs-stream-prefix" = "${var.project_name}-container"
        }
      },
    }
  ])
}

resource "aws_ecs_service" "app" {
  name            = "${var.project_name}-service"
  cluster         = aws_ecs_cluster.this.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = var.desired_count
  launch_type     = "FARGATE"

  lifecycle {
    ignore_changes = [desired_count]
  }

  network_configuration {
    subnets          = var.public_subnets
    security_groups  = var.security_groups
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = var.lb_settings.target_group.arn
    container_name   = "${var.project_name}-container"
    container_port   = var.container_settings.port
  }

}

resource "aws_iam_role" "ecs_task_execution" {
  name = "${var.project_name}-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy" "ssm_access" {
  name = "AllowSSMParameters"
  role = aws_iam_role.ecs_task_execution.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ssm:GetParameter",
          "ssm:GetParameters",
        ],
        Resource = "arn:aws:ssm:${var.region}:${var.account_id}:parameter/*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" : "arn:aws:logs:${var.region}:${var.account_id}:log-group:awslogs-stock-app-eb978ab2-default:*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_execution_policy" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}


module "services" {
  for_each = { for idx, service in var.services : idx => service }
  source   = "./modules/services"

  project_name = var.project_name
  region       = var.region
  cluster      = aws_ecs_cluster.this
  vpc_id       = var.vpc_id
  lb_listener  = var.lb_settings.listener
  ecs_task_execution = {
    iam_role = aws_iam_role.ecs_task_execution
  }

  name          = each.value.name
  port          = each.value.port
  protocol      = each.value.protocol
  health_check  = each.value.health_check
  rule_priority = each.value.rule_priority
  path_pattern  = each.value.path_pattern

  container_settings = each.value.container_settings
  desired_count      = each.value.desired_count
  public_subnets     = each.value.public_subnets
  security_groups    = each.value.security_groups
}
