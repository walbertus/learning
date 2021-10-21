locals {
  instance_name_prefix = "patroni"
  machine_type         = "e2-standard-2"
  subnetwork           = "asia-southeast2"
  network_tags         = ["allow-public-ssh"]

  boot_disk_size  = 10
  boot_disk_type  = "pd-standard"
  boot_disk_image = "ubuntu-2004-lts"
}

resource "google_compute_instance" "patroni_01" {
  name         = "${local.instance_name_prefix}-01"
  zone         = "asia-southeast2-a"
  machine_type = local.machine_type

  allow_stopping_for_update = false
  deletion_protection       = false

  boot_disk {
    auto_delete = true
    initialize_params {
      size  = local.boot_disk_size
      type  = local.boot_disk_type
      image = local.boot_disk_image
    }
  }

  network_interface {
    subnetwork = local.subnetwork
    access_config {}
  }
  tags = local.network_tags
}

resource "google_compute_instance" "patroni_02" {
  name         = "${local.instance_name_prefix}-02"
  zone         = "asia-southeast2-b"
  machine_type = local.machine_type

  allow_stopping_for_update = false
  deletion_protection       = false

  boot_disk {
    auto_delete = true
    initialize_params {
      size  = local.boot_disk_size
      type  = local.boot_disk_type
      image = local.boot_disk_image
    }
  }

  network_interface {
    subnetwork = local.subnetwork
    access_config {}
  }
  tags = local.network_tags
}
