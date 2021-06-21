package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/api"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
	"strconv"
)

func get(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Client)

	var diags diag.Diagnostics

	user, err := c.GetUser(ctx, d.Get("username").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Username)

	if err := d.Set("id", strconv.Itoa(user.Id)); err != nil {
		return diag.FromErr(err)
	}

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

	if err := d.Set("uid", user.Uid); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("gid", user.Gid); err != nil {
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

	if err := d.Set("upload_bandwidth", user.UploadBandwidth); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("download_bandwidth", user.DownloadBandwidth); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("filters", getFilters(user.Filters)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("filesystem", getFilesystem(user.Filesystem)); err != nil {
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

func getFilters(filters models.Filters) []interface{} {
	result := make(map[string]interface{})

	result["allowed_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.AllowedIp)
	result["denied_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedIp)
	result["denied_login_methods"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedLoginMethods)
	result["denied_protocols"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedProtocols)

	filePatterns := make([]interface{}, len(filters.FilePatterns))

	for i, v := range filters.FilePatterns {
		fileExtension := make(map[string]interface{})

		fileExtension["path"] = v.Path
		fileExtension["allowed_extensions"] = helpers.ConvertStringSliceToInterfaceSlice(v.AllowedPatterns)
		fileExtension["denied_extensions"] = helpers.ConvertStringSliceToInterfaceSlice(v.DeniedPatterns)

		filePatterns[i] = fileExtension
	}

	result["file_patterns"] = filePatterns

	return []interface{}{result}
}

func getFilesystem(filesystem models.Filesystem) []interface{} {
	result := make(map[string]interface{})

	result["provider"] = filesystem.Provider

	gcsconfig := make(map[string]interface{})
	gcsconfig["bucket"] = filesystem.Gcsconfig.Bucket
	gcsconfig["key_prefix"] = filesystem.Gcsconfig.KeyPrefix
	gcsconfig["automatic_credentials"] = filesystem.Gcsconfig.AutomaticCredentials
	gcsconfig["storage_class"] = filesystem.Gcsconfig.StorageClass

	result["gcsconfig"] = []interface{}{gcsconfig}

	return []interface{}{result}
}
