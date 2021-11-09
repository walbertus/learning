resource "google_service_account" "cluster_sa" {
  account_id   = "${local.cluster_name}-sa"
  display_name = "Service account for GKE cluster ${local.cluster_name}."
}

resource "google_project_iam_member" "iam_logging_log_writer" {
  project = var.project_name
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.cluster_sa.email}"
}

resource "google_project_iam_member" "iam_monitoring_metric_writer" {
  project = var.project_name
  role    = "roles/monitoring.metricWriter"
  member  = "serviceAccount:${google_service_account.cluster_sa.email}"
}

resource "google_project_iam_member" "iam_monitoring_viewer" {
  project = var.project_name
  role    = "roles/monitoring.viewer"
  member  = "serviceAccount:${google_service_account.cluster_sa.email}"
}

resource "google_project_iam_member" "iam_storage_object_viewer" {
  project = var.project_name
  role    = "roles/storage.objectViewer"
  member  = "serviceAccount:${google_service_account.cluster_sa.email}"
}
