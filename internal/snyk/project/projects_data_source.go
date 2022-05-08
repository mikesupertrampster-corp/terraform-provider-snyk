package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mikesupertrampster-corp/terraform-provider-snyk/internal/snyk/api"
)

func DataSourceProjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectsRead,
		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceProjectsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*api.Client)

	names, err := c.GetAllProjectNames(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(c.OrgId)

	if err := d.Set("names", names); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
