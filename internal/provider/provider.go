package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mikesupertrampster-corp/terraform-provider-snyk/internal/snyk/api"
	"github.com/mikesupertrampster-corp/terraform-provider-snyk/internal/snyk/project"
	"time"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SNYK_API_KEY", nil),
			},
			"host_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SNYK_HOST_URL", "https://snyk.io/api/v1"),
			},
			"org_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SNYK_ORG_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"snyk_project": project.DataSourceProject(),
		},
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return configure(ctx, d)
	}

	return p
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	c := api.NewClient(d.Get("host_url").(string), d.Get("api_key").(string), d.Get("org_id").(string), 10*time.Second)

	err := c.Validate(ctx)
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
