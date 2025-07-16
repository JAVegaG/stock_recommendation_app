locals {
  resource_suffix = "${var.project_name}-${random_id.suffix.hex}-${terraform.workspace}"
}
