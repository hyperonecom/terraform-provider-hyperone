package hyperone

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpCreate,
		Read:   resourceIpRead,
		Update: resourceIpUpdate,
		Delete: resourceIpDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ptr_record": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIpCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.IpCreate{}

	resource, _, err := client.IpApi.IpCreate(context.TODO(), options)

	if err != nil {
		return err
	}

	d.SetId(resource.Id)

	return resourceIpRead(d, m)
}

func resourceIpRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	resource, _, err := client.IpApi.IpShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("address", resource.Address)
	d.Set("ptr_record", resource.PtrRecord)

	return nil
}

func resourceIpUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("ptr_record") {
		options := openapi.IpUpdate{
			PtrRecord: d.Get("ptr_record").(string),
		}

		_, _, err := client.IpApi.IpUpdate(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("ptr_record")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceIpRead(d, m)
}

func resourceIpDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	_, err := client.IpApi.IpDelete(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	return nil
}
