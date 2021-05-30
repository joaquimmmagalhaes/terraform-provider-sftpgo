package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/api"
	datasources "github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/schemas/data-sources"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/schemas/resources"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFTP_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFTP_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SFTP_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sftp_admin": resources.ResourceAdmin(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sftp_admin": datasources.DataSourceAdmin(),
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
			Detail:   "Unable to authenticate user for authenticated STFPGo client",
		})

		return nil, diags
	}

	return c, diags
}
