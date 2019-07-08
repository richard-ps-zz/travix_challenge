variable "project_id" {}

provider "google" {
  credentials = "${file("sa_key2.json")}"
  project     = "${var.project_id}"
  region      = "us-central1"
}
