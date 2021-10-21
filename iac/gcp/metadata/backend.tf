terraform {
  backend "gcs" {
    bucket = "terraform-state-wad"
    prefix = "learning/metadata"
  }
}
