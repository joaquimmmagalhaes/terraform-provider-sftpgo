package models

type User struct {
	Id                int                 `json:"id,omitempty"`
	Status            int                 `json:"status"`
	Username          string              `json:"username"`
	Description       string              `json:"description"`
	ExpirationDate    float64             `json:"expiration_date"`
	Password          string              `json:"password,omitempty"`
	PublicKeys        []string            `json:"public_keys,omitempty"`
	HomeDir           string              `json:"home_dir"`
	Permissions       map[string][]string `json:"permissions"`
	Uid               int                 `json:"uid"`
	Gid               int                 `json:"gid"`
	MaxSessions       int                 `json:"max_sessions"`
	QuotaSize         float64             `json:"quota_size"`
	QuotaFiles        int                 `json:"quota_files"`
	VirtualFolders    []VirtualFolder     `json:"virtual_folders"`
	UploadBandwidth   int                 `json:"upload_bandwidth"`
	DownloadBandwidth int                 `json:"download_bandwidth"`
	Filters           *Filters            `json:"filters,omitempty"`
	Filesystem        *Filesystem         `json:"filesystem,omitempty"`
	AdditionalInfo    string              `json:"additional_info"`
}

type VirtualFolder struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	MappedPath  string   `json:"mapped_path"`
	Users       []string `json:"users"`
	Description string   `json:"description"`
	QuotaSize   int      `json:"quota_size"`
	QuotaFiles  int      `json:"quota_files"`
	VirtualPath string   `json:"virtual_path"`
	// Just for mapping. Not used
	UsedQuotaSize int `json:"used_quota_size"`
	// Just for mapping. Not used
	UsedQuotaFiles int `json:"used_quota_files"`
	// Just for mapping. Not used
	LastQuotaUpdate int         `json:"last_quota_update"`
	Filesystem      *Filesystem `json:"filesystem,omitempty"`
}

type Filesystem struct {
	Provider     int           `json:"provider"`
	S3Config     *S3Config     `json:"s3config,omitempty"`
	Gcsconfig    *Gcsconfig    `json:"gcsconfig,omitempty"`
	Azblobconfig *Azblobconfig `json:"azblobconfig,omitempty"`
	Cryptconfig  *Cryptconfig  `json:"cryptconfig,omitempty"`
	Sftpconfig   *Sftpconfig   `json:"sftpconfig,omitempty"`
}

type SftpPassword struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type SftpPrivateKey struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type Sftpconfig struct {
	Endpoint               string         `json:"endpoint"`
	Username               string         `json:"username"`
	Password               SftpPassword   `json:"password"`
	PrivateKey             SftpPrivateKey `json:"private_key"`
	Fingerprints           []string       `json:"fingerprints"`
	Prefix                 string         `json:"prefix"`
	DisableConcurrentReads bool           `json:"disable_concurrent_reads"`
	BufferSize             int            `json:"buffer_size"`
}

type CryptPassphrase struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type Cryptconfig struct {
	Passphrase CryptPassphrase `json:"passphrase"`
}

type AzAccountKey struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type Azblobconfig struct {
	Container         string       `json:"container"`
	AccountName       string       `json:"account_name"`
	AccountKey        AzAccountKey `json:"account_key"`
	SasUrl            string       `json:"sas_url"`
	Endpoint          string       `json:"endpoint"`
	UploadPartSize    int          `json:"upload_part_size"`
	UploadConcurrency int          `json:"upload_concurrency"`
	AccessTier        string       `json:"access_tier"`
	KeyPrefix         string       `json:"key_prefix"`
	UseEmulator       bool         `json:"use_emulator"`
}

type GcsCredentials struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type Gcsconfig struct {
	Bucket               string         `json:"bucket"`
	Credentials          GcsCredentials `json:"credentials,omitempty"`
	AutomaticCredentials int            `json:"automatic_credentials"`
	StorageClass         string         `json:"storage_class"`
	KeyPrefix            string         `json:"key_prefix"`
}

type S3AccessSecret struct {
	Status         string `json:"status"`
	Payload        string `json:"payload"`
	Key            string `json:"key"`
	AdditionalData string `json:"additional_data"`
	Mode           int    `json:"mode"`
}

type S3Config struct {
	Bucket            string         `json:"bucket"`
	Region            string         `json:"region"`
	AccessKey         string         `json:"access_key"`
	AccessSecret      S3AccessSecret `json:"access_secret"`
	Endpoint          string         `json:"endpoint"`
	StorageClass      string         `json:"storage_class"`
	UploadPartSize    int            `json:"upload_part_size"`
	UploadConcurrency int            `json:"upload_concurrency"`
	KeyPrefix         string         `json:"key_prefix"`
}

type FilePatterns struct {
	Path            string   `json:"path"`
	AllowedPatterns []string `json:"allowed_patterns"`
	DeniedPatterns  []string `json:"denied_patterns"`
}

type Filters struct {
	AllowedIp          []string       `json:"allowed_ip"`
	DeniedIp           []string       `json:"denied_ip"`
	DeniedLoginMethods []string       `json:"denied_login_methods"`
	DeniedProtocols    []string       `json:"denied_protocols"`
	FilePatterns       []FilePatterns `json:"file_patterns"`
}
