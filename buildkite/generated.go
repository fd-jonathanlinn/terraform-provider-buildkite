// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package buildkite

import (
	"github.com/Khan/genqlient/graphql"
)

// The roles a user can be within a team
type GTeamMemberRole string

const (
	// The user is a regular member of the team
	GTeamMemberRoleMember GTeamMemberRole = "MEMBER"
	// The user can manage pipelines and users within the team
	GTeamMemberRoleMaintainer GTeamMemberRole = "MAINTAINER"
)

// Whether a team is visible or secret within an organization
type GenqlientTeamPrivacy string

const (
	// Visible to all members of the organization
	GenqlientTeamPrivacyVisible GenqlientTeamPrivacy = "VISIBLE"
	// Visible to organization administrators and members
	GenqlientTeamPrivacySecret GenqlientTeamPrivacy = "SECRET"
)

// The access levels that can be assigned to a pipeline
type PipelineAccessLevels string

const (
	// Allows edits, builds and reads
	PipelineAccessLevelsManageBuildAndRead PipelineAccessLevels = "MANAGE_BUILD_AND_READ"
	// Allows builds and read only
	PipelineAccessLevelsBuildAndRead PipelineAccessLevels = "BUILD_AND_READ"
	// Read only - no builds or edits
	PipelineAccessLevelsReadOnly PipelineAccessLevels = "READ_ONLY"
)

// Used to assign teams to pipelines
type PipelineTeamAssignmentInput struct {
	// Used to assign teams to pipelines
	Id string `json:"id"`
	// Used to assign teams to pipelines
	AccessLevel PipelineAccessLevels `json:"accessLevel"`
}

// GetId returns PipelineTeamAssignmentInput.Id, and is useful for accessing the field via an interface.
func (v *PipelineTeamAssignmentInput) GetId() string { return v.Id }

// GetAccessLevel returns PipelineTeamAssignmentInput.AccessLevel, and is useful for accessing the field via an interface.
func (v *PipelineTeamAssignmentInput) GetAccessLevel() PipelineAccessLevels { return v.AccessLevel }

// __createPipelineInput is used internally by genqlient
type __createPipelineInput struct {
	OrganizationId                       string                        `json:"organizationId"`
	Name                                 string                        `json:"name"`
	Repository                           string                        `json:"repository"`
	Steps                                map[string]string             `json:"steps"`
	Teams                                []PipelineTeamAssignmentInput `json:"teams"`
	ClusterId                            string                        `json:"clusterId"`
	Description                          string                        `json:"description"`
	SkipIntermediateBuilds               bool                          `json:"skipIntermediateBuilds"`
	SkipIntermediateBuildsBranchFilter   string                        `json:"skipIntermediateBuildsBranchFilter"`
	CancelIntermediateBuilds             bool                          `json:"cancelIntermediateBuilds"`
	CancelIntermediateBuildsBranchFilter string                        `json:"cancelIntermediateBuildsBranchFilter"`
	Visibility                           string                        `json:"visibility"`
	AllowRebuilds                        bool                          `json:"allowRebuilds"`
	DefaultTimeoutInMinutes              int                           `json:"defaultTimeoutInMinutes"`
	MaximumTimeoutInMinutes              int                           `json:"maximumTimeoutInMinutes"`
	DefaultBranch                        string                        `json:"defaultBranch"`
	BranchConfiguration                  string                        `json:"branchConfiguration"`
}

// GetOrganizationId returns __createPipelineInput.OrganizationId, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetOrganizationId() string { return v.OrganizationId }

// GetName returns __createPipelineInput.Name, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetName() string { return v.Name }

// GetRepository returns __createPipelineInput.Repository, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetRepository() string { return v.Repository }

// GetSteps returns __createPipelineInput.Steps, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetSteps() map[string]string { return v.Steps }

// GetTeams returns __createPipelineInput.Teams, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetTeams() []PipelineTeamAssignmentInput { return v.Teams }

// GetClusterId returns __createPipelineInput.ClusterId, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetClusterId() string { return v.ClusterId }

// GetDescription returns __createPipelineInput.Description, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetDescription() string { return v.Description }

