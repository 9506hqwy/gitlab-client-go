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
)

type PostApiV4OrganizationsJSONBody struct {
	// Avatar The avatar image for the organization
	Avatar *string `json:"avatar,omitempty"`

	// Description The description of the organization
	Description *string `json:"description,omitempty"`

	// Name The name of the organization
	Name string `json:"name"`

	// Path The path of the organization
	Path string `json:"path"`
}
type PostApiV4OrganizationsJSONRequestBody PostApiV4OrganizationsJSONBody
type PostApiV4OrganizationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		AvatarUrl   *string    `json:"avatar_url,omitempty"`
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		Description *string    `json:"description,omitempty"`
		Id          *int32     `json:"id,omitempty"`
		Name        *string    `json:"name,omitempty"`
		Path        *string    `json:"path,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		WebUrl      *string    `json:"web_url,omitempty"`
	}
}

func (c *Client) PostApiV4OrganizationsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4OrganizationsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Organizations(ctx context.Context, body PostApiV4OrganizationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4OrganizationsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4OrganizationsRequest(server string, body PostApiV4OrganizationsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4OrganizationsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4OrganizationsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/organizations")
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
func (r PostApiV4OrganizationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4OrganizationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4OrganizationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4OrganizationsResponse, error) {
	rsp, err := c.PostApiV4OrganizationsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4OrganizationsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4OrganizationsWithResponse(ctx context.Context, body PostApiV4OrganizationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4OrganizationsResponse, error) {
	rsp, err := c.PostApiV4Organizations(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4OrganizationsResponse(rsp)
}
func ParsePostApiV4OrganizationsResponse(rsp *http.Response) (*PostApiV4OrganizationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4OrganizationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			AvatarUrl   *string    `json:"avatar_url,omitempty"`
			CreatedAt   *time.Time `json:"created_at,omitempty"`
			Description *string    `json:"description,omitempty"`
			Id          *int32     `json:"id,omitempty"`
			Name        *string    `json:"name,omitempty"`
			Path        *string    `json:"path,omitempty"`
			UpdatedAt   *time.Time `json:"updated_at,omitempty"`
			WebUrl      *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
