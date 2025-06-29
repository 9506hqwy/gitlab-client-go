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
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type GetApiV4NamespacesParams struct {
	// Search Returns a list of namespaces the user is authorized to view based on the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// OwnedOnly In GitLab 14.2 and later, returns a list of owned namespaces only
	OwnedOnly *bool `form:"owned_only,omitempty" json:"owned_only,omitempty"`

	// TopLevelOnly Only include top level namespaces
	TopLevelOnly *bool `form:"top_level_only,omitempty" json:"top_level_only,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// RequestedHostedPlan Name of the hosted plan requested by the customer
	RequestedHostedPlan *string `form:"requested_hosted_plan,omitempty" json:"requested_hosted_plan,omitempty"`
}
type GetApiV4NamespacesStorageLimitExclusionsParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PutApiV4NamespacesIdJSONBody struct {
	// AdditionalPurchasedStorageEndsOn End of subscription of the additional purchased storage
	AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`

	// AdditionalPurchasedStorageSize Additional storage size for this namespace
	AdditionalPurchasedStorageSize *int32 `json:"additional_purchased_storage_size,omitempty"`

	// ExtraSharedRunnersMinutesLimit Extra compute minutes for this namespace
	ExtraSharedRunnersMinutesLimit *int32 `json:"extra_shared_runners_minutes_limit,omitempty"`
	GitlabSubscriptionAttributes   *struct {
		// AutoRenew Whether subscription will auto renew on end date
		AutoRenew *bool `json:"auto_renew,omitempty"`

		// EndDate End date of subscription
		EndDate *openapi_types.Date `json:"end_date,omitempty"`

		// MaxSeatsUsed Highest number of active users in the last month
		MaxSeatsUsed *int32 `json:"max_seats_used,omitempty"`

		// PlanCode Subscription tier code
		PlanCode *string `json:"plan_code,omitempty"`

		// Seats Number of seats in subscription
		Seats *int32 `json:"seats,omitempty"`

		// StartDate Start date of subscription
		StartDate *openapi_types.Date `json:"start_date,omitempty"`

		// Trial Whether the subscription is a trial
		Trial *bool `json:"trial,omitempty"`

		// TrialEndsOn End date of trial
		TrialEndsOn *openapi_types.Date `json:"trial_ends_on,omitempty"`

		// TrialExtensionType Whether subscription is an extended or reactivated trial
		TrialExtensionType *int32 `json:"trial_extension_type,omitempty"`

		// TrialStartsOn Start date of trial
		TrialStartsOn *openapi_types.Date `json:"trial_starts_on,omitempty"`
	} `json:"gitlab_subscription_attributes,omitempty"`

	// SharedRunnersMinutesLimit Compute minutes quota for this namespace
	SharedRunnersMinutesLimit *int32 `json:"shared_runners_minutes_limit,omitempty"`
}
type GetApiV4NamespacesIdExistsParams struct {
	// ParentId The ID of the parent namespace. If no ID is specified, only top-level namespaces are considered.
	ParentId *int32 `form:"parent_id,omitempty" json:"parent_id,omitempty"`
}
type PostApiV4NamespacesIdStorageLimitExclusionJSONBody struct {
	// Reason The reason the Namespace is being excluded
	Reason string `json:"reason"`
}
type PutApiV4NamespacesIdJSONRequestBody PutApiV4NamespacesIdJSONBody
type PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody PostApiV4NamespacesIdStorageLimitExclusionJSONBody
type GetApiV4NamespacesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
		AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
		AvatarUrl                        *string             `json:"avatar_url,omitempty"`
		BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
		EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
		ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
		FullPath                         *string             `json:"full_path,omitempty"`
		Id                               *int32              `json:"id,omitempty"`
		Kind                             *string             `json:"kind,omitempty"`
		MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
		MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
		MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
		Name                             *string             `json:"name,omitempty"`
		ParentId                         *int32              `json:"parent_id,omitempty"`
		Path                             *string             `json:"path,omitempty"`
		Plan                             *string             `json:"plan,omitempty"`
		ProjectsCount                    *int32              `json:"projects_count,omitempty"`
		RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
		SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
		SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
		Trial                            *bool               `json:"trial,omitempty"`
		TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
		WebUrl                           *string             `json:"web_url,omitempty"`
	}
}
type GetApiV4NamespacesStorageLimitExclusionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Id            *int32  `json:"id,omitempty"`
		NamespaceId   *int32  `json:"namespace_id,omitempty"`
		NamespaceName *string `json:"namespace_name,omitempty"`
		Reason        *string `json:"reason,omitempty"`
	}
}
type GetApiV4NamespacesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
		AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
		AvatarUrl                        *string             `json:"avatar_url,omitempty"`
		BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
		EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
		ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
		FullPath                         *string             `json:"full_path,omitempty"`
		Id                               *int32              `json:"id,omitempty"`
		Kind                             *string             `json:"kind,omitempty"`
		MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
		MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
		MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
		Name                             *string             `json:"name,omitempty"`
		ParentId                         *int32              `json:"parent_id,omitempty"`
		Path                             *string             `json:"path,omitempty"`
		Plan                             *string             `json:"plan,omitempty"`
		ProjectsCount                    *int32              `json:"projects_count,omitempty"`
		RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
		SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
		SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
		Trial                            *bool               `json:"trial,omitempty"`
		TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
		WebUrl                           *string             `json:"web_url,omitempty"`
	}
}
type PutApiV4NamespacesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
		AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
		AvatarUrl                        *string             `json:"avatar_url,omitempty"`
		BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
		EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
		ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
		FullPath                         *string             `json:"full_path,omitempty"`
		Id                               *int32              `json:"id,omitempty"`
		Kind                             *string             `json:"kind,omitempty"`
		MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
		MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
		MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
		Name                             *string             `json:"name,omitempty"`
		ParentId                         *int32              `json:"parent_id,omitempty"`
		Path                             *string             `json:"path,omitempty"`
		Plan                             *string             `json:"plan,omitempty"`
		ProjectsCount                    *int32              `json:"projects_count,omitempty"`
		RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
		SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
		SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
		Trial                            *bool               `json:"trial,omitempty"`
		TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
		WebUrl                           *string             `json:"web_url,omitempty"`
	}
}
type GetApiV4NamespacesIdExistsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Exists   *bool     `json:"exists,omitempty"`
		Suggests *[]string `json:"suggests,omitempty"`
	}
}
type GetApiV4NamespacesIdGitlabSubscriptionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Billing *struct {
			SubscriptionEndDate   *string `json:"subscription_end_date,omitempty"`
			SubscriptionStartDate *string `json:"subscription_start_date,omitempty"`
			TrialEndsOn           *string `json:"trial_ends_on,omitempty"`
		} `json:"billing,omitempty"`
		Plan *struct {
			AutoRenew     *string `json:"auto_renew,omitempty"`
			Code          *string `json:"code,omitempty"`
			ExcludeGuests *string `json:"exclude_guests,omitempty"`
			Name          *string `json:"name,omitempty"`
			Trial         *string `json:"trial,omitempty"`
			Upgradable    *string `json:"upgradable,omitempty"`
		} `json:"plan,omitempty"`
		Usage *struct {
			MaxSeatsUsed        *string `json:"max_seats_used,omitempty"`
			SeatsInSubscription *string `json:"seats_in_subscription,omitempty"`
			SeatsInUse          *string `json:"seats_in_use,omitempty"`
			SeatsOwed           *string `json:"seats_owed,omitempty"`
		} `json:"usage,omitempty"`
	}
}
type DeleteApiV4NamespacesIdStorageLimitExclusionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4NamespacesIdStorageLimitExclusionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Id            *int32  `json:"id,omitempty"`
		NamespaceId   *int32  `json:"namespace_id,omitempty"`
		NamespaceName *string `json:"namespace_name,omitempty"`
		Reason        *string `json:"reason,omitempty"`
	}
}

