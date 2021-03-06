package main

import (
	"github.com/gessnerfl/terraform-provider-instana/instana"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return instana.Provider()
		},
	})
}
