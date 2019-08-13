package hyperone

import (
	openapi "github.com/hyperonecom/h1-client-go"
)

// Config hold client configuration
type Config struct {
	Token   string
	Project string
	client  *openapi.APIClient
}

// Load h1-client-go initializatin
func (c *Config) Load() (*Config, error) {

	cfg := openapi.NewConfiguration()
	cfg.AddDefaultHeader("authorization", "Bearer "+c.Token)
	cfg.AddDefaultHeader("x-project", c.Project)
	cfg.AddDefaultHeader("Prefer", "respond-async,wait=3600")

	c.client = openapi.NewAPIClient(cfg)

	return c, nil
}
