package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PostApiV4SlackTriggerJSONBody struct {
	// Text Text of the slack command
	Text string `json:"text"`
}
type PostApiV4SlackTriggerJSONRequestBody PostApiV4SlackTriggerJSONBody
type PostApiV4SlackTriggerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) PostApiV4SlackTriggerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4SlackTriggerRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4SlackTrigger(ctx context.Context, body PostApiV4SlackTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4SlackTriggerRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4SlackTriggerRequest(server string, body PostApiV4SlackTriggerJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4SlackTriggerRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4SlackTriggerRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/slack/trigger")
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
func (r PostApiV4SlackTriggerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4SlackTriggerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4SlackTriggerWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4SlackTriggerResponse, error) {
	rsp, err := c.PostApiV4SlackTriggerWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4SlackTriggerResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4SlackTriggerWithResponse(ctx context.Context, body PostApiV4SlackTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4SlackTriggerResponse, error) {
	rsp, err := c.PostApiV4SlackTrigger(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4SlackTriggerResponse(rsp)
}
func ParsePostApiV4SlackTriggerResponse(rsp *http.Response) (*PostApiV4SlackTriggerResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4SlackTriggerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
