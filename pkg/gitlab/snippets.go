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

type GetApiV4SnippetsParams struct {
	// CreatedAfter Return snippets created after the specified time
	CreatedAfter *time.Time `form:"created_after,omitempty" json:"created_after,omitempty"`

	// CreatedBefore Return snippets created before the specified time
	CreatedBefore *time.Time `form:"created_before,omitempty" json:"created_before,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4SnippetsJSONBody struct {
	// Content The content of a snippet
	Content *string `json:"content,omitempty"`

	// Description The description of a snippet
	Description *string `json:"description,omitempty"`

	// FileName The name of a snippet file
	FileName string `json:"file_name"`

	// Files An array of files
	Files *[]struct {
		// Content The content of a snippet file
		Content string `json:"content"`

		// FilePath The path of a snippet file
		FilePath string `json:"file_path"`
	} `json:"files,omitempty"`

	// Title The title of a snippet
	Title string `json:"title"`

	// Visibility The visibility of the snippet
	Visibility *string `json:"visibility,omitempty"`
}
type GetApiV4SnippetsAllParams struct {
	// CreatedAfter Return snippets created after the specified time
	CreatedAfter *time.Time `form:"created_after,omitempty" json:"created_after,omitempty"`

	// CreatedBefore Return snippets created before the specified time
	CreatedBefore *time.Time `form:"created_before,omitempty" json:"created_before,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// RepositoryStorage Filter by repository storage used by the snippet
	RepositoryStorage *string `form:"repository_storage,omitempty" json:"repository_storage,omitempty"`
}
type GetApiV4SnippetsPublicParams struct {
	// CreatedAfter Return snippets created after the specified time
	CreatedAfter *time.Time `form:"created_after,omitempty" json:"created_after,omitempty"`

	// CreatedBefore Return snippets created before the specified time
	CreatedBefore *time.Time `form:"created_before,omitempty" json:"created_before,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PutApiV4SnippetsIdJSONBody struct {
	// Content The content of a snippet
	Content *string `json:"content,omitempty"`

	// Description The description of a snippet
	Description *string `json:"description,omitempty"`

	// FileName The name of a snippet file
	FileName *string `json:"file_name,omitempty"`

	// Files An array of files to update
	Files *[]struct {
		// Action The type of action to perform on the file, must be one of: create, update, delete, move
		Action string `json:"action"`

		// Content The content of a snippet
		Content *string `json:"content,omitempty"`

		// FilePath The file path of a snippet file
		FilePath *string `json:"file_path,omitempty"`

		// PreviousPath The previous path of a snippet file
		PreviousPath *string `json:"previous_path,omitempty"`
	} `json:"files,omitempty"`

	// Title The title of a snippet
	Title *string `json:"title,omitempty"`

	// Visibility The visibility of the snippet
	Visibility *string `json:"visibility,omitempty"`
}
type PostApiV4SnippetsJSONRequestBody PostApiV4SnippetsJSONBody
type PutApiV4SnippetsIdJSONRequestBody PutApiV4SnippetsIdJSONBody
type GetApiV4SnippetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type PostApiV4SnippetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type GetApiV4SnippetsAllResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type GetApiV4SnippetsPublicResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type DeleteApiV4SnippetsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type GetApiV4SnippetsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type PutApiV4SnippetsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Author API_Entities_UserBasic model
		Author *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		Description       *string    `json:"description,omitempty"`
		FileName          *string    `json:"file_name,omitempty"`
		Files             *[]string  `json:"files,omitempty"`
		HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
		Id                *int32     `json:"id,omitempty"`
		Imported          *bool      `json:"imported,omitempty"`
		ImportedFrom      *string    `json:"imported_from,omitempty"`
		ProjectId         *int32     `json:"project_id,omitempty"`
		RawUrl            *string    `json:"raw_url,omitempty"`
		RepositoryStorage *string    `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
		Title             *string    `json:"title,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
		Visibility        *string    `json:"visibility,omitempty"`
		WebUrl            *string    `json:"web_url,omitempty"`
	}
}
type GetApiV4SnippetsIdFilesRefFilePathRawResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4SnippetsIdRawResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4SnippetsIdUserAgentDetailResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AkismetSubmitted *bool   `json:"akismet_submitted,omitempty"`
		IpAddress        *string `json:"ip_address,omitempty"`
		UserAgent        *string `json:"user_agent,omitempty"`
	}
}

func (c *Client) GetApiV4Snippets(ctx context.Context, params *GetApiV4SnippetsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4SnippetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4SnippetsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Snippets(ctx context.Context, body PostApiV4SnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4SnippetsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsAll(ctx context.Context, params *GetApiV4SnippetsAllParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsAllRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsPublic(ctx context.Context, params *GetApiV4SnippetsPublicParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsPublicRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4SnippetsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4SnippetsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4SnippetsIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SnippetsIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4SnippetsId(ctx context.Context, id int32, body PutApiV4SnippetsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4SnippetsIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsIdFilesRefFilePathRaw(ctx context.Context, id int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsIdFilesRefFilePathRawRequest(c.Server, id, ref, filePath)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsIdRaw(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsIdRawRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4SnippetsIdUserAgentDetail(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4SnippetsIdUserAgentDetailRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4SnippetsRequest(server string, params *GetApiV4SnippetsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/snippets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.CreatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_after", runtime.ParamLocationQuery, *params.CreatedAfter); err != nil {
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

		if params.CreatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_before", runtime.ParamLocationQuery, *params.CreatedBefore); err != nil {
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
func NewPostApiV4SnippetsRequest(server string, body PostApiV4SnippetsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4SnippetsRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4SnippetsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/snippets")
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
func NewGetApiV4SnippetsAllRequest(server string, params *GetApiV4SnippetsAllParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/snippets/all")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.CreatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_after", runtime.ParamLocationQuery, *params.CreatedAfter); err != nil {
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

		if params.CreatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_before", runtime.ParamLocationQuery, *params.CreatedBefore); err != nil {
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

		if params.RepositoryStorage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "repository_storage", runtime.ParamLocationQuery, *params.RepositoryStorage); err != nil {
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
func NewGetApiV4SnippetsPublicRequest(server string, params *GetApiV4SnippetsPublicParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/snippets/public")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.CreatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_after", runtime.ParamLocationQuery, *params.CreatedAfter); err != nil {
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

		if params.CreatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "created_before", runtime.ParamLocationQuery, *params.CreatedBefore); err != nil {
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
func NewDeleteApiV4SnippetsIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/snippets/%s", pathParam0)
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
func NewGetApiV4SnippetsIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/snippets/%s", pathParam0)
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
func NewPutApiV4SnippetsIdRequest(server string, id int32, body PutApiV4SnippetsIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4SnippetsIdRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4SnippetsIdRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/snippets/%s", pathParam0)
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
func NewGetApiV4SnippetsIdFilesRefFilePathRawRequest(server string, id int32, ref string, filePath string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "ref", runtime.ParamLocationPath, ref)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "file_path", runtime.ParamLocationPath, filePath)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/snippets/%s/files/%s/%s/raw", pathParam0, pathParam1, pathParam2)
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
func NewGetApiV4SnippetsIdRawRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/snippets/%s/raw", pathParam0)
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
func NewGetApiV4SnippetsIdUserAgentDetailRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/snippets/%s/user_agent_detail", pathParam0)
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
func (r GetApiV4SnippetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4SnippetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4SnippetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsAllResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsAllResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsPublicResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsPublicResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4SnippetsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4SnippetsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4SnippetsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4SnippetsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsIdFilesRefFilePathRawResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsIdFilesRefFilePathRawResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsIdRawResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsIdRawResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4SnippetsIdUserAgentDetailResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4SnippetsIdUserAgentDetailResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4SnippetsWithResponse(ctx context.Context, params *GetApiV4SnippetsParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsResponse, error) {
	rsp, err := c.GetApiV4Snippets(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4SnippetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4SnippetsResponse, error) {
	rsp, err := c.PostApiV4SnippetsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4SnippetsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4SnippetsWithResponse(ctx context.Context, body PostApiV4SnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4SnippetsResponse, error) {
	rsp, err := c.PostApiV4Snippets(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4SnippetsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsAllWithResponse(ctx context.Context, params *GetApiV4SnippetsAllParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsAllResponse, error) {
	rsp, err := c.GetApiV4SnippetsAll(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsAllResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsPublicWithResponse(ctx context.Context, params *GetApiV4SnippetsPublicParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsPublicResponse, error) {
	rsp, err := c.GetApiV4SnippetsPublic(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsPublicResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4SnippetsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4SnippetsIdResponse, error) {
	rsp, err := c.DeleteApiV4SnippetsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4SnippetsIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdResponse, error) {
	rsp, err := c.GetApiV4SnippetsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4SnippetsIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SnippetsIdResponse, error) {
	rsp, err := c.PutApiV4SnippetsIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SnippetsIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4SnippetsIdWithResponse(ctx context.Context, id int32, body PutApiV4SnippetsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SnippetsIdResponse, error) {
	rsp, err := c.PutApiV4SnippetsId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4SnippetsIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsIdFilesRefFilePathRawWithResponse(ctx context.Context, id int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdFilesRefFilePathRawResponse, error) {
	rsp, err := c.GetApiV4SnippetsIdFilesRefFilePathRaw(ctx, id, ref, filePath, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsIdFilesRefFilePathRawResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsIdRawWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdRawResponse, error) {
	rsp, err := c.GetApiV4SnippetsIdRaw(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsIdRawResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4SnippetsIdUserAgentDetailWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdUserAgentDetailResponse, error) {
	rsp, err := c.GetApiV4SnippetsIdUserAgentDetail(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4SnippetsIdUserAgentDetailResponse(rsp)
}
func ParseGetApiV4SnippetsResponse(rsp *http.Response) (*GetApiV4SnippetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4SnippetsResponse(rsp *http.Response) (*PostApiV4SnippetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4SnippetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseGetApiV4SnippetsAllResponse(rsp *http.Response) (*GetApiV4SnippetsAllResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsAllResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4SnippetsPublicResponse(rsp *http.Response) (*GetApiV4SnippetsPublicResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsPublicResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4SnippetsIdResponse(rsp *http.Response) (*DeleteApiV4SnippetsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4SnippetsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	}

	return response, nil
}
func ParseGetApiV4SnippetsIdResponse(rsp *http.Response) (*GetApiV4SnippetsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4SnippetsIdResponse(rsp *http.Response) (*PutApiV4SnippetsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4SnippetsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Author API_Entities_UserBasic model
			Author *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			FileName          *string    `json:"file_name,omitempty"`
			Files             *[]string  `json:"files,omitempty"`
			HttpUrlToRepo     *string    `json:"http_url_to_repo,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Imported          *bool      `json:"imported,omitempty"`
			ImportedFrom      *string    `json:"imported_from,omitempty"`
			ProjectId         *int32     `json:"project_id,omitempty"`
			RawUrl            *string    `json:"raw_url,omitempty"`
			RepositoryStorage *string    `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string    `json:"ssh_url_to_repo,omitempty"`
			Title             *string    `json:"title,omitempty"`
			UpdatedAt         *time.Time `json:"updated_at,omitempty"`
			Visibility        *string    `json:"visibility,omitempty"`
			WebUrl            *string    `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4SnippetsIdFilesRefFilePathRawResponse(rsp *http.Response) (*GetApiV4SnippetsIdFilesRefFilePathRawResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsIdFilesRefFilePathRawResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4SnippetsIdRawResponse(rsp *http.Response) (*GetApiV4SnippetsIdRawResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsIdRawResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4SnippetsIdUserAgentDetailResponse(rsp *http.Response) (*GetApiV4SnippetsIdUserAgentDetailResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4SnippetsIdUserAgentDetailResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AkismetSubmitted *bool   `json:"akismet_submitted,omitempty"`
			IpAddress        *string `json:"ip_address,omitempty"`
			UserAgent        *string `json:"user_agent,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
