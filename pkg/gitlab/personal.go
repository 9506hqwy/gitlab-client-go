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
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type GetApiV4PersonalAccessTokensParams struct {
	// UserId Filter PATs by User ID
	UserId *int32 `form:"user_id,omitempty" json:"user_id,omitempty"`

	// Revoked Filter tokens where revoked state matches parameter
	Revoked *bool `form:"revoked,omitempty" json:"revoked,omitempty"`

	// State Filter tokens which are either active or not
	State *string `form:"state,omitempty" json:"state,omitempty"`

	// CreatedBefore Filter tokens which were created before given datetime
	CreatedBefore *time.Time `form:"created_before,omitempty" json:"created_before,omitempty"`

	// CreatedAfter Filter tokens which were created after given datetime
	CreatedAfter *time.Time `form:"created_after,omitempty" json:"created_after,omitempty"`

	// LastUsedBefore Filter tokens which were used before given datetime
	LastUsedBefore *time.Time `form:"last_used_before,omitempty" json:"last_used_before,omitempty"`

	// LastUsedAfter Filter tokens which were used after given datetime
	LastUsedAfter *time.Time `form:"last_used_after,omitempty" json:"last_used_after,omitempty"`

	// ExpiresBefore Filter tokens which expire before given datetime
	ExpiresBefore *openapi_types.Date `form:"expires_before,omitempty" json:"expires_before,omitempty"`

	// ExpiresAfter Filter tokens which expire after given datetime
	ExpiresAfter *openapi_types.Date `form:"expires_after,omitempty" json:"expires_after,omitempty"`

	// Search Filters tokens by name
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// Sort Sort tokens
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4PersonalAccessTokensSelfAssociationsParams struct {
	// MinAccessLevel Limit by minimum access level of authenticated user
	MinAccessLevel *int32 `form:"min_access_level,omitempty" json:"min_access_level,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4PersonalAccessTokensSelfRotateJSONBody struct {
	// ExpiresAt The expiration date of the token
	ExpiresAt *openapi_types.Date `json:"expires_at,omitempty"`
}
type PostApiV4PersonalAccessTokensIdRotateJSONBody struct {
	// ExpiresAt The expiration date of the token
	ExpiresAt *openapi_types.Date `json:"expires_at,omitempty"`
}
type PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody PostApiV4PersonalAccessTokensSelfRotateJSONBody
type PostApiV4PersonalAccessTokensIdRotateJSONRequestBody PostApiV4PersonalAccessTokensIdRotateJSONBody
type GetApiV4PersonalAccessTokensResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}
type DeleteApiV4PersonalAccessTokensSelfResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4PersonalAccessTokensSelfResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}
type GetApiV4PersonalAccessTokensSelfAssociationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}
type PostApiV4PersonalAccessTokensSelfRotateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		Token       *string        `json:"token,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}
type DeleteApiV4PersonalAccessTokensIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4PersonalAccessTokensIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}
type PostApiV4PersonalAccessTokensIdRotateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Active      *bool          `json:"active,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		Description *string        `json:"description,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
		Id          *int32         `json:"id,omitempty"`
		LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
		Name        *string        `json:"name,omitempty"`
		Revoked     *bool          `json:"revoked,omitempty"`
		Scopes      *[]interface{} `json:"scopes,omitempty"`
		Token       *string        `json:"token,omitempty"`
		UserId      *int32         `json:"user_id,omitempty"`
	}
}