func (c *Client) GetApiV4Namespaces(ctx context.Context, params *GetApiV4NamespacesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4NamespacesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4NamespacesStorageLimitExclusions(ctx context.Context, params *GetApiV4NamespacesStorageLimitExclusionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4NamespacesStorageLimitExclusionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4NamespacesId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4NamespacesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4NamespacesIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4NamespacesIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4NamespacesId(ctx context.Context, id int32, body PutApiV4NamespacesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4NamespacesIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4NamespacesIdExists(ctx context.Context, id string, params *GetApiV4NamespacesIdExistsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4NamespacesIdExistsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4NamespacesIdGitlabSubscription(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4NamespacesIdGitlabSubscriptionRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4NamespacesIdStorageLimitExclusion(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4NamespacesIdStorageLimitExclusionRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4NamespacesIdStorageLimitExclusionWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4NamespacesIdStorageLimitExclusionRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4NamespacesIdStorageLimitExclusion(ctx context.Context, id int32, body PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4NamespacesIdStorageLimitExclusionRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4NamespacesRequest(server string, params *GetApiV4NamespacesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/namespaces")
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

		if params.OwnedOnly != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "owned_only", runtime.ParamLocationQuery, *params.OwnedOnly); err != nil {
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

		if params.TopLevelOnly != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "top_level_only", runtime.ParamLocationQuery, *params.TopLevelOnly); err != nil {
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

		if params.RequestedHostedPlan != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "requested_hosted_plan", runtime.ParamLocationQuery, *params.RequestedHostedPlan); err != nil {
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
func NewGetApiV4NamespacesStorageLimitExclusionsRequest(server string, params *GetApiV4NamespacesStorageLimitExclusionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/namespaces/storage/limit_exclusions")
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
func NewGetApiV4NamespacesIdRequest(server string, id string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s", pathParam0)
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
func NewPutApiV4NamespacesIdRequest(server string, id int32, body PutApiV4NamespacesIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4NamespacesIdRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4NamespacesIdRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s", pathParam0)
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
func NewGetApiV4NamespacesIdExistsRequest(server string, id string, params *GetApiV4NamespacesIdExistsParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s/exists", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ParentId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "parent_id", runtime.ParamLocationQuery, *params.ParentId); err != nil {
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
func NewGetApiV4NamespacesIdGitlabSubscriptionRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s/gitlab_subscription", pathParam0)
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
func NewDeleteApiV4NamespacesIdStorageLimitExclusionRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s/storage/limit_exclusion", pathParam0)
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
func NewPostApiV4NamespacesIdStorageLimitExclusionRequest(server string, id int32, body PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4NamespacesIdStorageLimitExclusionRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPostApiV4NamespacesIdStorageLimitExclusionRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/namespaces/%s/storage/limit_exclusion", pathParam0)
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
func (r GetApiV4NamespacesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4NamespacesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4NamespacesStorageLimitExclusionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4NamespacesStorageLimitExclusionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4NamespacesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4NamespacesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4NamespacesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4NamespacesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4NamespacesIdExistsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4NamespacesIdExistsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4NamespacesIdGitlabSubscriptionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4NamespacesIdGitlabSubscriptionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4NamespacesIdStorageLimitExclusionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4NamespacesIdStorageLimitExclusionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4NamespacesIdStorageLimitExclusionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4NamespacesIdStorageLimitExclusionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4NamespacesWithResponse(ctx context.Context, params *GetApiV4NamespacesParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesResponse, error) {
	rsp, err := c.GetApiV4Namespaces(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4NamespacesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4NamespacesStorageLimitExclusionsWithResponse(ctx context.Context, params *GetApiV4NamespacesStorageLimitExclusionsParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesStorageLimitExclusionsResponse, error) {
	rsp, err := c.GetApiV4NamespacesStorageLimitExclusions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4NamespacesStorageLimitExclusionsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4NamespacesIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdResponse, error) {
	rsp, err := c.GetApiV4NamespacesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4NamespacesIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4NamespacesIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4NamespacesIdResponse, error) {
	rsp, err := c.PutApiV4NamespacesIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4NamespacesIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4NamespacesIdWithResponse(ctx context.Context, id int32, body PutApiV4NamespacesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4NamespacesIdResponse, error) {
	rsp, err := c.PutApiV4NamespacesId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4NamespacesIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4NamespacesIdExistsWithResponse(ctx context.Context, id string, params *GetApiV4NamespacesIdExistsParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdExistsResponse, error) {
	rsp, err := c.GetApiV4NamespacesIdExists(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4NamespacesIdExistsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4NamespacesIdGitlabSubscriptionWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdGitlabSubscriptionResponse, error) {
	rsp, err := c.GetApiV4NamespacesIdGitlabSubscription(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4NamespacesIdGitlabSubscriptionResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4NamespacesIdStorageLimitExclusionWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4NamespacesIdStorageLimitExclusionResponse, error) {
	rsp, err := c.DeleteApiV4NamespacesIdStorageLimitExclusion(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4NamespacesIdStorageLimitExclusionResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4NamespacesIdStorageLimitExclusionWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4NamespacesIdStorageLimitExclusionResponse, error) {
	rsp, err := c.PostApiV4NamespacesIdStorageLimitExclusionWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4NamespacesIdStorageLimitExclusionResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4NamespacesIdStorageLimitExclusionWithResponse(ctx context.Context, id int32, body PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4NamespacesIdStorageLimitExclusionResponse, error) {
	rsp, err := c.PostApiV4NamespacesIdStorageLimitExclusion(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4NamespacesIdStorageLimitExclusionResponse(rsp)
}
func ParseGetApiV4NamespacesResponse(rsp *http.Response) (*GetApiV4NamespacesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4NamespacesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
			AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
			AvatarUrl                        *string             `json:"avatar_url,omitempty"`
			BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
			EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
			ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
			FullPath                         *string             `json:"full_path,omitempty"`
			Id                               *int32              `json:"id,omitempty"`
			Kind                             *string             `json:"kind,omitempty"`
			MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
			MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
			MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
			Name                             *string             `json:"name,omitempty"`
			ParentId                         *int32              `json:"parent_id,omitempty"`
			Path                             *string             `json:"path,omitempty"`
			Plan                             *string             `json:"plan,omitempty"`
			ProjectsCount                    *int32              `json:"projects_count,omitempty"`
			RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
			SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
			SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
			Trial                            *bool               `json:"trial,omitempty"`
			TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
			WebUrl                           *string             `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4NamespacesStorageLimitExclusionsResponse(rsp *http.Response) (*GetApiV4NamespacesStorageLimitExclusionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4NamespacesStorageLimitExclusionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Id            *int32  `json:"id,omitempty"`
			NamespaceId   *int32  `json:"namespace_id,omitempty"`
			NamespaceName *string `json:"namespace_name,omitempty"`
			Reason        *string `json:"reason,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4NamespacesIdResponse(rsp *http.Response) (*GetApiV4NamespacesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4NamespacesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
			AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
			AvatarUrl                        *string             `json:"avatar_url,omitempty"`
			BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
			EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
			ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
			FullPath                         *string             `json:"full_path,omitempty"`
			Id                               *int32              `json:"id,omitempty"`
			Kind                             *string             `json:"kind,omitempty"`
			MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
			MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
			MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
			Name                             *string             `json:"name,omitempty"`
			ParentId                         *int32              `json:"parent_id,omitempty"`
			Path                             *string             `json:"path,omitempty"`
			Plan                             *string             `json:"plan,omitempty"`
			ProjectsCount                    *int32              `json:"projects_count,omitempty"`
			RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
			SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
			SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
			Trial                            *bool               `json:"trial,omitempty"`
			TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
			WebUrl                           *string             `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4NamespacesIdResponse(rsp *http.Response) (*PutApiV4NamespacesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4NamespacesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AdditionalPurchasedStorageEndsOn *openapi_types.Date `json:"additional_purchased_storage_ends_on,omitempty"`
			AdditionalPurchasedStorageSize   *int32              `json:"additional_purchased_storage_size,omitempty"`
			AvatarUrl                        *string             `json:"avatar_url,omitempty"`
			BillableMembersCount             *int32              `json:"billable_members_count,omitempty"`
			EndDate                          *openapi_types.Date `json:"end_date,omitempty"`
			ExtraSharedRunnersMinutesLimit   *int32              `json:"extra_shared_runners_minutes_limit,omitempty"`
			FullPath                         *string             `json:"full_path,omitempty"`
			Id                               *int32              `json:"id,omitempty"`
			Kind                             *string             `json:"kind,omitempty"`
			MaxSeatsUsed                     *int32              `json:"max_seats_used,omitempty"`
			MaxSeatsUsedChangedAt            *openapi_types.Date `json:"max_seats_used_changed_at,omitempty"`
			MembersCountWithDescendants      *int32              `json:"members_count_with_descendants,omitempty"`
			Name                             *string             `json:"name,omitempty"`
			ParentId                         *int32              `json:"parent_id,omitempty"`
			Path                             *string             `json:"path,omitempty"`
			Plan                             *string             `json:"plan,omitempty"`
			ProjectsCount                    *int32              `json:"projects_count,omitempty"`
			RootRepositorySize               *int32              `json:"root_repository_size,omitempty"`
			SeatsInUse                       *int32              `json:"seats_in_use,omitempty"`
			SharedRunnersMinutesLimit        *int32              `json:"shared_runners_minutes_limit,omitempty"`
			Trial                            *bool               `json:"trial,omitempty"`
			TrialEndsOn                      *openapi_types.Date `json:"trial_ends_on,omitempty"`
			WebUrl                           *string             `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4NamespacesIdExistsResponse(rsp *http.Response) (*GetApiV4NamespacesIdExistsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4NamespacesIdExistsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Exists   *bool     `json:"exists,omitempty"`
			Suggests *[]string `json:"suggests,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4NamespacesIdGitlabSubscriptionResponse(rsp *http.Response) (*GetApiV4NamespacesIdGitlabSubscriptionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4NamespacesIdGitlabSubscriptionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Billing *struct {
				SubscriptionEndDate   *string `json:"subscription_end_date,omitempty"`
				SubscriptionStartDate *string `json:"subscription_start_date,omitempty"`
				TrialEndsOn           *string `json:"trial_ends_on,omitempty"`
			} `json:"billing,omitempty"`
			Plan *struct {
				AutoRenew     *string `json:"auto_renew,omitempty"`
				Code          *string `json:"code,omitempty"`
				ExcludeGuests *string `json:"exclude_guests,omitempty"`
				Name          *string `json:"name,omitempty"`
				Trial         *string `json:"trial,omitempty"`
				Upgradable    *string `json:"upgradable,omitempty"`
			} `json:"plan,omitempty"`
			Usage *struct {
				MaxSeatsUsed        *string `json:"max_seats_used,omitempty"`
				SeatsInSubscription *string `json:"seats_in_subscription,omitempty"`
				SeatsInUse          *string `json:"seats_in_use,omitempty"`
				SeatsOwed           *string `json:"seats_owed,omitempty"`
			} `json:"usage,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4NamespacesIdStorageLimitExclusionResponse(rsp *http.Response) (*DeleteApiV4NamespacesIdStorageLimitExclusionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4NamespacesIdStorageLimitExclusionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4NamespacesIdStorageLimitExclusionResponse(rsp *http.Response) (*PostApiV4NamespacesIdStorageLimitExclusionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4NamespacesIdStorageLimitExclusionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Id            *int32  `json:"id,omitempty"`
			NamespaceId   *int32  `json:"namespace_id,omitempty"`
			NamespaceName *string `json:"namespace_name,omitempty"`
			Reason        *string `json:"reason,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
