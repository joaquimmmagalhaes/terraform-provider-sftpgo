package models

type User struct {
	ID                int                 `json:"id"`
	Status            int                 `json:"status"`
	Username          string              `json:"username"`
	ExpirationDate    int                 `json:"expiration_date"`
	Password          string              `json:"password,omitempty"`
	PublicKeys        []string            `json:"public_keys,omitempty"`
	HomeDir           string              `json:"home_dir"`
	VirtualFolders    []VirtualFolder     `json:"virtual_folders,omitempty"`
	UID               int                 `json:"uid"`
	GID               int                 `json:"gid"`
	MaxSessions       int                 `json:"max_sessions"`
	QuotaSize         int                 `json:"quota_size"`
	QuotaFiles        int                 `json:"quota_files"`
	Permissions       map[string][]string `json:"permissions"`
	UsedQuotaSize     int                 `json:"used_quota_size"`
	UsedQuotaFiles    int                 `json:"used_quota_files"`
	LastQuotaUpdate   int                 `json:"last_quota_update"`
	UploadBandwidth   int                 `json:"upload_bandwidth"`
	DownloadBandwidth int                 `json:"download_bandwidth"`
	LastLogin         int                 `json:"last_login"`
	Filters           UserFilters         `json:"filters"`
	FsConfig          Filesystem          `json:"filesystem"`
	AdditionalInfo    string              `json:"additional_info,omitempty"`
}

type ExtensionsFilter struct {
	Path              string   `json:"path"`
	AllowedExtensions []string `json:"allowed_extensions,omitempty"`
	DeniedExtensions  []string `json:"denied_extensions,omitempty"`
}

type UserFilters struct {
	AllowedIP          []string           `json:"allowed_ip,omitempty"`
	DeniedIP           []string           `json:"denied_ip,omitempty"`
	DeniedLoginMethods []string           `json:"denied_login_methods,omitempty"`
	FileExtensions     []ExtensionsFilter `json:"file_extensions,omitempty"`
}

type S3FsConfig struct {
	Bucket            string      `json:"bucket,omitempty"`
	KeyPrefix         string      `json:"key_prefix,omitempty"`
	Region            string      `json:"region,omitempty"`
	AccessKey         string      `json:"access_key,omitempty"`
	AccessSecret      interface{} `json:"access_secret,omitempty"`
	Endpoint          string      `json:"endpoint,omitempty"`
	StorageClass      string      `json:"storage_class,omitempty"`
	UploadPartSize    int         `json:"upload_part_size,omitempty"`
	UploadConcurrency int         `json:"upload_concurrency,omitempty"`
}

type SecretStatus = string

type FileSystemCredentials struct {
	Status         SecretStatus `json:"status,omitempty"`
	Payload        string       `json:"payload,omitempty"`
	Key            string       `json:"key,omitempty"`
	AdditionalData string       `json:"additional_data,omitempty"`
	// 1 means encrypted using a master key
	Mode int `json:"mode,omitempty"`
}

// TODO FIX
type GCSFsConfig struct {
	Bucket               string                 `json:"bucket,omitempty"`
	KeyPrefix            string                 `json:"key_prefix,omitempty"`
	Credentials          *FileSystemCredentials `json:"credentials,omitempty"`
	AutomaticCredentials int                    `json:"automatic_credentials,omitempty"`
	StorageClass         string                 `json:"storage_class,omitempty"`
}

type Filesystem struct {
	Provider  int         `json:"provider"`
	S3Config  S3FsConfig  `json:"s3config,omitempty"`
	GCSConfig GCSFsConfig `json:"gcsconfig,omitempty"`
}

type BaseVirtualFolder struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	MappedPath      string   `json:"mapped_path,omitempty"`
	UsedQuotaSize   int      `json:"used_quota_size"`
	UsedQuotaFiles  int      `json:"used_quota_files"`
	LastQuotaUpdate int      `json:"last_quota_update"`
	Users           []string `json:"users,omitempty"`
}

type VirtualFolder struct {
	BaseVirtualFolder
	VirtualPath string `json:"virtual_path"`
	QuotaSize   int    `json:"quota_size"`
	QuotaFiles  int    `json:"quota_files"`
}
