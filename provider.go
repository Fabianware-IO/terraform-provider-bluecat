package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider :: create a new provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        map[string]*schema.Schema{},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
