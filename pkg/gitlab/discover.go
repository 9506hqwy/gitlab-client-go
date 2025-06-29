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

type GetApiV4DiscoverCertBasedClustersParams struct {
	// GroupId The group ID to find all certificate-based clusters in the hierarchy
	GroupId int32 `form:"group_id" json:"group_id"`
}
type GetApiV4DiscoverCertBasedClustersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Groups   *string `json:"groups,omitempty"`
		Projects *string `json:"projects,omitempty"`
	}
}

func (c *Client) GetApiV4DiscoverCertBasedClusters(ctx context.Context, params *GetApiV4DiscoverCertBasedClustersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4DiscoverCertBasedClustersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4DiscoverCertBasedClustersRequest(server string, params *GetApiV4DiscoverCertBasedClustersParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/discover-cert-based-clusters")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "group_id", runtime.ParamLocationQuery, params.GroupId); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (r GetApiV4DiscoverCertBasedClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4DiscoverCertBasedClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4DiscoverCertBasedClustersWithResponse(ctx context.Context, params *GetApiV4DiscoverCertBasedClustersParams, reqEditors ...RequestEditorFn) (*GetApiV4DiscoverCertBasedClustersResponse, error) {
	rsp, err := c.GetApiV4DiscoverCertBasedClusters(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4DiscoverCertBasedClustersResponse(rsp)
}
func ParseGetApiV4DiscoverCertBasedClustersResponse(rsp *http.Response) (*GetApiV4DiscoverCertBasedClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4DiscoverCertBasedClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Groups   *string `json:"groups,omitempty"`
			Projects *string `json:"projects,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
