package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

type GetApiV4AvatarParams struct {
	// Email Public email address of the user
	Email string `form:"email" json:"email"`

	// Size Single pixel dimension for Gravatar images
	Size *int32 `form:"size,omitempty" json:"size,omitempty"`
}
type GetApiV4AvatarResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AvatarUrl *string `json:"avatar_url,omitempty"`
	}
}

func (c *Client) GetApiV4Avatar(ctx context.Context, params *GetApiV4AvatarParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AvatarRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4AvatarRequest(server string, params *GetApiV4AvatarParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/avatar")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "email", runtime.ParamLocationQuery, params.Email); err != nil {
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
func (r GetApiV4AvatarResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AvatarResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4AvatarWithResponse(ctx context.Context, params *GetApiV4AvatarParams, reqEditors ...RequestEditorFn) (*GetApiV4AvatarResponse, error) {
	rsp, err := c.GetApiV4Avatar(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AvatarResponse(rsp)
}
func ParseGetApiV4AvatarResponse(rsp *http.Response) (*GetApiV4AvatarResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AvatarResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
