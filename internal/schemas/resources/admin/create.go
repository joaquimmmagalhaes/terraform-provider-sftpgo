package admin

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
)

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	admin, err := c.CreateAdmin(ctx, convertFromMapToAdminStruct(d))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(admin.Username)

	return get(ctx, d, m)
}
