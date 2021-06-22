---
page_title: "Provider: SFTPGo"
subcategory: ""
description: |-
Terraform provider for interacting with [drakkan/sftpgo](https://github.com/drakkan/sftpgo) API.
---

# SFTPGo Provider

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "sftpgo" {
  username = "test"
  password = "test123"
}
```

## Schema

### Optional

- **username** (String, Optional) Username to authenticate to SFTPGo API
- **password** (String, Optional) Password to authenticate to SFTPGo API
- **host** (String, Optional) SFTPGo API address
