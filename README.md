# [SFTPGo](https://github.com/drakkan/sftpgo) Terraform Provider
[![release](https://github.com/joaquimmmagalhaes/terraform-provider-sftpgo/actions/workflows/release.yml/badge.svg)](https://github.com/joaquimmmagalhaes/terraform-provider-sftpgo/actions/workflows/release.yml)

Terraform provider for [drakkan/sftpgo](drakkan/sftpgo).

### Usage example
```
terraform {
  required_providers {
    sftpgo = {
      source = "joaquimmmagalhaes/sftpgo"
      version = "0.0.2"
    }
  }
}

provider "sftpgo" {
  # Should another way to handle this values. Like vault secret or environment variable.
  host = "example.com"
  username = "dos"
  password = "test123"
}

resource "sftpgo_admin" "test" {
  status = 1
  username = "test"
  description = "test admin"
  password = "123456789"
  email = "text@example.com"
  permissions = [ "*" ]
  additional_info = "Terraform test user"

  filters {
    allow_list = ["0.0.0.0/0"]
  }
}
```

### Note
This not production ready yet. This is still in the early stages of development and will receive more updates that will not be B/C.