package resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/api"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
	"strconv"
	"time"
)

func ResourceAdmin() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAdminCreate,
		ReadContext:   resourceAdminGet,
		UpdateContext: resourceAdminUpdate,
		DeleteContext: resourceAdminDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"item": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password": {
							Type:     schema.TypeString,
							Required: true,
						},
						"email": {
							Type:     schema.TypeString,
							Required: true,
						},
						"permissions": {
							Type:     schema.TypeList,
							Required: true,
						},
						"filters": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_list": {
										Type:     schema.TypeList,
										Required: true,
									},
								},
							},
						},
						"additional_info": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func resourceAdminGet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	adminID := strconv.Itoa(d.Get("id").(int))

	admin, err := c.GetAdmin(ctx, adminID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("item", flattenAdmin(admin)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(adminID)

	return diags
}

func resourceAdminCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	item := d.Get("item").(map[string]interface{})

	admin, err := c.CreateAdmin(ctx, convertFromMapToAdminStruct(item))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(admin.ID))

	return resourceAdminGet(ctx, d, m)
}

func resourceAdminUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	adminID := d.Id()

	if d.HasChange("item") {
		item := d.Get("item").(map[string]interface{})

		err := c.UpdateAdmin(ctx, adminID, convertFromMapToAdminStruct(item))
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceAdminGet(ctx, d, m)
}

func resourceAdminDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	adminID := d.Id()

	err := c.DeleteAdmin(ctx, adminID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func convertFromMapToAdminStruct(item map[string]interface{}) models.Admin {
	var admin models.Admin

	admin.ID = item["id"].(int)
	admin.Status = item["status"].(int)
	admin.Username = item["username"].(string)
	admin.Description = item["description"].(string)
	admin.Password = item["password"].(string)
	admin.Email = item["email"].(string)
	admin.Permissions = item["permissions"].([]string)
	admin.AdditionalInfo = item["additional_info"].(string)

	// Map filters
	filters := item["filters"].(map[string]interface{})
	admin.Filters = models.AdminFilters{
		AllowList: filters["allow_list"].([]string),
	}

	return admin
}

func flattenAdmin(admin *models.Admin) []interface{} {
	c := make(map[string]interface{})

	c["id"] = admin.ID
	c["status"] = admin.Status
	c["username"] = admin.Username
	c["description"] = admin.Description
	c["password"] = admin.Password
	c["email"] = admin.Email
	c["permissions"] = admin.Permissions
	c["filters"] = flattenFilters(admin.Filters)
	c["additional_info"] = admin.AdditionalInfo

	return []interface{}{c}
}

func flattenFilters(filters models.AdminFilters) []interface{} {
	c := make(map[string]interface{})

	c["allow_list"] = filters.AllowList

	return []interface{}{c}
}
