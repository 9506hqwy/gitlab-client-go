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

	"github.com/oapi-codegen/runtime"
)

// Defines values for PostApiV4BroadcastMessagesJSONBodyBroadcastType.
const (
	PostApiV4BroadcastMessagesJSONBodyBroadcastTypeBanner       PostApiV4BroadcastMessagesJSONBodyBroadcastType = "banner"
	PostApiV4BroadcastMessagesJSONBodyBroadcastTypeNotification PostApiV4BroadcastMessagesJSONBodyBroadcastType = "notification"
)

// Defines values for PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels.
const (
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN10 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 10
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN15 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 15
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN20 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 20
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN30 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 30
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN40 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 40
	PostApiV4BroadcastMessagesJSONBodyTargetAccessLevelsN50 PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels = 50
)

// Defines values for PostApiV4BroadcastMessagesJSONBodyTheme.
const (
	PostApiV4BroadcastMessagesJSONBodyThemeBlue        PostApiV4BroadcastMessagesJSONBodyTheme = "blue"
	PostApiV4BroadcastMessagesJSONBodyThemeDark        PostApiV4BroadcastMessagesJSONBodyTheme = "dark"
	PostApiV4BroadcastMessagesJSONBodyThemeGreen       PostApiV4BroadcastMessagesJSONBodyTheme = "green"
	PostApiV4BroadcastMessagesJSONBodyThemeIndigo      PostApiV4BroadcastMessagesJSONBodyTheme = "indigo"
	PostApiV4BroadcastMessagesJSONBodyThemeLight       PostApiV4BroadcastMessagesJSONBodyTheme = "light"
	PostApiV4BroadcastMessagesJSONBodyThemeLightBlue   PostApiV4BroadcastMessagesJSONBodyTheme = "light-blue"
	PostApiV4BroadcastMessagesJSONBodyThemeLightGreen  PostApiV4BroadcastMessagesJSONBodyTheme = "light-green"
	PostApiV4BroadcastMessagesJSONBodyThemeLightIndigo PostApiV4BroadcastMessagesJSONBodyTheme = "light-indigo"
	PostApiV4BroadcastMessagesJSONBodyThemeLightRed    PostApiV4BroadcastMessagesJSONBodyTheme = "light-red"
	PostApiV4BroadcastMessagesJSONBodyThemeRed         PostApiV4BroadcastMessagesJSONBodyTheme = "red"
)

// Defines values for PutApiV4BroadcastMessagesIdJSONBodyBroadcastType.
const (
	PutApiV4BroadcastMessagesIdJSONBodyBroadcastTypeBanner       PutApiV4BroadcastMessagesIdJSONBodyBroadcastType = "banner"
	PutApiV4BroadcastMessagesIdJSONBodyBroadcastTypeNotification PutApiV4BroadcastMessagesIdJSONBodyBroadcastType = "notification"
)

// Defines values for PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels.
const (
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN10 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 10
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN15 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 15
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN20 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 20
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN30 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 30
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN40 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 40
	PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevelsN50 PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels = 50
)

// Defines values for PutApiV4BroadcastMessagesIdJSONBodyTheme.
const (
	PutApiV4BroadcastMessagesIdJSONBodyThemeBlue        PutApiV4BroadcastMessagesIdJSONBodyTheme = "blue"
	PutApiV4BroadcastMessagesIdJSONBodyThemeDark        PutApiV4BroadcastMessagesIdJSONBodyTheme = "dark"
	PutApiV4BroadcastMessagesIdJSONBodyThemeGreen       PutApiV4BroadcastMessagesIdJSONBodyTheme = "green"
	PutApiV4BroadcastMessagesIdJSONBodyThemeIndigo      PutApiV4BroadcastMessagesIdJSONBodyTheme = "indigo"
	PutApiV4BroadcastMessagesIdJSONBodyThemeLight       PutApiV4BroadcastMessagesIdJSONBodyTheme = "light"
	PutApiV4BroadcastMessagesIdJSONBodyThemeLightBlue   PutApiV4BroadcastMessagesIdJSONBodyTheme = "light-blue"
	PutApiV4BroadcastMessagesIdJSONBodyThemeLightGreen  PutApiV4BroadcastMessagesIdJSONBodyTheme = "light-green"
	PutApiV4BroadcastMessagesIdJSONBodyThemeLightIndigo PutApiV4BroadcastMessagesIdJSONBodyTheme = "light-indigo"
	PutApiV4BroadcastMessagesIdJSONBodyThemeLightRed    PutApiV4BroadcastMessagesIdJSONBodyTheme = "light-red"
	PutApiV4BroadcastMessagesIdJSONBodyThemeRed         PutApiV4BroadcastMessagesIdJSONBodyTheme = "red"
)

