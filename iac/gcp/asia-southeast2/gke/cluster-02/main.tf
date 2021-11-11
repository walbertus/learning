locals {
  cluster_name = "cluster-02"
  location     = "${var.region}-b"
  machine_type = "e2-medium"
  oauth_scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/logging.write",
    "https://www.googleapis.com/auth/monitoring",
    "https://www.googleapis.com/auth/devstorage.read_only",
    "https://www.googleapis.com/auth/compute"
  ]
  cluster_ipv4_cidr_block  = "10.102.0.0/17"
  services_ipv4_cidr_block = "10.102.128.0/17"
  master_ipv4_cidr_block   = "10.202.0.0/28"
}

resource "google_container_cluster" "cluster_01" {
  provider   = google-beta
  name       = local.cluster_name
  location   = local.location
  network    = "main"
  subnetwork = "asia-southeast2"

  default_max_pods_per_node = 100

  ip_allocation_policy {
    cluster_ipv4_cidr_block  = local.cluster_ipv4_cidr_block
    services_ipv4_cidr_block = local.services_ipv4_cidr_block
  }

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = false
    master_ipv4_cidr_block  = local.master_ipv4_cidr_block
  }

  addons_config {
    horizontal_pod_autoscaling {
      disabled = true
    }
  }

  vertical_pod_autoscaling {
    enabled = false
  }

  initial_node_count       = 1
  remove_default_node_pool = true
  lifecycle {
    ignore_changes = [
      node_config
    ]
  }
  node_config {
    preemptible  = true
    machine_type = "e2-micro"
    disk_size_gb = 10
    disk_type    = "pd-standard"
    image_type   = "COS_CONTAINERD"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.cluster_sa.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/compute"
    ]
  }
}


resource "google_container_node_pool" "primary_preemptible_nodes" {
  name              = "${local.cluster_name}-np-01"
  location          = local.location
  cluster           = google_container_cluster.cluster_01.name
  max_pods_per_node = 100

  initial_node_count = 1
  lifecycle {
    ignore_changes = [
      initial_node_count
    ]
  }

  node_config {
    preemptible  = true
    machine_type = local.machine_type
    disk_size_gb = 10
    disk_type    = "pd-standard"
    image_type   = "COS_CONTAINERD"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.cluster_sa.email
    oauth_scopes    = local.oauth_scopes
  }

  autoscaling {
    max_node_count = 3
    min_node_count = 1
  }
}
