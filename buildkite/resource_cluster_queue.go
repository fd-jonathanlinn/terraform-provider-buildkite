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
	ID             types.String `tfsdk:"id"`
	UUID           types.String `tfsdk:"uuid"`
	OrganizationID types.String `tfsdk:"organizationId"`
	ClusterID      types.String `tfsdk:"clusterId"`
	Key            types.String `tfsdk:"key"`
	Description    types.String `tfsdk:"description"`
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
				Computed:            true,
				MarkdownDescription: "The public UUID of the Cluster Queue.",
			},
			"organizationId": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the Organization that the Cluster belongs to.",
			},
			"clusterId": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the Cluster that this Cluster Queue belongs to.",
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

	createReq := ClusterQueueCreateInput{
		OrganizationId: data.OrganizationID.ValueString(),
		ClusterId:      data.ClusterID.ValueString(),
		Description:    data.Description.ValueString(),
		Key:            data.Key.ValueString(),
	}

	r, err := createClusterQueue(c.client.genqlient, createReq)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Cluster Queue",
			fmt.Sprintf("Unable to create Cluster: %s", err.Error()),
		)
		return
	}

	data.ID = types.StringValue(r.ClusterQueueCreate.ClusterQueue.Id)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (c *ClusterQueueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	/* ... */
}

func (c *ClusterQueueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

} /* ... */

func (c *ClusterQueueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	/* ... */
}
