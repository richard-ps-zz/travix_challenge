
provider "google" {
  credentials = "${file("sa_key.json")}"
  project     = "${var.project_id}"
  region      = "us-central1"
}
