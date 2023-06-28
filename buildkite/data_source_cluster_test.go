package buildkite

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCluster_read(t *testing.T) {
	t.Parallel()
	var c ClusterDataSourceModel

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccClusterBasic("foo"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckClusterRemoteValues(c.ID.ValueString()),
					// Confirm the Cluster data source has the correct values in terraform state
					resource.TestCheckResourceAttr("data.buildkite_cluster", "name", "foos_test_cluster"),
					resource.TestCheckResourceAttr("data.buildkite_cluster", "description", "Test cluster"),
				),
			},
		},
	})
}