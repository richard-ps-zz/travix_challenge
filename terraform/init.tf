provider "google" {
  credentials = "${file("../sa_key.json")}"
  project     = "starry-antonym-226601"
  region      = "us-central1"
}
