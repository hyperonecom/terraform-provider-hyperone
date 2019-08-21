package hyperone

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	openapi "github.com/hyperonecom/h1-client-go"
	"testing"
)

func TestAccHyperoneIp_basic(t *testing.T) {
	var n openapi.Ip

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneIpConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccIp("hyperone_ip.foo", &n),
				),
			},
		},
	})
}

func testAccIp(n string, ip *openapi.Ip) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*Config).client
		resource, _, err := client.IpApi.IpShow(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}
		*ip = resource
		return nil
	}
}

const testAccHyperoneIpConfig = `
resource "hyperone_ip" "foo" {
}
`

func TestAccHyperoneIp_ptr(t *testing.T) {
	var n openapi.Ip

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneIpPtrConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccIp("hyperone_ip.foo", &n),
					resource.TestCheckResourceAttr("hyperone_ip.foo", "ptr_record", "ptr.example.com"),
				),
			},
		},
	})
}

const testAccHyperoneIpPtrConfig = `
resource "hyperone_ip" "foo" {
  ptr_record = "ptr.example.com"
}
`
