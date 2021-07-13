package admin

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/models"
)

func convertFromMapToAdminStruct(d *schema.ResourceData) models.Admin {
	var admin models.Admin

	admin.Status = d.Get("status").(int)
	admin.Username = d.Get("username").(string)
	admin.Description = d.Get("description").(string)

	if d.HasChanges("password") {
		admin.Password = d.Get("password").(string)
	}

	admin.Email = d.Get("email").(string)
	admin.AdditionalInfo = d.Get("additional_info").(string)

	permissions := d.Get("permissions").([]interface{})
	admin.Permissions = make([]string, len(permissions))
	for i, permission := range permissions {
		admin.Permissions[i] = permission.(string)
	}

	filters := d.Get("filters").([]interface{})
	if len(filters) > 0 {
		filters := filters[0].(map[string]interface{})
		admin.Filters = models.AdminFilters{}

		if v, ok := filters["allow_list"]; ok {
			admin.Filters.AllowList = helpers.ConvertFromInterfaceSliceToStringSlice(v.([]interface{}))
		}
	}

	return admin
}
