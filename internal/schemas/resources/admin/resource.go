package admin

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/api"
)

func Get() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   get,
		UpdateContext: update,
		DeleteContext: delete,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				StateFunc: hashSum,
				Sensitive: true,
				Type:      schema.TypeString,
				Required:  true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filters": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_list": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"additional_info": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func get(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	var diags diag.Diagnostics

	admin, err := c.GetAdmin(ctx, d.Get("username").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(admin.Username)

	if err := d.Set("status", admin.Status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("username", admin.Username); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", admin.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("email", admin.Email); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("permissions", admin.Permissions); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("additional_info", admin.AdditionalInfo); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func hashSum(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}