func (c *Client) GetApiV4PersonalAccessTokens(ctx context.Context, params *GetApiV4PersonalAccessTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4PersonalAccessTokensRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4PersonalAccessTokensSelf(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4PersonalAccessTokensSelfRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4PersonalAccessTokensSelf(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4PersonalAccessTokensSelfRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4PersonalAccessTokensSelfAssociations(ctx context.Context, params *GetApiV4PersonalAccessTokensSelfAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4PersonalAccessTokensSelfAssociationsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4PersonalAccessTokensSelfRotateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4PersonalAccessTokensSelfRotateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4PersonalAccessTokensSelfRotate(ctx context.Context, body PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4PersonalAccessTokensSelfRotateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4PersonalAccessTokensId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4PersonalAccessTokensIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4PersonalAccessTokensId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4PersonalAccessTokensIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4PersonalAccessTokensIdRotateWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4PersonalAccessTokensIdRotateRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4PersonalAccessTokensIdRotate(ctx context.Context, id int32, body PostApiV4PersonalAccessTokensIdRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4PersonalAccessTokensIdRotateRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4PersonalAccessTokensRequest(server string, params *GetApiV4PersonalAccessTokensParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.UserId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "user_id", runtime.ParamLocationQuery, *params.UserId); err != nil {
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

		if params.Revoked != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "revoked", runtime.ParamLocationQuery, *params.Revoked); err != nil {
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

		if params.State != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "state", runtime.ParamLocationQuery, *params.State); err != nil {
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

		if params.CreatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_before", runtime.ParamLocationQuery, *params.CreatedBefore); err != nil {
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

		if params.CreatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_after", runtime.ParamLocationQuery, *params.CreatedAfter); err != nil {
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

		if params.LastUsedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_used_before", runtime.ParamLocationQuery, *params.LastUsedBefore); err != nil {
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

		if params.LastUsedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_used_after", runtime.ParamLocationQuery, *params.LastUsedAfter); err != nil {
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

		if params.ExpiresBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expires_before", runtime.ParamLocationQuery, *params.ExpiresBefore); err != nil {
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

		if params.ExpiresAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "expires_after", runtime.ParamLocationQuery, *params.ExpiresAfter); err != nil {
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

		if params.Search != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "search", runtime.ParamLocationQuery, *params.Search); err != nil {
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

		if params.Sort != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
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
func NewDeleteApiV4PersonalAccessTokensSelfRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/self")
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
func NewGetApiV4PersonalAccessTokensSelfRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/self")
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
func NewGetApiV4PersonalAccessTokensSelfAssociationsRequest(server string, params *GetApiV4PersonalAccessTokensSelfAssociationsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/self/associations")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.MinAccessLevel != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "min_access_level", runtime.ParamLocationQuery, *params.MinAccessLevel); err != nil {
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
func NewPostApiV4PersonalAccessTokensSelfRotateRequest(server string, body PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4PersonalAccessTokensSelfRotateRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4PersonalAccessTokensSelfRotateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/self/rotate")
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
func NewDeleteApiV4PersonalAccessTokensIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/%s", pathParam0)
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
func NewGetApiV4PersonalAccessTokensIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/%s", pathParam0)
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
func NewPostApiV4PersonalAccessTokensIdRotateRequest(server string, id int32, body PostApiV4PersonalAccessTokensIdRotateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4PersonalAccessTokensIdRotateRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPostApiV4PersonalAccessTokensIdRotateRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/personal_access_tokens/%s/rotate", pathParam0)
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
func (r GetApiV4PersonalAccessTokensResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4PersonalAccessTokensResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4PersonalAccessTokensSelfResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4PersonalAccessTokensSelfResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4PersonalAccessTokensSelfResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4PersonalAccessTokensSelfResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4PersonalAccessTokensSelfAssociationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4PersonalAccessTokensSelfAssociationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4PersonalAccessTokensSelfRotateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4PersonalAccessTokensSelfRotateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4PersonalAccessTokensIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4PersonalAccessTokensIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4PersonalAccessTokensIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4PersonalAccessTokensIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4PersonalAccessTokensIdRotateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4PersonalAccessTokensIdRotateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4PersonalAccessTokensWithResponse(ctx context.Context, params *GetApiV4PersonalAccessTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensResponse, error) {
	rsp, err := c.GetApiV4PersonalAccessTokens(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4PersonalAccessTokensResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4PersonalAccessTokensSelfWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteApiV4PersonalAccessTokensSelfResponse, error) {
	rsp, err := c.DeleteApiV4PersonalAccessTokensSelf(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4PersonalAccessTokensSelfResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4PersonalAccessTokensSelfWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensSelfResponse, error) {
	rsp, err := c.GetApiV4PersonalAccessTokensSelf(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4PersonalAccessTokensSelfResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4PersonalAccessTokensSelfAssociationsWithResponse(ctx context.Context, params *GetApiV4PersonalAccessTokensSelfAssociationsParams, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensSelfAssociationsResponse, error) {
	rsp, err := c.GetApiV4PersonalAccessTokensSelfAssociations(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4PersonalAccessTokensSelfAssociationsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4PersonalAccessTokensSelfRotateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensSelfRotateResponse, error) {
	rsp, err := c.PostApiV4PersonalAccessTokensSelfRotateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4PersonalAccessTokensSelfRotateResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4PersonalAccessTokensSelfRotateWithResponse(ctx context.Context, body PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensSelfRotateResponse, error) {
	rsp, err := c.PostApiV4PersonalAccessTokensSelfRotate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4PersonalAccessTokensSelfRotateResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4PersonalAccessTokensIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4PersonalAccessTokensIdResponse, error) {
	rsp, err := c.DeleteApiV4PersonalAccessTokensId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4PersonalAccessTokensIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4PersonalAccessTokensIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensIdResponse, error) {
	rsp, err := c.GetApiV4PersonalAccessTokensId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4PersonalAccessTokensIdResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4PersonalAccessTokensIdRotateWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensIdRotateResponse, error) {
	rsp, err := c.PostApiV4PersonalAccessTokensIdRotateWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4PersonalAccessTokensIdRotateResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4PersonalAccessTokensIdRotateWithResponse(ctx context.Context, id int32, body PostApiV4PersonalAccessTokensIdRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensIdRotateResponse, error) {
	rsp, err := c.PostApiV4PersonalAccessTokensIdRotate(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4PersonalAccessTokensIdRotateResponse(rsp)
}
func ParseGetApiV4PersonalAccessTokensResponse(rsp *http.Response) (*GetApiV4PersonalAccessTokensResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4PersonalAccessTokensResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4PersonalAccessTokensSelfResponse(rsp *http.Response) (*DeleteApiV4PersonalAccessTokensSelfResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4PersonalAccessTokensSelfResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4PersonalAccessTokensSelfResponse(rsp *http.Response) (*GetApiV4PersonalAccessTokensSelfResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4PersonalAccessTokensSelfResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4PersonalAccessTokensSelfAssociationsResponse(rsp *http.Response) (*GetApiV4PersonalAccessTokensSelfAssociationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4PersonalAccessTokensSelfAssociationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4PersonalAccessTokensSelfRotateResponse(rsp *http.Response) (*PostApiV4PersonalAccessTokensSelfRotateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4PersonalAccessTokensSelfRotateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			Token       *string        `json:"token,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4PersonalAccessTokensIdResponse(rsp *http.Response) (*DeleteApiV4PersonalAccessTokensIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4PersonalAccessTokensIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4PersonalAccessTokensIdResponse(rsp *http.Response) (*GetApiV4PersonalAccessTokensIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4PersonalAccessTokensIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			LastUsedIps *[]interface{} `json:"last_used_ips,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4PersonalAccessTokensIdRotateResponse(rsp *http.Response) (*PostApiV4PersonalAccessTokensIdRotateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4PersonalAccessTokensIdRotateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Active      *bool          `json:"active,omitempty"`
			CreatedAt   *time.Time     `json:"created_at,omitempty"`
			Description *string        `json:"description,omitempty"`
			ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
			Id          *int32         `json:"id,omitempty"`
			LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
			Name        *string        `json:"name,omitempty"`
			Revoked     *bool          `json:"revoked,omitempty"`
			Scopes      *[]interface{} `json:"scopes,omitempty"`
			Token       *string        `json:"token,omitempty"`
			UserId      *int32         `json:"user_id,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
