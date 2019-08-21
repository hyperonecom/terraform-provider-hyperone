package hyperone

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	configPath = "~/.h1-cli/conf.json"
)

// Provider Init provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("HYPERONE_ACCESS_TOKEN_SECRET"); v != "" {
						return v, nil
					}

					cliConfig, err := loadCLIConfig()
					if err != nil {
						return "", nil
					}

					return cliConfig.Profile.APIKey, nil
				},
				Description: "The token key for API operations.",
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("HYPERONE_PROJECT"); v != "" {
						return v, nil
					}

					cliConfig, err := loadCLIConfig()
					if err != nil {
						return "", nil
					}

					return cliConfig.Profile.Project.ID, nil
				},
				Description: "Select project for API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hyperone_disk":     resourceDisk(),
			"hyperone_vm":       resourceVM(),
			"hyperone_ip":       resourceIP(),
			"hyperone_firewall": resourceFirewall(),
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

type cliConfig struct {
	Profile struct {
		APIKey  string `json:"apiKey"`
		Project struct {
			ID string `json:"_id"`
		} `json:"project"`
	} `json:"profile"`
}

func loadCLIConfig() (cliConfig, error) {
	path, err := homedir.Expand(configPath)
	if err != nil {
		return cliConfig{}, err
	}

	_, err = os.Stat(path)
	if err != nil {
		// Config not found
		return cliConfig{}, nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return cliConfig{}, err
	}

	var c cliConfig
	err = json.Unmarshal(content, &c)
	if err != nil {
		return cliConfig{}, err
	}

	return c, nil
}
