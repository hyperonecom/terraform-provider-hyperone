package hyperone

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPCreate,
		Read:   resourceIPRead,
		Update: resourceIPUpdate,
		Delete: resourceIPDelete,

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

func resourceIPCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.IpCreate{}

	if v, ok := d.GetOk("ptr_record"); ok {
		options.PtrRecord = v.(string)
	}

	resource, _, err := client.IpApi.IpCreate(context.TODO(), options)

	if err != nil {
		return err
	}

	d.SetId(resource.Id)

	return resourceIPRead(d, m)
}

func resourceIPRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	resource, _, err := client.IpApi.IpShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("address", resource.Address)
	d.Set("ptr_record", resource.PtrRecord)

	return nil
}

func resourceIPUpdate(d *schema.ResourceData, m interface{}) error {
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

	return resourceIPRead(d, m)
}

func resourceIPDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	_, err := client.IpApi.IpDelete(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	return nil
}
