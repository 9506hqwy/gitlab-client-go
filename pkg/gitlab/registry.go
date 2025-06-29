package gitlab

import (
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

type GetApiV4RegistryRepositoriesIdParams struct {
	// Tags Determines if tags should be included
	Tags *bool `form:"tags,omitempty" json:"tags,omitempty"`

	// TagsCount Determines if the tags count should be included
	TagsCount *bool `form:"tags_count,omitempty" json:"tags_count,omitempty"`

	// Size Determines if the size should be included
	Size *bool `form:"size,omitempty" json:"size,omitempty"`
}
type GetApiV4RegistryRepositoriesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CleanupPolicyStartedAt *time.Time `json:"cleanup_policy_started_at,omitempty"`
		CreatedAt              *time.Time `json:"created_at,omitempty"`
		DeleteApiPath          *string    `json:"delete_api_path,omitempty"`
		Id                     *int32     `json:"id,omitempty"`
		Location               *string    `json:"location,omitempty"`
		Name                   *string    `json:"name,omitempty"`
		Path                   *string    `json:"path,omitempty"`
		ProjectId              *int32     `json:"project_id,omitempty"`
		Size                   *int32     `json:"size,omitempty"`
		Status                 *string    `json:"status,omitempty"`

		// Tags API_Entities_ContainerRegistry_Tag model
		Tags *struct {
			Location *string `json:"location,omitempty"`
			Name     *string `json:"name,omitempty"`
			Path     *string `json:"path,omitempty"`
		} `json:"tags,omitempty"`
		TagsCount *int32 `json:"tags_count,omitempty"`
	}
}

func (c *Client) GetApiV4RegistryRepositoriesId(ctx context.Context, id string, params *GetApiV4RegistryRepositoriesIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RegistryRepositoriesIdRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4RegistryRepositoriesIdRequest(server string, id string, params *GetApiV4RegistryRepositoriesIdParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/registry/repositories/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Tags != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "tags", runtime.ParamLocationQuery, *params.Tags); err != nil {
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

		if params.TagsCount != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "tags_count", runtime.ParamLocationQuery, *params.TagsCount); err != nil {
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

		if params.Size != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "size", runtime.ParamLocationQuery, *params.Size); err != nil {
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
func (r GetApiV4RegistryRepositoriesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RegistryRepositoriesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4RegistryRepositoriesIdWithResponse(ctx context.Context, id string, params *GetApiV4RegistryRepositoriesIdParams, reqEditors ...RequestEditorFn) (*GetApiV4RegistryRepositoriesIdResponse, error) {
	rsp, err := c.GetApiV4RegistryRepositoriesId(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RegistryRepositoriesIdResponse(rsp)
}
func ParseGetApiV4RegistryRepositoriesIdResponse(rsp *http.Response) (*GetApiV4RegistryRepositoriesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RegistryRepositoriesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CleanupPolicyStartedAt *time.Time `json:"cleanup_policy_started_at,omitempty"`
			CreatedAt              *time.Time `json:"created_at,omitempty"`
			DeleteApiPath          *string    `json:"delete_api_path,omitempty"`
			Id                     *int32     `json:"id,omitempty"`
			Location               *string    `json:"location,omitempty"`
			Name                   *string    `json:"name,omitempty"`
			Path                   *string    `json:"path,omitempty"`
			ProjectId              *int32     `json:"project_id,omitempty"`
			Size                   *int32     `json:"size,omitempty"`
			Status                 *string    `json:"status,omitempty"`

			// Tags API_Entities_ContainerRegistry_Tag model
			Tags *struct {
				Location *string `json:"location,omitempty"`
				Name     *string `json:"name,omitempty"`
				Path     *string `json:"path,omitempty"`
			} `json:"tags,omitempty"`
			TagsCount *int32 `json:"tags_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
