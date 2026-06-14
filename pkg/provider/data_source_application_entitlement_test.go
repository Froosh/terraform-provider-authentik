package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceApplicationEntitlement(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	appName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApplicationEntitlementConfig(appName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_application_entitlement.test", "name", rName),
				),
			},
		},
	})
}

func testAccDataSourceApplicationEntitlementConfig(appName string, entitlementName string) string {
	return fmt.Sprintf(`
resource "authentik_application" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "authentik_application_entitlement" "test" {
  application = authentik_application.test.uuid
  name        = "%[2]s"
}

data "authentik_application_entitlement" "test" {
  app = authentik_application_entitlement.test.application
  name = authentik_application_entitlement.test.name
}
`, appName, entitlementName)
}
