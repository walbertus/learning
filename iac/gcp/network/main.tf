resource "google_compute_network" "main_network" {
  name                            = "main"
  auto_create_subnetworks         = false
  delete_default_routes_on_create = false
  routing_mode                    = "REGIONAL"
}

resource "google_compute_subnetwork" "asia_southeast2_subnetwork" {
  name          = "asia-southeast2"
  ip_cidr_range = "10.11.0.0/16"
  region        = "asia-southeast2"
  network       = google_compute_network.main_network.id
}
