package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type GetApiV4EventsParams struct {
	// Scope Include all events across a user’s projects
	Scope *string `form:"scope,omitempty" json:"scope,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Action Event action to filter on
	Action *string `form:"action,omitempty" json:"action,omitempty"`

	// TargetType Event target type to filter on
	TargetType *string `form:"target_type,omitempty" json:"target_type,omitempty"`

	// Before Include only events created before this date
	Before *openapi_types.Date `form:"before,omitempty" json:"before,omitempty"`

	// After Include only events created after this date
	After *openapi_types.Date `form:"after,omitempty" json:"after,omitempty"`

	// Sort Return events sorted in ascending and descending order
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`
}
type GetApiV4EventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		ActionName *string `json:"action_name,omitempty"`

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
		AuthorId       *int32  `json:"author_id,omitempty"`
		AuthorUsername *string `json:"author_username,omitempty"`
		CreatedAt      *string `json:"created_at,omitempty"`
		Id             *int32  `json:"id,omitempty"`
		Imported       *bool   `json:"imported,omitempty"`
		ImportedFrom   *string `json:"imported_from,omitempty"`
		Note           *struct {
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
			Body            *string `json:"body,omitempty"`
			CommandsChanges *string `json:"commands_changes,omitempty"`
			CommitId        *string `json:"commit_id,omitempty"`
			Confidential    *string `json:"confidential,omitempty"`
			CreatedAt       *string `json:"created_at,omitempty"`
			Id              *string `json:"id,omitempty"`
			Imported        *string `json:"imported,omitempty"`
			ImportedFrom    *string `json:"imported_from,omitempty"`
			Internal        *string `json:"internal,omitempty"`
			NoteableId      *string `json:"noteable_id,omitempty"`
			NoteableIid     *string `json:"noteable_iid,omitempty"`
			NoteableType    *string `json:"noteable_type,omitempty"`
			Position        *string `json:"position,omitempty"`
			ProjectId       *string `json:"project_id,omitempty"`
			Resolvable      *string `json:"resolvable,omitempty"`
			Resolved        *string `json:"resolved,omitempty"`
			ResolvedAt      *string `json:"resolved_at,omitempty"`

			// ResolvedBy API_Entities_UserBasic model
			ResolvedBy *struct {
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
			} `json:"resolved_by,omitempty"`

			// Suggestions API_Entities_Suggestion model
			Suggestions *struct {
				Appliable   *string `json:"appliable,omitempty"`
				Applied     *string `json:"applied,omitempty"`
				FromContent *string `json:"from_content,omitempty"`
				FromLine    *string `json:"from_line,omitempty"`
				Id          *string `json:"id,omitempty"`
				ToContent   *string `json:"to_content,omitempty"`
				ToLine      *string `json:"to_line,omitempty"`
			} `json:"suggestions,omitempty"`
			System    *string `json:"system,omitempty"`
			Type      *string `json:"type,omitempty"`
			UpdatedAt *string `json:"updated_at,omitempty"`
		} `json:"note,omitempty"`
		ProjectId *int32 `json:"project_id,omitempty"`
		PushData  *struct {
			Action      *string `json:"action,omitempty"`
			CommitCount *int32  `json:"commit_count,omitempty"`
			CommitFrom  *string `json:"commit_from,omitempty"`
			CommitTitle *string `json:"commit_title,omitempty"`
			CommitTo    *string `json:"commit_to,omitempty"`
			Ref         *string `json:"ref,omitempty"`
			RefCount    *int32  `json:"ref_count,omitempty"`
			RefType     *string `json:"ref_type,omitempty"`
		} `json:"push_data,omitempty"`
		TargetId    *int32  `json:"target_id,omitempty"`
		TargetIid   *int32  `json:"target_iid,omitempty"`
		TargetTitle *string `json:"target_title,omitempty"`
		TargetType  *string `json:"target_type,omitempty"`

		// WikiPage API_Entities_WikiPageBasic model
		WikiPage *struct {
			Format         *string `json:"format,omitempty"`
			Slug           *string `json:"slug,omitempty"`
			Title          *string `json:"title,omitempty"`
			WikiPageMetaId *int32  `json:"wiki_page_meta_id,omitempty"`
		} `json:"wiki_page,omitempty"`
	}
}

