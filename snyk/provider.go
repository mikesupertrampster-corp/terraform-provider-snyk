package snyk

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return configure(ctx, d)
	}

	return p
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	c := NewClient(ctx, d.Get("hostUrl").(string), d.Get("apiKey").(string), d.Get("orgId").(string), 10*time.Second)

	err := c.Validate()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Snyk client",
			Detail:   "Unable to authenticate user for authenticated Snyk client",
		})

		return nil, diags
	}

	return c, diags
}
