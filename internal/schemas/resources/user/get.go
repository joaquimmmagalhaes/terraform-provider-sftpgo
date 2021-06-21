package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/api"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
)

func get(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	var diags diag.Diagnostics

	user, err := c.GetUser(ctx, d.Get("username").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Username)

	if err := d.Set("status", user.Status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("username", user.Username); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("expiration_date", user.ExpirationDate); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("public_keys", user.PublicKeys); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("home_dir", user.HomeDir); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("virtual_folders", getVirtualFolders(user.VirtualFolders)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("uid", user.UID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("gid", user.GID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("max_sessions", user.MaxSessions); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("quota_size", user.QuotaSize); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("quota_files", user.QuotaFiles); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("permissions", getPermissions(user.Permissions)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("used_quota_size", user.UsedQuotaSize); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("used_quota_files", user.UsedQuotaFiles); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("last_quota_update", user.LastQuotaUpdate); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("upload_bandwidth", user.UploadBandwidth); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("download_bandwidth", user.DownloadBandwidth); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("last_login", user.LastLogin); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("filters", getFilters(user.Filters)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("filesystem", getFilesystem(user.FsConfig)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("additional_info", user.AdditionalInfo); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func getVirtualFolders(virtualFolders []models.VirtualFolder) []interface{} {
	result := make([]interface{}, len(virtualFolders))

	for i, v := range virtualFolders {
		result[i] = v
	}

	return result
}

func getPermissions(permissions map[string][]string) []interface{} {
	result := make(map[string][]interface{}, len(permissions))
	var subDirs []map[string]interface{}

	for k, v := range permissions {
		if k == "/" {
			result["global"] = helpers.ConvertStringSliceToInterfaceSlice(v)
		} else {
			subDirs = append(subDirs, map[string]interface{}{
				"folder":     k,
				"permission": helpers.ConvertStringSliceToInterfaceSlice(v),
			})
		}
	}

	result["sub_dirs"] = []interface{}{subDirs}

	return []interface{}{result}
}

func getFilters(filters models.UserFilters) interface{} {
	result := make(map[string]interface{})

	result["allowed_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.AllowedIP)
	result["denied_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedIP)
	result["denied_login_methods"] = helpers.ConvertStringSliceToInterfaceSlice(filters.AllowedIP)

	fileExtensions := make([]interface{}, len(filters.FileExtensions))

	for i, v := range filters.FileExtensions {
		fileExtension := make(map[string]interface{})

		fileExtension["path"] = v.Path
		fileExtension["allowed_extensions"] = helpers.ConvertStringSliceToInterfaceSlice(v.AllowedExtensions)
		fileExtension["denied_extensions"] = helpers.ConvertStringSliceToInterfaceSlice(v.DeniedExtensions)

		fileExtensions[i] = fileExtension
	}

	result["file_extensions"] = fileExtensions

	return []interface{}{result}
}

func getFilesystem(filesystem models.Filesystem) []interface{} {
	result := make(map[string]interface{})

	result["provider"] = filesystem.Provider

	gcsconfig := make(map[string]interface{})
	gcsconfig["bucket"] = filesystem.GCSConfig.Bucket
	gcsconfig["key_prefix"] = filesystem.GCSConfig.KeyPrefix
	gcsconfig["automatic_credentials"] = filesystem.GCSConfig.AutomaticCredentials
	gcsconfig["storage_class"] = filesystem.GCSConfig.StorageClass

	result["gcsconfig"] = []interface{}{gcsconfig}

	return []interface{}{result}
}
