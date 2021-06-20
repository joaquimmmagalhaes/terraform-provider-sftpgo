package user

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
)

func convertToStruct(d *schema.ResourceData) models.User {
	var user models.User

	user.Status = d.Get("status").(int)
	user.Username = d.Get("username").(string)
	user.ExpirationDate = d.Get("expiration_date").(int)
	user.Password = d.Get("password").(string)
	user.PublicKeys = convertFromInterfaceSliceToStringSlice(d.Get("public_keys"))
	user.HomeDir = d.Get("home_dir").(string)
	user.VirtualFolders = flattenVirtualFolders(d.Get("virtual_folders"))
	user.UID = d.Get("uid").(int)
	user.GID = d.Get("gid").(int)
	user.MaxSessions = d.Get("max_sessions").(int)
	user.QuotaSize = d.Get("quota_size").(int)
	user.QuotaFiles = d.Get("quota_files").(int)
	user.Permissions = flattenPermissions(d.Get("permissions"))
	user.UsedQuotaSize = d.Get("used_quota_size").(int)
	user.UsedQuotaFiles = d.Get("used_quota_files").(int)
	user.LastQuotaUpdate = d.Get("last_quota_update").(int)
	user.UploadBandwidth = d.Get("upload_bandwidth").(int)
	user.DownloadBandwidth = d.Get("download_bandwidth").(int)
	user.LastLogin = d.Get("last_login").(int)
	// user.Filters = flattenFilters(d.Get("filters"))
	// user.FsConfig = d.Get("filesystem").(int)

	return user
}

func convertFromInterfaceSliceToStringSlice(val interface{}) []string {
	slice := val.([]interface{})
	result := make([]string, len(slice))

	for i, v := range slice {
		result[i] = fmt.Sprint(v)
	}

	return result
}

func flattenVirtualFolders(data interface{}) []models.VirtualFolder {
	items := data.([]interface{})
	result := make([]models.VirtualFolder, len(items))

	for _, item := range items {
		folder := item.(map[string]interface{})
		var entry models.VirtualFolder

		if v, ok := folder["name"]; ok {
			entry.Name = v.(string)
		}

		if v, ok := folder["mapped_path"]; ok {
			entry.MappedPath = v.(string)
		}

		if v, ok := folder["used_quota_size"]; ok {
			entry.UsedQuotaSize = v.(int)
		}

		if v, ok := folder["used_quota_files"]; ok {
			entry.UsedQuotaFiles = v.(int)
		}

		if v, ok := folder["last_quota_update"]; ok {
			entry.LastQuotaUpdate = v.(int)
		}

		if v, ok := folder["users"]; ok {
			entry.Users = convertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := folder["virtual_path"]; ok {
			entry.VirtualPath = v.(string)
		}

		if v, ok := folder["quota_size"]; ok {
			entry.QuotaSize = v.(int)
		}

		if v, ok := folder["quota_files"]; ok {
			entry.QuotaFiles = v.(int)
		}

		result = append(result, entry)
	}

	return result
}

func flattenPermissions(data interface{}) map[string][]string {
	items := data.(map[string]interface{})
	result := make(map[string][]string)

	for k, v := range items {
		permissions := v.([]string)
		result[k] = permissions
	}

	return result
}

// func flattenFilters(data interface{}) models.UserFilters {
// 	var result models.UserFilters
//
// 	items := data.(map[string]interface{})
//
// 	if v, ok := items["allowed_ip"]; ok {
// 		entry.Name = v.([]string)
// 	}
//
// 	return result
// }
