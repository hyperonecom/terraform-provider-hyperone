package hyperone

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	openapi "github.com/hyperonecom/h1-client-go"
	"testing"
)

func TestAccHyperoneDisk_basic(t *testing.T) {
	var n openapi.Disk

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneDiskConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDisk("hyperone_disk.foo", &n),
					resource.TestCheckResourceAttr("hyperone_disk.foo", "name", "acceptable_disk"),
					resource.TestCheckResourceAttr("hyperone_disk.foo", "type", "ssd"),
				),
			},
		},
	})
}

func testAccDisk(n string, disk *openapi.Disk) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*Config).client
		resource, _, err := client.DiskApi.DiskShow(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}
		*disk = resource
		return nil
	}
}

const testAccHyperoneDiskConfig = `
resource "hyperone_disk" "foo" {
	name = "acceptable_disk"
	type = "ssd"
	size = 10
}
`

func TestAccHyperoneDisk_updateSize(t *testing.T) {
	var n openapi.Disk

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneDiskConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDisk("hyperone_disk.foo", &n),
					resource.TestCheckResourceAttr("hyperone_disk.foo", "size", "10"),
				),
			},
			{
				Config: testAccHyperoneDiskResizedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDisk("hyperone_disk.foo", &n),
					resource.TestCheckResourceAttr("hyperone_disk.foo", "size", "15"),
				),
			},
		},
	})
}

const testAccHyperoneDiskResizedConfig = `
resource "hyperone_disk" "foo" {
  name = "acceptable_disk"
  type = "ssd"
  size = 15
}
`
