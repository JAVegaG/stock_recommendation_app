data "aws_ssm_parameter" "port" {
  name = "/port"
}

data "aws_ssm_parameter" "seed_api" {
  name = "/seed_api"
}

data "aws_ssm_parameter" "api_url" {
  name = "/api_url"
}

data "aws_ssm_parameter" "api_auth_token" {
  name = "/api_auth_token"
}

data "aws_ssm_parameter" "db_host" {
  name = "/db_host"
}

data "aws_ssm_parameter" "db_port" {
  name = "/db_port"
}

data "aws_ssm_parameter" "db_user" {
  name = "/db_user"
}

data "aws_ssm_parameter" "db_password" {
  name = "/db_password"
}

data "aws_ssm_parameter" "db_ssl" {
  name = "/db_ssl"
}

resource "aws_ssm_parameter" "app_base_url" {
  name  = "/app_base_url"
  type  = "SecureString"
  value = "${aws_lb_listener.http.protocol}://${aws_lb.app.dns_name}:${aws_lb_listener.http.port}"
}

resource "aws_ssm_parameter" "db_name" {
  name  = "/db_name"
  type  = "SecureString"
  value = module.database.database_name
}
