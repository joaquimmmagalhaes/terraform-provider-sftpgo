package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/api"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/models"
	"sort"
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

	if err := d.Set("description", user.Description); err != nil {
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

	// Validate if struct isn't empty
	if user.Filters != nil && (len(user.Filters.AllowedIp) > 0 ||
		len(user.Filters.DeniedIp) > 0 ||
		len(user.Filters.DeniedLoginMethods) > 0 ||
		len(user.Filters.DeniedProtocols) > 0 ||
		len(user.Filters.FilePatterns) > 0) {
		if err := d.Set("filters", getFilters(user.Filters)); err != nil {
			return diag.FromErr(err)
		}
	}

	if user.Filesystem != nil {
		if err := d.Set("filesystem", getFilesystem(*user.Filesystem)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("additional_info", user.AdditionalInfo); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func getVirtualFolders(virtualFolders []models.VirtualFolder) []interface{} {
	result := make([]interface{}, len(virtualFolders))

	for i, virtualFolder := range virtualFolders {
		entry := make(map[string]interface{})

		entry["name"] = virtualFolder.Name
		entry["mapped_path"] = virtualFolder.MappedPath
		entry["description"] = virtualFolder.Description
		entry["virtual_path"] = virtualFolder.VirtualPath
		entry["quota_size"] = virtualFolder.QuotaSize
		entry["quota_files"] = virtualFolder.QuotaFiles
		entry["users"] = helpers.ConvertStringSliceToInterfaceSlice(virtualFolder.Users)

		if virtualFolder.Filesystem != nil {
			entry["filesystem"] = getFilesystem(*virtualFolder.Filesystem)
		}

		result[i] = entry
	}

	return result
}

// TODO Improve and add sort by folder to prevent changes due to sort change.
func getPermissions(permissions map[string][]string) []interface{} {
	result := make(map[string][]interface{}, len(permissions))
	var subDirs []map[string]interface{}

	for k, v := range permissions {
		if k == "/" {
			sort.Strings(v)
			result["global"] = helpers.ConvertStringSliceToInterfaceSlice(v)
		} else {
			sort.Strings(v)
			subDirs = append(subDirs, map[string]interface{}{
				"folder":     k,
				"permission": helpers.ConvertStringSliceToInterfaceSlice(v),
			})
		}
	}

	if len(subDirs) > 0 {
		sort.Slice(subDirs, func(i, j int) bool {
			return subDirs[i]["folder"].(string) < subDirs[j]["folder"].(string)
		})

		result["sub_dirs"] = []interface{}{subDirs}
	}

	return []interface{}{result}
}

func getFilters(filters *models.Filters) []interface{} {
	result := make(map[string]interface{})

	result["allowed_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.AllowedIp)
	result["denied_ip"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedIp)
	result["denied_login_methods"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedLoginMethods)
	result["denied_protocols"] = helpers.ConvertStringSliceToInterfaceSlice(filters.DeniedProtocols)

	filePatterns := make([]interface{}, len(filters.FilePatterns))

	for i, v := range filters.FilePatterns {
		fileExtension := make(map[string]interface{})

		fileExtension["path"] = v.Path
		fileExtension["allowed_patterns"] = helpers.ConvertStringSliceToInterfaceSlice(v.AllowedPatterns)
		fileExtension["denied_patterns"] = helpers.ConvertStringSliceToInterfaceSlice(v.DeniedPatterns)

		filePatterns[i] = fileExtension
	}

	result["file_patterns"] = filePatterns

	return []interface{}{result}
}

// TODO Add missing filesystems
func getFilesystem(filesystem models.Filesystem) []interface{} {
	result := make(map[string]interface{})

	result["provider"] = filesystem.Provider

	gcsconfig := make(map[string]interface{})
	gcsconfig["bucket"] = filesystem.Gcsconfig.Bucket
	gcsconfig["key_prefix"] = filesystem.Gcsconfig.KeyPrefix
	gcsconfig["automatic_credentials"] = filesystem.Gcsconfig.AutomaticCredentials
	gcsconfig["storage_class"] = filesystem.Gcsconfig.StorageClass

	// Ugly fix to handle empty struct response.
	if filesystem.Gcsconfig.Credentials.Key != "" ||
		filesystem.Gcsconfig.Credentials.Status != "" ||
		filesystem.Gcsconfig.Credentials.Payload != "" ||
		filesystem.Gcsconfig.Credentials.AdditionalData != "" {

		credentials := make(map[string]interface{})
		credentials["status"] = filesystem.Gcsconfig.Credentials.Status
		credentials["payload"] = filesystem.Gcsconfig.Credentials.Payload
		credentials["key"] = filesystem.Gcsconfig.Credentials.Key
		credentials["additional_data"] = filesystem.Gcsconfig.Credentials.AdditionalData
		credentials["mode"] = filesystem.Gcsconfig.Credentials.Mode

		gcsconfig["credentials"] = []interface{}{credentials}
	}

	result["gcsconfig"] = []interface{}{gcsconfig}

	return []interface{}{result}
}
