package hyperone

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	openapi "github.com/hyperonecom/h1-client-go"
)

func TestAccHyperoneAgentCertificate_basic(t *testing.T) {
	var n openapi.CredentialCertificate

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHyperoneAgentCertificateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccAgentCertificate("hyperone_agent_certificate.foo", &n),
				),
			},
		},
	})
}

func testAccAgentCertificate(n string, AgentCertificate *openapi.CredentialCertificate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*Config).client
		log.Printf("[DEBUG] Reading test agent credential: %s", rs.Primary.Attributes["agent"])

		resource, _, err := client.AgentApi.AgentGetCredentialcertificateId(context.TODO(), rs.Primary.Attributes["agent"], rs.Primary.ID)
		if err != nil {
			return err
		}
		*AgentCertificate = resource
		return nil
	}
}

const testAccHyperoneAgentCertificateConfig = `
resource "hyperone_agent" "foo" {
	name = "test-agent"
	type = "container"
}

resource "hyperone_agent_certificate" "foo" {
	agent = "${hyperone_agent.foo.id}"
	name = "my-key"
	type = "ssh"
	value = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCm37y9eLbnJaWdhC/b6Ji5XH8kckgyLDvFuzhO0G4o8Mw09x7QsXYLBwvseutEfYBpgxui8GUHjSmfKS3QNRHpsKESkEjyrh6gpS/Cgk7K97eOOLii3JXLE4n8ETFIN2cxumIqtxqQ3/6oW2o/F/kBJb+Mb+tVS+u2WFU/R62PIP80XoWwB8udibkvbJIyVKogwV8R4yKGTqiOOeAq87hBCZRO4XS5gSKdWGKHdejOURZ5uHFos30Irq5xdecUVGab8wSwib7KD17XcfaTmKz3waxRQdJrXpVXXbhVXwdH4wLhTQ6OWKjok1QqeUJ65kOBTdUZKAOAblfTVrEV8d27"
}
`
