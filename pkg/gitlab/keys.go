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

type GetApiV4KeysParams struct {
	// Fingerprint The fingerprint of an SSH key
	Fingerprint string `form:"fingerprint" json:"fingerprint"`
}
type GetApiV4KeysResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AvatarPath       *string    `json:"avatar_path,omitempty"`
		AvatarUrl        *string    `json:"avatar_url,omitempty"`
		Bio              *string    `json:"bio,omitempty"`
		Bot              *string    `json:"bot,omitempty"`
		CanCreateGroup   *bool      `json:"can_create_group,omitempty"`
		CanCreateProject *bool      `json:"can_create_project,omitempty"`
		ColorSchemeId    *int32     `json:"color_scheme_id,omitempty"`
		CommitEmail      *string    `json:"commit_email,omitempty"`
		ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
		CreatedAt        *string    `json:"created_at,omitempty"`

		// CreatedBy API_Entities_UserBasic model
		CreatedBy *struct {
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
		} `json:"created_by,omitempty"`
		CurrentSignInAt  *time.Time `json:"current_sign_in_at,omitempty"`
		CustomAttributes *[]struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"custom_attributes,omitempty"`
		Discord                        *string `json:"discord,omitempty"`
		Email                          *string `json:"email,omitempty"`
		EmailResetOfferedAt            *string `json:"email_reset_offered_at,omitempty"`
		EnterpriseGroupAssociatedAt    *string `json:"enterprise_group_associated_at,omitempty"`
		EnterpriseGroupId              *string `json:"enterprise_group_id,omitempty"`
		External                       *string `json:"external,omitempty"`
		ExtraSharedRunnersMinutesLimit *string `json:"extra_shared_runners_minutes_limit,omitempty"`
		Followers                      *string `json:"followers,omitempty"`
		Following                      *string `json:"following,omitempty"`
		Github                         *string `json:"github,omitempty"`
		Id                             *int32  `json:"id,omitempty"`
		Identities                     *struct {
			ExternUid      *string `json:"extern_uid,omitempty"`
			Provider       *string `json:"provider,omitempty"`
			SamlProviderId *string `json:"saml_provider_id,omitempty"`
		} `json:"identities,omitempty"`
		IsAdmin              *string    `json:"is_admin,omitempty"`
		IsAuditor            *string    `json:"is_auditor,omitempty"`
		IsFollowed           *string    `json:"is_followed,omitempty"`
		JobTitle             *string    `json:"job_title,omitempty"`
		LastActivityOn       *time.Time `json:"last_activity_on,omitempty"`
		LastSignInAt         *time.Time `json:"last_sign_in_at,omitempty"`
		Linkedin             *string    `json:"linkedin,omitempty"`
		LocalTime            *string    `json:"local_time,omitempty"`
		Location             *string    `json:"location,omitempty"`
		Locked               *bool      `json:"locked,omitempty"`
		Name                 *string    `json:"name,omitempty"`
		NamespaceId          *string    `json:"namespace_id,omitempty"`
		Note                 *string    `json:"note,omitempty"`
		Organization         *string    `json:"organization,omitempty"`
		PrivateProfile       *bool      `json:"private_profile,omitempty"`
		ProjectsLimit        *int32     `json:"projects_limit,omitempty"`
		Pronouns             *string    `json:"pronouns,omitempty"`
		ProvisionedByGroupId *string    `json:"provisioned_by_group_id,omitempty"`
		PublicEmail          *string    `json:"public_email,omitempty"`
		ScimIdentities       *struct {
			Active    *string `json:"active,omitempty"`
			ExternUid *string `json:"extern_uid,omitempty"`
			GroupId   *string `json:"group_id,omitempty"`
		} `json:"scim_identities,omitempty"`
		SharedRunnersMinutesLimit *string `json:"shared_runners_minutes_limit,omitempty"`
		Skype                     *string `json:"skype,omitempty"`
		State                     *string `json:"state,omitempty"`
		ThemeId                   *int32  `json:"theme_id,omitempty"`
		Twitter                   *string `json:"twitter,omitempty"`
		TwoFactorEnabled          *bool   `json:"two_factor_enabled,omitempty"`
		Username                  *string `json:"username,omitempty"`
		UsingLicenseSeat          *string `json:"using_license_seat,omitempty"`
		WebUrl                    *string `json:"web_url,omitempty"`
		WebsiteUrl                *string `json:"website_url,omitempty"`
		WorkInformation           *string `json:"work_information,omitempty"`
	}
}
type GetApiV4KeysIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		ExpiresAt  *time.Time `json:"expires_at,omitempty"`
		Id         *int32     `json:"id,omitempty"`
		Key        *string    `json:"key,omitempty"`
		LastUsedAt *time.Time `json:"last_used_at,omitempty"`
		Title      *string    `json:"title,omitempty"`
		UsageType  *string    `json:"usage_type,omitempty"`

		// User API_Entities_UserPublic model
		User *struct {
			AvatarPath       *string    `json:"avatar_path,omitempty"`
			AvatarUrl        *string    `json:"avatar_url,omitempty"`
			Bio              *string    `json:"bio,omitempty"`
			Bot              *string    `json:"bot,omitempty"`
			CanCreateGroup   *bool      `json:"can_create_group,omitempty"`
			CanCreateProject *bool      `json:"can_create_project,omitempty"`
			ColorSchemeId    *int32     `json:"color_scheme_id,omitempty"`
			CommitEmail      *string    `json:"commit_email,omitempty"`
			ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
			CreatedAt        *string    `json:"created_at,omitempty"`
			CurrentSignInAt  *time.Time `json:"current_sign_in_at,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Discord                        *string `json:"discord,omitempty"`
			Email                          *string `json:"email,omitempty"`
			External                       *string `json:"external,omitempty"`
			ExtraSharedRunnersMinutesLimit *string `json:"extra_shared_runners_minutes_limit,omitempty"`
			Followers                      *string `json:"followers,omitempty"`
			Following                      *string `json:"following,omitempty"`
			Github                         *string `json:"github,omitempty"`
			Id                             *int32  `json:"id,omitempty"`
			Identities                     *struct {
				ExternUid      *string `json:"extern_uid,omitempty"`
				Provider       *string `json:"provider,omitempty"`
				SamlProviderId *string `json:"saml_provider_id,omitempty"`
			} `json:"identities,omitempty"`
			IsFollowed     *string    `json:"is_followed,omitempty"`
			JobTitle       *string    `json:"job_title,omitempty"`
			LastActivityOn *time.Time `json:"last_activity_on,omitempty"`
			LastSignInAt   *time.Time `json:"last_sign_in_at,omitempty"`
			Linkedin       *string    `json:"linkedin,omitempty"`
			LocalTime      *string    `json:"local_time,omitempty"`
			Location       *string    `json:"location,omitempty"`
			Locked         *bool      `json:"locked,omitempty"`
			Name           *string    `json:"name,omitempty"`
			Organization   *string    `json:"organization,omitempty"`
			PrivateProfile *bool      `json:"private_profile,omitempty"`
			ProjectsLimit  *int32     `json:"projects_limit,omitempty"`
			Pronouns       *string    `json:"pronouns,omitempty"`
			PublicEmail    *string    `json:"public_email,omitempty"`
			ScimIdentities *struct {
				Active    *string `json:"active,omitempty"`
				ExternUid *string `json:"extern_uid,omitempty"`
				GroupId   *string `json:"group_id,omitempty"`
			} `json:"scim_identities,omitempty"`
			SharedRunnersMinutesLimit *string `json:"shared_runners_minutes_limit,omitempty"`
			Skype                     *string `json:"skype,omitempty"`
			State                     *string `json:"state,omitempty"`
			ThemeId                   *int32  `json:"theme_id,omitempty"`
			Twitter                   *string `json:"twitter,omitempty"`
			TwoFactorEnabled          *bool   `json:"two_factor_enabled,omitempty"`
			Username                  *string `json:"username,omitempty"`
			WebUrl                    *string `json:"web_url,omitempty"`
			WebsiteUrl                *string `json:"website_url,omitempty"`
			WorkInformation           *string `json:"work_information,omitempty"`
		} `json:"user,omitempty"`
	}
}

