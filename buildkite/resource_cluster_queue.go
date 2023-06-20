package buildkite

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClusterQueueResource struct {
	client *Client
}

type ClusterQueueResourceModel struct {
	ID          types.String `tfsdk:"id"`
	UUID		types.String `tfsdk:"uuid"`
	Key	        types.String `tfsdk:"key"`
	Description types.String `tfsdk:"description"`
}

func (c *ClusterQueueResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "_clusterqueue"
}

func (c *ClusterQueueResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A Cluster Queue is a queue belonging to a specific Cluster for its Agents to target builds on. ",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"uuid": schema.StringAttribute{
				Computed: 			 true,
				MarkdownDescription: "The public UUID of the Cluster Queue.",
			},
			"key": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The key of the Cluster Queue.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A description for the Cluster Queue. ",
			},
		},
	}
}

func (c *ClusterQueueResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ClusterQueueResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (c *ClusterQueueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	/* ... */
}

func (c *ClusterQueueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	
}	/* ... */

func (c *ClusterQueueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	/* ... */
}