// GetSkipIntermediateBuilds returns __createPipelineInput.SkipIntermediateBuilds, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetSkipIntermediateBuilds() bool { return v.SkipIntermediateBuilds }

// GetSkipIntermediateBuildsBranchFilter returns __createPipelineInput.SkipIntermediateBuildsBranchFilter, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetSkipIntermediateBuildsBranchFilter() string {
	return v.SkipIntermediateBuildsBranchFilter
}

// GetCancelIntermediateBuilds returns __createPipelineInput.CancelIntermediateBuilds, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetCancelIntermediateBuilds() bool { return v.CancelIntermediateBuilds }

// GetCancelIntermediateBuildsBranchFilter returns __createPipelineInput.CancelIntermediateBuildsBranchFilter, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetCancelIntermediateBuildsBranchFilter() string {
	return v.CancelIntermediateBuildsBranchFilter
}

// GetVisibility returns __createPipelineInput.Visibility, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetVisibility() string { return v.Visibility }

// GetAllowRebuilds returns __createPipelineInput.AllowRebuilds, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetAllowRebuilds() bool { return v.AllowRebuilds }

// GetDefaultTimeoutInMinutes returns __createPipelineInput.DefaultTimeoutInMinutes, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetDefaultTimeoutInMinutes() int { return v.DefaultTimeoutInMinutes }

// GetMaximumTimeoutInMinutes returns __createPipelineInput.MaximumTimeoutInMinutes, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetMaximumTimeoutInMinutes() int { return v.MaximumTimeoutInMinutes }

// GetDefaultBranch returns __createPipelineInput.DefaultBranch, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetDefaultBranch() string { return v.DefaultBranch }

// GetBranchConfiguration returns __createPipelineInput.BranchConfiguration, and is useful for accessing the field via an interface.
func (v *__createPipelineInput) GetBranchConfiguration() string { return v.BranchConfiguration }

// __deletePipelineInput is used internally by genqlient
type __deletePipelineInput struct {
	Id string `json:"id"`
}

// GetId returns __deletePipelineInput.Id, and is useful for accessing the field via an interface.
func (v *__deletePipelineInput) GetId() string { return v.Id }

// __getOrganizationInput is used internally by genqlient
type __getOrganizationInput struct {
	Slug string `json:"slug"`
}

// GetSlug returns __getOrganizationInput.Slug, and is useful for accessing the field via an interface.
func (v *__getOrganizationInput) GetSlug() string { return v.Slug }

// __getPipelineInput is used internally by genqlient
type __getPipelineInput struct {
	Slug string `json:"slug"`
}

// GetSlug returns __getPipelineInput.Slug, and is useful for accessing the field via an interface.
func (v *__getPipelineInput) GetSlug() string { return v.Slug }

// __getTeamInput is used internally by genqlient
type __getTeamInput struct {
	Slug string `json:"slug"`
}

// GetSlug returns __getTeamInput.Slug, and is useful for accessing the field via an interface.
func (v *__getTeamInput) GetSlug() string { return v.Slug }

// __setApiIpAddressesInput is used internally by genqlient
type __setApiIpAddressesInput struct {
	OrganizationID string `json:"organizationID"`
	IpAddresses    string `json:"ipAddresses"`
}

// GetOrganizationID returns __setApiIpAddressesInput.OrganizationID, and is useful for accessing the field via an interface.
func (v *__setApiIpAddressesInput) GetOrganizationID() string { return v.OrganizationID }

// GetIpAddresses returns __setApiIpAddressesInput.IpAddresses, and is useful for accessing the field via an interface.
func (v *__setApiIpAddressesInput) GetIpAddresses() string { return v.IpAddresses }

// createPipelinePipelineCreatePipelineCreatePayload includes the requested fields of the GraphQL type PipelineCreatePayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of PipelineCreate.
type createPipelinePipelineCreatePipelineCreatePayload struct {
	Pipeline createPipelinePipelineCreatePipelineCreatePayloadPipeline `json:"pipeline"`
}

// GetPipeline returns createPipelinePipelineCreatePipelineCreatePayload.Pipeline, and is useful for accessing the field via an interface.
func (v *createPipelinePipelineCreatePipelineCreatePayload) GetPipeline() createPipelinePipelineCreatePipelineCreatePayloadPipeline {
	return v.Pipeline
}

