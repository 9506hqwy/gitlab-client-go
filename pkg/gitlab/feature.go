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

type GetApiV4FeatureFlagsUnleashProjectIdParams struct {
	// InstanceId The instance ID of Unleash Client
	InstanceId *string `form:"instance_id,omitempty" json:"instance_id,omitempty"`

	// AppName The application name of Unleash Client
	AppName *string `form:"app_name,omitempty" json:"app_name,omitempty"`
}
type GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams struct {
	// InstanceId The instance ID of Unleash Client
	InstanceId *string `form:"instance_id,omitempty" json:"instance_id,omitempty"`

	// AppName The application name of Unleash Client
	AppName *string `form:"app_name,omitempty" json:"app_name,omitempty"`
}
type PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONBody struct {
	// AppName The application name of Unleash Client
	AppName *string `json:"app_name,omitempty"`

	// InstanceId The instance ID of Unleash Client
	InstanceId *string `json:"instance_id,omitempty"`
}
type PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONBody struct {
	// AppName The application name of Unleash Client
	AppName *string `json:"app_name,omitempty"`

	// InstanceId The instance ID of Unleash Client
	InstanceId *string `json:"instance_id,omitempty"`
}
type GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams struct {
	// InstanceId The instance ID of Unleash Client
	InstanceId *string `form:"instance_id,omitempty" json:"instance_id,omitempty"`

	// AppName The application name of Unleash Client
	AppName *string `form:"app_name,omitempty" json:"app_name,omitempty"`
}
type PostApiV4FeaturesNameJSONBody struct {
	// FeatureGroup A Feature group name
	FeatureGroup *string `json:"feature_group,omitempty"`

	// Force Skip feature flag validation checks, such as a YAML definition
	Force *bool `json:"force,omitempty"`

	// Group A GitLab group's path, for example `gitlab-org`, or comma-separated multiple group paths
	Group *string `json:"group,omitempty"`

	// Key `percentage_of_actors` or `percentage_of_time` (default)
	Key *string `json:"key,omitempty"`

	// Namespace A GitLab group or user namespace's path, for example `john-doe`, or comma-separated multiple namespace paths. Introduced in GitLab 15.0.
	Namespace *string `json:"namespace,omitempty"`

	// Project A projects path, for example `gitlab-org/gitlab-foss`, or comma-separated multiple project paths
	Project *string `json:"project,omitempty"`

	// Repository A repository path, for example `gitlab-org/gitlab-test.git`, `gitlab-org/gitlab-test.wiki.git`, `snippets/21.git`, to name a few. Use comma to separate multiple repository paths
	Repository *string `json:"repository,omitempty"`

	// User A GitLab username or comma-separated multiple usernames
	User *string `json:"user,omitempty"`

	// Value `true` or `false` to enable/disable, or an integer for percentage of time
	Value string `json:"value"`
}
type PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONBody
type PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONBody
type PostApiV4FeaturesNameJSONRequestBody PostApiV4FeaturesNameJSONBody
type GetApiV4FeatureFlagsUnleashProjectIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4FeaturesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		// Definition API_Entities_Feature_Definition model
		Definition *struct {
			DefaultEnabled      *string `json:"default_enabled,omitempty"`
			FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
			Group               *string `json:"group,omitempty"`
			IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
			IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
			LogStateChanges     *string `json:"log_state_changes,omitempty"`
			Milestone           *string `json:"milestone,omitempty"`
			Name                *string `json:"name,omitempty"`
			RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
			Type                *string `json:"type,omitempty"`
		} `json:"definition,omitempty"`
		Gates *struct {
			Key   *string `json:"key,omitempty"`
			Value *int32  `json:"value,omitempty"`
		} `json:"gates,omitempty"`
		Name  *string `json:"name,omitempty"`
		State *string `json:"state,omitempty"`
	}
}
type GetApiV4FeaturesDefinitionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		DefaultEnabled      *string `json:"default_enabled,omitempty"`
		FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
		Group               *string `json:"group,omitempty"`
		IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
		IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
		LogStateChanges     *string `json:"log_state_changes,omitempty"`
		Milestone           *string `json:"milestone,omitempty"`
		Name                *string `json:"name,omitempty"`
		RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
		Type                *string `json:"type,omitempty"`
	}
}
type DeleteApiV4FeaturesNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4FeaturesNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Definition API_Entities_Feature_Definition model
		Definition *struct {
			DefaultEnabled      *string `json:"default_enabled,omitempty"`
			FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
			Group               *string `json:"group,omitempty"`
			IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
			IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
			LogStateChanges     *string `json:"log_state_changes,omitempty"`
			Milestone           *string `json:"milestone,omitempty"`
			Name                *string `json:"name,omitempty"`
			RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
			Type                *string `json:"type,omitempty"`
		} `json:"definition,omitempty"`
		Gates *struct {
			Key   *string `json:"key,omitempty"`
			Value *int32  `json:"value,omitempty"`
		} `json:"gates,omitempty"`
		Name  *string `json:"name,omitempty"`
		State *string `json:"state,omitempty"`
	}
}

