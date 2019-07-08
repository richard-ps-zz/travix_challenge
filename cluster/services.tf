resource "google_project_service" "container_api" {
  project = "${var.project_id}"
  service = "container.googleapis.com"
}

resource "google_project_service" "cloud_resource_manager_api" {
  project = "${var.project_id}"
  service = "cloudresourcemanager.googleapis.com"
}
