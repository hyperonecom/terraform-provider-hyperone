package hyperone

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HYPERONE_TOKEN", nil),
				Description: "The token key for API operations.",
			},
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HYPERONE_PROJECT", nil),
				Description: "Select project for API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hyperone_disk": resourceDisk(),
			"hyperone_vm":   resourceVM(),
			"hyperone_ip":   resourceIP(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("token").(string),
		Project: d.Get("project").(string),
	}

	return config.Load()
}
