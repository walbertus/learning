resource "google_compute_network" "main_network" {
  name                            = "main"
  auto_create_subnetworks         = false
  delete_default_routes_on_create = false
  routing_mode                    = "REGIONAL"
}

resource "google_compute_subnetwork" "asia_southeast2_subnetwork" {
  name                     = "asia-southeast2"
  ip_cidr_range            = "10.11.0.0/16"
  region                   = "asia-southeast2"
  network                  = google_compute_network.main_network.id
  private_ip_google_access = true
}

resource "google_compute_router" "asia_southeast2_router" {
  name    = "asia-southeast2-router"
  region  = google_compute_subnetwork.asia_southeast2_subnetwork.region
  network = google_compute_network.main_network.id

  bgp {
    asn = 64514
  }
}

resource "google_compute_router_nat" "nat" {
  name                               = "asia-southeast2-nat"
  router                             = google_compute_router.asia_southeast2_router.name
  region                             = google_compute_router.asia_southeast2_router.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"

  log_config {
    enable = false
    filter = "ERRORS_ONLY"
  }
}
