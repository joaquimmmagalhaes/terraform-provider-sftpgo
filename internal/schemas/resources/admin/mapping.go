package admin

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
)

func convertFromMapToAdminStruct(d *schema.ResourceData) models.Admin {
	var admin models.Admin

	// if d.Get("id") != nil {
	// 	admin.ID = d.Get("id").(int)
	// }

	admin.Status = d.Get("status").(int)
	admin.Username = d.Get("username").(string)
	admin.Description = d.Get("description").(string)
	admin.Password = d.Get("password").(string)
	admin.Email = d.Get("email").(string)

	itemsRaw := d.Get("permissions").([]interface{})
	admin.Permissions = make([]string, len(itemsRaw))
	for i, raw := range itemsRaw {
		admin.Permissions[i] = raw.(string)
	}

	admin.AdditionalInfo = d.Get("additional_info").(string)

	// Map filters
	// filters := d.Get("filters").(map[string]interface{})
	// admin.Filters = models.AdminFilters{
	// 	AllowList: filters["allow_list"].([]string),
	// }

	return admin
}
