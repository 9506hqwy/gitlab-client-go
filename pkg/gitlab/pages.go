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

type GetApiV4PagesDomainsParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4PagesDomainsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AutoSslEnabled        *string `json:"auto_ssl_enabled,omitempty"`
		CertificateExpiration *struct {
			Expiration *string `json:"expiration,omitempty"`
			Expired    *string `json:"expired,omitempty"`
		} `json:"certificate_expiration,omitempty"`
		Domain           *string `json:"domain,omitempty"`
		EnabledUntil     *string `json:"enabled_until,omitempty"`
		ProjectId        *string `json:"project_id,omitempty"`
		Url              *string `json:"url,omitempty"`
		VerificationCode *string `json:"verification_code,omitempty"`
		Verified         *string `json:"verified,omitempty"`
	}
}

func (c *Client) GetApiV4PagesDomains(ctx context.Context, params *GetApiV4PagesDomainsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4PagesDomainsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4PagesDomainsRequest(server string, params *GetApiV4PagesDomainsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/pages/domains")
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
func (r GetApiV4PagesDomainsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4PagesDomainsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4PagesDomainsWithResponse(ctx context.Context, params *GetApiV4PagesDomainsParams, reqEditors ...RequestEditorFn) (*GetApiV4PagesDomainsResponse, error) {
	rsp, err := c.GetApiV4PagesDomains(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4PagesDomainsResponse(rsp)
}
func ParseGetApiV4PagesDomainsResponse(rsp *http.Response) (*GetApiV4PagesDomainsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4PagesDomainsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AutoSslEnabled        *string `json:"auto_ssl_enabled,omitempty"`
			CertificateExpiration *struct {
				Expiration *string `json:"expiration,omitempty"`
				Expired    *string `json:"expired,omitempty"`
			} `json:"certificate_expiration,omitempty"`
			Domain           *string `json:"domain,omitempty"`
			EnabledUntil     *string `json:"enabled_until,omitempty"`
			ProjectId        *string `json:"project_id,omitempty"`
			Url              *string `json:"url,omitempty"`
			VerificationCode *string `json:"verification_code,omitempty"`
			Verified         *string `json:"verified,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
