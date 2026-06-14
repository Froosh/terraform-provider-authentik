package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicyBinding(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	appName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicyBindingGroupConfig(appName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_policy_binding.group-test1", "order", "10"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.group-test2-20", "order", "20"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.group-test2-30", "order", "30"),
				),
			},
			{
				Config: testAccDataSourcePolicyBindingPolicyConfig(appName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_policy_binding.policy-test1", "order", "10"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.policy-test2-20", "order", "20"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.policy-test2-30", "order", "30"),
				),
			},
			{
				Config: testAccDataSourcePolicyBindingUserConfig(appName, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_policy_binding.user-test1", "order", "10"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.user-test2-20", "order", "20"),
					resource.TestCheckResourceAttr("data.authentik_policy_binding.user-test2-30", "order", "30"),
				),
			},
		},
	})
}

func testAccDataSourcePolicyBindingGroupConfig(appName string, groupName string) string {
	return fmt.Sprintf(`
resource "authentik_application" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "authentik_group" "test1" {
  name = "%[2]s-1"
}

resource "authentik_group" "test2" {
  name = "%[2]s-2"
}

resource "authentik_policy_binding" "group-test1" {
  group  = authentik_group.test1.id
  target = authentik_application.test.uuid
  order  = 10
}

resource "authentik_policy_binding" "group-test2-20" {
  group  = authentik_group.test2.id
  target = authentik_application.test.uuid
  order  = 20
}

resource "authentik_policy_binding" "group-test2-30" {
  group  = authentik_group.test2.id
  target = authentik_application.test.uuid
  order  = 30
}

data "authentik_policy_binding" "group-test1" {
  group  = authentik_policy_binding.group-test1.group
  target = authentik_policy_binding.group-test1.target
}

data "authentik_policy_binding" "group-test2-20" {
  group  = authentik_policy_binding.group-test2-20.group
  target = authentik_policy_binding.group-test2-20.target
  order  = authentik_policy_binding.group-test2-20.order
}

data "authentik_policy_binding" "group-test2-30" {
  group  = authentik_policy_binding.group-test2-30.group
  target = authentik_policy_binding.group-test2-30.target
  order  = authentik_policy_binding.group-test2-30.order
}
`, appName, groupName)
}

func testAccDataSourcePolicyBindingPolicyConfig(appName string, policyName string) string {
	return fmt.Sprintf(`
resource "authentik_application" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "authentik_policy_expression" "test1" {
  name = "%[2]s-1"
  expression = "return True"
}

resource "authentik_policy_expression" "test2" {
  name = "%[2]s-2"
  expression = "return True"
}

resource "authentik_policy_binding" "policy-test1" {
  policy = authentik_policy_expression.test1.id
  target = authentik_application.test.uuid
  order  = 10
}

resource "authentik_policy_binding" "policy-test2-20" {
  policy = authentik_policy_expression.test2.id
  target = authentik_application.test.uuid
  order  = 20
}

resource "authentik_policy_binding" "policy-test2-30" {
  policy = authentik_policy_expression.test2.id
  target = authentik_application.test.uuid
  order  = 30
}

data "authentik_policy_binding" "policy-test1" {
  policy = authentik_policy_binding.policy-test1.policy
  target = authentik_policy_binding.policy-test1.target
  order  = authentik_policy_binding.policy-test1.order
}

data "authentik_policy_binding" "policy-test2-20" {
  policy = authentik_policy_binding.policy-test2-20.policy
  target = authentik_policy_binding.policy-test2-20.target
  order  = authentik_policy_binding.policy-test2-20.order
}

data "authentik_policy_binding" "policy-test2-30" {
  policy = authentik_policy_binding.policy-test2-30.policy
  target = authentik_policy_binding.policy-test2-30.target
  order  = authentik_policy_binding.policy-test2-30.order
}
`, appName, policyName)
}

func testAccDataSourcePolicyBindingUserConfig(appName string, userName string) string {
	return fmt.Sprintf(`
resource "authentik_application" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "authentik_user" "test1" {
  username = "%[2]s"
}

resource "authentik_user" "test2" {
  username = "%[2]s-2"
}

resource "authentik_policy_binding" "user-test1" {
  user   = authentik_user.test1.id
  target = authentik_application.test.uuid
  order  = 10
}

resource "authentik_policy_binding" "user-test2-20" {
  user   = authentik_user.test2.id
  target = authentik_application.test.uuid
  order  = 20
}

resource "authentik_policy_binding" "user-test2-30" {
  user   = authentik_user.test2.id
  target = authentik_application.test.uuid
  order  = 30
}

data "authentik_policy_binding" "user-test1" {
  user   = authentik_policy_binding.user-test1.user
  target = authentik_policy_binding.user-test1.target
  order  = authentik_policy_binding.user-test1.order
}

data "authentik_policy_binding" "user-test2-20" {
  user   = authentik_policy_binding.user-test2-20.user
  target = authentik_policy_binding.user-test2-20.target
  order  = authentik_policy_binding.user-test2-20.order
}

data "authentik_policy_binding" "user-test2-30" {
  user   = authentik_policy_binding.user-test2-30.user
  target = authentik_policy_binding.user-test2-30.target
  order  = authentik_policy_binding.user-test2-30.order
}
`, appName, userName)
}