type GetApiV4BroadcastMessagesParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4BroadcastMessagesJSONBody struct {
	// BroadcastType Broadcast type. Defaults to banner
	BroadcastType *PostApiV4BroadcastMessagesJSONBodyBroadcastType `json:"broadcast_type,omitempty"`

	// Color Background color (Deprecated. Use "theme" instead.)
	Color *string `json:"color,omitempty"`

	// Dismissable Is dismissable
	Dismissable *bool `json:"dismissable,omitempty"`

	// EndsAt Ending time
	EndsAt *time.Time `json:"ends_at,omitempty"`

	// Font Foreground color (Deprecated. Use "theme" instead.)
	Font *string `json:"font,omitempty"`

	// Message Message to display
	Message string `json:"message"`

	// StartsAt Starting time
	StartsAt *time.Time `json:"starts_at,omitempty"`

	// TargetAccessLevels Target user roles
	TargetAccessLevels *[]PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels `json:"target_access_levels,omitempty"`

	// TargetPath Target path
	TargetPath *string `json:"target_path,omitempty"`

	// Theme The theme for the message
	Theme *PostApiV4BroadcastMessagesJSONBodyTheme `json:"theme,omitempty"`
}
type PostApiV4BroadcastMessagesJSONBodyBroadcastType string
type PostApiV4BroadcastMessagesJSONBodyTargetAccessLevels int32
type PostApiV4BroadcastMessagesJSONBodyTheme string
type PutApiV4BroadcastMessagesIdJSONBody struct {
	// BroadcastType Broadcast Type
	BroadcastType *PutApiV4BroadcastMessagesIdJSONBodyBroadcastType `json:"broadcast_type,omitempty"`

	// Color Background color (Deprecated. Use "theme" instead.)
	Color *string `json:"color,omitempty"`

	// Dismissable Is dismissable
	Dismissable *bool `json:"dismissable,omitempty"`

	// EndsAt Ending time
	EndsAt *time.Time `json:"ends_at,omitempty"`

	// Font Foreground color (Deprecated. Use "theme" instead.)
	Font *string `json:"font,omitempty"`

	// Message Message to display
	Message *string `json:"message,omitempty"`

	// StartsAt Starting time
	StartsAt *time.Time `json:"starts_at,omitempty"`

	// TargetAccessLevels Target user roles
	TargetAccessLevels *[]PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels `json:"target_access_levels,omitempty"`

	// TargetPath Target path
	TargetPath *string `json:"target_path,omitempty"`

	// Theme The theme for the message
	Theme *PutApiV4BroadcastMessagesIdJSONBodyTheme `json:"theme,omitempty"`
}
type PutApiV4BroadcastMessagesIdJSONBodyBroadcastType string
type PutApiV4BroadcastMessagesIdJSONBodyTargetAccessLevels int32
type PutApiV4BroadcastMessagesIdJSONBodyTheme string
type PostApiV4BroadcastMessagesJSONRequestBody PostApiV4BroadcastMessagesJSONBody
type PutApiV4BroadcastMessagesIdJSONRequestBody PutApiV4BroadcastMessagesIdJSONBody
type GetApiV4BroadcastMessagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active             *string `json:"active,omitempty"`
		BroadcastType      *string `json:"broadcast_type,omitempty"`
		Color              *string `json:"color,omitempty"`
		Dismissable        *string `json:"dismissable,omitempty"`
		EndsAt             *string `json:"ends_at,omitempty"`
		Font               *string `json:"font,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Message            *string `json:"message,omitempty"`
		StartsAt           *string `json:"starts_at,omitempty"`
		TargetAccessLevels *string `json:"target_access_levels,omitempty"`
		TargetPath         *string `json:"target_path,omitempty"`
		Theme              *string `json:"theme,omitempty"`
	}
}
type PostApiV4BroadcastMessagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Active             *string `json:"active,omitempty"`
		BroadcastType      *string `json:"broadcast_type,omitempty"`
		Color              *string `json:"color,omitempty"`
		Dismissable        *string `json:"dismissable,omitempty"`
		EndsAt             *string `json:"ends_at,omitempty"`
		Font               *string `json:"font,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Message            *string `json:"message,omitempty"`
		StartsAt           *string `json:"starts_at,omitempty"`
		TargetAccessLevels *string `json:"target_access_levels,omitempty"`
		TargetPath         *string `json:"target_path,omitempty"`
		Theme              *string `json:"theme,omitempty"`
	}
}
type DeleteApiV4BroadcastMessagesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active             *string `json:"active,omitempty"`
		BroadcastType      *string `json:"broadcast_type,omitempty"`
		Color              *string `json:"color,omitempty"`
		Dismissable        *string `json:"dismissable,omitempty"`
		EndsAt             *string `json:"ends_at,omitempty"`
		Font               *string `json:"font,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Message            *string `json:"message,omitempty"`
		StartsAt           *string `json:"starts_at,omitempty"`
		TargetAccessLevels *string `json:"target_access_levels,omitempty"`
		TargetPath         *string `json:"target_path,omitempty"`
		Theme              *string `json:"theme,omitempty"`
	}
}
type GetApiV4BroadcastMessagesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active             *string `json:"active,omitempty"`
		BroadcastType      *string `json:"broadcast_type,omitempty"`
		Color              *string `json:"color,omitempty"`
		Dismissable        *string `json:"dismissable,omitempty"`
		EndsAt             *string `json:"ends_at,omitempty"`
		Font               *string `json:"font,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Message            *string `json:"message,omitempty"`
		StartsAt           *string `json:"starts_at,omitempty"`
		TargetAccessLevels *string `json:"target_access_levels,omitempty"`
		TargetPath         *string `json:"target_path,omitempty"`
		Theme              *string `json:"theme,omitempty"`
	}
}
type PutApiV4BroadcastMessagesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active             *string `json:"active,omitempty"`
		BroadcastType      *string `json:"broadcast_type,omitempty"`
		Color              *string `json:"color,omitempty"`
		Dismissable        *string `json:"dismissable,omitempty"`
		EndsAt             *string `json:"ends_at,omitempty"`
		Font               *string `json:"font,omitempty"`
		Id                 *string `json:"id,omitempty"`
		Message            *string `json:"message,omitempty"`
		StartsAt           *string `json:"starts_at,omitempty"`
		TargetAccessLevels *string `json:"target_access_levels,omitempty"`
		TargetPath         *string `json:"target_path,omitempty"`
		Theme              *string `json:"theme,omitempty"`
	}
}

