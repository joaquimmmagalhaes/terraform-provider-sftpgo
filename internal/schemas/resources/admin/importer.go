package admin

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func importer(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("username", d.Id()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
