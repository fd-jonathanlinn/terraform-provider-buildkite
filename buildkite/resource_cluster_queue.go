package buildkite

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type ClusterQueueResourceModel struct {
	ID             types.String `tfsdk:"id"`
	UUID           types.String `tfsdk:"uuid"`
	OrganizationID types.String `tfsdk:"organizationId"`
	ClusterID      types.String `tfsdk:"clusterId"`
	Key            types.String `tfsdk:"key"`
	Description    types.String `tfsdk:"description"`
}

type CluseterQueue struct {
	client *Client
}

func NewClusterQueue() resource.Resource {
	return &CluseterQueue{}
}

func (CluseterQueue) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_schema.Schema{
		MarkdownDescription: "A Cluster Queue is a queue belonging to a specific Cluster for its Agents to target builds on. ",
		Attributes: map[string]resource_schema.Attribute{
			"id": resource_schema.StringAttribute{
				Computed: true,
			},
			"uuid": resource_schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The public UUID of the Cluster Queue.",
			},
			"organizationId": resource_schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the Organization that the Cluster belongs to.",
			},
			"clusterId": resource_schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the Cluster that this Cluster Queue belongs to.",
			},
			"key": resource_schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The key of the Cluster Queue.",
			},
			"description": resource_schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A description for the Cluster Queue. ",
			},
		},
	}
}

func (CluseterQueue) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "buildkite_cluster_queue"
}

func (cq *CluseterQueue) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	cq.client = req.ProviderData.(*Client)
}

func (cq *CluseterQueue) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

	r, err := createClusterQueue(cq.client.genqlient, createReq)

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

func (cq *CluseterQueue) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	/* ... */
}

func (cq *CluseterQueue) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

} /* ... */

func (cq *CluseterQueue) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	/* ... */
}
