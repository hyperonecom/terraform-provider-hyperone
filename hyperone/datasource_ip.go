package hyperone

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceIP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPRead,
		Schema: map[string]*schema.Schema{

			"address": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "ip address",
				ValidateFunc: validation.NoZeroValues,
			},
			// computed attributes
			"ptr_record": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIPRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	address := d.Get("address").(string)

	resource, resp, err := client.IpApi.IpShow(context.TODO(), address)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("ip not found: %s", err)
		}
		return fmt.Errorf("Error retrieving ip: %s", err)
	}

	d.SetId(resource.Id)
	d.Set("address", resource.Address)
	d.Set("ptr_record", resource.PtrRecord)
	d.Set("network", resource.Network)

	return nil
}
