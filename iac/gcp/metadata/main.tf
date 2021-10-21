resource "google_compute_project_metadata" "ssh_key" {
  metadata = {
    ssh-keys = <<EOF
      william.dembo:ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFaocyRaNwK68PUdDjUuOZQd+66q8dUyavSBRnLq18ND 29192168+walbertus@users.noreply.github.com
    EOF
  }
}
