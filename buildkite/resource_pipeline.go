package buildkite

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const defaultSteps = `steps:
- label: ':pipeline: Pipeline Upload'
  command: buildkite-agent pipeline upload`

// resourcePipeline represents the terraform pipeline resource schema
func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreatePipeline,
		ReadContext:   ReadPipeline,
		UpdateContext: CreatePipeline,
		DeleteContext: DeletePipeline,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"allow_rebuilds": {
				Optional: true,
				Type:     schema.TypeBool,
			},
			"cancel_intermediate_builds": {
				Optional: true,
				Type:     schema.TypeBool,
			},
			"cancel_intermediate_builds_branch_filter": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"branch_configuration": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"cluster_id": {
				Optional: true,
				Default:  nil,
				Type:     schema.TypeString,
			},
			"default_branch": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"default_timeout_in_minutes": {
				Optional: true,
				Default:  20,
				Type:     schema.TypeInt,
			},
			"deletion_protection": {
				Optional:    true,
				Default:     false,
				Type:        schema.TypeBool,
				Description: "If set to 'true', deletion of a pipeline via `terraform destroy` will fail, until set to 'false'.",
			},
			"maximum_timeout_in_minutes": {
				Optional: true,
				Default:  60,
				Type:     schema.TypeInt,
			},
			"description": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"name": {
				Required: true,
				Type:     schema.TypeString,
			},
			"repository": {
				Required: true,
				Type:     schema.TypeString,
			},
			"skip_intermediate_builds": {
				Optional: true,
				Type:     schema.TypeBool,
			},
			"skip_intermediate_builds_branch_filter": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"slug": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"steps": {
				Optional: true,
				Default:  defaultSteps,
				Type:     schema.TypeString,
			},
			"team": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Required: true,
							Type:     schema.TypeString,
						},
						"access_level": {
							Optional: true,
							Type:     schema.TypeString,
							Default:  "READ_ONLY",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								switch v := val.(string); v {
								case "READ_ONLY":
								case "BUILD_AND_READ":
								case "MANAGE_BUILD_AND_READ":
									return
								default:
									errs = append(errs, fmt.Errorf("%q must be one of READ_ONLY, BUILD_AND_READ or MANAGE_BUILD_AND_READ, got: %s", key, v))
									return
								}
								return
							},
						},
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"provider_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trigger_mode": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								switch v := val.(string); v {
								case "code":
								case "deployment":
								case "fork":
								case "none":
									return
								default:
									errs = append(errs, fmt.Errorf("%q must be one of code, deployment, fork or none, got: %s", key, v))
									return
								}
								return
							},
						},
						"build_pull_requests": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"pull_request_branch_filter_enabled": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"pull_request_branch_filter_configuration": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
						"skip_pull_request_builds_for_existing_commits": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"build_pull_request_ready_for_review": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"build_pull_request_labels_changed": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"build_pull_request_forks": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"prefix_pull_request_fork_branch_names": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"build_branches": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"build_tags": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"cancel_deleted_branch_builds": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"filter_enabled": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"filter_condition": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeString,
						},
						"publish_commit_status": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"publish_blocked_as_pending": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"publish_commit_status_per_step": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
						"separate_pull_request_statuses": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeBool,
						},
					},
				},
			},
			"webhook_url": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"visibility": {
				Optional: true,
				Type:     schema.TypeString,
				Default:  "PRIVATE",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					switch v := val.(string); v {
					case "PRIVATE":
					case "PUBLIC":
						return
					default:
						errs = append(errs, fmt.Errorf("%q must be one of PRIVATE or PUBLIC, got: %s", key, v))
						return
					}
					return
				},
			},
			"badge_url": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

// CreatePipeline creates a Buildkite pipeline
func CreatePipeline(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Client)

	fmt.Printf("Teams data: %#v", d.Get("team"))

	steps := map[string]string{
		"yaml": d.Get("steps").(string),
	}

	teams := d.Get("team").(*schema.Set).List()

	pipelineTeams := make([]PipelineTeamAssignmentInput, len(teams))

	for i, team := range teams {
		teamMap := team.(map[string]interface{})

		pipelineTeams[i] = PipelineTeamAssignmentInput{
			Id:          teamMap["id"].(string),
			AccessLevel: PipelineAccessLevels(teamMap["access_level"].(string)),
		}
	}

	apiResponse, err := createPipeline(
		client.genqlient,
		client.organizationId,
		d.Get("name").(string),
		d.Get("repository").(string),
		steps,
		pipelineTeams,
		d.Get("cluster_id").(string),
		d.Get("description").(string),
		d.Get("skip_intermediate_builds").(bool),
		d.Get("skip_intermediate_builds_branch_filter").(string),
		d.Get("cancel_intermediate_builds").(bool),
		d.Get("cancel_intermediate_builds_branch_filter").(string),
		d.Get("visibility").(string),
		d.Get("allow_rebuilds").(bool),
		d.Get("default_timeout_in_minutes").(int),
		d.Get("maximum_timeout_in_minutes").(int),
		d.Get("default_branch").(string),
		d.Get("branch_configuration").(string),
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(apiResponse.PipelineCreate.Pipeline.Id)
	d.Set("name", apiResponse.PipelineCreate.Pipeline.Name)
	d.Set("slug", apiResponse.PipelineCreate.Pipeline.Slug)
	d.Set("url", apiResponse.PipelineCreate.Pipeline.Url)

	return diags
}

func ReadPipeline(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := m.(*Client)

	orgPipelineSlug := fmt.Sprintf("%s/%s", client.organization, d.Get("slug").(string))
	pipeline, err := getPipeline(client.genqlient, orgPipelineSlug)

	if err != nil {
		return diag.FromErr(err)
	}

	if pipeline.Pipeline.Id == "" {
		return diag.FromErr(errors.New("pipeline not found"))
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

func DeletePipeline(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*Client)

	resp, err := getPipeline(client.genqlient, fmt.Sprintf("%s/%s", client.organization, d.Get("slug").(string)))
	if err != nil {
		return diag.FromErr(err)
	}

	pipelineId := resp.Pipeline.Id

	_, err = deletePipeline(client.genqlient, pipelineId)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
