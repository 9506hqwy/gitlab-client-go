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

type PutApiV4SuggestionsBatchApplyJSONBody struct {
	// CommitMessage A custom commit message to use instead of the default generated message or the project's default message
	CommitMessage *string `json:"commit_message,omitempty"`

	// Ids An array of the suggestion IDs
	Ids []int32 `json:"ids"`
}
type PutApiV4SuggestionsIdApplyJSONBody struct {
	// CommitMessage A custom commit message to use instead of the default generated message or the project's default message
	CommitMessage *string `json:"commit_message,omitempty"`
}
type PutApiV4SuggestionsBatchApplyJSONRequestBody PutApiV4SuggestionsBatchApplyJSONBody
type PutApiV4SuggestionsIdApplyJSONRequestBody PutApiV4SuggestionsIdApplyJSONBody
type PutApiV4SuggestionsBatchApplyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Appliable   *string `json:"appliable,omitempty"`
		Applied     *string `json:"applied,omitempty"`
		FromContent *string `json:"from_content,omitempty"`
		FromLine    *string `json:"from_line,omitempty"`
		Id          *string `json:"id,omitempty"`
		ToContent   *string `json:"to_content,omitempty"`
		ToLine      *string `json:"to_line,omitempty"`
	}
}
type PutApiV4SuggestionsIdApplyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Appliable   *string `json:"appliable,omitempty"`
		Applied     *string `json:"applied,omitempty"`
		FromContent *string `json:"from_content,omitempty"`
		FromLine    *string `json:"from_line,omitempty"`
		Id          *string `json:"id,omitempty"`
		ToContent   *string `json:"to_content,omitempty"`
		ToLine      *string `json:"to_line,omitempty"`
	}
}

func (c *Client) PutApiV4SuggestionsBatchApplyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SuggestionsBatchApplyRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4SuggestionsBatchApply(ctx context.Context, body PutApiV4SuggestionsBatchApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SuggestionsBatchApplyRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4SuggestionsIdApplyWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SuggestionsIdApplyRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4SuggestionsIdApply(ctx context.Context, id int32, body PutApiV4SuggestionsIdApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SuggestionsIdApplyRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPutApiV4SuggestionsBatchApplyRequest(server string, body PutApiV4SuggestionsBatchApplyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4SuggestionsBatchApplyRequestWithBody(server, "application/json", bodyReader)
}
func NewPutApiV4SuggestionsBatchApplyRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/suggestions/batch_apply")
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
func NewPutApiV4SuggestionsIdApplyRequest(server string, id int32, body PutApiV4SuggestionsIdApplyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4SuggestionsIdApplyRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4SuggestionsIdApplyRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/suggestions/%s/apply", pathParam0)
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
func (r PutApiV4SuggestionsBatchApplyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4SuggestionsBatchApplyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4SuggestionsIdApplyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4SuggestionsIdApplyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PutApiV4SuggestionsBatchApplyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsBatchApplyResponse, error) {
	rsp, err := c.PutApiV4SuggestionsBatchApplyWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SuggestionsBatchApplyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4SuggestionsBatchApplyWithResponse(ctx context.Context, body PutApiV4SuggestionsBatchApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsBatchApplyResponse, error) {
	rsp, err := c.PutApiV4SuggestionsBatchApply(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SuggestionsBatchApplyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4SuggestionsIdApplyWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsIdApplyResponse, error) {
	rsp, err := c.PutApiV4SuggestionsIdApplyWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SuggestionsIdApplyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4SuggestionsIdApplyWithResponse(ctx context.Context, id int32, body PutApiV4SuggestionsIdApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsIdApplyResponse, error) {
	rsp, err := c.PutApiV4SuggestionsIdApply(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SuggestionsIdApplyResponse(rsp)
}
func ParsePutApiV4SuggestionsBatchApplyResponse(rsp *http.Response) (*PutApiV4SuggestionsBatchApplyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4SuggestionsBatchApplyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Appliable   *string `json:"appliable,omitempty"`
			Applied     *string `json:"applied,omitempty"`
			FromContent *string `json:"from_content,omitempty"`
			FromLine    *string `json:"from_line,omitempty"`
			Id          *string `json:"id,omitempty"`
			ToContent   *string `json:"to_content,omitempty"`
			ToLine      *string `json:"to_line,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4SuggestionsIdApplyResponse(rsp *http.Response) (*PutApiV4SuggestionsIdApplyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4SuggestionsIdApplyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Appliable   *string `json:"appliable,omitempty"`
			Applied     *string `json:"applied,omitempty"`
			FromContent *string `json:"from_content,omitempty"`
			FromLine    *string `json:"from_line,omitempty"`
			Id          *string `json:"id,omitempty"`
			ToContent   *string `json:"to_content,omitempty"`
			ToLine      *string `json:"to_line,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
