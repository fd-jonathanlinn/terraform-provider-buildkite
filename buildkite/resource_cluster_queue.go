package buildkite

import (
	"context"
	//"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClusterQueueStateModel struct {
	Id             types.String `tfsdk:"id"`
	Uuid           types.String `tfsdk:"uuid"`
	OrganizationId types.String `tfsdk:"organization_id"`
	ClusterId      types.String `tfsdk:"cluster_id"`
	Key            types.String `tfsdk:"key"`
	Description    types.String `tfsdk:"description"`
}

type CluseterQueueResource struct {
	client *Client
}

func NewClusterQueueResource() resource.Resource {
	return &CluseterQueueResource{}
}

func (CluseterQueueResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"organization_id": resource_schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The ID of the Organization that the Cluster belongs to.",
			},
			"cluster_id": resource_schema.StringAttribute{
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

func (CluseterQueueResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "buildkite_cluster_queue"
}

func (cq *CluseterQueueResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	cq.client = req.ProviderData.(*Client)
}

func (cq *CluseterQueueResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state ClusterQueueStateModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	
	apiResponse, err := createClusterQueue(
		cq.client.genqlient,
		cq.client.organizationId,
		plan.ClusterId.ValueString(),
		plan.Key.ValueString(),
		plan.Description.ValueStringPointer(),
	)

	if err != nil {
		resp.Diagnostics.AddError(err.Error(), err.Error())
	}

	state.Id = types.StringValue(apiResponse.ClusterQueueCreate.ClusterQueue.Id)
	state.Uuid = types.StringValue(apiResponse.ClusterQueueCreate.ClusterQueue.Uuid)
	state.OrganizationId = types.StringValue(cq.client.organizationId)
	state.ClusterId = plan.ClusterId
	state.Key = types.StringValue(apiResponse.ClusterQueueCreate.ClusterQueue.Key)
	state.Description = types.StringPointerValue(&apiResponse.ClusterQueueCreate.ClusterQueue.Description)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (cq *CluseterQueueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// To implement
}

func (cq *CluseterQueueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state ClusterQueueStateModel
	var description string

	//Load state and ontain description from plan (singularly)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("description"), &description)...)

	apiResponse, err := updateClusterQueue(
		cq.client.genqlient,
		cq.client.organizationId,
		state.Id.ValueString(),
		&description,
	)

	if err != nil {
		resp.Diagnostics.AddError(err.Error(), err.Error())
	}

	state.Description = types.StringPointerValue(&apiResponse.ClusterQueueUpdate.ClusterQueue.Description)
	
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (cq *CluseterQueueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan ClusterQueueStateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	_, err := deleteClusterQueue(
		cq.client.genqlient,
		cq.client.organizationId,
		plan.Id.ValueString(),
	)

	if err != nil {
		resp.Diagnostics.AddError(err.Error(), err.Error())
	}
}
