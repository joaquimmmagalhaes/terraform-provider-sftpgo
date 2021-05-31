package models

type User struct {
	ID                int64               `json:"id"`
	Status            int                 `json:"status"`
	Username          string              `json:"username"`
	ExpirationDate    int64               `json:"expiration_date"`
	Password          string              `json:"password,omitempty"`
	PublicKeys        []string            `json:"public_keys,omitempty"`
	HomeDir           string              `json:"home_dir"`
	VirtualFolders    []VirtualFolder     `json:"virtual_folders,omitempty"`
	UID               int                 `json:"uid"`
	GID               int                 `json:"gid"`
	MaxSessions       int                 `json:"max_sessions"`
	QuotaSize         int64               `json:"quota_size"`
	QuotaFiles        int                 `json:"quota_files"`
	Permissions       map[string][]string `json:"permissions"`
	UsedQuotaSize     int64               `json:"used_quota_size"`
	UsedQuotaFiles    int                 `json:"used_quota_files"`
	LastQuotaUpdate   int64               `json:"last_quota_update"`
	UploadBandwidth   int64               `json:"upload_bandwidth"`
	DownloadBandwidth int64               `json:"download_bandwidth"`
	LastLogin         int64               `json:"last_login"`
	Filters           UserFilters         `json:"filters"`
	FsConfig          Filesystem          `json:"filesystem"`
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
	Bucket            string `json:"bucket,omitempty"`
	KeyPrefix         string `json:"key_prefix,omitempty"`
	Region            string `json:"region,omitempty"`
	AccessKey         string `json:"access_key,omitempty"`
	AccessSecret      string `json:"access_secret,omitempty"`
	Endpoint          string `json:"endpoint,omitempty"`
	StorageClass      string `json:"storage_class,omitempty"`
	UploadPartSize    int64  `json:"upload_part_size,omitempty"`
	UploadConcurrency int    `json:"upload_concurrency,omitempty"`
}

type GCSFsConfig struct {
	Bucket               string `json:"bucket,omitempty"`
	KeyPrefix            string `json:"key_prefix,omitempty"`
	Credentials          string `json:"credentials,omitempty"`
	AutomaticCredentials int    `json:"automatic_credentials,omitempty"`
	StorageClass         string `json:"storage_class,omitempty"`
}

type Filesystem struct {
	Provider  int         `json:"provider"`
	S3Config  S3FsConfig  `json:"s3config,omitempty"`
	GCSConfig GCSFsConfig `json:"gcsconfig,omitempty"`
}

type VirtualFolder struct {
	VirtualPath string `json:"virtual_path"`
	MappedPath  string `json:"mapped_path"`
}
