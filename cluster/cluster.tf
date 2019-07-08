
resource "google_container_cluster" "primary" {
  name     = "${var.cluster_name}"
  location = "us-central1"
  project  = "${var.project_id}"

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  maintenance_policy {
    daily_maintenance_window {
      start_time = "03:00"
    }
  }

  depends_on = [google_project_service.container_api,
  google_project_service.cloud_resource_manager_api]
}
