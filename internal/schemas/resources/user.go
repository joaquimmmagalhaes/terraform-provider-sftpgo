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

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserGet,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
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

func resourceUserGet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	adminID := strconv.Itoa(d.Get("id").(int))

	admin, err := c.GetUser(ctx, adminID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("item", flattenUser(admin)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(adminID)

	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	item := d.Get("item").(map[string]interface{})

	user, err := c.CreateUser(ctx, convertFromMapToUserStruct(item))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(int(user.ID)))

	return resourceUserGet(ctx, d, m)
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	adminID := d.Id()

	if d.HasChange("item") {
		item := d.Get("item").(map[string]interface{})

		err := c.UpdateUser(ctx, adminID, convertFromMapToUserStruct(item))
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceUserGet(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	adminID := d.Id()

	err := c.DeleteUser(ctx, adminID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func convertFromMapToUserStruct(item map[string]interface{}) models.User {
	admin := models.User{
		ID:                item["id"].(int64),
		Status:            item["status"].(int),
		Username:          item["username"].(string),
		ExpirationDate:    item["expiration_date"].(int64),
		Password:          item["password"].(string),
		PublicKeys:        item["public_keys"].([]string),
		HomeDir:           item["home_dir"].(string),
		VirtualFolders:    nil,
		UID:               item["uid"].(int),
		GID:               item["gid"].(int),
		MaxSessions:       item["max_sessions"].(int),
		QuotaSize:         item["quota_size"].(int64),
		QuotaFiles:        item["quota_files"].(int),
		Permissions:       nil,
		UsedQuotaSize:     item["used_quota_size"].(int64),
		UsedQuotaFiles:    item["used_quota_files"].(int),
		LastQuotaUpdate:   item["last_quota_update"].(int64),
		UploadBandwidth:   item["upload_bandwidth"].(int64),
		DownloadBandwidth: item["download_bandwidth"].(int64),
		LastLogin:         item["last_login"].(int64),
		Filters:           models.UserFilters{},
		FsConfig:          models.Filesystem{},
	}

	return admin
}

func flattenUser(admin *models.User) []interface{} {
	c := make(map[string]interface{})

	return []interface{}{c}
}