// createPipelinePipelineCreatePipelineCreatePayloadPipeline includes the requested fields of the GraphQL type Pipeline.
// The GraphQL type's documentation follows.
//
// A pipeline
type createPipelinePipelineCreatePipelineCreatePayloadPipeline struct {
	Id string `json:"id"`
	// The name of the pipeline
	Name string `json:"name"`
	// The URL for the pipeline
	Url string `json:"url"`
	// The slug of the pipeline
	Slug string `json:"slug"`
}

// GetId returns createPipelinePipelineCreatePipelineCreatePayloadPipeline.Id, and is useful for accessing the field via an interface.
func (v *createPipelinePipelineCreatePipelineCreatePayloadPipeline) GetId() string { return v.Id }

// GetName returns createPipelinePipelineCreatePipelineCreatePayloadPipeline.Name, and is useful for accessing the field via an interface.
func (v *createPipelinePipelineCreatePipelineCreatePayloadPipeline) GetName() string { return v.Name }

// GetUrl returns createPipelinePipelineCreatePipelineCreatePayloadPipeline.Url, and is useful for accessing the field via an interface.
func (v *createPipelinePipelineCreatePipelineCreatePayloadPipeline) GetUrl() string { return v.Url }

// GetSlug returns createPipelinePipelineCreatePipelineCreatePayloadPipeline.Slug, and is useful for accessing the field via an interface.
func (v *createPipelinePipelineCreatePipelineCreatePayloadPipeline) GetSlug() string { return v.Slug }

// createPipelineResponse is returned by createPipeline on success.
type createPipelineResponse struct {
	// Create a pipeline.
	PipelineCreate createPipelinePipelineCreatePipelineCreatePayload `json:"pipelineCreate"`
}

// GetPipelineCreate returns createPipelineResponse.PipelineCreate, and is useful for accessing the field via an interface.
func (v *createPipelineResponse) GetPipelineCreate() createPipelinePipelineCreatePipelineCreatePayload {
	return v.PipelineCreate
}

// deletePipelinePipelineDeletePipelineDeletePayload includes the requested fields of the GraphQL type PipelineDeletePayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of PipelineDelete.
type deletePipelinePipelineDeletePipelineDeletePayload struct {
	DeletedPipelineID string `json:"deletedPipelineID"`
	// A unique identifier for the client performing the mutation.
	ClientMutationId string `json:"clientMutationId"`
}

// GetDeletedPipelineID returns deletePipelinePipelineDeletePipelineDeletePayload.DeletedPipelineID, and is useful for accessing the field via an interface.
func (v *deletePipelinePipelineDeletePipelineDeletePayload) GetDeletedPipelineID() string {
	return v.DeletedPipelineID
}

// GetClientMutationId returns deletePipelinePipelineDeletePipelineDeletePayload.ClientMutationId, and is useful for accessing the field via an interface.
func (v *deletePipelinePipelineDeletePipelineDeletePayload) GetClientMutationId() string {
	return v.ClientMutationId
}

// deletePipelineResponse is returned by deletePipeline on success.
type deletePipelineResponse struct {
	// Delete a pipeline.
	PipelineDelete deletePipelinePipelineDeletePipelineDeletePayload `json:"pipelineDelete"`
}

// GetPipelineDelete returns deletePipelineResponse.PipelineDelete, and is useful for accessing the field via an interface.
func (v *deletePipelineResponse) GetPipelineDelete() deletePipelinePipelineDeletePipelineDeletePayload {
	return v.PipelineDelete
}

// getOrganizationOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An organization
type getOrganizationOrganization struct {
	// A space-separated allowlist of IP addresses that can access the organization via the GraphQL or REST API
	AllowedApiIpAddresses string `json:"allowedApiIpAddresses"`
	Id                    string `json:"id"`
	// The public UUID for this organization
	Uuid string `json:"uuid"`
}

// GetAllowedApiIpAddresses returns getOrganizationOrganization.AllowedApiIpAddresses, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetAllowedApiIpAddresses() string {
	return v.AllowedApiIpAddresses
}

// GetId returns getOrganizationOrganization.Id, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetId() string { return v.Id }

