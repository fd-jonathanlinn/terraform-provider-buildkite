package buildkite

import (
	"context"
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
		// DeleteContext: DeletePipeline,
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
				Default:  nil,
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
				Default:  nil,
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
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slug": {
							Required: true,
							Type:     schema.TypeString,
						},
						"access_level": {
							Required: true,
							Type:     schema.TypeString,
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
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"provider_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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

	response, err := getOrganization(client.genqlient, client.organization)

	if err != nil {
		return diag.FromErr(err)
	}

	if response.Organization.Id == "" {
		return diag.FromErr(fmt.Errorf("organization not found: '%s'", client.organization))
	}

	apiResponse, err := createPipeline(
		client.genqlient,
		response.Organization.Id,
		d.Get("name").(string),
		d.Get("repository").(string),
		d.Get("steps").(PipelineStepsInput),
		d.Get("teams").([]PipelineTeamAssignmentInput),
		d.Get("cluster_id").(string),
		d.Get("description").(string),
		d.Get("skip_intermediate_builds").(bool),
		d.Get("skip_intermediate_builds_branch_filter").(string),
		d.Get("cancel_intermediate_builds").(bool),
		d.Get("cancel_intermediate_builds_branch_filter").(string),
		d.Get("visibility").(PipelineVisibility),
		d.Get("allow_rebuilds").(bool),
		d.Get("default_timeout_in_minutes").(int),
		d.Get("maximum_timeout_in_minutes").(int),
		d.Get("default_branch").(string),
		d.Get("tags").([]PipelineTagInput),
		d.Get("branch_configuration").(string),
	)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(apiResponse.PipelineCreate.Pipeline.Id)
	d.Set("slug", apiResponse.PipelineCreate.Pipeline.Slug)
	d.Set("url", apiResponse.PipelineCreate.Pipeline.Url)

	return diags
}

func ReadPipeline() {
	fmt.Println("hello")
}
