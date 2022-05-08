package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mikesupertrampster-corp/terraform-provider-snyk/internal/snyk/api"
)

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*api.Client)
	name := d.Get("name").(string)

	project, err := c.GetProject(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(project.Id)

	if err := d.Set("id", project.Id); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
