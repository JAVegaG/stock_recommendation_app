resource "cockroach_database" "this" {
  name       = var.database_name
  cluster_id = var.cluster_id

  lifecycle {
    prevent_destroy = true
  }
}
