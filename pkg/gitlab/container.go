package gitlab

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type PostApiV4ContainerRegistryEventEventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) PostApiV4ContainerRegistryEventEvents(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ContainerRegistryEventEventsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4ContainerRegistryEventEventsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/container_registry_event/events")
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
func (r PostApiV4ContainerRegistryEventEventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ContainerRegistryEventEventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4ContainerRegistryEventEventsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4ContainerRegistryEventEventsResponse, error) {
	rsp, err := c.PostApiV4ContainerRegistryEventEvents(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ContainerRegistryEventEventsResponse(rsp)
}
func ParsePostApiV4ContainerRegistryEventEventsResponse(rsp *http.Response) (*PostApiV4ContainerRegistryEventEventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ContainerRegistryEventEventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
