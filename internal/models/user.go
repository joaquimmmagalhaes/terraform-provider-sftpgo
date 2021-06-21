package models

type User struct {
	ID                int                 `json:"id"`
	Status            int                 `json:"status"`
	Username          string              `json:"username"`
	ExpirationDate    int                 `json:"expiration_date"`
	Password          string              `json:"password"`
	PublicKeys        []string            `json:"public_keys"`
	HomeDir           string              `json:"home_dir"`
	VirtualFolders    []VirtualFolder     `json:"virtual_folders"`
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
	AdditionalInfo    string              `json:"additional_info"`
}

type ExtensionsFilter struct {
	Path              string   `json:"path"`
	AllowedExtensions []string `json:"allowed_extensions"`
	DeniedExtensions  []string `json:"denied_extensions"`
}

type UserFilters struct {
	AllowedIP          []string           `json:"allowed_ip"`
	DeniedIP           []string           `json:"denied_ip"`
	DeniedLoginMethods []string           `json:"denied_login_methods"`
	FileExtensions     []ExtensionsFilter `json:"file_extensions"`
}

type S3FsConfig struct {
	Bucket            string      `json:"bucket"`
	KeyPrefix         string      `json:"key_prefix"`
	Region            string      `json:"region"`
	AccessKey         string      `json:"access_key"`
	AccessSecret      interface{} `json:"access_secret"`
	Endpoint          string      `json:"endpoint"`
	StorageClass      string      `json:"storage_class"`
	UploadPartSize    int         `json:"upload_part_size"`
	UploadConcurrency int         `json:"upload_concurrency"`
}

type SecretStatus = string

type FileSystemCredentials struct {
	Status         SecretStatus `json:"status"`
	Payload        string       `json:"payload"`
	Key            string       `json:"key"`
	AdditionalData string       `json:"additional_data"`
	// 1 means encrypted using a master key
	Mode int `json:"mode"`
}

// TODO FIX
type GCSFsConfig struct {
	Bucket               string                 `json:"bucket"`
	KeyPrefix            string                 `json:"key_prefix"`
	Credentials          *FileSystemCredentials `json:"credentials"`
	AutomaticCredentials int                    `json:"automatic_credentials"`
	StorageClass         string                 `json:"storage_class"`
}

type Filesystem struct {
	Provider  int         `json:"provider"`
	S3Config  S3FsConfig  `json:"s3config"`
	GCSConfig GCSFsConfig `json:"gcsconfig"`
}

type BaseVirtualFolder struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	MappedPath      string   `json:"mapped_path"`
	UsedQuotaSize   int      `json:"used_quota_size"`
	UsedQuotaFiles  int      `json:"used_quota_files"`
	LastQuotaUpdate int      `json:"last_quota_update"`
	Users           []string `json:"users"`
}

type VirtualFolder struct {
	BaseVirtualFolder
	VirtualPath string `json:"virtual_path"`
	QuotaSize   int    `json:"quota_size"`
	QuotaFiles  int    `json:"quota_files"`
}
