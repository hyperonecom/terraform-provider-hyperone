package hyperone

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sshkeys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disk": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"netadp": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.VmCreate{
		Service: d.Get("type").(string),
		Name:    d.Get("name").(string),
		Image:   d.Get("image").(string),
		SshKeys: expandSet(d.Get("sshkeys").([]interface{})),
		Disk:    expandDisk(d.Get("disk").([]interface{})),
		Netadp:  expandNetadp(d.Get("netadp").([]interface{})),
	}

	resource, _, err := client.VmApi.VmCreate(context.TODO(), options)
	if err != nil {
		return handleError(err)
	}

	d.SetId(resource.Id)

	// Initialize the connection info
	d.SetConnInfo(map[string]string{
		"type": "ssh",
		"host": resource.Fqdn,
	})

	return resourceVMRead(d, m)
}

func expandDisk(config []interface{}) []openapi.VmCreateDisk {
	arr := make([]openapi.VmCreateDisk, len(config))
	for i, v := range config {
		disk := v.(map[string]interface{})

		arr[i] = openapi.VmCreateDisk{
			Name:    disk["name"].(string),
			Service: disk["type"].(string),
			Size:    float32(disk["size"].(int)),
		}
	}

	return arr
}
func expandNetadp(config []interface{}) []openapi.VmCreateNetadp {
	arr := make([]openapi.VmCreateNetadp, len(config))
	for i, v := range config {
		netadp := v.(map[string]interface{})

		arr[i] = openapi.VmCreateNetadp{
			Ip:      expandSet(netadp["ip"].([]interface{})),
			Service: "public",
		}
	}

	return arr
}
func expandSet(tags []interface{}) []string {
	arr := make([]string, len(tags))
	for i, v := range tags {
		arr[i] = v.(string)
	}

	return arr
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	resource, _, err := client.VmApi.VmShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("name", resource.Name)
	d.Set("type", resource.Flavour)
	d.Set("fqdn", resource.Fqdn)

	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("name") {
		options := openapi.VmUpdate{
			Name: d.Get("name").(string),
		}

		_, _, err := client.VmApi.VmUpdate(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("name")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceVMRead(d, m)
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.VmDelete{
		RemoveDisks: []string{},
	}

	_, err := client.VmApi.VmDelete(context.TODO(), d.Id(), options)

	if err != nil {
		return err
	}

	return nil
}

func handleError(err error) error {
	if e, ok := err.(openapi.GenericOpenAPIError); ok {
		if m, ok := e.Model().(openapi.InlineResponse400); ok {
			return fmt.Errorf("Error creating: %s", m.Message)
		}
	}
	return err
}
