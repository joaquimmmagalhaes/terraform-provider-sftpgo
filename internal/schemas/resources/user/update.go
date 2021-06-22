package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
)

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	adminID := d.Id()

	if d.HasChanges("status") ||
		d.HasChanges("username") ||
		d.HasChanges("description") ||
		d.HasChanges("expiration_date") ||
		d.HasChanges("password") ||
		d.HasChanges("public_keys") ||
		d.HasChanges("home_dir") ||
		d.HasChanges("permissions") ||
		d.HasChanges("uid") ||
		d.HasChanges("gid") ||
		d.HasChanges("max_sessions") ||
		d.HasChanges("quota_size") ||
		d.HasChanges("quota_files") ||
		d.HasChanges("virtual_folders") ||
		d.HasChanges("upload_bandwidth") ||
		d.HasChanges("download_bandwidth") ||
		d.HasChanges("filters") ||
		d.HasChanges("filesystem") ||
		d.HasChanges("additional_info") {
		err := c.UpdateUser(ctx, adminID, convertToStruct(d))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return get(ctx, d, m)
}
