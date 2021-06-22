package admin

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
		d.HasChanges("description") ||
		d.HasChanges("password") ||
		d.HasChanges("email") ||
		d.HasChanges("permissions") ||
		d.HasChanges("filters") ||
		d.HasChanges("additional_info") {
		err := c.UpdateAdmin(ctx, adminID, convertFromMapToAdminStruct(d))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return get(ctx, d, m)
}
