ephemeral "aws_ssm_parameter" "port" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/port"
}

ephemeral "aws_ssm_parameter" "seed_api" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/seed_api"
}

ephemeral "aws_ssm_parameter" "api_url" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/api_url"
}

ephemeral "aws_ssm_parameter" "api_auth_token" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/api_auth_token"
}

ephemeral "aws_ssm_parameter" "db_host" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/db_host"
}

ephemeral "aws_ssm_parameter" "db_port" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/db_port"
}

ephemeral "aws_ssm_parameter" "db_user" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/db_user"
}

ephemeral "aws_ssm_parameter" "db_password" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/db_password"
}

ephemeral "aws_ssm_parameter" "db_ssl" {
  arn = "arn:aws:ssm:${var.aws_region}:${var.aws_account_id}:parameter/db_ssl"
}
