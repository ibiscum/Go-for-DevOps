package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	petstore "github.com/ibiscum/Go-for-DevOps/chapter13/petstore-provider/internal"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return petstore.Provider()
		},
	})
}
