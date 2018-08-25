package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mdibaiee/terraform-provider-pingdom/pingdom"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})
}
