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
	"time"

	"github.com/oapi-codegen/runtime"
)

type GetApiV4HooksParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4HooksJSONBody struct {
	// BranchFilterStrategy Filter push events by branch. Possible values are `wildcard` (default), `regex`, and `all_branches`
	BranchFilterStrategy *string `json:"branch_filter_strategy,omitempty"`

	// CustomHeaders Custom headers
	CustomHeaders *[]struct {
		// Key Name of the header
		Key string `json:"key"`

		// Value Value of the header
		Value string `json:"value"`
	} `json:"custom_headers,omitempty"`

	// Description Description of the hook
	Description *string `json:"description,omitempty"`

	// EnableSslVerification Do SSL verification when triggering the hook
	EnableSslVerification *bool `json:"enable_ssl_verification,omitempty"`

	// MergeRequestsEvents Trigger hook on merge requests events
	MergeRequestsEvents *bool `json:"merge_requests_events,omitempty"`

	// Name Name of the hook
	Name *string `json:"name,omitempty"`

	// PushEvents When true, the hook fires on push events
	PushEvents *bool `json:"push_events,omitempty"`

	// PushEventsBranchFilter Trigger hook on specified branch only
	PushEventsBranchFilter *string `json:"push_events_branch_filter,omitempty"`

	// RepositoryUpdateEvents Trigger hook on repository update events
	RepositoryUpdateEvents *bool `json:"repository_update_events,omitempty"`

	// TagPushEvents When true, the hook fires on new tags being pushed
	TagPushEvents *bool `json:"tag_push_events,omitempty"`

	// Token Secret token to validate received payloads; this isn't returned in the response
	Token *string `json:"token,omitempty"`

	// Url The URL to send the request to
	Url string `json:"url"`

	// UrlVariables URL variables for interpolation
	UrlVariables *[]struct {
		// Key Name of the variable
		Key string `json:"key"`

		// Value Value of the variable
		Value string `json:"value"`
	} `json:"url_variables,omitempty"`
}
type PutApiV4HooksHookIdJSONBody struct {
	// BranchFilterStrategy Filter push events by branch. Possible values are `wildcard` (default), `regex`, and `all_branches`
	BranchFilterStrategy *string `json:"branch_filter_strategy,omitempty"`

	// CustomHeaders Custom headers
	CustomHeaders *[]struct {
		// Key Name of the header
		Key string `json:"key"`

		// Value Value of the header
		Value string `json:"value"`
	} `json:"custom_headers,omitempty"`

	// Description Description of the hook
	Description *string `json:"description,omitempty"`

	// EnableSslVerification Do SSL verification when triggering the hook
	EnableSslVerification *bool `json:"enable_ssl_verification,omitempty"`

	// MergeRequestsEvents Trigger hook on merge requests events
	MergeRequestsEvents *bool `json:"merge_requests_events,omitempty"`

	// Name Name of the hook
	Name *string `json:"name,omitempty"`

	// PushEvents When true, the hook fires on push events
	PushEvents *bool `json:"push_events,omitempty"`

	// PushEventsBranchFilter Trigger hook on specified branch only
	PushEventsBranchFilter *string `json:"push_events_branch_filter,omitempty"`

	// RepositoryUpdateEvents Trigger hook on repository update events
	RepositoryUpdateEvents *bool `json:"repository_update_events,omitempty"`

	// TagPushEvents When true, the hook fires on new tags being pushed
	TagPushEvents *bool `json:"tag_push_events,omitempty"`

	// Token Secret token to validate received payloads; this isn't returned in the response
	Token *string `json:"token,omitempty"`

	// Url The URL to send the request to
	Url *string `json:"url,omitempty"`

	// UrlVariables URL variables for interpolation
	UrlVariables *[]struct {
		// Key Name of the variable
		Key string `json:"key"`

		// Value Value of the variable
		Value string `json:"value"`
	} `json:"url_variables,omitempty"`
}
type PutApiV4HooksHookIdCustomHeadersKeyJSONBody struct {
	// Value The value of the custom header
	Value string `json:"value"`
}
type PutApiV4HooksHookIdUrlVariablesKeyJSONBody struct {
	// Value The value of the variable
	Value string `json:"value"`
}
type PostApiV4HooksJSONRequestBody PostApiV4HooksJSONBody
type PutApiV4HooksHookIdJSONRequestBody PutApiV4HooksHookIdJSONBody
type PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody PutApiV4HooksHookIdCustomHeadersKeyJSONBody
type PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody PutApiV4HooksHookIdUrlVariablesKeyJSONBody
type GetApiV4HooksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		AlertStatus            *string                   `json:"alert_status,omitempty"`
		BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
		CreatedAt              *time.Time                `json:"created_at,omitempty"`
		CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
		CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
		Description            *string                   `json:"description,omitempty"`
		DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
		EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
		Id                     *string                   `json:"id,omitempty"`
		MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
		Name                   *string                   `json:"name,omitempty"`
		PushEvents             *bool                     `json:"push_events,omitempty"`
		PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
		RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
		TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
		Url                    *string                   `json:"url,omitempty"`
		UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
	}
}
type PostApiV4HooksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		AlertStatus            *string                   `json:"alert_status,omitempty"`
		BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
		CreatedAt              *time.Time                `json:"created_at,omitempty"`
		CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
		CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
		Description            *string                   `json:"description,omitempty"`
		DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
		EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
		Id                     *string                   `json:"id,omitempty"`
		MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
		Name                   *string                   `json:"name,omitempty"`
		PushEvents             *bool                     `json:"push_events,omitempty"`
		PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
		RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
		TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
		Url                    *string                   `json:"url,omitempty"`
		UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
	}
}
type DeleteApiV4HooksHookIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *struct {
		AlertStatus            *string                   `json:"alert_status,omitempty"`
		BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
		CreatedAt              *time.Time                `json:"created_at,omitempty"`
		CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
		CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
		Description            *string                   `json:"description,omitempty"`
		DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
		EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
		Id                     *string                   `json:"id,omitempty"`
		MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
		Name                   *string                   `json:"name,omitempty"`
		PushEvents             *bool                     `json:"push_events,omitempty"`
		PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
		RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
		TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
		Url                    *string                   `json:"url,omitempty"`
		UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
	}
}
type GetApiV4HooksHookIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AlertStatus            *string                   `json:"alert_status,omitempty"`
		BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
		CreatedAt              *time.Time                `json:"created_at,omitempty"`
		CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
		CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
		Description            *string                   `json:"description,omitempty"`
		DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
		EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
		Id                     *string                   `json:"id,omitempty"`
		MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
		Name                   *string                   `json:"name,omitempty"`
		PushEvents             *bool                     `json:"push_events,omitempty"`
		PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
		RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
		TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
		Url                    *string                   `json:"url,omitempty"`
		UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
	}
}
type PostApiV4HooksHookIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PutApiV4HooksHookIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AlertStatus            *string                   `json:"alert_status,omitempty"`
		BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
		CreatedAt              *time.Time                `json:"created_at,omitempty"`
		CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
		CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
		Description            *string                   `json:"description,omitempty"`
		DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
		EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
		Id                     *string                   `json:"id,omitempty"`
		MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
		Name                   *string                   `json:"name,omitempty"`
		PushEvents             *bool                     `json:"push_events,omitempty"`
		PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
		RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
		TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
		Url                    *string                   `json:"url,omitempty"`
		UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
	}
}
type DeleteApiV4HooksHookIdCustomHeadersKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PutApiV4HooksHookIdCustomHeadersKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type DeleteApiV4HooksHookIdUrlVariablesKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PutApiV4HooksHookIdUrlVariablesKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) GetApiV4Hooks(ctx context.Context, params *GetApiV4HooksParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4HooksRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4HooksWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4HooksRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Hooks(ctx context.Context, body PostApiV4HooksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4HooksRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4HooksHookIdRequest(c.Server, hookId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4HooksHookIdRequest(c.Server, hookId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4HooksHookIdRequest(c.Server, hookId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookIdWithBody(ctx context.Context, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdRequestWithBody(c.Server, hookId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookId(ctx context.Context, hookId int32, body PutApiV4HooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdRequest(c.Server, hookId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4HooksHookIdCustomHeadersKey(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4HooksHookIdCustomHeadersKeyRequest(c.Server, hookId, key)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookIdCustomHeadersKeyWithBody(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdCustomHeadersKeyRequestWithBody(c.Server, hookId, key, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookIdCustomHeadersKey(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdCustomHeadersKeyRequest(c.Server, hookId, key, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4HooksHookIdUrlVariablesKey(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4HooksHookIdUrlVariablesKeyRequest(c.Server, hookId, key)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookIdUrlVariablesKeyWithBody(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdUrlVariablesKeyRequestWithBody(c.Server, hookId, key, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4HooksHookIdUrlVariablesKey(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4HooksHookIdUrlVariablesKeyRequest(c.Server, hookId, key, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4HooksRequest(server string, params *GetApiV4HooksParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
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

		if params.PerPage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
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
func NewPostApiV4HooksRequest(server string, body PostApiV4HooksJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4HooksRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4HooksRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks")
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
func NewDeleteApiV4HooksHookIdRequest(server string, hookId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s", pathParam0)
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
func NewGetApiV4HooksHookIdRequest(server string, hookId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s", pathParam0)
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
func NewPostApiV4HooksHookIdRequest(server string, hookId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s", pathParam0)
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
func NewPutApiV4HooksHookIdRequest(server string, hookId int32, body PutApiV4HooksHookIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4HooksHookIdRequestWithBody(server, hookId, "application/json", bodyReader)
}
func NewPutApiV4HooksHookIdRequestWithBody(server string, hookId int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s", pathParam0)
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
func NewDeleteApiV4HooksHookIdCustomHeadersKeyRequest(server string, hookId int32, key string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s/custom_headers/%s", pathParam0, pathParam1)
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
func NewPutApiV4HooksHookIdCustomHeadersKeyRequest(server string, hookId int32, key string, body PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4HooksHookIdCustomHeadersKeyRequestWithBody(server, hookId, key, "application/json", bodyReader)
}
func NewPutApiV4HooksHookIdCustomHeadersKeyRequestWithBody(server string, hookId int32, key string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s/custom_headers/%s", pathParam0, pathParam1)
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
func NewDeleteApiV4HooksHookIdUrlVariablesKeyRequest(server string, hookId int32, key string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s/url_variables/%s", pathParam0, pathParam1)
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
func NewPutApiV4HooksHookIdUrlVariablesKeyRequest(server string, hookId int32, key string, body PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4HooksHookIdUrlVariablesKeyRequestWithBody(server, hookId, key, "application/json", bodyReader)
}
func NewPutApiV4HooksHookIdUrlVariablesKeyRequestWithBody(server string, hookId int32, key string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "hook_id", runtime.ParamLocationPath, hookId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/hooks/%s/url_variables/%s", pathParam0, pathParam1)
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
func (r GetApiV4HooksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4HooksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4HooksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4HooksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4HooksHookIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4HooksHookIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4HooksHookIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4HooksHookIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4HooksHookIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4HooksHookIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4HooksHookIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4HooksHookIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4HooksHookIdCustomHeadersKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4HooksHookIdCustomHeadersKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4HooksHookIdCustomHeadersKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4HooksHookIdCustomHeadersKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4HooksHookIdUrlVariablesKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4HooksHookIdUrlVariablesKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4HooksHookIdUrlVariablesKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4HooksHookIdUrlVariablesKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4HooksWithResponse(ctx context.Context, params *GetApiV4HooksParams, reqEditors ...RequestEditorFn) (*GetApiV4HooksResponse, error) {
	rsp, err := c.GetApiV4Hooks(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4HooksResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4HooksWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4HooksResponse, error) {
	rsp, err := c.PostApiV4HooksWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4HooksResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4HooksWithResponse(ctx context.Context, body PostApiV4HooksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4HooksResponse, error) {
	rsp, err := c.PostApiV4Hooks(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4HooksResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdResponse, error) {
	rsp, err := c.DeleteApiV4HooksHookId(ctx, hookId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4HooksHookIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*GetApiV4HooksHookIdResponse, error) {
	rsp, err := c.GetApiV4HooksHookId(ctx, hookId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4HooksHookIdResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*PostApiV4HooksHookIdResponse, error) {
	rsp, err := c.PostApiV4HooksHookId(ctx, hookId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4HooksHookIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdWithBodyWithResponse(ctx context.Context, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdResponse, error) {
	rsp, err := c.PutApiV4HooksHookIdWithBody(ctx, hookId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, body PutApiV4HooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdResponse, error) {
	rsp, err := c.PutApiV4HooksHookId(ctx, hookId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4HooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdCustomHeadersKeyResponse, error) {
	rsp, err := c.DeleteApiV4HooksHookIdCustomHeadersKey(ctx, hookId, key, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4HooksHookIdCustomHeadersKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdCustomHeadersKeyWithBodyWithResponse(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdCustomHeadersKeyResponse, error) {
	rsp, err := c.PutApiV4HooksHookIdCustomHeadersKeyWithBody(ctx, hookId, key, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdCustomHeadersKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdCustomHeadersKeyResponse, error) {
	rsp, err := c.PutApiV4HooksHookIdCustomHeadersKey(ctx, hookId, key, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdCustomHeadersKeyResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4HooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdUrlVariablesKeyResponse, error) {
	rsp, err := c.DeleteApiV4HooksHookIdUrlVariablesKey(ctx, hookId, key, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4HooksHookIdUrlVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdUrlVariablesKeyWithBodyWithResponse(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdUrlVariablesKeyResponse, error) {
	rsp, err := c.PutApiV4HooksHookIdUrlVariablesKeyWithBody(ctx, hookId, key, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdUrlVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4HooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdUrlVariablesKeyResponse, error) {
	rsp, err := c.PutApiV4HooksHookIdUrlVariablesKey(ctx, hookId, key, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4HooksHookIdUrlVariablesKeyResponse(rsp)
}
func ParseGetApiV4HooksResponse(rsp *http.Response) (*GetApiV4HooksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4HooksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			AlertStatus            *string                   `json:"alert_status,omitempty"`
			BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
			CreatedAt              *time.Time                `json:"created_at,omitempty"`
			CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
			CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
			Description            *string                   `json:"description,omitempty"`
			DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
			EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
			Id                     *string                   `json:"id,omitempty"`
			MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
			Name                   *string                   `json:"name,omitempty"`
			PushEvents             *bool                     `json:"push_events,omitempty"`
			PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
			RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
			TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
			Url                    *string                   `json:"url,omitempty"`
			UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4HooksResponse(rsp *http.Response) (*PostApiV4HooksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4HooksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			AlertStatus            *string                   `json:"alert_status,omitempty"`
			BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
			CreatedAt              *time.Time                `json:"created_at,omitempty"`
			CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
			CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
			Description            *string                   `json:"description,omitempty"`
			DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
			EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
			Id                     *string                   `json:"id,omitempty"`
			MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
			Name                   *string                   `json:"name,omitempty"`
			PushEvents             *bool                     `json:"push_events,omitempty"`
			PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
			RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
			TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
			Url                    *string                   `json:"url,omitempty"`
			UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4HooksHookIdResponse(rsp *http.Response) (*DeleteApiV4HooksHookIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4HooksHookIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest struct {
			AlertStatus            *string                   `json:"alert_status,omitempty"`
			BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
			CreatedAt              *time.Time                `json:"created_at,omitempty"`
			CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
			CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
			Description            *string                   `json:"description,omitempty"`
			DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
			EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
			Id                     *string                   `json:"id,omitempty"`
			MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
			Name                   *string                   `json:"name,omitempty"`
			PushEvents             *bool                     `json:"push_events,omitempty"`
			PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
			RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
			TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
			Url                    *string                   `json:"url,omitempty"`
			UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	}

	return response, nil
}
func ParseGetApiV4HooksHookIdResponse(rsp *http.Response) (*GetApiV4HooksHookIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4HooksHookIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AlertStatus            *string                   `json:"alert_status,omitempty"`
			BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
			CreatedAt              *time.Time                `json:"created_at,omitempty"`
			CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
			CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
			Description            *string                   `json:"description,omitempty"`
			DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
			EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
			Id                     *string                   `json:"id,omitempty"`
			MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
			Name                   *string                   `json:"name,omitempty"`
			PushEvents             *bool                     `json:"push_events,omitempty"`
			PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
			RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
			TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
			Url                    *string                   `json:"url,omitempty"`
			UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4HooksHookIdResponse(rsp *http.Response) (*PostApiV4HooksHookIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4HooksHookIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePutApiV4HooksHookIdResponse(rsp *http.Response) (*PutApiV4HooksHookIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4HooksHookIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AlertStatus            *string                   `json:"alert_status,omitempty"`
			BranchFilterStrategy   *string                   `json:"branch_filter_strategy,omitempty"`
			CreatedAt              *time.Time                `json:"created_at,omitempty"`
			CustomHeaders          *[]map[string]interface{} `json:"custom_headers,omitempty"`
			CustomWebhookTemplate  *string                   `json:"custom_webhook_template,omitempty"`
			Description            *string                   `json:"description,omitempty"`
			DisabledUntil          *time.Time                `json:"disabled_until,omitempty"`
			EnableSslVerification  *bool                     `json:"enable_ssl_verification,omitempty"`
			Id                     *string                   `json:"id,omitempty"`
			MergeRequestsEvents    *bool                     `json:"merge_requests_events,omitempty"`
			Name                   *string                   `json:"name,omitempty"`
			PushEvents             *bool                     `json:"push_events,omitempty"`
			PushEventsBranchFilter *string                   `json:"push_events_branch_filter,omitempty"`
			RepositoryUpdateEvents *bool                     `json:"repository_update_events,omitempty"`
			TagPushEvents          *bool                     `json:"tag_push_events,omitempty"`
			Url                    *string                   `json:"url,omitempty"`
			UrlVariables           *[]map[string]interface{} `json:"url_variables,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4HooksHookIdCustomHeadersKeyResponse(rsp *http.Response) (*DeleteApiV4HooksHookIdCustomHeadersKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4HooksHookIdCustomHeadersKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePutApiV4HooksHookIdCustomHeadersKeyResponse(rsp *http.Response) (*PutApiV4HooksHookIdCustomHeadersKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4HooksHookIdCustomHeadersKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseDeleteApiV4HooksHookIdUrlVariablesKeyResponse(rsp *http.Response) (*DeleteApiV4HooksHookIdUrlVariablesKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4HooksHookIdUrlVariablesKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePutApiV4HooksHookIdUrlVariablesKeyResponse(rsp *http.Response) (*PutApiV4HooksHookIdUrlVariablesKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4HooksHookIdUrlVariablesKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
