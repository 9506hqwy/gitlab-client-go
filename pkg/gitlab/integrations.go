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

type PostApiV4IntegrationsJiraConnectSubscriptionsJSONBody struct {
	// Jwt JWT token for authorization with the Jira Connect installation
	Jwt string `json:"jwt"`

	// NamespacePath Path for the namespace that should be subscribed
	NamespacePath string `json:"namespace_path"`
}
type PostApiV4IntegrationsSlackEventsJSONBody struct {
	// ApiAppId The Slack app ID
	ApiAppId *string `json:"api_app_id,omitempty"`

	// AuthedUsers (Deprecated by Slack) An array of Slack user IDs
	AuthedUsers *[]string `json:"authed_users,omitempty"`

	// Event The event object with variable properties
	Event *map[string]interface{} `json:"event,omitempty"`

	// EventId A unique identifier for this specific event
	EventId *string `json:"event_id,omitempty"`

	// EventTime The epoch timestamp in seconds when this event was dispatched
	EventTime *int32 `json:"event_time,omitempty"`

	// TeamId The Slack workspace ID of where the event occurred
	TeamId *string `json:"team_id,omitempty"`

	// Token (Deprecated by Slack) The request token, unused by GitLab
	Token *string `json:"token,omitempty"`

	// Type The kind of event this is, usually `event_callback`
	Type *string `json:"type,omitempty"`
}
type PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody PostApiV4IntegrationsJiraConnectSubscriptionsJSONBody
type PostApiV4IntegrationsSlackEventsJSONRequestBody PostApiV4IntegrationsSlackEventsJSONBody
type PostApiV4IntegrationsJiraConnectSubscriptionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Success *map[string]interface{} `json:"success,omitempty"`
	}
}
type PostApiV4IntegrationsSlackEventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4IntegrationsSlackInteractionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4IntegrationsSlackOptionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) PostApiV4IntegrationsJiraConnectSubscriptionsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsJiraConnectSubscriptionsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4IntegrationsJiraConnectSubscriptions(ctx context.Context, body PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsJiraConnectSubscriptionsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4IntegrationsSlackEventsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsSlackEventsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4IntegrationsSlackEvents(ctx context.Context, body PostApiV4IntegrationsSlackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsSlackEventsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4IntegrationsSlackInteractions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsSlackInteractionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4IntegrationsSlackOptions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4IntegrationsSlackOptionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4IntegrationsJiraConnectSubscriptionsRequest(server string, body PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4IntegrationsJiraConnectSubscriptionsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4IntegrationsJiraConnectSubscriptionsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/integrations/jira_connect/subscriptions")
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
func NewPostApiV4IntegrationsSlackEventsRequest(server string, body PostApiV4IntegrationsSlackEventsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4IntegrationsSlackEventsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4IntegrationsSlackEventsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/integrations/slack/events")
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
func NewPostApiV4IntegrationsSlackInteractionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/integrations/slack/interactions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4IntegrationsSlackOptionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/integrations/slack/options")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (r PostApiV4IntegrationsJiraConnectSubscriptionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4IntegrationsJiraConnectSubscriptionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4IntegrationsSlackEventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4IntegrationsSlackEventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4IntegrationsSlackInteractionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4IntegrationsSlackInteractionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4IntegrationsSlackOptionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4IntegrationsSlackOptionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4IntegrationsJiraConnectSubscriptionsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsJiraConnectSubscriptionsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsJiraConnectSubscriptionsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsJiraConnectSubscriptionsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4IntegrationsJiraConnectSubscriptionsWithResponse(ctx context.Context, body PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsJiraConnectSubscriptionsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsJiraConnectSubscriptions(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsJiraConnectSubscriptionsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4IntegrationsSlackEventsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackEventsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsSlackEventsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsSlackEventsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4IntegrationsSlackEventsWithResponse(ctx context.Context, body PostApiV4IntegrationsSlackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackEventsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsSlackEvents(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsSlackEventsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4IntegrationsSlackInteractionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackInteractionsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsSlackInteractions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsSlackInteractionsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4IntegrationsSlackOptionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackOptionsResponse, error) {
	rsp, err := c.PostApiV4IntegrationsSlackOptions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4IntegrationsSlackOptionsResponse(rsp)
}
func ParsePostApiV4IntegrationsJiraConnectSubscriptionsResponse(rsp *http.Response) (*PostApiV4IntegrationsJiraConnectSubscriptionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4IntegrationsJiraConnectSubscriptionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Success *map[string]interface{} `json:"success,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4IntegrationsSlackEventsResponse(rsp *http.Response) (*PostApiV4IntegrationsSlackEventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4IntegrationsSlackEventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4IntegrationsSlackInteractionsResponse(rsp *http.Response) (*PostApiV4IntegrationsSlackInteractionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4IntegrationsSlackInteractionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4IntegrationsSlackOptionsResponse(rsp *http.Response) (*PostApiV4IntegrationsSlackOptionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4IntegrationsSlackOptionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
