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

	"github.com/oapi-codegen/runtime"
)

type GetApiV4TopicsParams struct {
	// Search Return list of topics matching the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// WithoutProjects Return list of topics without assigned projects
	WithoutProjects *bool `form:"without_projects,omitempty" json:"without_projects,omitempty"`

	// OrganizationId The organization id for the topics
	OrganizationId *int32 `form:"organization_id,omitempty" json:"organization_id,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4TopicsJSONBody struct {
	// Avatar Avatar image for topic
	Avatar *string `json:"avatar,omitempty"`

	// Description Description
	Description *string `json:"description,omitempty"`

	// Name Slug (name)
	Name string `json:"name"`

	// OrganizationId The organization id for the topic
	OrganizationId *int32 `json:"organization_id,omitempty"`

	// Title Title
	Title string `json:"title"`
}
type PostApiV4TopicsMergeJSONBody struct {
	// SourceTopicId ID of source project topic
	SourceTopicId int32 `json:"source_topic_id"`

	// TargetTopicId ID of target project topic
	TargetTopicId int32 `json:"target_topic_id"`
}
type PutApiV4TopicsIdJSONBody struct {
	// Avatar Avatar image for topic
	Avatar *string `json:"avatar,omitempty"`

	// Description Description
	Description *string `json:"description,omitempty"`

	// Name Slug (name)
	Name *string `json:"name,omitempty"`

	// Title Title
	Title *string `json:"title,omitempty"`
}
type PostApiV4TopicsJSONRequestBody PostApiV4TopicsJSONBody
type PostApiV4TopicsMergeJSONRequestBody PostApiV4TopicsMergeJSONBody
type PutApiV4TopicsIdJSONRequestBody PutApiV4TopicsIdJSONBody
type GetApiV4TopicsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AvatarUrl          *string `json:"avatar_url,omitempty"`
		Description        *string `json:"description,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		OrganizationId     *string `json:"organization_id,omitempty"`
		Title              *string `json:"title,omitempty"`
		TotalProjectsCount *string `json:"total_projects_count,omitempty"`
	}
}
type PostApiV4TopicsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		AvatarUrl          *string `json:"avatar_url,omitempty"`
		Description        *string `json:"description,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		OrganizationId     *string `json:"organization_id,omitempty"`
		Title              *string `json:"title,omitempty"`
		TotalProjectsCount *string `json:"total_projects_count,omitempty"`
	}
}
type PostApiV4TopicsMergeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		AvatarUrl          *string `json:"avatar_url,omitempty"`
		Description        *string `json:"description,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		OrganizationId     *string `json:"organization_id,omitempty"`
		Title              *string `json:"title,omitempty"`
		TotalProjectsCount *string `json:"total_projects_count,omitempty"`
	}
}
type DeleteApiV4TopicsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4TopicsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AvatarUrl          *string `json:"avatar_url,omitempty"`
		Description        *string `json:"description,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		OrganizationId     *string `json:"organization_id,omitempty"`
		Title              *string `json:"title,omitempty"`
		TotalProjectsCount *string `json:"total_projects_count,omitempty"`
	}
}
type PutApiV4TopicsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AvatarUrl          *string `json:"avatar_url,omitempty"`
		Description        *string `json:"description,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Name               *string `json:"name,omitempty"`
		OrganizationId     *string `json:"organization_id,omitempty"`
		Title              *string `json:"title,omitempty"`
		TotalProjectsCount *string `json:"total_projects_count,omitempty"`
	}
}

