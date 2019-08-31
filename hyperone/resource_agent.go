package hyperone

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgentCreate,
		Read:   resourceAgentRead,
		Update: resourceAgentUpdate,
		Delete: resourceAgentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Computed: true,
				Optional: true,
			},
			"tag": tagsSchema(),
		},
	}
}

func resourceAgentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.AgentCreate{
		Name:    d.Get("name").(string),
		Service: d.Get("type").(string),
	}

	if v, ok := d.GetOk("tag"); ok {
		options.Tag = v.(map[string]interface{})
	}
	log.Printf("[DEBUG] Creating Agent")

	resource, _, err := client.AgentApi.AgentCreate(context.TODO(), options)

	if err != nil {
		return err
	}

	d.SetId(resource.Id)

	return resourceAgentRead(d, m)
}

func resourceAgentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	log.Printf("[DEBUG] Reading Agent: %s", d.Id())
	resource, _, err := client.AgentApi.AgentShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("name", resource.Name)
	d.Set("type", resource.Flavour)
	d.Set("tag", resource.Tag)

	return nil
}

func resourceAgentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("name") {
		options := openapi.AgentUpdate{
			Name: d.Get("name").(string),
		}

		_, _, err := client.AgentApi.AgentUpdate(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("name")
	}

	if d.HasChange("tag") {
		_, _, err := client.AgentApi.AgentPatchTag(context.TODO(), d.Id(),
			d.Get("name").(map[string]string),
		)
		if err != nil {
			return err
		}

		d.SetPartial("tag")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceAgentRead(d, m)
}

func resourceAgentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	log.Printf("[DEBUG] Deleting Agent: %s", d.Id())

	_, err := client.AgentApi.AgentDelete(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	return nil
}