// GetUuid returns getOrganizationOrganization.Uuid, and is useful for accessing the field via an interface.
func (v *getOrganizationOrganization) GetUuid() string { return v.Uuid }

// getOrganizationResponse is returned by getOrganization on success.
type getOrganizationResponse struct {
	// Find an organization
	Organization getOrganizationOrganization `json:"organization"`
}

// GetOrganization returns getOrganizationResponse.Organization, and is useful for accessing the field via an interface.
func (v *getOrganizationResponse) GetOrganization() getOrganizationOrganization {
	return v.Organization
}

// getPipelinePipeline includes the requested fields of the GraphQL type Pipeline.
// The GraphQL type's documentation follows.
//
// A pipeline
type getPipelinePipeline struct {
	Id string `json:"id"`
	// The default branch for this pipeline
	DefaultBranch string `json:"defaultBranch"`
	// The short description of the pipeline
	Description string `json:"description"`
	// The name of the pipeline
	Name string `json:"name"`
	// The repository for this pipeline
	Repository getPipelinePipelineRepository `json:"repository"`
	// The slug of the pipeline
	Slug string `json:"slug"`
	// The URL to use in your repository settings for commit webhooks
	WebhookURL string `json:"webhookURL"`
}

// GetId returns getPipelinePipeline.Id, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetId() string { return v.Id }

// GetDefaultBranch returns getPipelinePipeline.DefaultBranch, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetDefaultBranch() string { return v.DefaultBranch }

// GetDescription returns getPipelinePipeline.Description, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetDescription() string { return v.Description }

// GetName returns getPipelinePipeline.Name, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetName() string { return v.Name }

// GetRepository returns getPipelinePipeline.Repository, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetRepository() getPipelinePipelineRepository { return v.Repository }

// GetSlug returns getPipelinePipeline.Slug, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetSlug() string { return v.Slug }

// GetWebhookURL returns getPipelinePipeline.WebhookURL, and is useful for accessing the field via an interface.
func (v *getPipelinePipeline) GetWebhookURL() string { return v.WebhookURL }

// getPipelinePipelineRepository includes the requested fields of the GraphQL type Repository.
// The GraphQL type's documentation follows.
//
// A repository associated with a pipeline
type getPipelinePipelineRepository struct {
	// The git URL for this repository
	Url string `json:"url"`
}

// GetUrl returns getPipelinePipelineRepository.Url, and is useful for accessing the field via an interface.
func (v *getPipelinePipelineRepository) GetUrl() string { return v.Url }

// getPipelineResponse is returned by getPipeline on success.
type getPipelineResponse struct {
	// Find a pipeline
	Pipeline getPipelinePipeline `json:"pipeline"`
}

// GetPipeline returns getPipelineResponse.Pipeline, and is useful for accessing the field via an interface.
func (v *getPipelineResponse) GetPipeline() getPipelinePipeline { return v.Pipeline }

// getTeamResponse is returned by getTeam on success.
type getTeamResponse struct {
	// Find a team
	Team getTeamTeam `json:"team"`
}

// GetTeam returns getTeamResponse.Team, and is useful for accessing the field via an interface.
func (v *getTeamResponse) GetTeam() getTeamTeam { return v.Team }

// getTeamTeam includes the requested fields of the GraphQL type Team.
// The GraphQL type's documentation follows.
//
// An organization team
type getTeamTeam struct {
	// New organization members will be granted this role on this team
	DefaultMemberRole GTeamMemberRole `json:"defaultMemberRole"`
	// A description of the team
	Description string `json:"description"`
	Id          string `json:"id"`
	// Add new organization members to this team by default
	IsDefaultTeam bool `json:"isDefaultTeam"`
	// Whether or not team members can create new pipelines in this team
	MembersCanCreatePipelines bool `json:"membersCanCreatePipelines"`
	// The name of the team
	Name string `json:"name"`
	// The privacy setting for this team
	Privacy GenqlientTeamPrivacy `json:"privacy"`
	// The slug of the team
	Slug string `json:"slug"`
	// The public UUID for this team
	Uuid string `json:"uuid"`
}

