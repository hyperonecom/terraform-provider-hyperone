package hyperone

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiskCreate,
		Read:   resourceDiskRead,
		Update: resourceDiskUpdate,
		Delete: resourceDiskDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDiskCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.DiskCreate{
		Size:    float32(d.Get("size").(int)),
		Service: d.Get("type").(string),
		Name:    d.Get("name").(string),
	}

	disk, _, err := client.DiskApi.DiskCreate(context.TODO(), options)

	if err != nil {
		return err
	}

	d.SetId(disk.Id)

	return resourceDiskRead(d, m)
}

func resourceDiskRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	disk, _, err := client.DiskApi.DiskShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("size", disk.Size)
	d.Set("name", disk.Name)
	d.Set("type", disk.Type)

	return nil
}

func resourceDiskUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("name") {
		options := openapi.DiskUpdate{
			Name: d.Get("name").(string),
		}

		_, _, err := client.DiskApi.DiskUpdate(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("name")
	}

	if d.HasChange("size") {
		options := openapi.DiskActionResize{
			Size: float32(d.Get("size").(int)),
		}

		_, _, err := client.DiskApi.DiskActionResize(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("size")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceDiskRead(d, m)
}

func resourceDiskDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	_, err := client.DiskApi.DiskDelete(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	return nil
}
