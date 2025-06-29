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

type GetApiV4DeployKeysParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Public Only return deploy keys that are public
	Public *bool `form:"public,omitempty" json:"public,omitempty"`
}
type PostApiV4DeployKeysJSONBody struct {
	// ExpiresAt The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// Key New deploy key
	Key string `json:"key"`

	// Title New deploy key's title
	Title string `json:"title"`
}
type GetApiV4DeployTokensParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Active Limit by active status
	Active *bool `form:"active,omitempty" json:"active,omitempty"`
}
type PostApiV4DeployKeysJSONRequestBody PostApiV4DeployKeysJSONBody
type GetApiV4DeployKeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		CreatedAt                  *time.Time `json:"created_at,omitempty"`
		ExpiresAt                  *time.Time `json:"expires_at,omitempty"`
		Fingerprint                *string    `json:"fingerprint,omitempty"`
		FingerprintSha256          *string    `json:"fingerprint_sha256,omitempty"`
		Id                         *int32     `json:"id,omitempty"`
		Key                        *string    `json:"key,omitempty"`
		LastUsedAt                 *time.Time `json:"last_used_at,omitempty"`
		ProjectsWithReadonlyAccess *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"projects_with_readonly_access,omitempty"`
		ProjectsWithWriteAccess *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"projects_with_write_access,omitempty"`
		Title     *string `json:"title,omitempty"`
		UsageType *string `json:"usage_type,omitempty"`
	}
}
type PostApiV4DeployKeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		CreatedAt                  *time.Time `json:"created_at,omitempty"`
		ExpiresAt                  *time.Time `json:"expires_at,omitempty"`
		Fingerprint                *string    `json:"fingerprint,omitempty"`
		FingerprintSha256          *string    `json:"fingerprint_sha256,omitempty"`
		Id                         *int32     `json:"id,omitempty"`
		Key                        *string    `json:"key,omitempty"`
		LastUsedAt                 *time.Time `json:"last_used_at,omitempty"`
		ProjectsWithReadonlyAccess *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"projects_with_readonly_access,omitempty"`
		ProjectsWithWriteAccess *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"projects_with_write_access,omitempty"`
		Title     *string `json:"title,omitempty"`
		UsageType *string `json:"usage_type,omitempty"`
	}
}
type GetApiV4DeployTokensResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		Expired   *bool          `json:"expired,omitempty"`
		ExpiresAt *time.Time     `json:"expires_at,omitempty"`
		Id        *int32         `json:"id,omitempty"`
		Name      *string        `json:"name,omitempty"`
		Revoked   *bool          `json:"revoked,omitempty"`
		Scopes    *[]interface{} `json:"scopes,omitempty"`
		Username  *string        `json:"username,omitempty"`
	}
}

func (c *Client) GetApiV4DeployKeys(ctx context.Context, params *GetApiV4DeployKeysParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4DeployKeysRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4DeployKeysWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4DeployKeysRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4DeployKeys(ctx context.Context, body PostApiV4DeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4DeployKeysRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4DeployTokens(ctx context.Context, params *GetApiV4DeployTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4DeployTokensRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4DeployKeysRequest(server string, params *GetApiV4DeployKeysParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/deploy_keys")
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

		if params.Public != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "public", runtime.ParamLocationQuery, *params.Public); err != nil {
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
func NewPostApiV4DeployKeysRequest(server string, body PostApiV4DeployKeysJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4DeployKeysRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4DeployKeysRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/deploy_keys")
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
func NewGetApiV4DeployTokensRequest(server string, params *GetApiV4DeployTokensParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/deploy_tokens")
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

		if params.Active != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "active", runtime.ParamLocationQuery, *params.Active); err != nil {
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
func (r GetApiV4DeployKeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4DeployKeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4DeployKeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4DeployKeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4DeployTokensResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4DeployTokensResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4DeployKeysWithResponse(ctx context.Context, params *GetApiV4DeployKeysParams, reqEditors ...RequestEditorFn) (*GetApiV4DeployKeysResponse, error) {
	rsp, err := c.GetApiV4DeployKeys(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4DeployKeysResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4DeployKeysWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4DeployKeysResponse, error) {
	rsp, err := c.PostApiV4DeployKeysWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4DeployKeysResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4DeployKeysWithResponse(ctx context.Context, body PostApiV4DeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4DeployKeysResponse, error) {
	rsp, err := c.PostApiV4DeployKeys(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4DeployKeysResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4DeployTokensWithResponse(ctx context.Context, params *GetApiV4DeployTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4DeployTokensResponse, error) {
	rsp, err := c.GetApiV4DeployTokens(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4DeployTokensResponse(rsp)
}
func ParseGetApiV4DeployKeysResponse(rsp *http.Response) (*GetApiV4DeployKeysResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4DeployKeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			CreatedAt                  *time.Time `json:"created_at,omitempty"`
			ExpiresAt                  *time.Time `json:"expires_at,omitempty"`
			Fingerprint                *string    `json:"fingerprint,omitempty"`
			FingerprintSha256          *string    `json:"fingerprint_sha256,omitempty"`
			Id                         *int32     `json:"id,omitempty"`
			Key                        *string    `json:"key,omitempty"`
			LastUsedAt                 *time.Time `json:"last_used_at,omitempty"`
			ProjectsWithReadonlyAccess *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"projects_with_readonly_access,omitempty"`
			ProjectsWithWriteAccess *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"projects_with_write_access,omitempty"`
			Title     *string `json:"title,omitempty"`
			UsageType *string `json:"usage_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4DeployKeysResponse(rsp *http.Response) (*PostApiV4DeployKeysResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4DeployKeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			CreatedAt                  *time.Time `json:"created_at,omitempty"`
			ExpiresAt                  *time.Time `json:"expires_at,omitempty"`
			Fingerprint                *string    `json:"fingerprint,omitempty"`
			FingerprintSha256          *string    `json:"fingerprint_sha256,omitempty"`
			Id                         *int32     `json:"id,omitempty"`
			Key                        *string    `json:"key,omitempty"`
			LastUsedAt                 *time.Time `json:"last_used_at,omitempty"`
			ProjectsWithReadonlyAccess *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"projects_with_readonly_access,omitempty"`
			ProjectsWithWriteAccess *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"projects_with_write_access,omitempty"`
			Title     *string `json:"title,omitempty"`
			UsageType *string `json:"usage_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseGetApiV4DeployTokensResponse(rsp *http.Response) (*GetApiV4DeployTokensResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4DeployTokensResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			Expired   *bool          `json:"expired,omitempty"`
			ExpiresAt *time.Time     `json:"expires_at,omitempty"`
			Id        *int32         `json:"id,omitempty"`
			Name      *string        `json:"name,omitempty"`
			Revoked   *bool          `json:"revoked,omitempty"`
			Scopes    *[]interface{} `json:"scopes,omitempty"`
			Username  *string        `json:"username,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