// GetDefaultMemberRole returns getTeamTeam.DefaultMemberRole, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetDefaultMemberRole() GTeamMemberRole { return v.DefaultMemberRole }

// GetDescription returns getTeamTeam.Description, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetDescription() string { return v.Description }

// GetId returns getTeamTeam.Id, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetId() string { return v.Id }

// GetIsDefaultTeam returns getTeamTeam.IsDefaultTeam, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetIsDefaultTeam() bool { return v.IsDefaultTeam }

// GetMembersCanCreatePipelines returns getTeamTeam.MembersCanCreatePipelines, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetMembersCanCreatePipelines() bool { return v.MembersCanCreatePipelines }

// GetName returns getTeamTeam.Name, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetName() string { return v.Name }

// GetPrivacy returns getTeamTeam.Privacy, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetPrivacy() GenqlientTeamPrivacy { return v.Privacy }

// GetSlug returns getTeamTeam.Slug, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetSlug() string { return v.Slug }

// GetUuid returns getTeamTeam.Uuid, and is useful for accessing the field via an interface.
func (v *getTeamTeam) GetUuid() string { return v.Uuid }

// setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload includes the requested fields of the GraphQL type OrganizationAPIIPAllowlistUpdateMutationPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of OrganizationAPIIPAllowlistUpdateMutation.
type setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload struct {
	Organization setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization `json:"organization"`
}

// GetOrganization returns setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload.Organization, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload) GetOrganization() setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization {
	return v.Organization
}

// setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An organization
type setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization struct {
	// A space-separated allowlist of IP addresses that can access the organization via the GraphQL or REST API
	AllowedApiIpAddresses string `json:"allowedApiIpAddresses"`
}

// GetAllowedApiIpAddresses returns setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization.AllowedApiIpAddresses, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayloadOrganization) GetAllowedApiIpAddresses() string {
	return v.AllowedApiIpAddresses
}

// setApiIpAddressesResponse is returned by setApiIpAddresses on success.
type setApiIpAddressesResponse struct {
	// Sets an allowlist of IP addresses for API access to an organization. Please note that this is a beta feature and is not yet available to all organizations.
	OrganizationApiIpAllowlistUpdate setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload `json:"organizationApiIpAllowlistUpdate"`
}

// GetOrganizationApiIpAllowlistUpdate returns setApiIpAddressesResponse.OrganizationApiIpAllowlistUpdate, and is useful for accessing the field via an interface.
func (v *setApiIpAddressesResponse) GetOrganizationApiIpAllowlistUpdate() setApiIpAddressesOrganizationApiIpAllowlistUpdateOrganizationAPIIPAllowlistUpdateMutationPayload {
	return v.OrganizationApiIpAllowlistUpdate
}

