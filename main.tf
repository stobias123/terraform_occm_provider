provider "occm" {
  occm_endpoint = ""
  occm_username = ""
  occm_password = ""
}

resource "occm_instance" "my_manager" {}
