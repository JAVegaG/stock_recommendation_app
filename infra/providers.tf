
terraform {
  required_version = ">= 1.3.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
    cockroach = {
      source  = "cockroachdb/cockroach"
      version = "~> 1.12.0"
    }
  }

  backend "s3" {
    bucket       = "terraform-states-aws-vega"
    key          = "stock-app/terraform.tfstate"
    region       = "us-east-1"
    encrypt      = true
    use_lockfile = true
  }
}

provider "aws" {
  region  = var.aws_region
  profile = var.aws_profile
}

provider "cockroach" {
  # Instructions for using the CockroachDB Cloud API
  # https://www.cockroachlabs.com/docs/cockroachcloud/cloud-api.html
  #
  # Instructions for getting an API Key
  # https://www.cockroachlabs.com/docs/cockroachcloud/console-access-management.html#api-access
  #
  # The Terraform provider requires either the COCKROACH_API_KEY or COCKROACH_API_JWT environment variable for performing authentication.
  # export COCKROACH_API_KEY="the API Key value here"
  # export COCKROACH_API_JWT="the JWT value here"
}
