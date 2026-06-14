package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicyExpression(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicyExpressionSimple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_policy_expression.default_user_settings_authorization", "name", "default-user-settings-authorization"),
				),
			},
		},
	})
}

const testAccDataSourcePolicyExpressionSimple = `
data "authentik_policy_expression" "default_user_settings_authorization" {
  name = "default-user-settings-authorization"
}
`
