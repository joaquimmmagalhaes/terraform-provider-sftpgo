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

	// TODO Improve this to prevent invalid sort order
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

	for i, virtualFolder := range virtualFolders {
		entry := make(map[string]interface{})

		entry["name"] = virtualFolder.Name
		entry["mapped_path"] = virtualFolder.MappedPath
		entry["description"] = virtualFolder.Description
		entry["virtual_path"] = virtualFolder.VirtualPath
		entry["quota_size"] = virtualFolder.QuotaSize
		entry["quota_files"] = virtualFolder.QuotaFiles
		entry["users"] = helpers.ConvertStringSliceToInterfaceSlice(virtualFolder.Users)
		entry["filesystem"] = getFilesystem(virtualFolder.Filesystem)

		result[i] = entry
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
	result["max_upload_file_size"] = filters.MaxUploadFileSize
	result["tls_username"] = filters.TlsUsername
	result["disable_fs_checks"] = filters.DisableFsChecks
	result["web_client"] = helpers.ConvertStringSliceToInterfaceSlice(filters.WebClient)

	filePatterns := make([]interface{}, len(filters.FilePatterns))

	for i, v := range filters.FilePatterns {
		fileExtension := make(map[string]interface{})

		fileExtension["path"] = v.Path
		fileExtension["allowed_patterns"] = helpers.ConvertStringSliceToInterfaceSlice(v.AllowedPatterns)
		fileExtension["denied_patterns"] = helpers.ConvertStringSliceToInterfaceSlice(v.DeniedPatterns)

		filePatterns[i] = fileExtension
	}

	result["file_patterns"] = filePatterns

	hooks := make(map[string]interface{})
	hooks["external_auth_disabled"] = filters.Hooks.ExternalAuthDisabled
	hooks["pre_login_disabled"] = filters.Hooks.PreLoginDisabled
	hooks["check_password_disabled"] = filters.Hooks.CheckPasswordDisabled

	result["hooks"] = []interface{}{hooks}

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

	credentials := make(map[string]interface{})
	credentials["status"] = filesystem.Gcsconfig.Credentials.Status
	credentials["payload"] = filesystem.Gcsconfig.Credentials.Payload
	credentials["key"] = filesystem.Gcsconfig.Credentials.Key
	credentials["additional_data"] = filesystem.Gcsconfig.Credentials.AdditionalData
	credentials["mode"] = filesystem.Gcsconfig.Credentials.Mode

	gcsconfig["credentials"] = []interface{}{credentials}

	result["gcsconfig"] = []interface{}{gcsconfig}

	return []interface{}{result}
}
