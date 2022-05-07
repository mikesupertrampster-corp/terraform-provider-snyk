package snyk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
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

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*Client)
	name := d.Get("name").(string)

	project, err := c.GetProject(ctx, name)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("id", project.Id)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", project.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Id)

	return diags
}

func (c Client) GetProject(ctx context.Context, name string) (*Project, error) {
	res, err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("org/%s/projects", c.orgId), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw := map[string]json.RawMessage{}
	err = json.NewDecoder(res.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(raw["projects"], &projects)
	if err != nil {
		return nil, err
	}

	for _, element := range projects {
		if element.Name == name {
			return &element, nil
		}
	}

	return nil, ErrNotFound
}
