package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joaquimmmagalhaes/terraform-provider-sftpgo/internal/schemas/resources"
)

func Get() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   get,
		UpdateContext: update,
		DeleteContext: delete,
		Importer: &schema.ResourceImporter{
			StateContext: importer,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expiration_date": {
				Type:     schema.TypeFloat,
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
				Type:     schema.TypeFloat,
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
						"users": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
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
						"filesystem": &fileSystemSchema,
					},
				},
			},
			"additional_info": {
				Type:     schema.TypeString,
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
						"denied_protocols": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"file_patterns": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"allowed_patterns": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"denied_patterns": {
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
			"filesystem": &fileSystemSchema,
		},
	}
}

// TODO Add missing filesystems
var fileSystemSchema = schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"provider": {
				Type:     schema.TypeInt,
				Optional: true,
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
						"credentials": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"payload": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"additional_data": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mode": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
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
}
