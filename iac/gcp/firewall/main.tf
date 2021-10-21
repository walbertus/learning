resource "google_compute_firewall" "allow_public_ssh" {
  name    = "allow-public-ssh"
  network = "main"

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["allow-public-ssh"]
}
