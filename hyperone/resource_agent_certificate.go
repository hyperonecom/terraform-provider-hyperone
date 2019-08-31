package hyperone

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceAgentCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgentCertificateCreate,
		Read:   resourceAgentCertificateRead,
		Delete: resourceAgentCertificateDelete,

		Schema: map[string]*schema.Schema{
			"agent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceAgentCertificateCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.AgentPostCredential{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
	}
	log.Printf("[DEBUG] Creating Agent credential: %s", d.Get("agent").(string))
	resource, _, err := client.AgentApi.AgentPostCredential(context.TODO(), d.Get("agent").(string), options)

	if err != nil {
		return err
	}

	d.Set("agent", d.Get("agent").(string))
	d.SetId(resource.Id)

	return resourceAgentCertificateRead(d, m)
}

func resourceAgentCertificateRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	log.Printf("[DEBUG] Reading Agent credential: %s:%s", d.Get("agent").(string), d.Id())
	resource, _, err := client.AgentApi.AgentGetCredentialcertificateId(context.TODO(), d.Get("agent").(string), d.Id())

	if err != nil {
		return err
	}

	d.Set("name", resource.Name)
	d.Set("type", resource.Type)
	d.Set("value", resource.Value)

	return nil
}

func resourceAgentCertificateDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client
	log.Printf("[DEBUG] Deleting Agent credential: %s", d.Get("agent").(string))

	_, _, err := client.AgentApi.AgentDeleteCredentialcertificateId(context.TODO(), d.Get("agent").(string), d.Id())

	if err != nil {
		return err
	}

	return nil
}
