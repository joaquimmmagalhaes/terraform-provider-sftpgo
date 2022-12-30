---
page_title: "user Resource - terraform-provider-sftpgo"
subcategory: ""
description: |-
The user resource allows you to manage a SFTPGo users.
---

# Resource `sftpgo_user`

## Example Usage

```terraform
resource "sftpgo_user" "test" {
  status      = 1
  username    = "test"
  description = "test user"
  password    = "123456789"
  email       = "text@example.com"
  home_dir    = "/test"
  additional_info = "created via terraform"
  public_keys = [
    "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHQp02X10cc844T+EI+5zcF16kzMGeNZP0v7ASZtqwxb text@example.com\r\n",
  ]
  permissions {
    global = ["*"] # Permissions for / of home_dir
    sub_dirs = [[{
      folder     = "/mnt/example"
      permission = ["download"]
    }]]
  }
  filters {
    denied_protocols = [
      "FTP",
    ]
    denied_login_methods = [
      "password",
      "password-over-SSH",
    ]
  }
}
```

## Argument Reference
- `status` - (Required) User status (0: Disabled, 1: Enabled).
- `username` - (Required|Unique) username.
- `description` - (Optional) User description.
- `additional_info` - (Optional) Additional info.
- `password` - (Optional) User password
- `email` - (Optional) User email
- `permissions` - (Required) User permissions
- `filters` - (Optional) User filters.
