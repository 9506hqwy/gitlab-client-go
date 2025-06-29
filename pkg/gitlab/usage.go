package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/oapi-codegen/runtime"
)

type PostApiV4UsageDataIncrementCounterJSONBody struct {
	// Event The event name that should be tracked
	Event string `json:"event"`
}
type PostApiV4UsageDataIncrementUniqueUsersJSONBody struct {
	// Event The event name that should be tracked
	Event string `json:"event"`
}
type GetApiV4UsageDataMetricDefinitionsParams struct {
	// IncludePaths Include file paths in the metric definitions
	IncludePaths *bool `form:"include_paths,omitempty" json:"include_paths,omitempty"`
}
type PostApiV4UsageDataTrackEventJSONBody struct {
	// AdditionalProperties Additional properties to be tracked
	AdditionalProperties *map[string]interface{} `json:"additional_properties,omitempty"`

	// Event The event name that should be tracked
	Event string `json:"event"`

	// NamespaceId Namespace ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// ProjectId Project ID
	ProjectId *int32 `json:"project_id,omitempty"`

	// SendToSnowplow Send the tracked event to Snowplow
	SendToSnowplow *bool `json:"send_to_snowplow,omitempty"`
}
type PostApiV4UsageDataTrackEventsJSONBody struct {
	// Events An array of internal events. Maximum 50 events allowed.
	Events []struct {
		// AdditionalProperties Additional properties to be tracked
		AdditionalProperties *map[string]interface{} `json:"additional_properties,omitempty"`

		// Event The event name that should be tracked
		Event string `json:"event"`

		// NamespaceId Namespace ID
		NamespaceId *int32 `json:"namespace_id,omitempty"`

		// ProjectId Project ID
		ProjectId *int32 `json:"project_id,omitempty"`

		// SendToSnowplow Send the tracked event to Snowplow
		SendToSnowplow *bool `json:"send_to_snowplow,omitempty"`
	} `json:"events"`
}
type PostApiV4UsageDataIncrementCounterJSONRequestBody PostApiV4UsageDataIncrementCounterJSONBody
type PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody PostApiV4UsageDataIncrementUniqueUsersJSONBody
type PostApiV4UsageDataTrackEventJSONRequestBody PostApiV4UsageDataTrackEventJSONBody
type PostApiV4UsageDataTrackEventsJSONRequestBody PostApiV4UsageDataTrackEventsJSONBody
type PostApiV4UsageDataIncrementCounterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4UsageDataIncrementUniqueUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4UsageDataMetricDefinitionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4UsageDataNonSqlMetricsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4UsageDataQueriesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4UsageDataServicePingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4UsageDataTrackEventResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4UsageDataTrackEventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) PostApiV4UsageDataIncrementCounterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataIncrementCounterRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataIncrementCounter(ctx context.Context, body PostApiV4UsageDataIncrementCounterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataIncrementCounterRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataIncrementUniqueUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataIncrementUniqueUsersRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataIncrementUniqueUsers(ctx context.Context, body PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataIncrementUniqueUsersRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsageDataMetricDefinitions(ctx context.Context, params *GetApiV4UsageDataMetricDefinitionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsageDataMetricDefinitionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsageDataNonSqlMetrics(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsageDataNonSqlMetricsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsageDataQueries(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsageDataQueriesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsageDataServicePing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsageDataServicePingRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataTrackEventWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataTrackEventRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataTrackEvent(ctx context.Context, body PostApiV4UsageDataTrackEventJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataTrackEventRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataTrackEventsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataTrackEventsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UsageDataTrackEvents(ctx context.Context, body PostApiV4UsageDataTrackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UsageDataTrackEventsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4UsageDataIncrementCounterRequest(server string, body PostApiV4UsageDataIncrementCounterJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4UsageDataIncrementCounterRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4UsageDataIncrementCounterRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/increment_counter")
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
func NewPostApiV4UsageDataIncrementUniqueUsersRequest(server string, body PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4UsageDataIncrementUniqueUsersRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4UsageDataIncrementUniqueUsersRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/increment_unique_users")
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
func NewGetApiV4UsageDataMetricDefinitionsRequest(server string, params *GetApiV4UsageDataMetricDefinitionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/metric_definitions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.IncludePaths != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "include_paths", runtime.ParamLocationQuery, *params.IncludePaths); err != nil {
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
func NewGetApiV4UsageDataNonSqlMetricsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/non_sql_metrics")
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
func NewGetApiV4UsageDataQueriesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/queries")
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
func NewGetApiV4UsageDataServicePingRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/service_ping")
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
func NewPostApiV4UsageDataTrackEventRequest(server string, body PostApiV4UsageDataTrackEventJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4UsageDataTrackEventRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4UsageDataTrackEventRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/track_event")
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
func NewPostApiV4UsageDataTrackEventsRequest(server string, body PostApiV4UsageDataTrackEventsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4UsageDataTrackEventsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4UsageDataTrackEventsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/usage_data/track_events")
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
func (r PostApiV4UsageDataIncrementCounterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4UsageDataIncrementCounterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4UsageDataIncrementUniqueUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4UsageDataIncrementUniqueUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsageDataMetricDefinitionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsageDataMetricDefinitionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsageDataNonSqlMetricsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsageDataNonSqlMetricsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsageDataQueriesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsageDataQueriesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsageDataServicePingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsageDataServicePingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4UsageDataTrackEventResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4UsageDataTrackEventResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4UsageDataTrackEventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4UsageDataTrackEventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4UsageDataIncrementCounterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementCounterResponse, error) {
	rsp, err := c.PostApiV4UsageDataIncrementCounterWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataIncrementCounterResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataIncrementCounterWithResponse(ctx context.Context, body PostApiV4UsageDataIncrementCounterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementCounterResponse, error) {
	rsp, err := c.PostApiV4UsageDataIncrementCounter(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataIncrementCounterResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataIncrementUniqueUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementUniqueUsersResponse, error) {
	rsp, err := c.PostApiV4UsageDataIncrementUniqueUsersWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataIncrementUniqueUsersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataIncrementUniqueUsersWithResponse(ctx context.Context, body PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementUniqueUsersResponse, error) {
	rsp, err := c.PostApiV4UsageDataIncrementUniqueUsers(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataIncrementUniqueUsersResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsageDataMetricDefinitionsWithResponse(ctx context.Context, params *GetApiV4UsageDataMetricDefinitionsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataMetricDefinitionsResponse, error) {
	rsp, err := c.GetApiV4UsageDataMetricDefinitions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsageDataMetricDefinitionsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsageDataNonSqlMetricsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataNonSqlMetricsResponse, error) {
	rsp, err := c.GetApiV4UsageDataNonSqlMetrics(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsageDataNonSqlMetricsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsageDataQueriesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataQueriesResponse, error) {
	rsp, err := c.GetApiV4UsageDataQueries(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsageDataQueriesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsageDataServicePingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataServicePingResponse, error) {
	rsp, err := c.GetApiV4UsageDataServicePing(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsageDataServicePingResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataTrackEventWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventResponse, error) {
	rsp, err := c.PostApiV4UsageDataTrackEventWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataTrackEventResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataTrackEventWithResponse(ctx context.Context, body PostApiV4UsageDataTrackEventJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventResponse, error) {
	rsp, err := c.PostApiV4UsageDataTrackEvent(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataTrackEventResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataTrackEventsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventsResponse, error) {
	rsp, err := c.PostApiV4UsageDataTrackEventsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataTrackEventsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UsageDataTrackEventsWithResponse(ctx context.Context, body PostApiV4UsageDataTrackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventsResponse, error) {
	rsp, err := c.PostApiV4UsageDataTrackEvents(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UsageDataTrackEventsResponse(rsp)
}
func ParsePostApiV4UsageDataIncrementCounterResponse(rsp *http.Response) (*PostApiV4UsageDataIncrementCounterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4UsageDataIncrementCounterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4UsageDataIncrementUniqueUsersResponse(rsp *http.Response) (*PostApiV4UsageDataIncrementUniqueUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4UsageDataIncrementUniqueUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4UsageDataMetricDefinitionsResponse(rsp *http.Response) (*GetApiV4UsageDataMetricDefinitionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsageDataMetricDefinitionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4UsageDataNonSqlMetricsResponse(rsp *http.Response) (*GetApiV4UsageDataNonSqlMetricsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsageDataNonSqlMetricsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4UsageDataQueriesResponse(rsp *http.Response) (*GetApiV4UsageDataQueriesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsageDataQueriesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4UsageDataServicePingResponse(rsp *http.Response) (*GetApiV4UsageDataServicePingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsageDataServicePingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4UsageDataTrackEventResponse(rsp *http.Response) (*PostApiV4UsageDataTrackEventResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4UsageDataTrackEventResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4UsageDataTrackEventsResponse(rsp *http.Response) (*PostApiV4UsageDataTrackEventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4UsageDataTrackEventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
