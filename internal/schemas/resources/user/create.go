package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
)

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	user := convertToStruct(d)
	user.Password = d.Get("password").(string)

	result, err := c.CreateUser(ctx, user)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Username)

	return get(ctx, d, m)
}
