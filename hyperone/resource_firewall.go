package hyperone

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	openapi "github.com/hyperonecom/h1-client-go"
)

func resourceFirewall() *schema.Resource {

	rule := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"filter": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"external": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"internal": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	return &schema.Resource{
		Create: resourceFirewallCreate,
		Read:   resourceFirewallRead,
		Update: resourceFirewallUpdate,
		Delete: resourceFirewallDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ingress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: rule,
				},
			},
			"egress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: rule,
				},
			},
		},
	}
}

func expandRules(config []interface{}) []openapi.FirewallCreateIngress {
	arr := make([]openapi.FirewallCreateIngress, len(config))

	for i, v := range config {
		rule := v.(map[string]interface{})

		arr[i] = openapi.FirewallCreateIngress{
			Name:     rule["name"].(string),
			Action:   rule["action"].(string),
			Priority: float32(rule["priority"].(int)),
			Filter:   expandSet(rule["filter"].([]interface{})),
			External: expandSet(rule["external"].([]interface{})),
			Internal: expandSet(rule["internal"].([]interface{})),
		}
	}

	return arr
}

func resourceFirewallCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	options := openapi.FirewallCreate{
		Service: d.Get("service").(string),
		Name:    d.Get("name").(string),
		Ingress: expandRules(d.Get("ingress").([]interface{})),
		Egress:  expandRules(d.Get("egress").([]interface{})),
	}

	resource, _, err := client.FirewallApi.FirewallCreate(context.TODO(), options)

	if err != nil {
		return err
	}

	d.SetId(resource.Id)

	return resourceFirewallRead(d, m)
}

func resourceFirewallRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	resource, _, err := client.FirewallApi.FirewallShow(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	d.Set("name", resource.Name)

	return nil
}

func mapRules(arr []interface{}) []map[string]interface{} {

	ret := make([]map[string]interface{}, len(arr))

	for i, v := range arr {
		rule := v.(map[string]interface{})

		ret[i] = rule
		ret[i]["filter"] = expandSet(rule["filter"].([]interface{}))
		ret[i]["external"] = expandSet(rule["external"].([]interface{}))
		ret[i]["internal"] = expandSet(rule["internal"].([]interface{}))
	}

	return ret
}

func resourceFirewallUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("name") {
		options := openapi.FirewallUpdate{
			Name: d.Get("name").(string),
		}

		_, _, err := client.FirewallApi.FirewallUpdate(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("name")
	}

	if d.HasChange("ingress") {
		options := mapRules(d.Get("ingress").([]interface{}))

		_, _, err := client.FirewallApi.FirewallPutIngress(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("ingress")
	}

	if d.HasChange("egress") {
		options := mapRules(d.Get("egress").([]interface{}))

		_, _, err := client.FirewallApi.FirewallPutEgress(context.TODO(), d.Id(), options)
		if err != nil {
			return err
		}

		d.SetPartial("egress")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceFirewallRead(d, m)
}

func resourceFirewallDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).client

	_, err := client.FirewallApi.FirewallDelete(context.TODO(), d.Id())

	if err != nil {
		return err
	}

	return nil
}
