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
<<<<<<< HEAD
=======

// ReadPipeline retrieves a Buildkite pipeline
func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := m.(*Client)

	orgPipelineSlug := fmt.Sprintf("%s/%s", client.organization, d.Get("slug").(string))
	pipeline, err := getPipeline(client.genqlient, orgPipelineSlug)

	if err != nil {
		return diag.FromErr(err)
	}

	if pipeline.Pipeline.Id == "" {
		return diag.FromErr(errors.New("Pipeline not found"))
	}

	d.SetId(pipeline.Pipeline.Id)
	d.Set("default_branch", pipeline.Pipeline.DefaultBranch)
	d.Set("description", pipeline.Pipeline.Description)
	d.Set("name", pipeline.Pipeline.Name)
	d.Set("repository", pipeline.Pipeline.Repository.Url)
	d.Set("slug", pipeline.Pipeline.Slug)
	d.Set("webhook_url", pipeline.Pipeline.WebhookURL)

	return diags
}
>>>>>>> origin/main
