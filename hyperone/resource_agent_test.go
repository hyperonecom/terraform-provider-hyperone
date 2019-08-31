package hyperone

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	openapi "github.com/hyperonecom/h1-client-go"
)

func TestAccHyperoneAgent_basic(t *testing.T) {
	var n openapi.Agent

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneAgentConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccAgent("hyperone_agent.foo", &n),
				),
			},
		},
	})
}

func testAccAgent(n string, Agent *openapi.Agent) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*Config).client
		resource, _, err := client.AgentApi.AgentShow(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}
		*Agent = resource
		return nil
	}
}

const testAccHyperoneAgentConfig = `
resource "hyperone_agent" "foo" {
	name = "test-agent"
	type = "container"
}
`