func (c *Client) GetApiV4Events(ctx context.Context, params *GetApiV4EventsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4EventsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4EventsRequest(server string, params *GetApiV4EventsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/events")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Scope != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "scope", runtime.ParamLocationQuery, *params.Scope); err != nil {
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

		if params.Action != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "action", runtime.ParamLocationQuery, *params.Action); err != nil {
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

		if params.TargetType != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "target_type", runtime.ParamLocationQuery, *params.TargetType); err != nil {
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

		if params.Before != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "before", runtime.ParamLocationQuery, *params.Before); err != nil {
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

		if params.After != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "after", runtime.ParamLocationQuery, *params.After); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (r GetApiV4EventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4EventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4EventsWithResponse(ctx context.Context, params *GetApiV4EventsParams, reqEditors ...RequestEditorFn) (*GetApiV4EventsResponse, error) {
	rsp, err := c.GetApiV4Events(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4EventsResponse(rsp)
}
func ParseGetApiV4EventsResponse(rsp *http.Response) (*GetApiV4EventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4EventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			ActionName *string `json:"action_name,omitempty"`

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
			AuthorId       *int32  `json:"author_id,omitempty"`
			AuthorUsername *string `json:"author_username,omitempty"`
			CreatedAt      *string `json:"created_at,omitempty"`
			Id             *int32  `json:"id,omitempty"`
			Imported       *bool   `json:"imported,omitempty"`
			ImportedFrom   *string `json:"imported_from,omitempty"`
			Note           *struct {
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
				Body            *string `json:"body,omitempty"`
				CommandsChanges *string `json:"commands_changes,omitempty"`
				CommitId        *string `json:"commit_id,omitempty"`
				Confidential    *string `json:"confidential,omitempty"`
				CreatedAt       *string `json:"created_at,omitempty"`
				Id              *string `json:"id,omitempty"`
				Imported        *string `json:"imported,omitempty"`
				ImportedFrom    *string `json:"imported_from,omitempty"`
				Internal        *string `json:"internal,omitempty"`
				NoteableId      *string `json:"noteable_id,omitempty"`
				NoteableIid     *string `json:"noteable_iid,omitempty"`
				NoteableType    *string `json:"noteable_type,omitempty"`
				Position        *string `json:"position,omitempty"`
				ProjectId       *string `json:"project_id,omitempty"`
				Resolvable      *string `json:"resolvable,omitempty"`
				Resolved        *string `json:"resolved,omitempty"`
				ResolvedAt      *string `json:"resolved_at,omitempty"`

				// ResolvedBy API_Entities_UserBasic model
				ResolvedBy *struct {
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
				} `json:"resolved_by,omitempty"`

				// Suggestions API_Entities_Suggestion model
				Suggestions *struct {
					Appliable   *string `json:"appliable,omitempty"`
					Applied     *string `json:"applied,omitempty"`
					FromContent *string `json:"from_content,omitempty"`
					FromLine    *string `json:"from_line,omitempty"`
					Id          *string `json:"id,omitempty"`
					ToContent   *string `json:"to_content,omitempty"`
					ToLine      *string `json:"to_line,omitempty"`
				} `json:"suggestions,omitempty"`
				System    *string `json:"system,omitempty"`
				Type      *string `json:"type,omitempty"`
				UpdatedAt *string `json:"updated_at,omitempty"`
			} `json:"note,omitempty"`
			ProjectId *int32 `json:"project_id,omitempty"`
			PushData  *struct {
				Action      *string `json:"action,omitempty"`
				CommitCount *int32  `json:"commit_count,omitempty"`
				CommitFrom  *string `json:"commit_from,omitempty"`
				CommitTitle *string `json:"commit_title,omitempty"`
				CommitTo    *string `json:"commit_to,omitempty"`
				Ref         *string `json:"ref,omitempty"`
				RefCount    *int32  `json:"ref_count,omitempty"`
				RefType     *string `json:"ref_type,omitempty"`
			} `json:"push_data,omitempty"`
			TargetId    *int32  `json:"target_id,omitempty"`
			TargetIid   *int32  `json:"target_iid,omitempty"`
			TargetTitle *string `json:"target_title,omitempty"`
			TargetType  *string `json:"target_type,omitempty"`

			// WikiPage API_Entities_WikiPageBasic model
			WikiPage *struct {
				Format         *string `json:"format,omitempty"`
				Slug           *string `json:"slug,omitempty"`
				Title          *string `json:"title,omitempty"`
				WikiPageMetaId *int32  `json:"wiki_page_meta_id,omitempty"`
			} `json:"wiki_page,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
