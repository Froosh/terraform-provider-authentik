package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceStagePromptField(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceStagePromptFieldSimple,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_stage_prompt_field.default_user_settings_field_email", "name", "default-user-settings-field-email"),
				),
			},
		},
	})
}

const testAccDataSourceStagePromptFieldSimple = `
data "authentik_stage_prompt_field" "default_user_settings_field_email" {
  name = "default-user-settings-field-email"
}
`
