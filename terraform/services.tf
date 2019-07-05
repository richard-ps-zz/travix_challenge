resource "google_project_service" "container_api" {
  project = "starry-antonym-226601"
  service = "container.googleapis.com"
}
