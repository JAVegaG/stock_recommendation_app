resource "aws_ecs_task_definition" "app" {
  family                   = "${var.project_name}-task-${var.name}"
  cpu                      = 256
  memory                   = 512
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = var.ecs_task_execution.iam_role.arn

  container_definitions = jsonencode([
    {
      name  = "${var.project_name}-container-${var.name}",
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
          "awslogs-group"         = "awslogs-${var.project_name}",
          "awslogs-region"        = "${var.region}",
          "awslogs-stream-prefix" = "${var.project_name}-container-${var.name}"
        }
      },
    }
  ])
}

resource "aws_ecs_service" "app" {
  name            = "${var.project_name}-service-${var.name}"
  cluster         = var.cluster.id
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
    target_group_arn = aws_lb_target_group.svc.arn
    container_name   = "${var.project_name}-container-${var.name}"
    container_port   = var.container_settings.port
  }

}

resource "aws_lb_target_group" "svc" {
  name        = "${var.project_name}-${var.name}"
  port        = var.port
  protocol    = var.protocol
  vpc_id      = var.vpc_id
  target_type = "ip" # important for FARGATE

  health_check {
    path                = var.health_check.path
    protocol            = var.health_check.protocol
    port                = var.container_settings.port
    matcher             = "200"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
}

resource "aws_lb_listener_rule" "routing" {
  listener_arn = var.lb_listener.arn
  priority     = var.rule_priority

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.svc.arn
  }

  condition {
    path_pattern {
      values = [var.path_pattern]
    }
  }
}
