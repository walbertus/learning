terraform {
  backend "gcs" {
    bucket = "terraform-state-wad"
    prefix = "learning/asia-southeast2/gke/cluster-01"
  }
}
