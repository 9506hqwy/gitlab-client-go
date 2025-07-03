package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

type PutApiV4ApplicationAppearanceMultipartBody struct {
	// Description Markdown text shown on the sign in / sign up page
	Description *string `json:"description,omitempty"`

	// EmailHeaderAndFooterEnabled Add header and footer to all outgoing emails if enabled
	EmailHeaderAndFooterEnabled *bool `json:"email_header_and_footer_enabled,omitempty"`

	// Favicon Instance favicon in .ico/.png format
	Favicon *string `json:"favicon,omitempty"`

	// FooterMessage Message within the system footer bar
	FooterMessage *string `json:"footer_message,omitempty"`

	// HeaderLogo Instance image used for the main navigation bar
	HeaderLogo *string `json:"header_logo,omitempty"`

	// HeaderMessage Message within the system header bar
	HeaderMessage *string `json:"header_message,omitempty"`

	// Logo Instance image used on the sign in / sign up page
	Logo *string `json:"logo,omitempty"`

	// MemberGuidelines Markdown text shown on the members page of a group or project
	MemberGuidelines *string `json:"member_guidelines,omitempty"`

	// MessageBackgroundColor Background color for the system header / footer bar
	MessageBackgroundColor *string `json:"message_background_color,omitempty"`

	// MessageFontColor Font color for the system header / footer bar
	MessageFontColor *string `json:"message_font_color,omitempty"`

	// NewProjectGuidelines Markdown text shown on the new project page
	NewProjectGuidelines *string `json:"new_project_guidelines,omitempty"`

	// ProfileImageGuidelines Markdown text shown on the profile page below Public Avatar
	ProfileImageGuidelines *string `json:"profile_image_guidelines,omitempty"`

	// PwaDescription An explanation of what the Progressive Web App does
	PwaDescription *string `json:"pwa_description,omitempty"`

	// PwaIcon Icon used for Progressive Web App
	PwaIcon *string `json:"pwa_icon,omitempty"`

	// PwaName Name of the Progressive Web App
	PwaName *string `json:"pwa_name,omitempty"`

	// PwaShortName Optional, short name for Progressive Web App
	PwaShortName *string `json:"pwa_short_name,omitempty"`

	// Title Instance title on the sign in / sign up page
	Title *string `json:"title,omitempty"`
}
type GetApiV4ApplicationPlanLimitsParams struct {
	// PlanName Name of the plan to get the limits from. Default: default.
	PlanName *string `form:"plan_name,omitempty" json:"plan_name,omitempty"`
}
type PutApiV4ApplicationPlanLimitsJSONBody struct {
	// CiActiveJobs Total number of jobs in currently active pipelines
	CiActiveJobs *int32 `json:"ci_active_jobs,omitempty"`

	// CiInstanceLevelVariables Maximum number of Instance-level CI/CD variables that can be defined
	CiInstanceLevelVariables *int32 `json:"ci_instance_level_variables,omitempty"`

	// CiNeedsSizeLimit Maximum number of needs dependencies that a job can have
	CiNeedsSizeLimit *int32 `json:"ci_needs_size_limit,omitempty"`

	// CiPipelineSchedules Maximum number of pipeline schedules
	CiPipelineSchedules *int32 `json:"ci_pipeline_schedules,omitempty"`

	// CiPipelineSize Maximum number of jobs in a single pipeline
	CiPipelineSize *int32 `json:"ci_pipeline_size,omitempty"`

	// CiProjectSubscriptions Maximum number of pipeline subscriptions to and from a project
	CiProjectSubscriptions *int32 `json:"ci_project_subscriptions,omitempty"`

	// CiRegisteredGroupRunners Maximum number of runners created or active in a group during the past seven days
	CiRegisteredGroupRunners *int32 `json:"ci_registered_group_runners,omitempty"`

	// CiRegisteredProjectRunners Maximum number of runners created or active in a project during the past seven days
	CiRegisteredProjectRunners *int32 `json:"ci_registered_project_runners,omitempty"`

	// ConanMaxFileSize Maximum Conan package file size in bytes
	ConanMaxFileSize *int32 `json:"conan_max_file_size,omitempty"`

	// DotenvSize Maximum size of a dotenv artifact in bytes
	DotenvSize *int32 `json:"dotenv_size,omitempty"`

	// DotenvVariables Maximum number of variables in a dotenv artifact
	DotenvVariables *int32 `json:"dotenv_variables,omitempty"`

	// EnforcementLimit Maximum storage size for the root namespace enforcement in MiB
	EnforcementLimit *int32 `json:"enforcement_limit,omitempty"`

	// GenericPackagesMaxFileSize Maximum generic package file size in bytes
	GenericPackagesMaxFileSize *int32 `json:"generic_packages_max_file_size,omitempty"`

	// HelmMaxFileSize Maximum Helm chart file size in bytes
	HelmMaxFileSize *int32 `json:"helm_max_file_size,omitempty"`

	// MavenMaxFileSize Maximum Maven package file size in bytes
	MavenMaxFileSize *int32 `json:"maven_max_file_size,omitempty"`

	// NotificationLimit Maximum storage size for the root namespace notifications in MiB
	NotificationLimit *int32 `json:"notification_limit,omitempty"`

	// NpmMaxFileSize Maximum NPM package file size in bytes
	NpmMaxFileSize *int32 `json:"npm_max_file_size,omitempty"`

	// NugetMaxFileSize Maximum NuGet package file size in bytes
	NugetMaxFileSize *int32 `json:"nuget_max_file_size,omitempty"`

	// PipelineHierarchySize Maximum number of downstream pipelines in a pipeline's hierarchy tree
	PipelineHierarchySize *int32 `json:"pipeline_hierarchy_size,omitempty"`

	// PlanName Name of the plan to update
	PlanName string `json:"plan_name"`

	// PypiMaxFileSize Maximum PyPI package file size in bytes
	PypiMaxFileSize *int32 `json:"pypi_max_file_size,omitempty"`

	// StorageSizeLimit Maximum storage size for the root namespace in MiB
	StorageSizeLimit *int32 `json:"storage_size_limit,omitempty"`

	// TerraformModuleMaxFileSize Maximum Terraform Module package file size in bytes
	TerraformModuleMaxFileSize *int32 `json:"terraform_module_max_file_size,omitempty"`
}
type PostApiV4ApplicationsJSONBody struct {
	// Confidential The application is used where the client secret can be kept confidential. Native mobile apps \
	//
	//
	//
	//
	//                         and Single Page Apps are considered non-confidential. Defaults to true if not supplied
	Confidential *bool `json:"confidential,omitempty"`

	// Name Name of the application.
	Name string `json:"name"`

	// RedirectUri Redirect URI of the application.
	RedirectUri string `json:"redirect_uri"`

	// Scopes Scopes of the application. You can specify multiple scopes by separating\
	//
	//
	//
	//
	//                                  each scope using a space
	Scopes string `json:"scopes"`
}
type PutApiV4ApplicationAppearanceMultipartRequestBody PutApiV4ApplicationAppearanceMultipartBody
type PutApiV4ApplicationPlanLimitsJSONRequestBody PutApiV4ApplicationPlanLimitsJSONBody
type PostApiV4ApplicationsJSONRequestBody PostApiV4ApplicationsJSONBody
type GetApiV4ApplicationAppearanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Description                 *string `json:"description,omitempty"`
		EmailHeaderAndFooterEnabled *string `json:"email_header_and_footer_enabled,omitempty"`
		Favicon                     *string `json:"favicon,omitempty"`
		FooterMessage               *string `json:"footer_message,omitempty"`
		HeaderLogo                  *string `json:"header_logo,omitempty"`
		HeaderMessage               *string `json:"header_message,omitempty"`
		Logo                        *string `json:"logo,omitempty"`
		MemberGuidelines            *string `json:"member_guidelines,omitempty"`
		MessageBackgroundColor      *string `json:"message_background_color,omitempty"`
		MessageFontColor            *string `json:"message_font_color,omitempty"`
		NewProjectGuidelines        *string `json:"new_project_guidelines,omitempty"`
		ProfileImageGuidelines      *string `json:"profile_image_guidelines,omitempty"`
		PwaDescription              *string `json:"pwa_description,omitempty"`
		PwaIcon                     *string `json:"pwa_icon,omitempty"`
		PwaName                     *string `json:"pwa_name,omitempty"`
		PwaShortName                *string `json:"pwa_short_name,omitempty"`
		Title                       *string `json:"title,omitempty"`
	}
}
type PutApiV4ApplicationAppearanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Description                 *string `json:"description,omitempty"`
		EmailHeaderAndFooterEnabled *string `json:"email_header_and_footer_enabled,omitempty"`
		Favicon                     *string `json:"favicon,omitempty"`
		FooterMessage               *string `json:"footer_message,omitempty"`
		HeaderLogo                  *string `json:"header_logo,omitempty"`
		HeaderMessage               *string `json:"header_message,omitempty"`
		Logo                        *string `json:"logo,omitempty"`
		MemberGuidelines            *string `json:"member_guidelines,omitempty"`
		MessageBackgroundColor      *string `json:"message_background_color,omitempty"`
		MessageFontColor            *string `json:"message_font_color,omitempty"`
		NewProjectGuidelines        *string `json:"new_project_guidelines,omitempty"`
		ProfileImageGuidelines      *string `json:"profile_image_guidelines,omitempty"`
		PwaDescription              *string `json:"pwa_description,omitempty"`
		PwaIcon                     *string `json:"pwa_icon,omitempty"`
		PwaName                     *string `json:"pwa_name,omitempty"`
		PwaShortName                *string `json:"pwa_short_name,omitempty"`
		Title                       *string `json:"title,omitempty"`
	}
}
type GetApiV4ApplicationPlanLimitsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CiActiveJobs               *int32                  `json:"ci_active_jobs,omitempty"`
		CiInstanceLevelVariables   *int32                  `json:"ci_instance_level_variables,omitempty"`
		CiNeedsSizeLimit           *int32                  `json:"ci_needs_size_limit,omitempty"`
		CiPipelineSchedules        *int32                  `json:"ci_pipeline_schedules,omitempty"`
		CiPipelineSize             *int32                  `json:"ci_pipeline_size,omitempty"`
		CiProjectSubscriptions     *int32                  `json:"ci_project_subscriptions,omitempty"`
		CiRegisteredGroupRunners   *int32                  `json:"ci_registered_group_runners,omitempty"`
		CiRegisteredProjectRunners *int32                  `json:"ci_registered_project_runners,omitempty"`
		ConanMaxFileSize           *int32                  `json:"conan_max_file_size,omitempty"`
		DotenvSize                 *int32                  `json:"dotenv_size,omitempty"`
		DotenvVariables            *int32                  `json:"dotenv_variables,omitempty"`
		EnforcementLimit           *int32                  `json:"enforcement_limit,omitempty"`
		GenericPackagesMaxFileSize *int32                  `json:"generic_packages_max_file_size,omitempty"`
		HelmMaxFileSize            *int32                  `json:"helm_max_file_size,omitempty"`
		LimitsHistory              *map[string]interface{} `json:"limits_history,omitempty"`
		MavenMaxFileSize           *int32                  `json:"maven_max_file_size,omitempty"`
		NotificationLimit          *int32                  `json:"notification_limit,omitempty"`
		NpmMaxFileSize             *int32                  `json:"npm_max_file_size,omitempty"`
		NugetMaxFileSize           *int32                  `json:"nuget_max_file_size,omitempty"`
		PipelineHierarchySize      *int32                  `json:"pipeline_hierarchy_size,omitempty"`
		PypiMaxFileSize            *int32                  `json:"pypi_max_file_size,omitempty"`
		StorageSizeLimit           *int32                  `json:"storage_size_limit,omitempty"`
		TerraformModuleMaxFileSize *int32                  `json:"terraform_module_max_file_size,omitempty"`
	}
}
type PutApiV4ApplicationPlanLimitsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CiActiveJobs               *int32                  `json:"ci_active_jobs,omitempty"`
		CiInstanceLevelVariables   *int32                  `json:"ci_instance_level_variables,omitempty"`
		CiNeedsSizeLimit           *int32                  `json:"ci_needs_size_limit,omitempty"`
		CiPipelineSchedules        *int32                  `json:"ci_pipeline_schedules,omitempty"`
		CiPipelineSize             *int32                  `json:"ci_pipeline_size,omitempty"`
		CiProjectSubscriptions     *int32                  `json:"ci_project_subscriptions,omitempty"`
		CiRegisteredGroupRunners   *int32                  `json:"ci_registered_group_runners,omitempty"`
		CiRegisteredProjectRunners *int32                  `json:"ci_registered_project_runners,omitempty"`
		ConanMaxFileSize           *int32                  `json:"conan_max_file_size,omitempty"`
		DotenvSize                 *int32                  `json:"dotenv_size,omitempty"`
		DotenvVariables            *int32                  `json:"dotenv_variables,omitempty"`
		EnforcementLimit           *int32                  `json:"enforcement_limit,omitempty"`
		GenericPackagesMaxFileSize *int32                  `json:"generic_packages_max_file_size,omitempty"`
		HelmMaxFileSize            *int32                  `json:"helm_max_file_size,omitempty"`
		LimitsHistory              *map[string]interface{} `json:"limits_history,omitempty"`
		MavenMaxFileSize           *int32                  `json:"maven_max_file_size,omitempty"`
		NotificationLimit          *int32                  `json:"notification_limit,omitempty"`
		NpmMaxFileSize             *int32                  `json:"npm_max_file_size,omitempty"`
		NugetMaxFileSize           *int32                  `json:"nuget_max_file_size,omitempty"`
		PipelineHierarchySize      *int32                  `json:"pipeline_hierarchy_size,omitempty"`
		PypiMaxFileSize            *int32                  `json:"pypi_max_file_size,omitempty"`
		StorageSizeLimit           *int32                  `json:"storage_size_limit,omitempty"`
		TerraformModuleMaxFileSize *int32                  `json:"terraform_module_max_file_size,omitempty"`
	}
}
type GetApiV4ApplicationStatisticsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ActiveUsers Number of active users
		ActiveUsers *int32 `json:"active_users,omitempty"`

		// Forks Approximate number of repo forks
		Forks *int32 `json:"forks,omitempty"`

		// Groups Approximate number of projects
		Groups *int32 `json:"groups,omitempty"`

		// Issues Approximate number of issues
		Issues *int32 `json:"issues,omitempty"`

		// MergeRequests Approximate number of merge requests
		MergeRequests *int32 `json:"merge_requests,omitempty"`

		// Milestones Approximate number of milestones
		Milestones *int32 `json:"milestones,omitempty"`

		// Notes Approximate number of notes
		Notes *int32 `json:"notes,omitempty"`

		// Projects Approximate number of projects
		Projects *int32 `json:"projects,omitempty"`

		// Snippets Approximate number of snippets
		Snippets *int32 `json:"snippets,omitempty"`

		// SshKeys Approximate number of SSH keys
		SshKeys *int32 `json:"ssh_keys,omitempty"`

		// Users Approximate number of users
		Users *int32 `json:"users,omitempty"`
	}
}
type GetApiV4ApplicationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		ApplicationId   *string `json:"application_id,omitempty"`
		ApplicationName *string `json:"application_name,omitempty"`
		CallbackUrl     *string `json:"callback_url,omitempty"`
		Confidential    *bool   `json:"confidential,omitempty"`
		Id              *string `json:"id,omitempty"`
	}
}
type PostApiV4ApplicationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ApplicationId   *string `json:"application_id,omitempty"`
		ApplicationName *string `json:"application_name,omitempty"`
		CallbackUrl     *string `json:"callback_url,omitempty"`
		Confidential    *bool   `json:"confidential,omitempty"`
		Id              *string `json:"id,omitempty"`
		Secret          *string `json:"secret,omitempty"`
	}
}
type DeleteApiV4ApplicationsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4ApplicationsIdRenewSecretResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ApplicationId   *string `json:"application_id,omitempty"`
		ApplicationName *string `json:"application_name,omitempty"`
		CallbackUrl     *string `json:"callback_url,omitempty"`
		Confidential    *bool   `json:"confidential,omitempty"`
		Id              *string `json:"id,omitempty"`
		Secret          *string `json:"secret,omitempty"`
	}
}