func (c *Client) GetApiV4Topics(ctx context.Context, params *GetApiV4TopicsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4TopicsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4TopicsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4TopicsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Topics(ctx context.Context, body PostApiV4TopicsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4TopicsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4TopicsMergeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4TopicsMergeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4TopicsMerge(ctx context.Context, body PostApiV4TopicsMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4TopicsMergeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4TopicsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4TopicsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4TopicsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4TopicsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4TopicsIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4TopicsIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4TopicsId(ctx context.Context, id int32, body PutApiV4TopicsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4TopicsIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4TopicsRequest(server string, params *GetApiV4TopicsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Search != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "search", runtime.ParamLocationQuery, *params.Search); err != nil {
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

		if params.WithoutProjects != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "without_projects", runtime.ParamLocationQuery, *params.WithoutProjects); err != nil {
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

		if params.OrganizationId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "organization_id", runtime.ParamLocationQuery, *params.OrganizationId); err != nil {
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
func NewPostApiV4TopicsRequest(server string, body PostApiV4TopicsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4TopicsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4TopicsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics")
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
func NewPostApiV4TopicsMergeRequest(server string, body PostApiV4TopicsMergeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4TopicsMergeRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4TopicsMergeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics/merge")
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
func NewDeleteApiV4TopicsIdRequest(server string, id int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4TopicsIdRequest(server string, id int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics/%s", pathParam0)
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
func NewPutApiV4TopicsIdRequest(server string, id int32, body PutApiV4TopicsIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4TopicsIdRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4TopicsIdRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/topics/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func (r GetApiV4TopicsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4TopicsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4TopicsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4TopicsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4TopicsMergeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4TopicsMergeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4TopicsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4TopicsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4TopicsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4TopicsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4TopicsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4TopicsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4TopicsWithResponse(ctx context.Context, params *GetApiV4TopicsParams, reqEditors ...RequestEditorFn) (*GetApiV4TopicsResponse, error) {
	rsp, err := c.GetApiV4Topics(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4TopicsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4TopicsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4TopicsResponse, error) {
	rsp, err := c.PostApiV4TopicsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4TopicsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4TopicsWithResponse(ctx context.Context, body PostApiV4TopicsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4TopicsResponse, error) {
	rsp, err := c.PostApiV4Topics(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4TopicsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4TopicsMergeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4TopicsMergeResponse, error) {
	rsp, err := c.PostApiV4TopicsMergeWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4TopicsMergeResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4TopicsMergeWithResponse(ctx context.Context, body PostApiV4TopicsMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4TopicsMergeResponse, error) {
	rsp, err := c.PostApiV4TopicsMerge(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4TopicsMergeResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4TopicsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4TopicsIdResponse, error) {
	rsp, err := c.DeleteApiV4TopicsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4TopicsIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4TopicsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4TopicsIdResponse, error) {
	rsp, err := c.GetApiV4TopicsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4TopicsIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4TopicsIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4TopicsIdResponse, error) {
	rsp, err := c.PutApiV4TopicsIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4TopicsIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4TopicsIdWithResponse(ctx context.Context, id int32, body PutApiV4TopicsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4TopicsIdResponse, error) {
	rsp, err := c.PutApiV4TopicsId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4TopicsIdResponse(rsp)
}
func ParseGetApiV4TopicsResponse(rsp *http.Response) (*GetApiV4TopicsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4TopicsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AvatarUrl          *string `json:"avatar_url,omitempty"`
			Description        *string `json:"description,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Name               *string `json:"name,omitempty"`
			OrganizationId     *string `json:"organization_id,omitempty"`
			Title              *string `json:"title,omitempty"`
			TotalProjectsCount *string `json:"total_projects_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4TopicsResponse(rsp *http.Response) (*PostApiV4TopicsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4TopicsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			AvatarUrl          *string `json:"avatar_url,omitempty"`
			Description        *string `json:"description,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Name               *string `json:"name,omitempty"`
			OrganizationId     *string `json:"organization_id,omitempty"`
			Title              *string `json:"title,omitempty"`
			TotalProjectsCount *string `json:"total_projects_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4TopicsMergeResponse(rsp *http.Response) (*PostApiV4TopicsMergeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4TopicsMergeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			AvatarUrl          *string `json:"avatar_url,omitempty"`
			Description        *string `json:"description,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Name               *string `json:"name,omitempty"`
			OrganizationId     *string `json:"organization_id,omitempty"`
			Title              *string `json:"title,omitempty"`
			TotalProjectsCount *string `json:"total_projects_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4TopicsIdResponse(rsp *http.Response) (*DeleteApiV4TopicsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4TopicsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4TopicsIdResponse(rsp *http.Response) (*GetApiV4TopicsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4TopicsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AvatarUrl          *string `json:"avatar_url,omitempty"`
			Description        *string `json:"description,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Name               *string `json:"name,omitempty"`
			OrganizationId     *string `json:"organization_id,omitempty"`
			Title              *string `json:"title,omitempty"`
			TotalProjectsCount *string `json:"total_projects_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4TopicsIdResponse(rsp *http.Response) (*PutApiV4TopicsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4TopicsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AvatarUrl          *string `json:"avatar_url,omitempty"`
			Description        *string `json:"description,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Name               *string `json:"name,omitempty"`
			OrganizationId     *string `json:"organization_id,omitempty"`
			Title              *string `json:"title,omitempty"`
			TotalProjectsCount *string `json:"total_projects_count,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
