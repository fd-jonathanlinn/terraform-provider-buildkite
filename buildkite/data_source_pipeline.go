package buildkite

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePipeline() *schema.Resource {
	resource := resourcePipeline()
	return &schema.Resource{
		ReadContext: resource.ReadContext,
		Schema:      resource.Schema,
	}
}