func createPipeline(
	client graphql.Client,
	organizationId string,
	name string,
	repository string,
	steps map[string]string,
	teams []PipelineTeamAssignmentInput,
	clusterId string,
	description string,
	skipIntermediateBuilds bool,
	skipIntermediateBuildsBranchFilter string,
	cancelIntermediateBuilds bool,
	cancelIntermediateBuildsBranchFilter string,
	visibility string,
	allowRebuilds bool,
	defaultTimeoutInMinutes int,
	maximumTimeoutInMinutes int,
	defaultBranch string,
	branchConfiguration string,
) (*createPipelineResponse, error) {
	__input := __createPipelineInput{
		OrganizationId:                       organizationId,
		Name:                                 name,
		Repository:                           repository,
		Steps:                                steps,
		Teams:                                teams,
		ClusterId:                            clusterId,
		Description:                          description,
		SkipIntermediateBuilds:               skipIntermediateBuilds,
		SkipIntermediateBuildsBranchFilter:   skipIntermediateBuildsBranchFilter,
		CancelIntermediateBuilds:             cancelIntermediateBuilds,
		CancelIntermediateBuildsBranchFilter: cancelIntermediateBuildsBranchFilter,
		Visibility:                           visibility,
		AllowRebuilds:                        allowRebuilds,
		DefaultTimeoutInMinutes:              defaultTimeoutInMinutes,
		MaximumTimeoutInMinutes:              maximumTimeoutInMinutes,
		DefaultBranch:                        defaultBranch,
		BranchConfiguration:                  branchConfiguration,
	}
	var err error

	var retval createPipelineResponse
	err = client.MakeRequest(
		nil,
		"createPipeline",
		`
mutation createPipeline ($organizationId: ID!, $name: String!, $repository: String!, $steps: PipelineStepsInput!, $teams: [PipelineTeamAssignmentInput!], $clusterId: ID, $description: String, $skipIntermediateBuilds: Boolean, $skipIntermediateBuildsBranchFilter: String, $cancelIntermediateBuilds: Boolean, $cancelIntermediateBuildsBranchFilter: String, $visibility: PipelineVisibility, $allowRebuilds: Boolean, $defaultTimeoutInMinutes: Int, $maximumTimeoutInMinutes: Int, $defaultBranch: String, $branchConfiguration: String) {
	pipelineCreate(input: {organizationId:$organizationId,name:$name,repository:{url:$repository},description:$description,steps:$steps,teams:$teams,clusterId:$clusterId,skipIntermediateBuilds:$skipIntermediateBuilds,skipIntermediateBuildsBranchFilter:$skipIntermediateBuildsBranchFilter,cancelIntermediateBuilds:$cancelIntermediateBuilds,cancelIntermediateBuildsBranchFilter:$cancelIntermediateBuildsBranchFilter,visibility:$visibility,allowRebuilds:$allowRebuilds,defaultTimeoutInMinutes:$defaultTimeoutInMinutes,maximumTimeoutInMinutes:$maximumTimeoutInMinutes,defaultBranch:$defaultBranch,branchConfiguration:$branchConfiguration}) {
		pipeline {
			id
			name
			url
			slug
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func deletePipeline(
	client graphql.Client,
	id string,
) (*deletePipelineResponse, error) {
	__input := __deletePipelineInput{
		Id: id,
	}
	var err error

	var retval deletePipelineResponse
	err = client.MakeRequest(
		nil,
		"deletePipeline",
		`
mutation deletePipeline ($id: ID!) {
	pipelineDelete(input: {id:$id}) {
		deletedPipelineID
		clientMutationId
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func getOrganization(
	client graphql.Client,
	slug string,
) (*getOrganizationResponse, error) {
	__input := __getOrganizationInput{
		Slug: slug,
	}
	var err error

	var retval getOrganizationResponse
	err = client.MakeRequest(
		nil,
		"getOrganization",
		`
query getOrganization ($slug: ID!) {
	organization(slug: $slug) {
		allowedApiIpAddresses
		id
		uuid
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func getPipeline(
	client graphql.Client,
	slug string,
) (*getPipelineResponse, error) {
	__input := __getPipelineInput{
		Slug: slug,
	}
	var err error

	var retval getPipelineResponse
	err = client.MakeRequest(
		nil,
		"getPipeline",
		`
query getPipeline ($slug: ID!) {
	pipeline(slug: $slug) {
		id
		defaultBranch
		description
		name
		repository {
			url
		}
		slug
		webhookURL
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func getTeam(
	client graphql.Client,
	slug string,
) (*getTeamResponse, error) {
	__input := __getTeamInput{
		Slug: slug,
	}
	var err error

	var retval getTeamResponse
	err = client.MakeRequest(
		nil,
		"getTeam",
		`
query getTeam ($slug: ID!) {
	team(slug: $slug) {
		defaultMemberRole
		description
		id
		isDefaultTeam
		membersCanCreatePipelines
		name
		privacy
		slug
		uuid
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func setApiIpAddresses(
	client graphql.Client,
	organizationID string,
	ipAddresses string,
) (*setApiIpAddressesResponse, error) {
	__input := __setApiIpAddressesInput{
		OrganizationID: organizationID,
		IpAddresses:    ipAddresses,
	}
	var err error

	var retval setApiIpAddressesResponse
	err = client.MakeRequest(
		nil,
		"setApiIpAddresses",
		`
mutation setApiIpAddresses ($organizationID: ID!, $ipAddresses: String!) {
	organizationApiIpAllowlistUpdate(input: {organizationID:$organizationID,ipAddresses:$ipAddresses}) {
		organization {
			allowedApiIpAddresses
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
