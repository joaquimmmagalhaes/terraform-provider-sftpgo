package sftp

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
	adminResource "github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/schemas/resources/admin"
	userResource "github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/schemas/resources/user"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFTPGO_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFTPGO_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SFTPGO_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sftpgo_admin": adminResource.Get(),
			"sftpgo_user":  userResource.Get(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var host *string
	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c, err := api.NewClient(ctx, host, &username, &password)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create STFPGo client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	return c, diags
}
