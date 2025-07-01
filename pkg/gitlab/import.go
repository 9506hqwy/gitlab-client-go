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

type PostApiV4ImportBitbucketJSONBody struct {
	// BitbucketAppPassword BitBucket app password
	BitbucketAppPassword string `json:"bitbucket_app_password"`

	// BitbucketUsername BitBucket username
	BitbucketUsername string `json:"bitbucket_username"`

	// NewName New repository name
	NewName *string `json:"new_name,omitempty"`

	// RepoPath Repository path
	RepoPath string `json:"repo_path"`

	// TargetNamespace Target namespace
	TargetNamespace string `json:"target_namespace"`
}
type PostApiV4ImportBitbucketServerJSONBody struct {
	// BitbucketServerProject BitBucket Server Project Key
	BitbucketServerProject string `json:"bitbucket_server_project"`

	// BitbucketServerRepo BitBucket Server Repository Name
	BitbucketServerRepo string `json:"bitbucket_server_repo"`

	// BitbucketServerUrl Bitbucket Server URL
	BitbucketServerUrl string `json:"bitbucket_server_url"`

	// BitbucketServerUsername BitBucket Server Username
	BitbucketServerUsername string `json:"bitbucket_server_username"`

	// NewName New repo name
	NewName *string `json:"new_name,omitempty"`

	// NewNamespace Namespace to import repo into
	NewNamespace *string `json:"new_namespace,omitempty"`

	// PersonalAccessToken BitBucket Server personal access token/password
	PersonalAccessToken string `json:"personal_access_token"`

	// TimeoutStrategy Strategy for behavior on timeouts
	TimeoutStrategy *string `json:"timeout_strategy,omitempty"`
}
type PostApiV4ImportGithubJSONBody struct {
	// GithubHostname Custom GitHub enterprise hostname. For example: https://github.example.com. From GitLab 16.5 to GitLab 17.1, you must include the path `/api/v3`.
	GithubHostname *string `json:"github_hostname,omitempty"`

	// NewName New repo name
	NewName *string `json:"new_name,omitempty"`

	// OptionalStages Optional stages of import to be performed
	OptionalStages *map[string]interface{} `json:"optional_stages,omitempty"`

	// PaginationLimit Pagination limit
	PaginationLimit *int32 `json:"pagination_limit,omitempty"`

	// PersonalAccessToken GitHub personal access token
	PersonalAccessToken string `json:"personal_access_token"`

	// RepoId GitHub repository ID
	RepoId int32 `json:"repo_id"`

	// TargetNamespace Namespace or group to import repository into
	TargetNamespace string `json:"target_namespace"`

	// TimeoutStrategy Strategy for behavior on timeouts
	TimeoutStrategy *string `json:"timeout_strategy,omitempty"`
}
type PostApiV4ImportGithubCancelJSONBody struct {
	// ProjectId ID of importing project to be canceled
	ProjectId int32 `json:"project_id"`
}
type PostApiV4ImportGithubGistsJSONBody struct {
	// PersonalAccessToken GitHub personal access token
	PersonalAccessToken string `json:"personal_access_token"`
}
type PostApiV4ImportBitbucketJSONRequestBody PostApiV4ImportBitbucketJSONBody
type PostApiV4ImportBitbucketServerJSONRequestBody PostApiV4ImportBitbucketServerJSONBody
type PostApiV4ImportGithubJSONRequestBody PostApiV4ImportGithubJSONBody
type PostApiV4ImportGithubCancelJSONRequestBody PostApiV4ImportGithubCancelJSONBody
type PostApiV4ImportGithubGistsJSONRequestBody PostApiV4ImportGithubGistsJSONBody
type PostApiV4ImportBitbucketResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Forked                *bool   `json:"forked,omitempty"`
		FullName              *string `json:"full_name,omitempty"`
		FullPath              *string `json:"full_path,omitempty"`
		HumanImportStatusName *string `json:"human_import_status_name,omitempty"`
		Id                    *int32  `json:"id,omitempty"`
		ImportError           *string `json:"import_error,omitempty"`
		ImportSource          *string `json:"import_source,omitempty"`
		ImportStatus          *string `json:"import_status,omitempty"`
		ImportWarning         *string `json:"import_warning,omitempty"`
		Name                  *string `json:"name,omitempty"`
		ProviderLink          *string `json:"provider_link,omitempty"`
		RefsUrl               *string `json:"refs_url,omitempty"`
		RelationType          *string `json:"relation_type,omitempty"`
	}
}
type PostApiV4ImportBitbucketServerResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Forked   *bool   `json:"forked,omitempty"`
		FullName *string `json:"full_name,omitempty"`
		FullPath *string `json:"full_path,omitempty"`
		Id       *int32  `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		RefsUrl  *string `json:"refs_url,omitempty"`
	}
}
type PostApiV4ImportGithubResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Forked   *bool   `json:"forked,omitempty"`
		FullName *string `json:"full_name,omitempty"`
		FullPath *string `json:"full_path,omitempty"`
		Id       *int32  `json:"id,omitempty"`
		Name     *string `json:"name,omitempty"`
		RefsUrl  *string `json:"refs_url,omitempty"`
	}
}
type PostApiV4ImportGithubCancelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Forked                *bool   `json:"forked,omitempty"`
		FullName              *string `json:"full_name,omitempty"`
		FullPath              *string `json:"full_path,omitempty"`
		HumanImportStatusName *string `json:"human_import_status_name,omitempty"`
		Id                    *int32  `json:"id,omitempty"`
		ImportError           *string `json:"import_error,omitempty"`
		ImportSource          *string `json:"import_source,omitempty"`
		ImportStatus          *string `json:"import_status,omitempty"`
		ImportWarning         *string `json:"import_warning,omitempty"`
		Name                  *string `json:"name,omitempty"`
		ProviderLink          *string `json:"provider_link,omitempty"`
		RefsUrl               *string `json:"refs_url,omitempty"`
		RelationType          *string `json:"relation_type,omitempty"`
	}
}
type PostApiV4ImportGithubGistsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) PostApiV4ImportBitbucketWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportBitbucketRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportBitbucket(ctx context.Context, body PostApiV4ImportBitbucketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportBitbucketRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportBitbucketServerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportBitbucketServerRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportBitbucketServer(ctx context.Context, body PostApiV4ImportBitbucketServerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportBitbucketServerRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithubWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithub(ctx context.Context, body PostApiV4ImportGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithubCancelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubCancelRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithubCancel(ctx context.Context, body PostApiV4ImportGithubCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubCancelRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithubGistsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubGistsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4ImportGithubGists(ctx context.Context, body PostApiV4ImportGithubGistsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4ImportGithubGistsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4ImportBitbucketRequest(server string, body PostApiV4ImportBitbucketJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ImportBitbucketRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ImportBitbucketRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/import/bitbucket")
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
func NewPostApiV4ImportBitbucketServerRequest(server string, body PostApiV4ImportBitbucketServerJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ImportBitbucketServerRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ImportBitbucketServerRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/import/bitbucket_server")
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
func NewPostApiV4ImportGithubRequest(server string, body PostApiV4ImportGithubJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ImportGithubRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ImportGithubRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/import/github")
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
func NewPostApiV4ImportGithubCancelRequest(server string, body PostApiV4ImportGithubCancelJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ImportGithubCancelRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ImportGithubCancelRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/import/github/cancel")
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
func NewPostApiV4ImportGithubGistsRequest(server string, body PostApiV4ImportGithubGistsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4ImportGithubGistsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4ImportGithubGistsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/import/github/gists")
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
func (r PostApiV4ImportBitbucketResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ImportBitbucketResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ImportBitbucketServerResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ImportBitbucketServerResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ImportGithubResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ImportGithubResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ImportGithubCancelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ImportGithubCancelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4ImportGithubGistsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4ImportGithubGistsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4ImportBitbucketWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketResponse, error) {
	rsp, err := c.PostApiV4ImportBitbucketWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportBitbucketResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportBitbucketWithResponse(ctx context.Context, body PostApiV4ImportBitbucketJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketResponse, error) {
	rsp, err := c.PostApiV4ImportBitbucket(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportBitbucketResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportBitbucketServerWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketServerResponse, error) {
	rsp, err := c.PostApiV4ImportBitbucketServerWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportBitbucketServerResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportBitbucketServerWithResponse(ctx context.Context, body PostApiV4ImportBitbucketServerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketServerResponse, error) {
	rsp, err := c.PostApiV4ImportBitbucketServer(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportBitbucketServerResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubResponse, error) {
	rsp, err := c.PostApiV4ImportGithubWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubWithResponse(ctx context.Context, body PostApiV4ImportGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubResponse, error) {
	rsp, err := c.PostApiV4ImportGithub(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubCancelWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubCancelResponse, error) {
	rsp, err := c.PostApiV4ImportGithubCancelWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubCancelResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubCancelWithResponse(ctx context.Context, body PostApiV4ImportGithubCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubCancelResponse, error) {
	rsp, err := c.PostApiV4ImportGithubCancel(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubCancelResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubGistsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubGistsResponse, error) {
	rsp, err := c.PostApiV4ImportGithubGistsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubGistsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4ImportGithubGistsWithResponse(ctx context.Context, body PostApiV4ImportGithubGistsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubGistsResponse, error) {
	rsp, err := c.PostApiV4ImportGithubGists(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4ImportGithubGistsResponse(rsp)
}
func ParsePostApiV4ImportBitbucketResponse(rsp *http.Response) (*PostApiV4ImportBitbucketResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ImportBitbucketResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Forked                *bool   `json:"forked,omitempty"`
			FullName              *string `json:"full_name,omitempty"`
			FullPath              *string `json:"full_path,omitempty"`
			HumanImportStatusName *string `json:"human_import_status_name,omitempty"`
			Id                    *int32  `json:"id,omitempty"`
			ImportError           *string `json:"import_error,omitempty"`
			ImportSource          *string `json:"import_source,omitempty"`
			ImportStatus          *string `json:"import_status,omitempty"`
			ImportWarning         *string `json:"import_warning,omitempty"`
			Name                  *string `json:"name,omitempty"`
			ProviderLink          *string `json:"provider_link,omitempty"`
			RefsUrl               *string `json:"refs_url,omitempty"`
			RelationType          *string `json:"relation_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4ImportBitbucketServerResponse(rsp *http.Response) (*PostApiV4ImportBitbucketServerResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ImportBitbucketServerResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Forked   *bool   `json:"forked,omitempty"`
			FullName *string `json:"full_name,omitempty"`
			FullPath *string `json:"full_path,omitempty"`
			Id       *int32  `json:"id,omitempty"`
			Name     *string `json:"name,omitempty"`
			RefsUrl  *string `json:"refs_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4ImportGithubResponse(rsp *http.Response) (*PostApiV4ImportGithubResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ImportGithubResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Forked   *bool   `json:"forked,omitempty"`
			FullName *string `json:"full_name,omitempty"`
			FullPath *string `json:"full_path,omitempty"`
			Id       *int32  `json:"id,omitempty"`
			Name     *string `json:"name,omitempty"`
			RefsUrl  *string `json:"refs_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4ImportGithubCancelResponse(rsp *http.Response) (*PostApiV4ImportGithubCancelResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ImportGithubCancelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Forked                *bool   `json:"forked,omitempty"`
			FullName              *string `json:"full_name,omitempty"`
			FullPath              *string `json:"full_path,omitempty"`
			HumanImportStatusName *string `json:"human_import_status_name,omitempty"`
			Id                    *int32  `json:"id,omitempty"`
			ImportError           *string `json:"import_error,omitempty"`
			ImportSource          *string `json:"import_source,omitempty"`
			ImportStatus          *string `json:"import_status,omitempty"`
			ImportWarning         *string `json:"import_warning,omitempty"`
			Name                  *string `json:"name,omitempty"`
			ProviderLink          *string `json:"provider_link,omitempty"`
			RefsUrl               *string `json:"refs_url,omitempty"`
			RelationType          *string `json:"relation_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4ImportGithubGistsResponse(rsp *http.Response) (*PostApiV4ImportGithubGistsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4ImportGithubGistsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
