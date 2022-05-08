package api

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetAllProjectNames(ctx context.Context) ([]string, error) {
	res, err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("org/%s/projects", c.OrgId), nil)
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

	var names []string
	for _, element := range projects {
		names = append(names, element.Name)
	}

	return names, nil
}

func (c *Client) GetProject(ctx context.Context, name string) (*Project, error) {
	var body = []byte(fmt.Sprintf(`{"filter": {"name": "%s"}}`, name))
	res, err := c.doRequest(ctx, http.MethodPost, fmt.Sprintf("org/%s/projects", c.OrgId), body)
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
