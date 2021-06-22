package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/models"
)

func convertToStruct(d *schema.ResourceData) models.User {
	var user models.User

	user.Status = d.Get("status").(int)
	user.Username = d.Get("username").(string)
	user.Description = d.Get("description").(string)
	user.ExpirationDate = d.Get("expiration_date").(float64)

	publicKeys := helpers.ConvertFromInterfaceSliceToStringSlice(d.Get("public_keys"))
	if len(publicKeys) > 0 {
		user.PublicKeys = publicKeys
	}

	user.HomeDir = d.Get("home_dir").(string)
	user.VirtualFolders = flattenVirtualFolders(d.Get("virtual_folders"))
	user.Uid = d.Get("uid").(int)
	user.Gid = d.Get("gid").(int)
	user.MaxSessions = d.Get("max_sessions").(int)
	user.QuotaSize = d.Get("quota_size").(float64)
	user.QuotaFiles = d.Get("quota_files").(int)
	user.Permissions = flattenPermissions(d.Get("permissions"))
	user.UploadBandwidth = d.Get("upload_bandwidth").(int)
	user.DownloadBandwidth = d.Get("download_bandwidth").(int)
	user.Filters = flattenFilters(d.Get("filters"))
	user.Filesystem = flattenFileSystem(d.Get("filesystem"))

	return user
}

func flattenVirtualFolders(data interface{}) []models.VirtualFolder {
	items := data.([]interface{})
	result := make([]models.VirtualFolder, len(items))

	for _, item := range items {
		folder := item.(map[string]interface{})
		var entry models.VirtualFolder

		if v, ok := folder["id"]; ok {
			entry.Id = v.(int)
		}

		if v, ok := folder["name"]; ok {
			entry.Name = v.(string)
		}

		if v, ok := folder["mapped_path"]; ok {
			entry.MappedPath = v.(string)
		}

		if v, ok := folder["description"]; ok {
			entry.Description = v.(string)
		}

		if v, ok := folder["users"]; ok {
			entry.Users = helpers.ConvertFromInterfaceSliceToStringSlice(v)
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

		if v, ok := folder["filesystem"]; ok {
			entry.Filesystem = flattenFileSystem(v)
		}

		result = append(result, entry)
	}

	return result
}

func flattenPermissions(data interface{}) map[string][]string {
	permissions := data.([]interface{})
	result := make(map[string][]string)

	if len(permissions) > 0 {
		permissions := permissions[0].(map[string]interface{})

		if v, ok := permissions["global"]; ok {
			result["/"] = helpers.ConvertFromInterfaceSliceToStringSlice(v.([]interface{}))
		}

		if subDirs, ok := permissions["sub_dirs"]; ok {
			subDirs := subDirs.([]interface{})

			if len(subDirs) > 0 {
				subDirs := subDirs[0].([]interface{})

				if len(subDirs) > 0 {

					for _, v := range subDirs {
						subDir := v.(map[string]interface{})
						result[subDir["folder"].(string)] = helpers.ConvertFromInterfaceSliceToStringSlice(subDir["permission"].([]interface{}))
					}
				}
			}
		}
	}

	return result
}

func flattenFilters(data interface{}) models.Filters {
	var result models.Filters
	items := data.([]interface{})

	if len(items) > 0 {
		items := items[0].(map[string]interface{})

		if v, ok := items["allowed_ip"]; ok {
			result.AllowedIp = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["denied_ip"]; ok {
			result.DeniedIp = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["denied_login_methods"]; ok {
			result.DeniedLoginMethods = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["file_patterns"]; ok {
			fileExtensions := v.([]interface{})

			if len(fileExtensions) > 0 {
				fileExtensions := fileExtensions[0].(map[string]interface{})
				result.FilePatterns = make([]models.FilePatterns, len(fileExtensions))

				if v, ok = fileExtensions["path"]; ok {
					result.FilePatterns[0].Path = v.(string)
				}

				if v, ok = fileExtensions["allowed_patterns"]; ok {
					result.FilePatterns[0].AllowedPatterns = helpers.ConvertFromInterfaceSliceToStringSlice(v)
				}

				if v, ok = fileExtensions["denied_patterns"]; ok {
					result.FilePatterns[0].DeniedPatterns = helpers.ConvertFromInterfaceSliceToStringSlice(v)
				}
			}
		}
	}

	return result
}

func flattenFileSystem(data interface{}) models.Filesystem {
	var result models.Filesystem
	items := data.([]interface{})

	if len(items) > 0 {
		items := items[0].(map[string]interface{})

		if v, ok := items["provider"]; ok {
			result.Provider = v.(int)
		}

		if v, ok := items["gcsconfig"]; ok {
			gcsconfig := v.([]interface{})

			if len(gcsconfig) > 0 {
				gcsconfig := gcsconfig[0].(map[string]interface{})
				var config models.Gcsconfig

				if v, ok = gcsconfig["bucket"]; ok {
					config.Bucket = v.(string)
				}

				if v, ok = gcsconfig["key_prefix"]; ok {
					config.KeyPrefix = v.(string)
				}

				if v, ok := items["credentials"]; ok {
					credentials := v.([]interface{})

					if len(credentials) > 0 {
						credentials := credentials[0].(map[string]interface{})
						var credentialsConfig models.GcsCredentials

						if v, ok = credentials["status"]; ok {
							credentialsConfig.Status = v.(string)
						}

						if v, ok = credentials["payload"]; ok {
							credentialsConfig.Payload = v.(string)
						}

						if v, ok = credentials["key"]; ok {
							credentialsConfig.Key = v.(string)
						}

						if v, ok = credentials["additional_data"]; ok {
							credentialsConfig.AdditionalData = v.(string)
						}

						if v, ok = credentials["mode"]; ok {
							credentialsConfig.Mode = v.(int)
						}

						config.Credentials = credentialsConfig
					}
				}

				if v, ok = gcsconfig["automatic_credentials"]; ok {
					config.AutomaticCredentials = v.(int)
				}

				if v, ok = gcsconfig["storage_class"]; ok {
					config.StorageClass = v.(string)
				}

				result.Gcsconfig = &config
			}
		}
	}

	return result
}