func (c *Client) GetApiV4ApplicationAppearance(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4ApplicationAppearanceRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4ApplicationAppearanceWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4ApplicationAppearanceRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4ApplicationPlanLimits(ctx context.Context, params *GetApiV4ApplicationPlanLimitsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4ApplicationPlanLimitsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4ApplicationPlanLimitsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4ApplicationPlanLimitsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4ApplicationPlanLimits(ctx context.Context, body PutApiV4ApplicationPlanLimitsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4ApplicationPlanLimitsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4ApplicationStatistics(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4ApplicationStatisticsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4Applications(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4ApplicationsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ApplicationsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ApplicationsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Applications(ctx context.Context, body PostApiV4ApplicationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ApplicationsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4ApplicationsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4ApplicationsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ApplicationsIdRenewSecret(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ApplicationsIdRenewSecretRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4ApplicationAppearanceRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/application/appearance")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPutApiV4ApplicationAppearanceRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/application/appearance")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4ApplicationPlanLimitsRequest(server string, params *GetApiV4ApplicationPlanLimitsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/application/plan_limits")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.PlanName != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "plan_name", runtime.ParamLocationQuery, *params.PlanName); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPutApiV4ApplicationPlanLimitsRequest(server string, body PutApiV4ApplicationPlanLimitsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4ApplicationPlanLimitsRequestWithBody(server, "application/json", bodyReader)
}
func NewPutApiV4ApplicationPlanLimitsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/application/plan_limits")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4ApplicationStatisticsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/application/statistics")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4ApplicationsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/applications")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4ApplicationsRequest(server string, body PostApiV4ApplicationsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ApplicationsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ApplicationsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/applications")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewDeleteApiV4ApplicationsIdRequest(server string, id int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/applications/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4ApplicationsIdRenewSecretRequest(server string, id int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/applications/%s/renew-secret", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (r GetApiV4ApplicationAppearanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4ApplicationAppearanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4ApplicationAppearanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4ApplicationAppearanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4ApplicationPlanLimitsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4ApplicationPlanLimitsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4ApplicationPlanLimitsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4ApplicationPlanLimitsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4ApplicationStatisticsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4ApplicationStatisticsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4ApplicationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4ApplicationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ApplicationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ApplicationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4ApplicationsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4ApplicationsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ApplicationsIdRenewSecretResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ApplicationsIdRenewSecretResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4ApplicationAppearanceWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationAppearanceResponse, error) {
	rsp, err := c.GetApiV4ApplicationAppearance(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4ApplicationAppearanceResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4ApplicationAppearanceWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationAppearanceResponse, error) {
	rsp, err := c.PutApiV4ApplicationAppearanceWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4ApplicationAppearanceResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4ApplicationPlanLimitsWithResponse(ctx context.Context, params *GetApiV4ApplicationPlanLimitsParams, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationPlanLimitsResponse, error) {
	rsp, err := c.GetApiV4ApplicationPlanLimits(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4ApplicationPlanLimitsResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4ApplicationPlanLimitsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationPlanLimitsResponse, error) {
	rsp, err := c.PutApiV4ApplicationPlanLimitsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4ApplicationPlanLimitsResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4ApplicationPlanLimitsWithResponse(ctx context.Context, body PutApiV4ApplicationPlanLimitsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationPlanLimitsResponse, error) {
	rsp, err := c.PutApiV4ApplicationPlanLimits(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4ApplicationPlanLimitsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4ApplicationStatisticsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationStatisticsResponse, error) {
	rsp, err := c.GetApiV4ApplicationStatistics(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4ApplicationStatisticsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4ApplicationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationsResponse, error) {
	rsp, err := c.GetApiV4Applications(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4ApplicationsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ApplicationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsResponse, error) {
	rsp, err := c.PostApiV4ApplicationsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ApplicationsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ApplicationsWithResponse(ctx context.Context, body PostApiV4ApplicationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsResponse, error) {
	rsp, err := c.PostApiV4Applications(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ApplicationsResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4ApplicationsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ApplicationsIdResponse, error) {
	rsp, err := c.DeleteApiV4ApplicationsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4ApplicationsIdResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ApplicationsIdRenewSecretWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsIdRenewSecretResponse, error) {
	rsp, err := c.PostApiV4ApplicationsIdRenewSecret(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ApplicationsIdRenewSecretResponse(rsp)
}
func ParseGetApiV4ApplicationAppearanceResponse(rsp *http.Response) (*GetApiV4ApplicationAppearanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4ApplicationAppearanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Description                 *string `json:"description,omitempty"`
			EmailHeaderAndFooterEnabled *string `json:"email_header_and_footer_enabled,omitempty"`
			Favicon                     *string `json:"favicon,omitempty"`
			FooterMessage               *string `json:"footer_message,omitempty"`
			HeaderLogo                  *string `json:"header_logo,omitempty"`
			HeaderMessage               *string `json:"header_message,omitempty"`
			Logo                        *string `json:"logo,omitempty"`
			MemberGuidelines            *string `json:"member_guidelines,omitempty"`
			MessageBackgroundColor      *string `json:"message_background_color,omitempty"`
			MessageFontColor            *string `json:"message_font_color,omitempty"`
			NewProjectGuidelines        *string `json:"new_project_guidelines,omitempty"`
			ProfileImageGuidelines      *string `json:"profile_image_guidelines,omitempty"`
			PwaDescription              *string `json:"pwa_description,omitempty"`
			PwaIcon                     *string `json:"pwa_icon,omitempty"`
			PwaName                     *string `json:"pwa_name,omitempty"`
			PwaShortName                *string `json:"pwa_short_name,omitempty"`
			Title                       *string `json:"title,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4ApplicationAppearanceResponse(rsp *http.Response) (*PutApiV4ApplicationAppearanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4ApplicationAppearanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Description                 *string `json:"description,omitempty"`
			EmailHeaderAndFooterEnabled *string `json:"email_header_and_footer_enabled,omitempty"`
			Favicon                     *string `json:"favicon,omitempty"`
			FooterMessage               *string `json:"footer_message,omitempty"`
			HeaderLogo                  *string `json:"header_logo,omitempty"`
			HeaderMessage               *string `json:"header_message,omitempty"`
			Logo                        *string `json:"logo,omitempty"`
			MemberGuidelines            *string `json:"member_guidelines,omitempty"`
			MessageBackgroundColor      *string `json:"message_background_color,omitempty"`
			MessageFontColor            *string `json:"message_font_color,omitempty"`
			NewProjectGuidelines        *string `json:"new_project_guidelines,omitempty"`
			ProfileImageGuidelines      *string `json:"profile_image_guidelines,omitempty"`
			PwaDescription              *string `json:"pwa_description,omitempty"`
			PwaIcon                     *string `json:"pwa_icon,omitempty"`
			PwaName                     *string `json:"pwa_name,omitempty"`
			PwaShortName                *string `json:"pwa_short_name,omitempty"`
			Title                       *string `json:"title,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4ApplicationPlanLimitsResponse(rsp *http.Response) (*GetApiV4ApplicationPlanLimitsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4ApplicationPlanLimitsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CiActiveJobs               *int32                  `json:"ci_active_jobs,omitempty"`
			CiInstanceLevelVariables   *int32                  `json:"ci_instance_level_variables,omitempty"`
			CiNeedsSizeLimit           *int32                  `json:"ci_needs_size_limit,omitempty"`
			CiPipelineSchedules        *int32                  `json:"ci_pipeline_schedules,omitempty"`
			CiPipelineSize             *int32                  `json:"ci_pipeline_size,omitempty"`
			CiProjectSubscriptions     *int32                  `json:"ci_project_subscriptions,omitempty"`
			CiRegisteredGroupRunners   *int32                  `json:"ci_registered_group_runners,omitempty"`
			CiRegisteredProjectRunners *int32                  `json:"ci_registered_project_runners,omitempty"`
			ConanMaxFileSize           *int32                  `json:"conan_max_file_size,omitempty"`
			DotenvSize                 *int32                  `json:"dotenv_size,omitempty"`
			DotenvVariables            *int32                  `json:"dotenv_variables,omitempty"`
			EnforcementLimit           *int32                  `json:"enforcement_limit,omitempty"`
			GenericPackagesMaxFileSize *int32                  `json:"generic_packages_max_file_size,omitempty"`
			HelmMaxFileSize            *int32                  `json:"helm_max_file_size,omitempty"`
			LimitsHistory              *map[string]interface{} `json:"limits_history,omitempty"`
			MavenMaxFileSize           *int32                  `json:"maven_max_file_size,omitempty"`
			NotificationLimit          *int32                  `json:"notification_limit,omitempty"`
			NpmMaxFileSize             *int32                  `json:"npm_max_file_size,omitempty"`
			NugetMaxFileSize           *int32                  `json:"nuget_max_file_size,omitempty"`
			PipelineHierarchySize      *int32                  `json:"pipeline_hierarchy_size,omitempty"`
			PypiMaxFileSize            *int32                  `json:"pypi_max_file_size,omitempty"`
			StorageSizeLimit           *int32                  `json:"storage_size_limit,omitempty"`
			TerraformModuleMaxFileSize *int32                  `json:"terraform_module_max_file_size,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4ApplicationPlanLimitsResponse(rsp *http.Response) (*PutApiV4ApplicationPlanLimitsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4ApplicationPlanLimitsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CiActiveJobs               *int32                  `json:"ci_active_jobs,omitempty"`
			CiInstanceLevelVariables   *int32                  `json:"ci_instance_level_variables,omitempty"`
			CiNeedsSizeLimit           *int32                  `json:"ci_needs_size_limit,omitempty"`
			CiPipelineSchedules        *int32                  `json:"ci_pipeline_schedules,omitempty"`
			CiPipelineSize             *int32                  `json:"ci_pipeline_size,omitempty"`
			CiProjectSubscriptions     *int32                  `json:"ci_project_subscriptions,omitempty"`
			CiRegisteredGroupRunners   *int32                  `json:"ci_registered_group_runners,omitempty"`
			CiRegisteredProjectRunners *int32                  `json:"ci_registered_project_runners,omitempty"`
			ConanMaxFileSize           *int32                  `json:"conan_max_file_size,omitempty"`
			DotenvSize                 *int32                  `json:"dotenv_size,omitempty"`
			DotenvVariables            *int32                  `json:"dotenv_variables,omitempty"`
			EnforcementLimit           *int32                  `json:"enforcement_limit,omitempty"`
			GenericPackagesMaxFileSize *int32                  `json:"generic_packages_max_file_size,omitempty"`
			HelmMaxFileSize            *int32                  `json:"helm_max_file_size,omitempty"`
			LimitsHistory              *map[string]interface{} `json:"limits_history,omitempty"`
			MavenMaxFileSize           *int32                  `json:"maven_max_file_size,omitempty"`
			NotificationLimit          *int32                  `json:"notification_limit,omitempty"`
			NpmMaxFileSize             *int32                  `json:"npm_max_file_size,omitempty"`
			NugetMaxFileSize           *int32                  `json:"nuget_max_file_size,omitempty"`
			PipelineHierarchySize      *int32                  `json:"pipeline_hierarchy_size,omitempty"`
			PypiMaxFileSize            *int32                  `json:"pypi_max_file_size,omitempty"`
			StorageSizeLimit           *int32                  `json:"storage_size_limit,omitempty"`
			TerraformModuleMaxFileSize *int32                  `json:"terraform_module_max_file_size,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4ApplicationStatisticsResponse(rsp *http.Response) (*GetApiV4ApplicationStatisticsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4ApplicationStatisticsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ActiveUsers Number of active users
			ActiveUsers *int32 `json:"active_users,omitempty"`

			// Forks Approximate number of repo forks
			Forks *int32 `json:"forks,omitempty"`

			// Groups Approximate number of projects
			Groups *int32 `json:"groups,omitempty"`

			// Issues Approximate number of issues
			Issues *int32 `json:"issues,omitempty"`

			// MergeRequests Approximate number of merge requests
			MergeRequests *int32 `json:"merge_requests,omitempty"`

			// Milestones Approximate number of milestones
			Milestones *int32 `json:"milestones,omitempty"`

			// Notes Approximate number of notes
			Notes *int32 `json:"notes,omitempty"`

			// Projects Approximate number of projects
			Projects *int32 `json:"projects,omitempty"`

			// Snippets Approximate number of snippets
			Snippets *int32 `json:"snippets,omitempty"`

			// SshKeys Approximate number of SSH keys
			SshKeys *int32 `json:"ssh_keys,omitempty"`

			// Users Approximate number of users
			Users *int32 `json:"users,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4ApplicationsResponse(rsp *http.Response) (*GetApiV4ApplicationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4ApplicationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			ApplicationId   *string `json:"application_id,omitempty"`
			ApplicationName *string `json:"application_name,omitempty"`
			CallbackUrl     *string `json:"callback_url,omitempty"`
			Confidential    *bool   `json:"confidential,omitempty"`
			Id              *string `json:"id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4ApplicationsResponse(rsp *http.Response) (*PostApiV4ApplicationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ApplicationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ApplicationId   *string `json:"application_id,omitempty"`
			ApplicationName *string `json:"application_name,omitempty"`
			CallbackUrl     *string `json:"callback_url,omitempty"`
			Confidential    *bool   `json:"confidential,omitempty"`
			Id              *string `json:"id,omitempty"`
			Secret          *string `json:"secret,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4ApplicationsIdResponse(rsp *http.Response) (*DeleteApiV4ApplicationsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4ApplicationsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4ApplicationsIdRenewSecretResponse(rsp *http.Response) (*PostApiV4ApplicationsIdRenewSecretResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ApplicationsIdRenewSecretResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ApplicationId   *string `json:"application_id,omitempty"`
			ApplicationName *string `json:"application_name,omitempty"`
			CallbackUrl     *string `json:"callback_url,omitempty"`
			Confidential    *bool   `json:"confidential,omitempty"`
			Id              *string `json:"id,omitempty"`
			Secret          *string `json:"secret,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
