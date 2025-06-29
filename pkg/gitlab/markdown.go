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
)

type PostApiV4MarkdownJSONBody struct {
	// Gfm Render text using GitLab Flavored Markdown. Default is false
	Gfm *bool `json:"gfm,omitempty"`

	// Project Use project as a context when creating references using GitLab Flavored Markdown
	Project *string `json:"project,omitempty"`

	// Text The Markdown text to render
	Text string `json:"text"`
}
type PostApiV4MarkdownJSONRequestBody PostApiV4MarkdownJSONBody
type PostApiV4MarkdownResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Html *string `json:"html,omitempty"`
	}
}

func (c *Client) PostApiV4MarkdownWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4MarkdownRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Markdown(ctx context.Context, body PostApiV4MarkdownJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4MarkdownRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4MarkdownRequest(server string, body PostApiV4MarkdownJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4MarkdownRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4MarkdownRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/markdown")
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
func (r PostApiV4MarkdownResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4MarkdownResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4MarkdownWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4MarkdownResponse, error) {
	rsp, err := c.PostApiV4MarkdownWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4MarkdownResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4MarkdownWithResponse(ctx context.Context, body PostApiV4MarkdownJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4MarkdownResponse, error) {
	rsp, err := c.PostApiV4Markdown(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4MarkdownResponse(rsp)
}
func ParsePostApiV4MarkdownResponse(rsp *http.Response) (*PostApiV4MarkdownResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4MarkdownResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Html *string `json:"html,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
