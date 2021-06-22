package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
)

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := c.DeleteUser(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
