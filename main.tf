terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  host = "http://35.195.16.229"
  username = "jmagalhaes"
  password = "jmagalhaes"
}

resource "hashicups_admin" "test" {
  username = "terraform12"
  permissions = ["*"]
  email = "aaa10@example.pt"
  status = 1
  # description = "somethunbg"
  password = "bom dia"
  filters {
    allow_list = [
      "TEST"
    ]
  }

  additional_info = "BOM DIAAAAAAAAAAAAAA"
}