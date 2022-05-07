package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mikesupertrampster-corp/terraform-provider-snyk/snyk"
)

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: snyk.Provider}
	plugin.Serve(opts)
}