func (c *Client) GetApiV4FeatureFlagsUnleashProjectId(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4FeatureFlagsUnleashProjectIdRequest(c.Server, projectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesRequest(c.Server, projectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBody(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeatureFlagsUnleashProjectIdClientMetricsRequestWithBody(c.Server, projectId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeatureFlagsUnleashProjectIdClientMetrics(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeatureFlagsUnleashProjectIdClientMetricsRequest(c.Server, projectId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBody(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeatureFlagsUnleashProjectIdClientRegisterRequestWithBody(c.Server, projectId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeatureFlagsUnleashProjectIdClientRegister(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeatureFlagsUnleashProjectIdClientRegisterRequest(c.Server, projectId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4FeatureFlagsUnleashProjectIdFeaturesRequest(c.Server, projectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4Features(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4FeaturesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4FeaturesDefinitions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4FeaturesDefinitionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4FeaturesName(ctx context.Context, name int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4FeaturesNameRequest(c.Server, name)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeaturesNameWithBody(ctx context.Context, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeaturesNameRequestWithBody(c.Server, name, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4FeaturesName(ctx context.Context, name int32, body PostApiV4FeaturesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4FeaturesNameRequest(c.Server, name, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4FeatureFlagsUnleashProjectIdRequest(server string, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/feature_flags/unleash/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.InstanceId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "instance_id", runtime.ParamLocationQuery, *params.InstanceId); err != nil {
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

		if params.AppName != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_name", runtime.ParamLocationQuery, *params.AppName); err != nil {
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
func NewGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesRequest(server string, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/feature_flags/unleash/%s/client/features", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.InstanceId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "instance_id", runtime.ParamLocationQuery, *params.InstanceId); err != nil {
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

		if params.AppName != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_name", runtime.ParamLocationQuery, *params.AppName); err != nil {
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
func NewPostApiV4FeatureFlagsUnleashProjectIdClientMetricsRequest(server string, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4FeatureFlagsUnleashProjectIdClientMetricsRequestWithBody(server, projectId, "application/json", bodyReader)
}
func NewPostApiV4FeatureFlagsUnleashProjectIdClientMetricsRequestWithBody(server string, projectId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/feature_flags/unleash/%s/client/metrics", pathParam0)
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
func NewPostApiV4FeatureFlagsUnleashProjectIdClientRegisterRequest(server string, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4FeatureFlagsUnleashProjectIdClientRegisterRequestWithBody(server, projectId, "application/json", bodyReader)
}
func NewPostApiV4FeatureFlagsUnleashProjectIdClientRegisterRequestWithBody(server string, projectId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/feature_flags/unleash/%s/client/register", pathParam0)
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
func NewGetApiV4FeatureFlagsUnleashProjectIdFeaturesRequest(server string, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/feature_flags/unleash/%s/features", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.InstanceId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "instance_id", runtime.ParamLocationQuery, *params.InstanceId); err != nil {
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

		if params.AppName != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "app_name", runtime.ParamLocationQuery, *params.AppName); err != nil {
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
func NewGetApiV4FeaturesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/features")
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
func NewGetApiV4FeaturesDefinitionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/features/definitions")
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
func NewDeleteApiV4FeaturesNameRequest(server string, name int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/features/%s", pathParam0)
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
func NewPostApiV4FeaturesNameRequest(server string, name int32, body PostApiV4FeaturesNameJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4FeaturesNameRequestWithBody(server, name, "application/json", bodyReader)
}
func NewPostApiV4FeaturesNameRequestWithBody(server string, name int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/features/%s", pathParam0)
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
func (r GetApiV4FeatureFlagsUnleashProjectIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4FeatureFlagsUnleashProjectIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4FeaturesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4FeaturesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4FeaturesDefinitionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4FeaturesDefinitionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4FeaturesNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4FeaturesNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4FeaturesNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4FeaturesNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4FeatureFlagsUnleashProjectIdWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdResponse, error) {
	rsp, err := c.GetApiV4FeatureFlagsUnleashProjectId(ctx, projectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4FeatureFlagsUnleashProjectIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse, error) {
	rsp, err := c.GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx, projectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBodyWithResponse(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse, error) {
	rsp, err := c.PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBody(ctx, projectId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithResponse(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse, error) {
	rsp, err := c.PostApiV4FeatureFlagsUnleashProjectIdClientMetrics(ctx, projectId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBodyWithResponse(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse, error) {
	rsp, err := c.PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBody(ctx, projectId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithResponse(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse, error) {
	rsp, err := c.PostApiV4FeatureFlagsUnleashProjectIdClientRegister(ctx, projectId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4FeatureFlagsUnleashProjectIdFeaturesWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse, error) {
	rsp, err := c.GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx, projectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4FeaturesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4FeaturesResponse, error) {
	rsp, err := c.GetApiV4Features(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4FeaturesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4FeaturesDefinitionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4FeaturesDefinitionsResponse, error) {
	rsp, err := c.GetApiV4FeaturesDefinitions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4FeaturesDefinitionsResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4FeaturesNameWithResponse(ctx context.Context, name int32, reqEditors ...RequestEditorFn) (*DeleteApiV4FeaturesNameResponse, error) {
	rsp, err := c.DeleteApiV4FeaturesName(ctx, name, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4FeaturesNameResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeaturesNameWithBodyWithResponse(ctx context.Context, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeaturesNameResponse, error) {
	rsp, err := c.PostApiV4FeaturesNameWithBody(ctx, name, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeaturesNameResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4FeaturesNameWithResponse(ctx context.Context, name int32, body PostApiV4FeaturesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeaturesNameResponse, error) {
	rsp, err := c.PostApiV4FeaturesName(ctx, name, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4FeaturesNameResponse(rsp)
}
func ParseGetApiV4FeatureFlagsUnleashProjectIdResponse(rsp *http.Response) (*GetApiV4FeatureFlagsUnleashProjectIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4FeatureFlagsUnleashProjectIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse(rsp *http.Response) (*GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse(rsp *http.Response) (*PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse(rsp *http.Response) (*PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse(rsp *http.Response) (*GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4FeaturesResponse(rsp *http.Response) (*GetApiV4FeaturesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4FeaturesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			// Definition API_Entities_Feature_Definition model
			Definition *struct {
				DefaultEnabled      *string `json:"default_enabled,omitempty"`
				FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
				Group               *string `json:"group,omitempty"`
				IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
				IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
				LogStateChanges     *string `json:"log_state_changes,omitempty"`
				Milestone           *string `json:"milestone,omitempty"`
				Name                *string `json:"name,omitempty"`
				RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
				Type                *string `json:"type,omitempty"`
			} `json:"definition,omitempty"`
			Gates *struct {
				Key   *string `json:"key,omitempty"`
				Value *int32  `json:"value,omitempty"`
			} `json:"gates,omitempty"`
			Name  *string `json:"name,omitempty"`
			State *string `json:"state,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4FeaturesDefinitionsResponse(rsp *http.Response) (*GetApiV4FeaturesDefinitionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4FeaturesDefinitionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			DefaultEnabled      *string `json:"default_enabled,omitempty"`
			FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
			Group               *string `json:"group,omitempty"`
			IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
			IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
			LogStateChanges     *string `json:"log_state_changes,omitempty"`
			Milestone           *string `json:"milestone,omitempty"`
			Name                *string `json:"name,omitempty"`
			RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
			Type                *string `json:"type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4FeaturesNameResponse(rsp *http.Response) (*DeleteApiV4FeaturesNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4FeaturesNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4FeaturesNameResponse(rsp *http.Response) (*PostApiV4FeaturesNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4FeaturesNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Definition API_Entities_Feature_Definition model
			Definition *struct {
				DefaultEnabled      *string `json:"default_enabled,omitempty"`
				FeatureIssueUrl     *string `json:"feature_issue_url,omitempty"`
				Group               *string `json:"group,omitempty"`
				IntendedToRolloutBy *string `json:"intended_to_rollout_by,omitempty"`
				IntroducedByUrl     *string `json:"introduced_by_url,omitempty"`
				LogStateChanges     *string `json:"log_state_changes,omitempty"`
				Milestone           *string `json:"milestone,omitempty"`
				Name                *string `json:"name,omitempty"`
				RolloutIssueUrl     *string `json:"rollout_issue_url,omitempty"`
				Type                *string `json:"type,omitempty"`
			} `json:"definition,omitempty"`
			Gates *struct {
				Key   *string `json:"key,omitempty"`
				Value *int32  `json:"value,omitempty"`
			} `json:"gates,omitempty"`
			Name  *string `json:"name,omitempty"`
			State *string `json:"state,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
