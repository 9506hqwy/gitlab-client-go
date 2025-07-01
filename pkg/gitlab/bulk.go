package gitlab

import (
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

type GetApiV4BulkImportsParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Sort Return GitLab Migrations sorted in created by `asc` or `desc` order.
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`

	// Status Return GitLab Migrations with specified status
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}
type PostApiV4BulkImportsFormdataBody struct {
	// ConfigurationAccessToken Access token to the source GitLab instance
	ConfigurationAccessToken string `form:"configuration[access_token]" json:"configuration[access_token]"`

	// ConfigurationUrl Source GitLab instance URL
	ConfigurationUrl string `form:"configuration[url]" json:"configuration[url]"`

	// EntitiesDestinationName Deprecated: Use :destination_slug instead. Destination slug for the entity
	EntitiesDestinationName *[]string `form:"entities[destination_name],omitempty" json:"entities[destination_name],omitempty"`

	// EntitiesDestinationNamespace Destination namespace for the entity
	EntitiesDestinationNamespace []string `form:"entities[destination_namespace]" json:"entities[destination_namespace]"`

	// EntitiesDestinationSlug Destination slug for the entity
	EntitiesDestinationSlug *[]string `form:"entities[destination_slug],omitempty" json:"entities[destination_slug],omitempty"`

	// EntitiesMigrateMemberships The option to migrate memberships or not
	EntitiesMigrateMemberships *[]bool `form:"entities[migrate_memberships],omitempty" json:"entities[migrate_memberships],omitempty"`

	// EntitiesMigrateProjects Indicates group migration should include nested projects
	EntitiesMigrateProjects *[]bool `form:"entities[migrate_projects],omitempty" json:"entities[migrate_projects],omitempty"`

	// EntitiesSourceFullPath Relative path of the source entity to import
	EntitiesSourceFullPath []string `form:"entities[source_full_path]" json:"entities[source_full_path]"`

	// EntitiesSourceType Source entity type
	EntitiesSourceType []string `form:"entities[source_type]" json:"entities[source_type]"`
}
type GetApiV4BulkImportsEntitiesParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Sort Return GitLab Migrations sorted in created by `asc` or `desc` order.
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`

	// Status Return all GitLab Migrations' entities with specified status
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}
type GetApiV4BulkImportsImportIdEntitiesParams struct {
	// Status Return import entities with specified status
	Status *string `form:"status,omitempty" json:"status,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4BulkImportsFormdataRequestBody PostApiV4BulkImportsFormdataBody
type GetApiV4BulkImportsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		HasFailures *bool      `json:"has_failures,omitempty"`
		Id          *int32     `json:"id,omitempty"`
		SourceType  *string    `json:"source_type,omitempty"`
		SourceUrl   *string    `json:"source_url,omitempty"`
		Status      *string    `json:"status,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	}
}
type PostApiV4BulkImportsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		HasFailures *bool      `json:"has_failures,omitempty"`
		Id          *int32     `json:"id,omitempty"`
		SourceType  *string    `json:"source_type,omitempty"`
		SourceUrl   *string    `json:"source_url,omitempty"`
		Status      *string    `json:"status,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	}
}
type GetApiV4BulkImportsEntitiesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
		CreatedAt            *time.Time `json:"created_at,omitempty"`
		DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
		DestinationName      *string    `json:"destination_name,omitempty"`
		DestinationNamespace *string    `json:"destination_namespace,omitempty"`
		DestinationSlug      *string    `json:"destination_slug,omitempty"`
		EntityType           *string    `json:"entity_type,omitempty"`
		Failures             *[]struct {
			CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
			ExceptionClass     *string `json:"exception_class,omitempty"`
			ExceptionMessage   *string `json:"exception_message,omitempty"`
			Relation           *string `json:"relation,omitempty"`
			SourceTitle        *string `json:"source_title,omitempty"`
			SourceUrl          *string `json:"source_url,omitempty"`
		} `json:"failures,omitempty"`
		HasFailures        *bool                   `json:"has_failures,omitempty"`
		Id                 *int32                  `json:"id,omitempty"`
		MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
		MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
		NamespaceId        *int32                  `json:"namespace_id,omitempty"`
		ParentId           *int32                  `json:"parent_id,omitempty"`
		ProjectId          *int32                  `json:"project_id,omitempty"`
		SourceFullPath     *string                 `json:"source_full_path,omitempty"`
		Stats              *map[string]interface{} `json:"stats,omitempty"`
		Status             *string                 `json:"status,omitempty"`
		UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
	}
}
type GetApiV4BulkImportsImportIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		HasFailures *bool      `json:"has_failures,omitempty"`
		Id          *int32     `json:"id,omitempty"`
		SourceType  *string    `json:"source_type,omitempty"`
		SourceUrl   *string    `json:"source_url,omitempty"`
		Status      *string    `json:"status,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	}
}
type PostApiV4BulkImportsImportIdCancelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		HasFailures *bool      `json:"has_failures,omitempty"`
		Id          *int32     `json:"id,omitempty"`
		SourceType  *string    `json:"source_type,omitempty"`
		SourceUrl   *string    `json:"source_url,omitempty"`
		Status      *string    `json:"status,omitempty"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	}
}
type GetApiV4BulkImportsImportIdEntitiesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
		CreatedAt            *time.Time `json:"created_at,omitempty"`
		DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
		DestinationName      *string    `json:"destination_name,omitempty"`
		DestinationNamespace *string    `json:"destination_namespace,omitempty"`
		DestinationSlug      *string    `json:"destination_slug,omitempty"`
		EntityType           *string    `json:"entity_type,omitempty"`
		Failures             *[]struct {
			CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
			ExceptionClass     *string `json:"exception_class,omitempty"`
			ExceptionMessage   *string `json:"exception_message,omitempty"`
			Relation           *string `json:"relation,omitempty"`
			SourceTitle        *string `json:"source_title,omitempty"`
			SourceUrl          *string `json:"source_url,omitempty"`
		} `json:"failures,omitempty"`
		HasFailures        *bool                   `json:"has_failures,omitempty"`
		Id                 *int32                  `json:"id,omitempty"`
		MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
		MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
		NamespaceId        *int32                  `json:"namespace_id,omitempty"`
		ParentId           *int32                  `json:"parent_id,omitempty"`
		ProjectId          *int32                  `json:"project_id,omitempty"`
		SourceFullPath     *string                 `json:"source_full_path,omitempty"`
		Stats              *map[string]interface{} `json:"stats,omitempty"`
		Status             *string                 `json:"status,omitempty"`
		UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
	}
}
type GetApiV4BulkImportsImportIdEntitiesEntityIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
		CreatedAt            *time.Time `json:"created_at,omitempty"`
		DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
		DestinationName      *string    `json:"destination_name,omitempty"`
		DestinationNamespace *string    `json:"destination_namespace,omitempty"`
		DestinationSlug      *string    `json:"destination_slug,omitempty"`
		EntityType           *string    `json:"entity_type,omitempty"`
		Failures             *[]struct {
			CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
			ExceptionClass     *string `json:"exception_class,omitempty"`
			ExceptionMessage   *string `json:"exception_message,omitempty"`
			Relation           *string `json:"relation,omitempty"`
			SourceTitle        *string `json:"source_title,omitempty"`
			SourceUrl          *string `json:"source_url,omitempty"`
		} `json:"failures,omitempty"`
		HasFailures        *bool                   `json:"has_failures,omitempty"`
		Id                 *int32                  `json:"id,omitempty"`
		MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
		MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
		NamespaceId        *int32                  `json:"namespace_id,omitempty"`
		ParentId           *int32                  `json:"parent_id,omitempty"`
		ProjectId          *int32                  `json:"project_id,omitempty"`
		SourceFullPath     *string                 `json:"source_full_path,omitempty"`
		Stats              *map[string]interface{} `json:"stats,omitempty"`
		Status             *string                 `json:"status,omitempty"`
		UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
	}
}
type GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
		ExceptionClass     *string `json:"exception_class,omitempty"`
		ExceptionMessage   *string `json:"exception_message,omitempty"`
		Relation           *string `json:"relation,omitempty"`
		SourceTitle        *string `json:"source_title,omitempty"`
		SourceUrl          *string `json:"source_url,omitempty"`
	}
}

func (c *Client) GetApiV4BulkImports(ctx context.Context, params *GetApiV4BulkImportsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4BulkImportsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4BulkImportsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4BulkImportsWithFormdataBody(ctx context.Context, body PostApiV4BulkImportsFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4BulkImportsRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BulkImportsEntities(ctx context.Context, params *GetApiV4BulkImportsEntitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsEntitiesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BulkImportsImportId(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsImportIdRequest(c.Server, importId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4BulkImportsImportIdCancel(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4BulkImportsImportIdCancelRequest(c.Server, importId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BulkImportsImportIdEntities(ctx context.Context, importId int32, params *GetApiV4BulkImportsImportIdEntitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsImportIdEntitiesRequest(c.Server, importId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BulkImportsImportIdEntitiesEntityId(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsImportIdEntitiesEntityIdRequest(c.Server, importId, entityId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4BulkImportsImportIdEntitiesEntityIdFailuresRequest(c.Server, importId, entityId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4BulkImportsRequest(server string, params *GetApiV4BulkImportsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports")
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

		if params.Sort != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
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

		if params.Status != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
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
func NewPostApiV4BulkImportsRequestWithFormdataBody(server string, body PostApiV4BulkImportsFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewPostApiV4BulkImportsRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}
func NewPostApiV4BulkImportsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports")
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
func NewGetApiV4BulkImportsEntitiesRequest(server string, params *GetApiV4BulkImportsEntitiesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/entities")
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

		if params.Sort != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, *params.Sort); err != nil {
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

		if params.Status != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
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
func NewGetApiV4BulkImportsImportIdRequest(server string, importId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "import_id", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/%s", pathParam0)
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
func NewPostApiV4BulkImportsImportIdCancelRequest(server string, importId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "import_id", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/%s/cancel", pathParam0)
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
func NewGetApiV4BulkImportsImportIdEntitiesRequest(server string, importId int32, params *GetApiV4BulkImportsImportIdEntitiesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "import_id", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/%s/entities", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Status != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
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
func NewGetApiV4BulkImportsImportIdEntitiesEntityIdRequest(server string, importId int32, entityId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "import_id", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "entity_id", runtime.ParamLocationPath, entityId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/%s/entities/%s", pathParam0, pathParam1)
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
func NewGetApiV4BulkImportsImportIdEntitiesEntityIdFailuresRequest(server string, importId int32, entityId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "import_id", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "entity_id", runtime.ParamLocationPath, entityId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/bulk_imports/%s/entities/%s/failures", pathParam0, pathParam1)
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
func (r GetApiV4BulkImportsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4BulkImportsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4BulkImportsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BulkImportsEntitiesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsEntitiesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BulkImportsImportIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsImportIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4BulkImportsImportIdCancelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4BulkImportsImportIdCancelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BulkImportsImportIdEntitiesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsImportIdEntitiesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BulkImportsImportIdEntitiesEntityIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsImportIdEntitiesEntityIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4BulkImportsWithResponse(ctx context.Context, params *GetApiV4BulkImportsParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsResponse, error) {
	rsp, err := c.GetApiV4BulkImports(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4BulkImportsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsResponse, error) {
	rsp, err := c.PostApiV4BulkImportsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4BulkImportsResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4BulkImportsWithFormdataBodyWithResponse(ctx context.Context, body PostApiV4BulkImportsFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsResponse, error) {
	rsp, err := c.PostApiV4BulkImportsWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4BulkImportsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BulkImportsEntitiesWithResponse(ctx context.Context, params *GetApiV4BulkImportsEntitiesParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsEntitiesResponse, error) {
	rsp, err := c.GetApiV4BulkImportsEntities(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsEntitiesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BulkImportsImportIdWithResponse(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdResponse, error) {
	rsp, err := c.GetApiV4BulkImportsImportId(ctx, importId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsImportIdResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4BulkImportsImportIdCancelWithResponse(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsImportIdCancelResponse, error) {
	rsp, err := c.PostApiV4BulkImportsImportIdCancel(ctx, importId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4BulkImportsImportIdCancelResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BulkImportsImportIdEntitiesWithResponse(ctx context.Context, importId int32, params *GetApiV4BulkImportsImportIdEntitiesParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesResponse, error) {
	rsp, err := c.GetApiV4BulkImportsImportIdEntities(ctx, importId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsImportIdEntitiesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BulkImportsImportIdEntitiesEntityIdWithResponse(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesEntityIdResponse, error) {
	rsp, err := c.GetApiV4BulkImportsImportIdEntitiesEntityId(ctx, importId, entityId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsImportIdEntitiesEntityIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresWithResponse(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse, error) {
	rsp, err := c.GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx, importId, entityId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse(rsp)
}
func ParseGetApiV4BulkImportsResponse(rsp *http.Response) (*GetApiV4BulkImportsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			CreatedAt   *time.Time `json:"created_at,omitempty"`
			HasFailures *bool      `json:"has_failures,omitempty"`
			Id          *int32     `json:"id,omitempty"`
			SourceType  *string    `json:"source_type,omitempty"`
			SourceUrl   *string    `json:"source_url,omitempty"`
			Status      *string    `json:"status,omitempty"`
			UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4BulkImportsResponse(rsp *http.Response) (*PostApiV4BulkImportsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4BulkImportsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CreatedAt   *time.Time `json:"created_at,omitempty"`
			HasFailures *bool      `json:"has_failures,omitempty"`
			Id          *int32     `json:"id,omitempty"`
			SourceType  *string    `json:"source_type,omitempty"`
			SourceUrl   *string    `json:"source_url,omitempty"`
			Status      *string    `json:"status,omitempty"`
			UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BulkImportsEntitiesResponse(rsp *http.Response) (*GetApiV4BulkImportsEntitiesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsEntitiesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
			CreatedAt            *time.Time `json:"created_at,omitempty"`
			DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
			DestinationName      *string    `json:"destination_name,omitempty"`
			DestinationNamespace *string    `json:"destination_namespace,omitempty"`
			DestinationSlug      *string    `json:"destination_slug,omitempty"`
			EntityType           *string    `json:"entity_type,omitempty"`
			Failures             *[]struct {
				CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
				ExceptionClass     *string `json:"exception_class,omitempty"`
				ExceptionMessage   *string `json:"exception_message,omitempty"`
				Relation           *string `json:"relation,omitempty"`
				SourceTitle        *string `json:"source_title,omitempty"`
				SourceUrl          *string `json:"source_url,omitempty"`
			} `json:"failures,omitempty"`
			HasFailures        *bool                   `json:"has_failures,omitempty"`
			Id                 *int32                  `json:"id,omitempty"`
			MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
			MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
			NamespaceId        *int32                  `json:"namespace_id,omitempty"`
			ParentId           *int32                  `json:"parent_id,omitempty"`
			ProjectId          *int32                  `json:"project_id,omitempty"`
			SourceFullPath     *string                 `json:"source_full_path,omitempty"`
			Stats              *map[string]interface{} `json:"stats,omitempty"`
			Status             *string                 `json:"status,omitempty"`
			UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BulkImportsImportIdResponse(rsp *http.Response) (*GetApiV4BulkImportsImportIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsImportIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CreatedAt   *time.Time `json:"created_at,omitempty"`
			HasFailures *bool      `json:"has_failures,omitempty"`
			Id          *int32     `json:"id,omitempty"`
			SourceType  *string    `json:"source_type,omitempty"`
			SourceUrl   *string    `json:"source_url,omitempty"`
			Status      *string    `json:"status,omitempty"`
			UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4BulkImportsImportIdCancelResponse(rsp *http.Response) (*PostApiV4BulkImportsImportIdCancelResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4BulkImportsImportIdCancelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CreatedAt   *time.Time `json:"created_at,omitempty"`
			HasFailures *bool      `json:"has_failures,omitempty"`
			Id          *int32     `json:"id,omitempty"`
			SourceType  *string    `json:"source_type,omitempty"`
			SourceUrl   *string    `json:"source_url,omitempty"`
			Status      *string    `json:"status,omitempty"`
			UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BulkImportsImportIdEntitiesResponse(rsp *http.Response) (*GetApiV4BulkImportsImportIdEntitiesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsImportIdEntitiesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
			CreatedAt            *time.Time `json:"created_at,omitempty"`
			DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
			DestinationName      *string    `json:"destination_name,omitempty"`
			DestinationNamespace *string    `json:"destination_namespace,omitempty"`
			DestinationSlug      *string    `json:"destination_slug,omitempty"`
			EntityType           *string    `json:"entity_type,omitempty"`
			Failures             *[]struct {
				CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
				ExceptionClass     *string `json:"exception_class,omitempty"`
				ExceptionMessage   *string `json:"exception_message,omitempty"`
				Relation           *string `json:"relation,omitempty"`
				SourceTitle        *string `json:"source_title,omitempty"`
				SourceUrl          *string `json:"source_url,omitempty"`
			} `json:"failures,omitempty"`
			HasFailures        *bool                   `json:"has_failures,omitempty"`
			Id                 *int32                  `json:"id,omitempty"`
			MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
			MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
			NamespaceId        *int32                  `json:"namespace_id,omitempty"`
			ParentId           *int32                  `json:"parent_id,omitempty"`
			ProjectId          *int32                  `json:"project_id,omitempty"`
			SourceFullPath     *string                 `json:"source_full_path,omitempty"`
			Stats              *map[string]interface{} `json:"stats,omitempty"`
			Status             *string                 `json:"status,omitempty"`
			UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BulkImportsImportIdEntitiesEntityIdResponse(rsp *http.Response) (*GetApiV4BulkImportsImportIdEntitiesEntityIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsImportIdEntitiesEntityIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			BulkImportId         *int32     `json:"bulk_import_id,omitempty"`
			CreatedAt            *time.Time `json:"created_at,omitempty"`
			DestinationFullPath  *string    `json:"destination_full_path,omitempty"`
			DestinationName      *string    `json:"destination_name,omitempty"`
			DestinationNamespace *string    `json:"destination_namespace,omitempty"`
			DestinationSlug      *string    `json:"destination_slug,omitempty"`
			EntityType           *string    `json:"entity_type,omitempty"`
			Failures             *[]struct {
				CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
				ExceptionClass     *string `json:"exception_class,omitempty"`
				ExceptionMessage   *string `json:"exception_message,omitempty"`
				Relation           *string `json:"relation,omitempty"`
				SourceTitle        *string `json:"source_title,omitempty"`
				SourceUrl          *string `json:"source_url,omitempty"`
			} `json:"failures,omitempty"`
			HasFailures        *bool                   `json:"has_failures,omitempty"`
			Id                 *int32                  `json:"id,omitempty"`
			MigrateMemberships *bool                   `json:"migrate_memberships,omitempty"`
			MigrateProjects    *bool                   `json:"migrate_projects,omitempty"`
			NamespaceId        *int32                  `json:"namespace_id,omitempty"`
			ParentId           *int32                  `json:"parent_id,omitempty"`
			ProjectId          *int32                  `json:"project_id,omitempty"`
			SourceFullPath     *string                 `json:"source_full_path,omitempty"`
			Stats              *map[string]interface{} `json:"stats,omitempty"`
			Status             *string                 `json:"status,omitempty"`
			UpdatedAt          *time.Time              `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse(rsp *http.Response) (*GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CorrelationIdValue *string `json:"correlation_id_value,omitempty"`
			ExceptionClass     *string `json:"exception_class,omitempty"`
			ExceptionMessage   *string `json:"exception_message,omitempty"`
			Relation           *string `json:"relation,omitempty"`
			SourceTitle        *string `json:"source_title,omitempty"`
			SourceUrl          *string `json:"source_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
