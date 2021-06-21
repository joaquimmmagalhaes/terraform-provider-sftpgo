package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-drakkan-sftpgo/internal/schemas/resources"
)

func Get() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   get,
		UpdateContext: update,
		DeleteContext: delete,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"expiration_date": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"password": {
				StateFunc: resources.HashSum,
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"public_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"home_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sub_dirs": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"folder": {
											Type:     schema.TypeString,
											Required: true,
										},
										"permission": {
											Type:     schema.TypeList,
											Optional: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
									},
								},
							},
						},
					},
				},
				Optional: true,
			},
			"uid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sessions": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"quota_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"quota_files": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"virtual_folders": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mapped_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"used_quota_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"used_quota_files": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"last_quota_update": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"users": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     schema.TypeList,
						},
						"virtual_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"quota_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"quota_files": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"used_quota_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used_quota_files": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_quota_update": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upload_bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"download_bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"last_login": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_ip": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"denied_ip": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"denied_login_methods": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"file_extensions": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"allowed_extensions": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_extensions": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"filesystem": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeInt,
							Optional: true,
							// ExactlyOneOf: []string{"s3config", "gcsconfig"},
						},
						"s3config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"key_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"region": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"access_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"access_secret": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"endpoint": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"storage_class": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"upload_part_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"upload_concurrency": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"gcsconfig": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"key_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
									// TODO FIX THIS
									"credentials": {
										Default:  nil,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"automatic_credentials": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"storage_class": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
