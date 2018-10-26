provider "occm" {
  occm_username = "user@mail.com"
  occm_password = "password"
}

resource "occm_instance" "my_manager" {}