func (c *Client) GetApiV4Keys(ctx context.Context, params *GetApiV4KeysParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4KeysRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4KeysId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4KeysIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4KeysRequest(server string, params *GetApiV4KeysParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/keys")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "fingerprint", runtime.ParamLocationQuery, params.Fingerprint); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4KeysIdRequest(server string, id string) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/keys/%s", pathParam0)
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
func (r GetApiV4KeysResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4KeysResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4KeysIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4KeysIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4KeysWithResponse(ctx context.Context, params *GetApiV4KeysParams, reqEditors ...RequestEditorFn) (*GetApiV4KeysResponse, error) {
	rsp, err := c.GetApiV4Keys(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4KeysResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4KeysIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4KeysIdResponse, error) {
	rsp, err := c.GetApiV4KeysId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4KeysIdResponse(rsp)
}
func ParseGetApiV4KeysResponse(rsp *http.Response) (*GetApiV4KeysResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4KeysResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AvatarPath       *string    `json:"avatar_path,omitempty"`
			AvatarUrl        *string    `json:"avatar_url,omitempty"`
			Bio              *string    `json:"bio,omitempty"`
			Bot              *string    `json:"bot,omitempty"`
			CanCreateGroup   *bool      `json:"can_create_group,omitempty"`
			CanCreateProject *bool      `json:"can_create_project,omitempty"`
			ColorSchemeId    *int32     `json:"color_scheme_id,omitempty"`
			CommitEmail      *string    `json:"commit_email,omitempty"`
			ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
			CreatedAt        *string    `json:"created_at,omitempty"`

			// CreatedBy API_Entities_UserBasic model
			CreatedBy *struct {
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
			} `json:"created_by,omitempty"`
			CurrentSignInAt  *time.Time `json:"current_sign_in_at,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Discord                        *string `json:"discord,omitempty"`
			Email                          *string `json:"email,omitempty"`
			EmailResetOfferedAt            *string `json:"email_reset_offered_at,omitempty"`
			EnterpriseGroupAssociatedAt    *string `json:"enterprise_group_associated_at,omitempty"`
			EnterpriseGroupId              *string `json:"enterprise_group_id,omitempty"`
			External                       *string `json:"external,omitempty"`
			ExtraSharedRunnersMinutesLimit *string `json:"extra_shared_runners_minutes_limit,omitempty"`
			Followers                      *string `json:"followers,omitempty"`
			Following                      *string `json:"following,omitempty"`
			Github                         *string `json:"github,omitempty"`
			Id                             *int32  `json:"id,omitempty"`
			Identities                     *struct {
				ExternUid      *string `json:"extern_uid,omitempty"`
				Provider       *string `json:"provider,omitempty"`
				SamlProviderId *string `json:"saml_provider_id,omitempty"`
			} `json:"identities,omitempty"`
			IsAdmin              *string    `json:"is_admin,omitempty"`
			IsAuditor            *string    `json:"is_auditor,omitempty"`
			IsFollowed           *string    `json:"is_followed,omitempty"`
			JobTitle             *string    `json:"job_title,omitempty"`
			LastActivityOn       *time.Time `json:"last_activity_on,omitempty"`
			LastSignInAt         *time.Time `json:"last_sign_in_at,omitempty"`
			Linkedin             *string    `json:"linkedin,omitempty"`
			LocalTime            *string    `json:"local_time,omitempty"`
			Location             *string    `json:"location,omitempty"`
			Locked               *bool      `json:"locked,omitempty"`
			Name                 *string    `json:"name,omitempty"`
			NamespaceId          *string    `json:"namespace_id,omitempty"`
			Note                 *string    `json:"note,omitempty"`
			Organization         *string    `json:"organization,omitempty"`
			PrivateProfile       *bool      `json:"private_profile,omitempty"`
			ProjectsLimit        *int32     `json:"projects_limit,omitempty"`
			Pronouns             *string    `json:"pronouns,omitempty"`
			ProvisionedByGroupId *string    `json:"provisioned_by_group_id,omitempty"`
			PublicEmail          *string    `json:"public_email,omitempty"`
			ScimIdentities       *struct {
				Active    *string `json:"active,omitempty"`
				ExternUid *string `json:"extern_uid,omitempty"`
				GroupId   *string `json:"group_id,omitempty"`
			} `json:"scim_identities,omitempty"`
			SharedRunnersMinutesLimit *string `json:"shared_runners_minutes_limit,omitempty"`
			Skype                     *string `json:"skype,omitempty"`
			State                     *string `json:"state,omitempty"`
			ThemeId                   *int32  `json:"theme_id,omitempty"`
			Twitter                   *string `json:"twitter,omitempty"`
			TwoFactorEnabled          *bool   `json:"two_factor_enabled,omitempty"`
			Username                  *string `json:"username,omitempty"`
			UsingLicenseSeat          *string `json:"using_license_seat,omitempty"`
			WebUrl                    *string `json:"web_url,omitempty"`
			WebsiteUrl                *string `json:"website_url,omitempty"`
			WorkInformation           *string `json:"work_information,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4KeysIdResponse(rsp *http.Response) (*GetApiV4KeysIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4KeysIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CreatedAt  *time.Time `json:"created_at,omitempty"`
			ExpiresAt  *time.Time `json:"expires_at,omitempty"`
			Id         *int32     `json:"id,omitempty"`
			Key        *string    `json:"key,omitempty"`
			LastUsedAt *time.Time `json:"last_used_at,omitempty"`
			Title      *string    `json:"title,omitempty"`
			UsageType  *string    `json:"usage_type,omitempty"`

			// User API_Entities_UserPublic model
			User *struct {
				AvatarPath       *string    `json:"avatar_path,omitempty"`
				AvatarUrl        *string    `json:"avatar_url,omitempty"`
				Bio              *string    `json:"bio,omitempty"`
				Bot              *string    `json:"bot,omitempty"`
				CanCreateGroup   *bool      `json:"can_create_group,omitempty"`
				CanCreateProject *bool      `json:"can_create_project,omitempty"`
				ColorSchemeId    *int32     `json:"color_scheme_id,omitempty"`
				CommitEmail      *string    `json:"commit_email,omitempty"`
				ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
				CreatedAt        *string    `json:"created_at,omitempty"`
				CurrentSignInAt  *time.Time `json:"current_sign_in_at,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Discord                        *string `json:"discord,omitempty"`
				Email                          *string `json:"email,omitempty"`
				External                       *string `json:"external,omitempty"`
				ExtraSharedRunnersMinutesLimit *string `json:"extra_shared_runners_minutes_limit,omitempty"`
				Followers                      *string `json:"followers,omitempty"`
				Following                      *string `json:"following,omitempty"`
				Github                         *string `json:"github,omitempty"`
				Id                             *int32  `json:"id,omitempty"`
				Identities                     *struct {
					ExternUid      *string `json:"extern_uid,omitempty"`
					Provider       *string `json:"provider,omitempty"`
					SamlProviderId *string `json:"saml_provider_id,omitempty"`
				} `json:"identities,omitempty"`
				IsFollowed     *string    `json:"is_followed,omitempty"`
				JobTitle       *string    `json:"job_title,omitempty"`
				LastActivityOn *time.Time `json:"last_activity_on,omitempty"`
				LastSignInAt   *time.Time `json:"last_sign_in_at,omitempty"`
				Linkedin       *string    `json:"linkedin,omitempty"`
				LocalTime      *string    `json:"local_time,omitempty"`
				Location       *string    `json:"location,omitempty"`
				Locked         *bool      `json:"locked,omitempty"`
				Name           *string    `json:"name,omitempty"`
				Organization   *string    `json:"organization,omitempty"`
				PrivateProfile *bool      `json:"private_profile,omitempty"`
				ProjectsLimit  *int32     `json:"projects_limit,omitempty"`
				Pronouns       *string    `json:"pronouns,omitempty"`
				PublicEmail    *string    `json:"public_email,omitempty"`
				ScimIdentities *struct {
					Active    *string `json:"active,omitempty"`
					ExternUid *string `json:"extern_uid,omitempty"`
					GroupId   *string `json:"group_id,omitempty"`
				} `json:"scim_identities,omitempty"`
				SharedRunnersMinutesLimit *string `json:"shared_runners_minutes_limit,omitempty"`
				Skype                     *string `json:"skype,omitempty"`
				State                     *string `json:"state,omitempty"`
				ThemeId                   *int32  `json:"theme_id,omitempty"`
				Twitter                   *string `json:"twitter,omitempty"`
				TwoFactorEnabled          *bool   `json:"two_factor_enabled,omitempty"`
				Username                  *string `json:"username,omitempty"`
				WebUrl                    *string `json:"web_url,omitempty"`
				WebsiteUrl                *string `json:"website_url,omitempty"`
				WorkInformation           *string `json:"work_information,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
