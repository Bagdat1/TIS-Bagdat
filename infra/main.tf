# Google Cloud немесе Azure/AWS провайдері
provider "google" {
  project = "tis-22k-project"
  region  = "us-central1"
}

# Виртуалды серверді құру (IaC)
resource "google_compute_instance" "app_server" {
  name         = "tis-prod-server"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Сыртқы IP мекенжайы
    }
  }
}