func (c *Client) GetApiV4BroadcastMessages(ctx context.Context, params *GetApiV4BroadcastMessagesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BroadcastMessagesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4BroadcastMessagesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4BroadcastMessagesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4BroadcastMessages(ctx context.Context, body PostApiV4BroadcastMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4BroadcastMessagesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4BroadcastMessagesId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4BroadcastMessagesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BroadcastMessagesId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BroadcastMessagesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4BroadcastMessagesIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4BroadcastMessagesIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4BroadcastMessagesId(ctx context.Context, id int32, body PutApiV4BroadcastMessagesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4BroadcastMessagesIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4BroadcastMessagesRequest(server string, params *GetApiV4BroadcastMessagesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/broadcast_messages")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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
func NewPostApiV4BroadcastMessagesRequest(server string, body PostApiV4BroadcastMessagesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4BroadcastMessagesRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4BroadcastMessagesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/broadcast_messages")
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
func NewDeleteApiV4BroadcastMessagesIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/broadcast_messages/%s", pathParam0)
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
func NewGetApiV4BroadcastMessagesIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/broadcast_messages/%s", pathParam0)
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
func NewPutApiV4BroadcastMessagesIdRequest(server string, id int32, body PutApiV4BroadcastMessagesIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4BroadcastMessagesIdRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4BroadcastMessagesIdRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/broadcast_messages/%s", pathParam0)
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
func (r GetApiV4BroadcastMessagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BroadcastMessagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4BroadcastMessagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4BroadcastMessagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4BroadcastMessagesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4BroadcastMessagesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BroadcastMessagesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BroadcastMessagesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4BroadcastMessagesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4BroadcastMessagesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4BroadcastMessagesWithResponse(ctx context.Context, params *GetApiV4BroadcastMessagesParams, reqEditors ...RequestEditorFn) (*GetApiV4BroadcastMessagesResponse, error) {
	rsp, err := c.GetApiV4BroadcastMessages(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BroadcastMessagesResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4BroadcastMessagesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4BroadcastMessagesResponse, error) {
	rsp, err := c.PostApiV4BroadcastMessagesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4BroadcastMessagesResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4BroadcastMessagesWithResponse(ctx context.Context, body PostApiV4BroadcastMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4BroadcastMessagesResponse, error) {
	rsp, err := c.PostApiV4BroadcastMessages(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4BroadcastMessagesResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4BroadcastMessagesIdResponse, error) {
	rsp, err := c.DeleteApiV4BroadcastMessagesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4BroadcastMessagesIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4BroadcastMessagesIdResponse, error) {
	rsp, err := c.GetApiV4BroadcastMessagesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BroadcastMessagesIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4BroadcastMessagesIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4BroadcastMessagesIdResponse, error) {
	rsp, err := c.PutApiV4BroadcastMessagesIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4BroadcastMessagesIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, body PutApiV4BroadcastMessagesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4BroadcastMessagesIdResponse, error) {
	rsp, err := c.PutApiV4BroadcastMessagesId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4BroadcastMessagesIdResponse(rsp)
}
func ParseGetApiV4BroadcastMessagesResponse(rsp *http.Response) (*GetApiV4BroadcastMessagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BroadcastMessagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active             *string `json:"active,omitempty"`
			BroadcastType      *string `json:"broadcast_type,omitempty"`
			Color              *string `json:"color,omitempty"`
			Dismissable        *string `json:"dismissable,omitempty"`
			EndsAt             *string `json:"ends_at,omitempty"`
			Font               *string `json:"font,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Message            *string `json:"message,omitempty"`
			StartsAt           *string `json:"starts_at,omitempty"`
			TargetAccessLevels *string `json:"target_access_levels,omitempty"`
			TargetPath         *string `json:"target_path,omitempty"`
			Theme              *string `json:"theme,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4BroadcastMessagesResponse(rsp *http.Response) (*PostApiV4BroadcastMessagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4BroadcastMessagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Active             *string `json:"active,omitempty"`
			BroadcastType      *string `json:"broadcast_type,omitempty"`
			Color              *string `json:"color,omitempty"`
			Dismissable        *string `json:"dismissable,omitempty"`
			EndsAt             *string `json:"ends_at,omitempty"`
			Font               *string `json:"font,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Message            *string `json:"message,omitempty"`
			StartsAt           *string `json:"starts_at,omitempty"`
			TargetAccessLevels *string `json:"target_access_levels,omitempty"`
			TargetPath         *string `json:"target_path,omitempty"`
			Theme              *string `json:"theme,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4BroadcastMessagesIdResponse(rsp *http.Response) (*DeleteApiV4BroadcastMessagesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4BroadcastMessagesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active             *string `json:"active,omitempty"`
			BroadcastType      *string `json:"broadcast_type,omitempty"`
			Color              *string `json:"color,omitempty"`
			Dismissable        *string `json:"dismissable,omitempty"`
			EndsAt             *string `json:"ends_at,omitempty"`
			Font               *string `json:"font,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Message            *string `json:"message,omitempty"`
			StartsAt           *string `json:"starts_at,omitempty"`
			TargetAccessLevels *string `json:"target_access_levels,omitempty"`
			TargetPath         *string `json:"target_path,omitempty"`
			Theme              *string `json:"theme,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BroadcastMessagesIdResponse(rsp *http.Response) (*GetApiV4BroadcastMessagesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BroadcastMessagesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active             *string `json:"active,omitempty"`
			BroadcastType      *string `json:"broadcast_type,omitempty"`
			Color              *string `json:"color,omitempty"`
			Dismissable        *string `json:"dismissable,omitempty"`
			EndsAt             *string `json:"ends_at,omitempty"`
			Font               *string `json:"font,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Message            *string `json:"message,omitempty"`
			StartsAt           *string `json:"starts_at,omitempty"`
			TargetAccessLevels *string `json:"target_access_levels,omitempty"`
			TargetPath         *string `json:"target_path,omitempty"`
			Theme              *string `json:"theme,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4BroadcastMessagesIdResponse(rsp *http.Response) (*PutApiV4BroadcastMessagesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4BroadcastMessagesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active             *string `json:"active,omitempty"`
			BroadcastType      *string `json:"broadcast_type,omitempty"`
			Color              *string `json:"color,omitempty"`
			Dismissable        *string `json:"dismissable,omitempty"`
			EndsAt             *string `json:"ends_at,omitempty"`
			Font               *string `json:"font,omitempty"`
			Id                 *string `json:"id,omitempty"`
			Message            *string `json:"message,omitempty"`
			StartsAt           *string `json:"starts_at,omitempty"`
			TargetAccessLevels *string `json:"target_access_levels,omitempty"`
			TargetPath         *string `json:"target_path,omitempty"`
			Theme              *string `json:"theme,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
