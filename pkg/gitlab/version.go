package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type GetApiV4VersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Enterprise *bool `json:"enterprise,omitempty"`
		Kas        *struct {
			Enabled             *bool   `json:"enabled,omitempty"`
			ExternalK8sProxyUrl *string `json:"externalK8sProxyUrl,omitempty"`
			ExternalUrl         *string `json:"externalUrl,omitempty"`
			Version             *string `json:"version,omitempty"`
		} `json:"kas,omitempty"`
		Revision *string `json:"revision,omitempty"`
		Version  *string `json:"version,omitempty"`
	}
}

func (c *Client) GetApiV4Version(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4VersionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4VersionRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/version")
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
func (r GetApiV4VersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4VersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4VersionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4VersionResponse, error) {
	rsp, err := c.GetApiV4Version(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4VersionResponse(rsp)
}
func ParseGetApiV4VersionResponse(rsp *http.Response) (*GetApiV4VersionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4VersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Enterprise *bool `json:"enterprise,omitempty"`
			Kas        *struct {
				Enabled             *bool   `json:"enabled,omitempty"`
				ExternalK8sProxyUrl *string `json:"externalK8sProxyUrl,omitempty"`
				ExternalUrl         *string `json:"externalUrl,omitempty"`
				Version             *string `json:"version,omitempty"`
			} `json:"kas,omitempty"`
			Revision *string `json:"revision,omitempty"`
			Version  *string `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
