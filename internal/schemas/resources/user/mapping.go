package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/helpers"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/models"
)

func convertToStruct(d *schema.ResourceData) models.User {
	var user models.User

	user.Status = d.Get("status").(int)
	user.Username = d.Get("username").(string)
	user.ExpirationDate = d.Get("expiration_date").(int)
	user.Password = d.Get("password").(string)
	user.PublicKeys = helpers.ConvertFromInterfaceSliceToStringSlice(d.Get("public_keys"))
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
	user.Filters = flattenFilters(d.Get("filters"))
	user.FsConfig = flattenFileSystem(d.Get("filesystem"))

	return user
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

func flattenFilters(data interface{}) models.UserFilters {
	var result models.UserFilters
	items := data.([]interface{})

	if len(items) > 0 {
		items := items[0].(map[string]interface{})

		if v, ok := items["allowed_ip"]; ok {
			result.AllowedIP = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["denied_ip"]; ok {
			result.DeniedIP = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["denied_login_methods"]; ok {
			result.DeniedLoginMethods = helpers.ConvertFromInterfaceSliceToStringSlice(v)
		}

		if v, ok := items["file_extensions"]; ok {
			fileExtensions := v.([]interface{})

			if len(fileExtensions) > 0 {
				fileExtensions := fileExtensions[0].(map[string]interface{})
				result.FileExtensions = make([]models.ExtensionsFilter, len(fileExtensions))

				if v, ok = fileExtensions["path"]; ok {
					result.FileExtensions[0].Path = v.(string)
				}

				if v, ok = fileExtensions["allowed_extensions"]; ok {
					result.FileExtensions[0].AllowedExtensions = helpers.ConvertFromInterfaceSliceToStringSlice(v)
				}

				if v, ok = fileExtensions["denied_extensions"]; ok {
					result.FileExtensions[0].DeniedExtensions = helpers.ConvertFromInterfaceSliceToStringSlice(v)
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
		/*
			if v, ok := items["s3config"]; ok {
				s3config := v.(map[string]interface{})

				if v, ok = s3config["bucket"]; ok {
					result.S3Config.Bucket = v.(string)
				}

				if v, ok = s3config["key_prefix"]; ok {
					result.S3Config.KeyPrefix = v.(string)
				}

				if v, ok = s3config["region"]; ok {
					result.S3Config.Region = v.(string)
				}

				if v, ok = s3config["access_key"]; ok {
					result.S3Config.AccessKey = v.(string)
				}

				if v, ok = s3config["access_secret"]; ok {
					result.S3Config.AccessSecret = v.(string)
				}

				if v, ok = s3config["endpoint"]; ok {
					result.S3Config.Endpoint = v.(string)
				}

				if v, ok = s3config["storage_class"]; ok {
					result.S3Config.StorageClass = v.(string)
				}

				if v, ok = s3config["upload_part_size"]; ok {
					result.S3Config.UploadPartSize = v.(int)
				}

				if v, ok = s3config["upload_concurrency"]; ok {
					result.S3Config.UploadConcurrency = v.(int)
				}
			}
		*/
		if v, ok := items["gcsconfig"]; ok {
			gcsconfig := v.([]interface{})

			if len(gcsconfig) > 0 {
				gcsconfig := gcsconfig[0].(map[string]interface{})

				if v, ok = gcsconfig["bucket"]; ok {
					result.GCSConfig.Bucket = v.(string)
				}

				if v, ok = gcsconfig["key_prefix"]; ok {
					result.GCSConfig.KeyPrefix = v.(string)
				}

				// TODO FIX
				// if v, ok = gcsconfig["credentials"]; ok {
				// 	result.GCSConfig.Credentials = v.(*models.FileSystemCredentials)
				// } else {
				// 	result.GCSConfig.Credentials = nil
				// }
				result.GCSConfig.Credentials = nil

				if v, ok = gcsconfig["automatic_credentials"]; ok {
					result.GCSConfig.AutomaticCredentials = v.(int)
				}

				if v, ok = gcsconfig["storage_class"]; ok {
					result.GCSConfig.StorageClass = v.(string)
				}
			}
		}
	}

	return result
}
