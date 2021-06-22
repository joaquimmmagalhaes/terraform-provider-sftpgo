---
page_title: "admin Resource - terraform-provider-sftpgo"
subcategory: ""
description: |-
The admin resource allows you to manage a sftp admin.
---

# Resource `sftpgo_admin`

## Example Usage

```terraform
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

## Argument Reference
- `status` - (Required) User status (0: Disabled, 1: Enabled).
- `username` - (Required|Unique) username. This is going to be used to authenticate in the dashboard.
- `description` - (Optional) User description.
- `password` - (Required) User password
- `email` - (Required) User email
- `permissions` - (Required) User permissions
- `filters` - (Required) User filters. See [Filters](#filters) below for details.

### Filters
- `allow_list` - (Required) Only clients connecting from these IP/Mask are allowed. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example "192.0.2.0/24" or "2001:db8::/32"
