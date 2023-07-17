package buildkite

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccOrganizationSettings_create(t *testing.T) {
	var o OrganizationResourceModel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		CheckDestroy:             testCheckOrganizationSettingsResourceRemoved,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationSettingsConfigBasic([]string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Confirm the Organization exists in the Buildkite API
					testAccCheckOrganizationSettingExists("buildkite_organization.let_them_in", &o),
					// Confirm that the allowed IP addresses are set correctly in Buildkite's system
					testAccCheckOrganizationSettingsRemoteValues(&o, []string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
					// Check that the second IP added to the list is the one we expect, 0.0.0.0/0, this also ensures the length is greater than 1
					// allowing us to assert the first IP is also added correctly
					resource.TestCheckResourceAttr("buildkite_organization.let_them_in", "allowed_api_ip_addresses.1", "1.1.1.1/32"),
				),
			},
		},
	})
}


func TestAccOrganizationSettings_update(t *testing.T) {
	var o OrganizationResourceModel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		CheckDestroy:             testCheckOrganizationSettingsResourceRemoved,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationSettingsConfigBasic([]string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Confirm the Organization exists in the Buildkite API
					testAccCheckOrganizationSettingExists("buildkite_organization.let_them_in", &o),
					// Confirm that the allowed IP addresses are set correctly in Buildkite's system
					testAccCheckOrganizationSettingsRemoteValues(&o, []string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
					// Check that the second IP added to the list is the one we expect, 0.0.0.0/0, this also ensures the length is greater than 1
					// allowing us to assert the first IP is also added correctly
					resource.TestCheckResourceAttr("buildkite_organization.let_them_in", "allowed_api_ip_addresses.2", "1.0.0.1/32"),
				),
			},

			{
				Config: testAccOrganizationSettingsConfigBasic([]string{"0.0.0.0/0", "4.4.4.4/32"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Confirm the Organization exists in the Buildkite API
					testAccCheckOrganizationSettingExists("buildkite_organization.let_them_in", &o),
					// Confirm that the allowed IP addresses are set correctly in Buildkite's system
					testAccCheckOrganizationSettingsRemoteValues(&o, []string{"0.0.0.0/0", "4.4.4.4/32"}),
					// This check allows us to ensure that TF still has access (0.0.0.0/0) and that the new IP address is added correctly
					resource.TestCheckResourceAttr("buildkite_organization.let_them_in", "allowed_api_ip_addresses.1", "4.4.4.4/32"),
				),
			},
		},
	})
}

func TestAccOrganizationSettings_import(t *testing.T) {
	var o OrganizationResourceModel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		CheckDestroy:             testCheckOrganizationSettingsResourceRemoved,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationSettingsConfigBasic([]string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Confirm the Organization exists in the Buildkite API
					testAccCheckOrganizationSettingExists("buildkite_organization.let_them_in", &o),
					// Confirm that the allowed IP addresses are set correctly in Buildkite's system
					testAccCheckOrganizationSettingsRemoteValues(&o, []string{"0.0.0.0/0", "1.1.1.1/32", "1.0.0.1/32"}),
					// Check that the second IP added to the list is the one we expect, 0.0.0.0/0, this also ensures the length is greater than 1
					// allowing us to assert the first IP is also added correctly
					resource.TestCheckResourceAttr("buildkite_organization.let_them_in", "allowed_api_ip_addresses.2", "1.0.0.1/32"),
				),
			},
			{
				ResourceName:      "buildkite_organization.let_them_in",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckOrganizationSettingExists(resourceName string, o *OrganizationResourceModel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found in state: %s", resourceName)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("No ID is set in state")
		}

		resp, err := getOrganization(genqlientGraphql, getenv("BUILDKITE_ORGANIZATION_SLUG"))

		// If organiztion was not able to be fetched by Genqlient
		if err != nil {
			return fmt.Errorf("Error fetching Organization from the GraphQL API: %v", err)
		}

		o.ID = types.StringValue(resp.Organization.Id)
		o.UUID = types.StringValue(resp.Organization.Uuid)
		ips, diag := types.ListValueFrom(context.Background(), types.StringType, strings.Split(resp.Organization.AllowedApiIpAddresses, " "))
		o.AllowedApiIpAddresses = ips

		if diag.HasError() {
			return fmt.Errorf("Error creating Allowed IP address list")
		}

		return nil
	}
}

func testAccOrganizationSettingsConfigBasic(ip_addresses []string) string {
	config := `
	
	resource "buildkite_organization" "let_them_in" {
        allowed_api_ip_addresses = %v
	}
	`
	marshal, _ := json.Marshal(ip_addresses)

	return fmt.Sprintf(config, string(marshal))
}

func testCheckOrganizationSettingsResourceRemoved(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "buildkite_organization" {
			continue
		}

		var getOrganizationQuery struct {
			Organization struct {
				AllowedApiIpAddresses string
			}
		}

		err := graphqlClient.Query(context.Background(), &getOrganizationQuery, map[string]interface{}{
			"slug": rs.Primary.ID,
		})

		if err == nil {
			return fmt.Errorf("Organization still exists")
		}
		return nil
	}
	return nil
}

func testAccCheckOrganizationSettingsRemoteValues(o *OrganizationResourceModel, ip_addresses []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Create a slice of strings for the AllowedApiIpAddresses CIDRs
		cidrs := make([]string, len(o.AllowedApiIpAddresses.Elements()))
		for i, v := range o.AllowedApiIpAddresses.Elements() {
			cidrs[i] = strings.Trim(v.String(), "\"")
		}

		if strings.Join(cidrs, " ") != strings.Join(ip_addresses, " ") {
			return fmt.Errorf("Allowed IP addresses do not match. Expected: %s, got: %s", cidrs, ip_addresses)

		}

		return nil
	}
}
