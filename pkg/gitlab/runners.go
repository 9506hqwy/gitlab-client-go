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

// Defines values for GetApiV4RunnersParamsScope.
const (
	GetApiV4RunnersParamsScopeActive         GetApiV4RunnersParamsScope = "active"
	GetApiV4RunnersParamsScopeGroupType      GetApiV4RunnersParamsScope = "group_type"
	GetApiV4RunnersParamsScopeInstanceType   GetApiV4RunnersParamsScope = "instance_type"
	GetApiV4RunnersParamsScopeNeverContacted GetApiV4RunnersParamsScope = "never_contacted"
	GetApiV4RunnersParamsScopeOffline        GetApiV4RunnersParamsScope = "offline"
	GetApiV4RunnersParamsScopeOnline         GetApiV4RunnersParamsScope = "online"
	GetApiV4RunnersParamsScopePaused         GetApiV4RunnersParamsScope = "paused"
	GetApiV4RunnersParamsScopeProjectType    GetApiV4RunnersParamsScope = "project_type"
	GetApiV4RunnersParamsScopeShared         GetApiV4RunnersParamsScope = "shared"
	GetApiV4RunnersParamsScopeSpecific       GetApiV4RunnersParamsScope = "specific"
	GetApiV4RunnersParamsScopeStale          GetApiV4RunnersParamsScope = "stale"
)

// Defines values for GetApiV4RunnersParamsType.
const (
	GetApiV4RunnersParamsTypeGroupType    GetApiV4RunnersParamsType = "group_type"
	GetApiV4RunnersParamsTypeInstanceType GetApiV4RunnersParamsType = "instance_type"
	GetApiV4RunnersParamsTypeProjectType  GetApiV4RunnersParamsType = "project_type"
)

// Defines values for GetApiV4RunnersParamsStatus.
const (
	GetApiV4RunnersParamsStatusActive         GetApiV4RunnersParamsStatus = "active"
	GetApiV4RunnersParamsStatusNeverContacted GetApiV4RunnersParamsStatus = "never_contacted"
	GetApiV4RunnersParamsStatusOffline        GetApiV4RunnersParamsStatus = "offline"
	GetApiV4RunnersParamsStatusOnline         GetApiV4RunnersParamsStatus = "online"
	GetApiV4RunnersParamsStatusPaused         GetApiV4RunnersParamsStatus = "paused"
	GetApiV4RunnersParamsStatusStale          GetApiV4RunnersParamsStatus = "stale"
)

// Defines values for PostApiV4RunnersJSONBodyAccessLevel.
const (
	PostApiV4RunnersJSONBodyAccessLevelNotProtected PostApiV4RunnersJSONBodyAccessLevel = "not_protected"
	PostApiV4RunnersJSONBodyAccessLevelRefProtected PostApiV4RunnersJSONBodyAccessLevel = "ref_protected"
)

// Defines values for GetApiV4RunnersAllParamsScope.
const (
	GetApiV4RunnersAllParamsScopeActive         GetApiV4RunnersAllParamsScope = "active"
	GetApiV4RunnersAllParamsScopeGroupType      GetApiV4RunnersAllParamsScope = "group_type"
	GetApiV4RunnersAllParamsScopeInstanceType   GetApiV4RunnersAllParamsScope = "instance_type"
	GetApiV4RunnersAllParamsScopeNeverContacted GetApiV4RunnersAllParamsScope = "never_contacted"
	GetApiV4RunnersAllParamsScopeOffline        GetApiV4RunnersAllParamsScope = "offline"
	GetApiV4RunnersAllParamsScopeOnline         GetApiV4RunnersAllParamsScope = "online"
	GetApiV4RunnersAllParamsScopePaused         GetApiV4RunnersAllParamsScope = "paused"
	GetApiV4RunnersAllParamsScopeProjectType    GetApiV4RunnersAllParamsScope = "project_type"
	GetApiV4RunnersAllParamsScopeShared         GetApiV4RunnersAllParamsScope = "shared"
	GetApiV4RunnersAllParamsScopeSpecific       GetApiV4RunnersAllParamsScope = "specific"
	GetApiV4RunnersAllParamsScopeStale          GetApiV4RunnersAllParamsScope = "stale"
)

// Defines values for GetApiV4RunnersAllParamsType.
const (
	GetApiV4RunnersAllParamsTypeGroupType    GetApiV4RunnersAllParamsType = "group_type"
	GetApiV4RunnersAllParamsTypeInstanceType GetApiV4RunnersAllParamsType = "instance_type"
	GetApiV4RunnersAllParamsTypeProjectType  GetApiV4RunnersAllParamsType = "project_type"
)

// Defines values for GetApiV4RunnersAllParamsStatus.
const (
	Active         GetApiV4RunnersAllParamsStatus = "active"
	NeverContacted GetApiV4RunnersAllParamsStatus = "never_contacted"
	Offline        GetApiV4RunnersAllParamsStatus = "offline"
	Online         GetApiV4RunnersAllParamsStatus = "online"
	Paused         GetApiV4RunnersAllParamsStatus = "paused"
	Stale          GetApiV4RunnersAllParamsStatus = "stale"
)

// Defines values for PutApiV4RunnersIdJSONBodyAccessLevel.
const (
	PutApiV4RunnersIdJSONBodyAccessLevelNotProtected PutApiV4RunnersIdJSONBodyAccessLevel = "not_protected"
	PutApiV4RunnersIdJSONBodyAccessLevelRefProtected PutApiV4RunnersIdJSONBodyAccessLevel = "ref_protected"
)

// Defines values for GetApiV4RunnersIdJobsParamsStatus.
const (
	Canceled           GetApiV4RunnersIdJobsParamsStatus = "canceled"
	Canceling          GetApiV4RunnersIdJobsParamsStatus = "canceling"
	Created            GetApiV4RunnersIdJobsParamsStatus = "created"
	Failed             GetApiV4RunnersIdJobsParamsStatus = "failed"
	Manual             GetApiV4RunnersIdJobsParamsStatus = "manual"
	Pending            GetApiV4RunnersIdJobsParamsStatus = "pending"
	Preparing          GetApiV4RunnersIdJobsParamsStatus = "preparing"
	Running            GetApiV4RunnersIdJobsParamsStatus = "running"
	Scheduled          GetApiV4RunnersIdJobsParamsStatus = "scheduled"
	Skipped            GetApiV4RunnersIdJobsParamsStatus = "skipped"
	Success            GetApiV4RunnersIdJobsParamsStatus = "success"
	WaitingForCallback GetApiV4RunnersIdJobsParamsStatus = "waiting_for_callback"
	WaitingForResource GetApiV4RunnersIdJobsParamsStatus = "waiting_for_resource"
)

// Defines values for GetApiV4RunnersIdJobsParamsOrderBy.
const (
	GetApiV4RunnersIdJobsParamsOrderById GetApiV4RunnersIdJobsParamsOrderBy = "id"
)

// Defines values for GetApiV4RunnersIdJobsParamsSort.
const (
	GetApiV4RunnersIdJobsParamsSortAsc  GetApiV4RunnersIdJobsParamsSort = "asc"
	GetApiV4RunnersIdJobsParamsSortDesc GetApiV4RunnersIdJobsParamsSort = "desc"
)

type DeleteApiV4RunnersParams struct {
	// Token The runner's authentication token
	Token string `form:"token" json:"token"`
}
type GetApiV4RunnersParams struct {
	// Scope Deprecated: Use `type` or `status` instead. The scope of runners to return
	Scope *GetApiV4RunnersParamsScope `form:"scope,omitempty" json:"scope,omitempty"`

	// Type The type of runners to return
	Type *GetApiV4RunnersParamsType `form:"type,omitempty" json:"type,omitempty"`

	// Paused Whether to include only runners that are accepting or ignoring new jobs
	Paused *bool `form:"paused,omitempty" json:"paused,omitempty"`

	// Status The status of runners to return
	Status *GetApiV4RunnersParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// TagList A list of runner tags
	TagList *[]string `form:"tag_list,omitempty" json:"tag_list,omitempty"`

	// VersionPrefix The version prefix of runners to return
	VersionPrefix *string `form:"version_prefix,omitempty" json:"version_prefix,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4RunnersParamsScope string
type GetApiV4RunnersParamsType string
type GetApiV4RunnersParamsStatus string
type PostApiV4RunnersJSONBody struct {
	// AccessLevel The access level of the runner
	AccessLevel *PostApiV4RunnersJSONBodyAccessLevel `json:"access_level,omitempty"`

	// Active Deprecated: Use `paused` instead. Specifies if the runner is allowed to receive new jobs
	Active *bool `json:"active,omitempty"`

	// Description Description of the runner
	Description *string `json:"description,omitempty"`

	// Info Runner's metadata
	Info *struct {
		// Architecture Runner's architecture
		Architecture *string `json:"architecture,omitempty"`

		// Name Runner's name
		Name *string `json:"name,omitempty"`

		// Platform Runner's platform
		Platform *string `json:"platform,omitempty"`

		// Revision Runner's revision
		Revision *string `json:"revision,omitempty"`

		// Version Runner's version
		Version *string `json:"version,omitempty"`
	} `json:"info,omitempty"`

	// Locked Specifies if the runner should be locked for the current project
	Locked *bool `json:"locked,omitempty"`

	// MaintainerNote Deprecated: see `maintenance_note`
	MaintainerNote *string `json:"maintainer_note,omitempty"`

	// MaintenanceNote Free-form maintenance notes for the runner (1024 characters)
	MaintenanceNote *string `json:"maintenance_note,omitempty"`

	// MaximumTimeout Maximum timeout that limits the amount of time (in seconds) that runners can run jobs
	MaximumTimeout *int32 `json:"maximum_timeout,omitempty"`

	// Paused Specifies if the runner should ignore new jobs
	Paused *bool `json:"paused,omitempty"`

	// RunUntagged Specifies if the runner should handle untagged jobs
	RunUntagged *bool `json:"run_untagged,omitempty"`

	// TagList A list of runner tags
	TagList *[]string `json:"tag_list,omitempty"`

	// Token Registration token
	Token string `json:"token"`
}
type PostApiV4RunnersJSONBodyAccessLevel string
type GetApiV4RunnersAllParams struct {
	// Scope Deprecated: Use `type` or `status` instead. The scope of runners to return
	Scope *GetApiV4RunnersAllParamsScope `form:"scope,omitempty" json:"scope,omitempty"`

	// Type The type of runners to return
	Type *GetApiV4RunnersAllParamsType `form:"type,omitempty" json:"type,omitempty"`

	// Paused Whether to include only runners that are accepting or ignoring new jobs
	Paused *bool `form:"paused,omitempty" json:"paused,omitempty"`

	// Status The status of runners to return
	Status *GetApiV4RunnersAllParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// TagList A list of runner tags
	TagList *[]string `form:"tag_list,omitempty" json:"tag_list,omitempty"`

	// VersionPrefix The version prefix of runners to return
	VersionPrefix *string `form:"version_prefix,omitempty" json:"version_prefix,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4RunnersAllParamsScope string
type GetApiV4RunnersAllParamsType string
type GetApiV4RunnersAllParamsStatus string
type DeleteApiV4RunnersManagersParams struct {
	// Token The runner's authentication token
	Token string `form:"token" json:"token"`

	// SystemId The runner's system identifier.
	SystemId string `form:"system_id" json:"system_id"`
}
type PostApiV4RunnersResetAuthenticationTokenJSONBody struct {
	// Token The current authentication token of the runner
	Token string `json:"token"`
}
type PostApiV4RunnersVerifyJSONBody struct {
	// SystemId The runner's system identifier
	SystemId *string `json:"system_id,omitempty"`

	// Token The runner's authentication token
	Token string `json:"token"`
}
type PutApiV4RunnersIdJSONBody struct {
	// AccessLevel The access level of the runner
	AccessLevel *PutApiV4RunnersIdJSONBodyAccessLevel `json:"access_level,omitempty"`

	// Active Deprecated: Use `paused` instead. Flag indicating whether the runner is allowed to receive jobs
	Active *bool `json:"active,omitempty"`

	// Description The description of the runner
	Description *string `json:"description,omitempty"`

	// Locked Specifies if the runner is locked
	Locked *bool `json:"locked,omitempty"`

	// MaintenanceNote Free-form maintenance notes for the runner (1024 characters)
	MaintenanceNote *string `json:"maintenance_note,omitempty"`

	// MaximumTimeout Maximum timeout that limits the amount of time (in seconds) that runners can run jobs
	MaximumTimeout *int32 `json:"maximum_timeout,omitempty"`

	// Paused Specifies if the runner should ignore new jobs
	Paused *bool `json:"paused,omitempty"`

	// RunUntagged Specifies if the runner can execute untagged jobs
	RunUntagged *bool `json:"run_untagged,omitempty"`

	// TagList The list of tags for a runner
	TagList *[]string `json:"tag_list,omitempty"`
}
type PutApiV4RunnersIdJSONBodyAccessLevel string
type GetApiV4RunnersIdJobsParams struct {
	// SystemId System ID associated with the runner manager
	SystemId *string `form:"system_id,omitempty" json:"system_id,omitempty"`

	// Status Status of the job
	Status *GetApiV4RunnersIdJobsParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// OrderBy Order by `id`
	OrderBy *GetApiV4RunnersIdJobsParamsOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Sort Sort by `asc` or `desc` order. Specify `order_by` as well, including for `id`
	Sort *GetApiV4RunnersIdJobsParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Cursor Cursor for obtaining the next set of records
	Cursor *string `form:"cursor,omitempty" json:"cursor,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4RunnersIdJobsParamsStatus string
type GetApiV4RunnersIdJobsParamsOrderBy string
type GetApiV4RunnersIdJobsParamsSort string
type PostApiV4RunnersJSONRequestBody PostApiV4RunnersJSONBody
type PostApiV4RunnersResetAuthenticationTokenJSONRequestBody PostApiV4RunnersResetAuthenticationTokenJSONBody
type PostApiV4RunnersVerifyJSONRequestBody PostApiV4RunnersVerifyJSONBody
type PutApiV4RunnersIdJSONRequestBody PutApiV4RunnersIdJSONBody
type DeleteApiV4RunnersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4RunnersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active    *bool      `json:"active,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`

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
		Description *string                       `json:"description,omitempty"`
		Id          *int32                        `json:"id,omitempty"`
		IpAddress   *string                       `json:"ip_address,omitempty"`
		IsShared    *bool                         `json:"is_shared,omitempty"`
		Name        *string                       `json:"name,omitempty"`
		Online      *bool                         `json:"online,omitempty"`
		Paused      *bool                         `json:"paused,omitempty"`
		RunnerType  *GetApiV4Runners200RunnerType `json:"runner_type,omitempty"`
		Status      *string                       `json:"status,omitempty"`
	}
}
type GetApiV4Runners200RunnerType string
type PostApiV4RunnersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Id             *string `json:"id,omitempty"`
		Token          *string `json:"token,omitempty"`
		TokenExpiresAt *string `json:"token_expires_at,omitempty"`
	}
}
type GetApiV4RunnersAllResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Active    *bool      `json:"active,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`

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
		Description *string                          `json:"description,omitempty"`
		Id          *int32                           `json:"id,omitempty"`
		IpAddress   *string                          `json:"ip_address,omitempty"`
		IsShared    *bool                            `json:"is_shared,omitempty"`
		Name        *string                          `json:"name,omitempty"`
		Online      *bool                            `json:"online,omitempty"`
		Paused      *bool                            `json:"paused,omitempty"`
		RunnerType  *GetApiV4RunnersAll200RunnerType `json:"runner_type,omitempty"`
		Status      *string                          `json:"status,omitempty"`
	}
}
type GetApiV4RunnersAll200RunnerType string
type DeleteApiV4RunnersManagersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4RunnersResetAuthenticationTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Token          *string `json:"token,omitempty"`
		TokenExpiresAt *string `json:"token_expires_at,omitempty"`
	}
}
type PostApiV4RunnersResetRegistrationTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Token          *string `json:"token,omitempty"`
		TokenExpiresAt *string `json:"token_expires_at,omitempty"`
	}
}
type PostApiV4RunnersVerifyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type DeleteApiV4RunnersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *struct {
		Active    *bool      `json:"active,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`

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
		Description *string                            `json:"description,omitempty"`
		Id          *int32                             `json:"id,omitempty"`
		IpAddress   *string                            `json:"ip_address,omitempty"`
		IsShared    *bool                              `json:"is_shared,omitempty"`
		Name        *string                            `json:"name,omitempty"`
		Online      *bool                              `json:"online,omitempty"`
		Paused      *bool                              `json:"paused,omitempty"`
		RunnerType  *DeleteApiV4RunnersId204RunnerType `json:"runner_type,omitempty"`
		Status      *string                            `json:"status,omitempty"`
	}
}
type DeleteApiV4RunnersId204RunnerType string
type GetApiV4RunnersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AccessLevel  *string    `json:"access_level,omitempty"`
		Active       *bool      `json:"active,omitempty"`
		Architecture *string    `json:"architecture,omitempty"`
		ContactedAt  *string    `json:"contacted_at,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`

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
		Description *string `json:"description,omitempty"`

		// Groups API_Entities_BasicGroupDetails model
		Groups *struct {
			Id     *string `json:"id,omitempty"`
			Name   *string `json:"name,omitempty"`
			WebUrl *string `json:"web_url,omitempty"`
		} `json:"groups,omitempty"`
		Id              *int32  `json:"id,omitempty"`
		IpAddress       *string `json:"ip_address,omitempty"`
		IsShared        *bool   `json:"is_shared,omitempty"`
		Locked          *string `json:"locked,omitempty"`
		MaintenanceNote *string `json:"maintenance_note,omitempty"`
		MaximumTimeout  *string `json:"maximum_timeout,omitempty"`
		Name            *string `json:"name,omitempty"`
		Online          *bool   `json:"online,omitempty"`
		Paused          *bool   `json:"paused,omitempty"`
		Platform        *string `json:"platform,omitempty"`

		// Projects API_Entities_BasicProjectDetails model
		Projects *struct {
			AvatarUrl *string    `json:"avatar_url,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

			// CustomAttributes API_Entities_CustomAttribute model
			CustomAttributes *struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			DefaultBranch  *string    `json:"default_branch,omitempty"`
			Description    *string    `json:"description,omitempty"`
			ForksCount     *int32     `json:"forks_count,omitempty"`
			HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
			Id             *int32     `json:"id,omitempty"`
			LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
			License        *struct {
				HtmlUrl   *string `json:"html_url,omitempty"`
				Key       *string `json:"key,omitempty"`
				Name      *string `json:"name,omitempty"`
				Nickname  *string `json:"nickname,omitempty"`
				SourceUrl *string `json:"source_url,omitempty"`
			} `json:"license,omitempty"`
			LicenseUrl        *string `json:"license_url,omitempty"`
			Name              *string `json:"name,omitempty"`
			NameWithNamespace *string `json:"name_with_namespace,omitempty"`
			Namespace         *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				FullPath  *string `json:"full_path,omitempty"`
				Id        *int32  `json:"id,omitempty"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespace,omitempty"`
			Path              *string   `json:"path,omitempty"`
			PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
			ReadmeUrl         *string   `json:"readme_url,omitempty"`
			RepositoryStorage *string   `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
			StarCount         *int32    `json:"star_count,omitempty"`
			TagList           *[]string `json:"tag_list,omitempty"`
			Topics            *[]string `json:"topics,omitempty"`
			WebUrl            *string   `json:"web_url,omitempty"`
		} `json:"projects,omitempty"`
		Revision    *string                         `json:"revision,omitempty"`
		RunUntagged *string                         `json:"run_untagged,omitempty"`
		RunnerType  *GetApiV4RunnersId200RunnerType `json:"runner_type,omitempty"`
		Status      *string                         `json:"status,omitempty"`
		TagList     *string                         `json:"tag_list,omitempty"`
		Version     *string                         `json:"version,omitempty"`
	}
}
type GetApiV4RunnersId200RunnerType string
type PutApiV4RunnersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AccessLevel  *string    `json:"access_level,omitempty"`
		Active       *bool      `json:"active,omitempty"`
		Architecture *string    `json:"architecture,omitempty"`
		ContactedAt  *string    `json:"contacted_at,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`

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
		Description *string `json:"description,omitempty"`

		// Groups API_Entities_BasicGroupDetails model
		Groups *struct {
			Id     *string `json:"id,omitempty"`
			Name   *string `json:"name,omitempty"`
			WebUrl *string `json:"web_url,omitempty"`
		} `json:"groups,omitempty"`
		Id              *int32  `json:"id,omitempty"`
		IpAddress       *string `json:"ip_address,omitempty"`
		IsShared        *bool   `json:"is_shared,omitempty"`
		Locked          *string `json:"locked,omitempty"`
		MaintenanceNote *string `json:"maintenance_note,omitempty"`
		MaximumTimeout  *string `json:"maximum_timeout,omitempty"`
		Name            *string `json:"name,omitempty"`
		Online          *bool   `json:"online,omitempty"`
		Paused          *bool   `json:"paused,omitempty"`
		Platform        *string `json:"platform,omitempty"`

		// Projects API_Entities_BasicProjectDetails model
		Projects *struct {
			AvatarUrl *string    `json:"avatar_url,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

			// CustomAttributes API_Entities_CustomAttribute model
			CustomAttributes *struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			DefaultBranch  *string    `json:"default_branch,omitempty"`
			Description    *string    `json:"description,omitempty"`
			ForksCount     *int32     `json:"forks_count,omitempty"`
			HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
			Id             *int32     `json:"id,omitempty"`
			LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
			License        *struct {
				HtmlUrl   *string `json:"html_url,omitempty"`
				Key       *string `json:"key,omitempty"`
				Name      *string `json:"name,omitempty"`
				Nickname  *string `json:"nickname,omitempty"`
				SourceUrl *string `json:"source_url,omitempty"`
			} `json:"license,omitempty"`
			LicenseUrl        *string `json:"license_url,omitempty"`
			Name              *string `json:"name,omitempty"`
			NameWithNamespace *string `json:"name_with_namespace,omitempty"`
			Namespace         *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				FullPath  *string `json:"full_path,omitempty"`
				Id        *int32  `json:"id,omitempty"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespace,omitempty"`
			Path              *string   `json:"path,omitempty"`
			PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
			ReadmeUrl         *string   `json:"readme_url,omitempty"`
			RepositoryStorage *string   `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
			StarCount         *int32    `json:"star_count,omitempty"`
			TagList           *[]string `json:"tag_list,omitempty"`
			Topics            *[]string `json:"topics,omitempty"`
			WebUrl            *string   `json:"web_url,omitempty"`
		} `json:"projects,omitempty"`
		Revision    *string                         `json:"revision,omitempty"`
		RunUntagged *string                         `json:"run_untagged,omitempty"`
		RunnerType  *PutApiV4RunnersId200RunnerType `json:"runner_type,omitempty"`
		Status      *string                         `json:"status,omitempty"`
		TagList     *string                         `json:"tag_list,omitempty"`
		Version     *string                         `json:"version,omitempty"`
	}
}
type PutApiV4RunnersId200RunnerType string
type GetApiV4RunnersIdJobsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AllowFailure *bool `json:"allow_failure,omitempty"`

		// Commit API_Entities_Commit model
		Commit *struct {
			AuthorEmail      *string                 `json:"author_email,omitempty"`
			AuthorName       *string                 `json:"author_name,omitempty"`
			AuthoredDate     *time.Time              `json:"authored_date,omitempty"`
			CommittedDate    *time.Time              `json:"committed_date,omitempty"`
			CommitterEmail   *string                 `json:"committer_email,omitempty"`
			CommitterName    *string                 `json:"committer_name,omitempty"`
			CreatedAt        *time.Time              `json:"created_at,omitempty"`
			ExtendedTrailers *map[string]interface{} `json:"extended_trailers,omitempty"`
			Id               *string                 `json:"id,omitempty"`
			Message          *string                 `json:"message,omitempty"`
			ParentIds        *[]string               `json:"parent_ids,omitempty"`
			ShortId          *string                 `json:"short_id,omitempty"`
			Title            *string                 `json:"title,omitempty"`
			Trailers         *map[string]interface{} `json:"trailers,omitempty"`
			WebUrl           *string                 `json:"web_url,omitempty"`
		} `json:"commit,omitempty"`
		Coverage  *float32   `json:"coverage,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`

		// Duration Time spent running
		Duration      *float32   `json:"duration,omitempty"`
		ErasedAt      *time.Time `json:"erased_at,omitempty"`
		FailureReason *string    `json:"failure_reason,omitempty"`
		FinishedAt    *time.Time `json:"finished_at,omitempty"`
		Id            *int32     `json:"id,omitempty"`
		Name          *string    `json:"name,omitempty"`

		// Pipeline API_Entities_Ci_PipelineBasic model
		Pipeline *struct {
			CreatedAt *time.Time `json:"created_at,omitempty"`
			Id        *int32     `json:"id,omitempty"`
			Iid       *int32     `json:"iid,omitempty"`
			ProjectId *int32     `json:"project_id,omitempty"`
			Ref       *string    `json:"ref,omitempty"`
			Sha       *string    `json:"sha,omitempty"`
			Source    *string    `json:"source,omitempty"`
			Status    *string    `json:"status,omitempty"`
			UpdatedAt *time.Time `json:"updated_at,omitempty"`
			WebUrl    *string    `json:"web_url,omitempty"`
		} `json:"pipeline,omitempty"`
		Project *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"project,omitempty"`

		// QueuedDuration Time spent enqueued
		QueuedDuration *float32   `json:"queued_duration,omitempty"`
		Ref            *string    `json:"ref,omitempty"`
		Stage          *string    `json:"stage,omitempty"`
		StartedAt      *time.Time `json:"started_at,omitempty"`
		Status         *string    `json:"status,omitempty"`
		Tag            *bool      `json:"tag,omitempty"`
		User           *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			Bio              *string `json:"bio,omitempty"`
			Bot              *string `json:"bot,omitempty"`
			CreatedAt        *string `json:"created_at,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Discord         *string `json:"discord,omitempty"`
			Followers       *string `json:"followers,omitempty"`
			Following       *string `json:"following,omitempty"`
			Github          *string `json:"github,omitempty"`
			Id              *int32  `json:"id,omitempty"`
			IsFollowed      *string `json:"is_followed,omitempty"`
			JobTitle        *string `json:"job_title,omitempty"`
			Linkedin        *string `json:"linkedin,omitempty"`
			LocalTime       *string `json:"local_time,omitempty"`
			Location        *string `json:"location,omitempty"`
			Locked          *bool   `json:"locked,omitempty"`
			Name            *string `json:"name,omitempty"`
			Organization    *string `json:"organization,omitempty"`
			Pronouns        *string `json:"pronouns,omitempty"`
			PublicEmail     *string `json:"public_email,omitempty"`
			Skype           *string `json:"skype,omitempty"`
			State           *string `json:"state,omitempty"`
			Twitter         *string `json:"twitter,omitempty"`
			Username        *string `json:"username,omitempty"`
			WebUrl          *string `json:"web_url,omitempty"`
			WebsiteUrl      *string `json:"website_url,omitempty"`
			WorkInformation *string `json:"work_information,omitempty"`
		} `json:"user,omitempty"`
		WebUrl *string `json:"web_url,omitempty"`
	}
}
type GetApiV4RunnersIdManagersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		Architecture *string `json:"architecture,omitempty"`
		ContactedAt  *string `json:"contacted_at,omitempty"`
		CreatedAt    *string `json:"created_at,omitempty"`
		Id           *int32  `json:"id,omitempty"`
		IpAddress    *string `json:"ip_address,omitempty"`
		Platform     *string `json:"platform,omitempty"`
		Revision     *string `json:"revision,omitempty"`
		Status       *string `json:"status,omitempty"`
		SystemId     *string `json:"system_id,omitempty"`
		Version      *string `json:"version,omitempty"`
	}
}
type PostApiV4RunnersIdResetAuthenticationTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Token          *string `json:"token,omitempty"`
		TokenExpiresAt *string `json:"token_expires_at,omitempty"`
	}
}

func (c *Client) DeleteApiV4Runners(ctx context.Context, params *DeleteApiV4RunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4RunnersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4Runners(ctx context.Context, params *GetApiV4RunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RunnersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4Runners(ctx context.Context, body PostApiV4RunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4RunnersAll(ctx context.Context, params *GetApiV4RunnersAllParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RunnersAllRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4RunnersManagers(ctx context.Context, params *DeleteApiV4RunnersManagersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4RunnersManagersRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersResetAuthenticationTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersResetAuthenticationTokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersResetAuthenticationToken(ctx context.Context, body PostApiV4RunnersResetAuthenticationTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersResetAuthenticationTokenRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersResetRegistrationToken(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersResetRegistrationTokenRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersVerifyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersVerifyRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersVerify(ctx context.Context, body PostApiV4RunnersVerifyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersVerifyRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4RunnersId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4RunnersIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4RunnersId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RunnersIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4RunnersIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4RunnersIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4RunnersId(ctx context.Context, id int32, body PutApiV4RunnersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4RunnersIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4RunnersIdJobs(ctx context.Context, id int32, params *GetApiV4RunnersIdJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RunnersIdJobsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4RunnersIdManagers(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4RunnersIdManagersRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4RunnersIdResetAuthenticationToken(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4RunnersIdResetAuthenticationTokenRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewDeleteApiV4RunnersRequest(server string, params *DeleteApiV4RunnersParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "token", runtime.ParamLocationQuery, params.Token); err != nil {
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

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4RunnersRequest(server string, params *GetApiV4RunnersParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners")
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

		if params.Type != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
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

		if params.Paused != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "paused", runtime.ParamLocationQuery, *params.Paused); err != nil {
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

		if params.TagList != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "tag_list", runtime.ParamLocationQuery, *params.TagList); err != nil {
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

		if params.VersionPrefix != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version_prefix", runtime.ParamLocationQuery, *params.VersionPrefix); err != nil {
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
func NewPostApiV4RunnersRequest(server string, body PostApiV4RunnersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4RunnersRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4RunnersRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners")
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
func NewGetApiV4RunnersAllRequest(server string, params *GetApiV4RunnersAllParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners/all")
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

		if params.Type != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "type", runtime.ParamLocationQuery, *params.Type); err != nil {
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

		if params.Paused != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "paused", runtime.ParamLocationQuery, *params.Paused); err != nil {
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

		if params.TagList != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "tag_list", runtime.ParamLocationQuery, *params.TagList); err != nil {
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

		if params.VersionPrefix != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version_prefix", runtime.ParamLocationQuery, *params.VersionPrefix); err != nil {
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
func NewDeleteApiV4RunnersManagersRequest(server string, params *DeleteApiV4RunnersManagersParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners/managers")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "token", runtime.ParamLocationQuery, params.Token); err != nil {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "system_id", runtime.ParamLocationQuery, params.SystemId); err != nil {
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

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4RunnersResetAuthenticationTokenRequest(server string, body PostApiV4RunnersResetAuthenticationTokenJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4RunnersResetAuthenticationTokenRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4RunnersResetAuthenticationTokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners/reset_authentication_token")
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
func NewPostApiV4RunnersResetRegistrationTokenRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners/reset_registration_token")
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
func NewPostApiV4RunnersVerifyRequest(server string, body PostApiV4RunnersVerifyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4RunnersVerifyRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4RunnersVerifyRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/runners/verify")
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
func NewDeleteApiV4RunnersIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s", pathParam0)
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
func NewGetApiV4RunnersIdRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s", pathParam0)
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
func NewPutApiV4RunnersIdRequest(server string, id int32, body PutApiV4RunnersIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4RunnersIdRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4RunnersIdRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s", pathParam0)
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
func NewGetApiV4RunnersIdJobsRequest(server string, id int32, params *GetApiV4RunnersIdJobsParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s/jobs", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.SystemId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "system_id", runtime.ParamLocationQuery, *params.SystemId); err != nil {
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

		if params.OrderBy != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order_by", runtime.ParamLocationQuery, *params.OrderBy); err != nil {
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

		if params.Cursor != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "cursor", runtime.ParamLocationQuery, *params.Cursor); err != nil {
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
func NewGetApiV4RunnersIdManagersRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s/managers", pathParam0)
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
func NewPostApiV4RunnersIdResetAuthenticationTokenRequest(server string, id int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/runners/%s/reset_authentication_token", pathParam0)
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
func (r DeleteApiV4RunnersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4RunnersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4RunnersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RunnersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4RunnersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4RunnersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4RunnersAllResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RunnersAllResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4RunnersManagersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4RunnersManagersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4RunnersResetAuthenticationTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4RunnersResetAuthenticationTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4RunnersResetRegistrationTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4RunnersResetRegistrationTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4RunnersVerifyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4RunnersVerifyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4RunnersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4RunnersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4RunnersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RunnersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4RunnersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4RunnersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4RunnersIdJobsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RunnersIdJobsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4RunnersIdManagersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4RunnersIdManagersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4RunnersIdResetAuthenticationTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4RunnersIdResetAuthenticationTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) DeleteApiV4RunnersWithResponse(ctx context.Context, params *DeleteApiV4RunnersParams, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersResponse, error) {
	rsp, err := c.DeleteApiV4Runners(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4RunnersResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4RunnersWithResponse(ctx context.Context, params *GetApiV4RunnersParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersResponse, error) {
	rsp, err := c.GetApiV4Runners(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RunnersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResponse, error) {
	rsp, err := c.PostApiV4RunnersWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersWithResponse(ctx context.Context, body PostApiV4RunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResponse, error) {
	rsp, err := c.PostApiV4Runners(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4RunnersAllWithResponse(ctx context.Context, params *GetApiV4RunnersAllParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersAllResponse, error) {
	rsp, err := c.GetApiV4RunnersAll(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RunnersAllResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4RunnersManagersWithResponse(ctx context.Context, params *DeleteApiV4RunnersManagersParams, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersManagersResponse, error) {
	rsp, err := c.DeleteApiV4RunnersManagers(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4RunnersManagersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersResetAuthenticationTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetAuthenticationTokenResponse, error) {
	rsp, err := c.PostApiV4RunnersResetAuthenticationTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersResetAuthenticationTokenResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersResetAuthenticationTokenWithResponse(ctx context.Context, body PostApiV4RunnersResetAuthenticationTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetAuthenticationTokenResponse, error) {
	rsp, err := c.PostApiV4RunnersResetAuthenticationToken(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersResetAuthenticationTokenResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersResetRegistrationTokenWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetRegistrationTokenResponse, error) {
	rsp, err := c.PostApiV4RunnersResetRegistrationToken(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersResetRegistrationTokenResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersVerifyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersVerifyResponse, error) {
	rsp, err := c.PostApiV4RunnersVerifyWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersVerifyResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersVerifyWithResponse(ctx context.Context, body PostApiV4RunnersVerifyJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersVerifyResponse, error) {
	rsp, err := c.PostApiV4RunnersVerify(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersVerifyResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4RunnersIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersIdResponse, error) {
	rsp, err := c.DeleteApiV4RunnersId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4RunnersIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4RunnersIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdResponse, error) {
	rsp, err := c.GetApiV4RunnersId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RunnersIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4RunnersIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4RunnersIdResponse, error) {
	rsp, err := c.PutApiV4RunnersIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4RunnersIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4RunnersIdWithResponse(ctx context.Context, id int32, body PutApiV4RunnersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4RunnersIdResponse, error) {
	rsp, err := c.PutApiV4RunnersId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4RunnersIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4RunnersIdJobsWithResponse(ctx context.Context, id int32, params *GetApiV4RunnersIdJobsParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdJobsResponse, error) {
	rsp, err := c.GetApiV4RunnersIdJobs(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RunnersIdJobsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4RunnersIdManagersWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdManagersResponse, error) {
	rsp, err := c.GetApiV4RunnersIdManagers(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4RunnersIdManagersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4RunnersIdResetAuthenticationTokenWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4RunnersIdResetAuthenticationTokenResponse, error) {
	rsp, err := c.PostApiV4RunnersIdResetAuthenticationToken(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4RunnersIdResetAuthenticationTokenResponse(rsp)
}
func ParseDeleteApiV4RunnersResponse(rsp *http.Response) (*DeleteApiV4RunnersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4RunnersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4RunnersResponse(rsp *http.Response) (*GetApiV4RunnersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RunnersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active    *bool      `json:"active,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

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
			Description *string                       `json:"description,omitempty"`
			Id          *int32                        `json:"id,omitempty"`
			IpAddress   *string                       `json:"ip_address,omitempty"`
			IsShared    *bool                         `json:"is_shared,omitempty"`
			Name        *string                       `json:"name,omitempty"`
			Online      *bool                         `json:"online,omitempty"`
			Paused      *bool                         `json:"paused,omitempty"`
			RunnerType  *GetApiV4Runners200RunnerType `json:"runner_type,omitempty"`
			Status      *string                       `json:"status,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4RunnersResponse(rsp *http.Response) (*PostApiV4RunnersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4RunnersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Id             *string `json:"id,omitempty"`
			Token          *string `json:"token,omitempty"`
			TokenExpiresAt *string `json:"token_expires_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseGetApiV4RunnersAllResponse(rsp *http.Response) (*GetApiV4RunnersAllResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RunnersAllResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Active    *bool      `json:"active,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

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
			Description *string                          `json:"description,omitempty"`
			Id          *int32                           `json:"id,omitempty"`
			IpAddress   *string                          `json:"ip_address,omitempty"`
			IsShared    *bool                            `json:"is_shared,omitempty"`
			Name        *string                          `json:"name,omitempty"`
			Online      *bool                            `json:"online,omitempty"`
			Paused      *bool                            `json:"paused,omitempty"`
			RunnerType  *GetApiV4RunnersAll200RunnerType `json:"runner_type,omitempty"`
			Status      *string                          `json:"status,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4RunnersManagersResponse(rsp *http.Response) (*DeleteApiV4RunnersManagersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4RunnersManagersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4RunnersResetAuthenticationTokenResponse(rsp *http.Response) (*PostApiV4RunnersResetAuthenticationTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4RunnersResetAuthenticationTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Token          *string `json:"token,omitempty"`
			TokenExpiresAt *string `json:"token_expires_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4RunnersResetRegistrationTokenResponse(rsp *http.Response) (*PostApiV4RunnersResetRegistrationTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4RunnersResetRegistrationTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Token          *string `json:"token,omitempty"`
			TokenExpiresAt *string `json:"token_expires_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParsePostApiV4RunnersVerifyResponse(rsp *http.Response) (*PostApiV4RunnersVerifyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4RunnersVerifyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseDeleteApiV4RunnersIdResponse(rsp *http.Response) (*DeleteApiV4RunnersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4RunnersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest struct {
			Active    *bool      `json:"active,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

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
			Description *string                            `json:"description,omitempty"`
			Id          *int32                             `json:"id,omitempty"`
			IpAddress   *string                            `json:"ip_address,omitempty"`
			IsShared    *bool                              `json:"is_shared,omitempty"`
			Name        *string                            `json:"name,omitempty"`
			Online      *bool                              `json:"online,omitempty"`
			Paused      *bool                              `json:"paused,omitempty"`
			RunnerType  *DeleteApiV4RunnersId204RunnerType `json:"runner_type,omitempty"`
			Status      *string                            `json:"status,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	}

	return response, nil
}
func ParseGetApiV4RunnersIdResponse(rsp *http.Response) (*GetApiV4RunnersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RunnersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AccessLevel  *string    `json:"access_level,omitempty"`
			Active       *bool      `json:"active,omitempty"`
			Architecture *string    `json:"architecture,omitempty"`
			ContactedAt  *string    `json:"contacted_at,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`

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
			Description *string `json:"description,omitempty"`

			// Groups API_Entities_BasicGroupDetails model
			Groups *struct {
				Id     *string `json:"id,omitempty"`
				Name   *string `json:"name,omitempty"`
				WebUrl *string `json:"web_url,omitempty"`
			} `json:"groups,omitempty"`
			Id              *int32  `json:"id,omitempty"`
			IpAddress       *string `json:"ip_address,omitempty"`
			IsShared        *bool   `json:"is_shared,omitempty"`
			Locked          *string `json:"locked,omitempty"`
			MaintenanceNote *string `json:"maintenance_note,omitempty"`
			MaximumTimeout  *string `json:"maximum_timeout,omitempty"`
			Name            *string `json:"name,omitempty"`
			Online          *bool   `json:"online,omitempty"`
			Paused          *bool   `json:"paused,omitempty"`
			Platform        *string `json:"platform,omitempty"`

			// Projects API_Entities_BasicProjectDetails model
			Projects *struct {
				AvatarUrl *string    `json:"avatar_url,omitempty"`
				CreatedAt *time.Time `json:"created_at,omitempty"`

				// CustomAttributes API_Entities_CustomAttribute model
				CustomAttributes *struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				DefaultBranch  *string    `json:"default_branch,omitempty"`
				Description    *string    `json:"description,omitempty"`
				ForksCount     *int32     `json:"forks_count,omitempty"`
				HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
				Id             *int32     `json:"id,omitempty"`
				LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
				License        *struct {
					HtmlUrl   *string `json:"html_url,omitempty"`
					Key       *string `json:"key,omitempty"`
					Name      *string `json:"name,omitempty"`
					Nickname  *string `json:"nickname,omitempty"`
					SourceUrl *string `json:"source_url,omitempty"`
				} `json:"license,omitempty"`
				LicenseUrl        *string `json:"license_url,omitempty"`
				Name              *string `json:"name,omitempty"`
				NameWithNamespace *string `json:"name_with_namespace,omitempty"`
				Namespace         *struct {
					AvatarUrl *string `json:"avatar_url,omitempty"`
					FullPath  *string `json:"full_path,omitempty"`
					Id        *int32  `json:"id,omitempty"`
					Kind      *string `json:"kind,omitempty"`
					Name      *string `json:"name,omitempty"`
					ParentId  *int32  `json:"parent_id,omitempty"`
					Path      *string `json:"path,omitempty"`
					WebUrl    *string `json:"web_url,omitempty"`
				} `json:"namespace,omitempty"`
				Path              *string   `json:"path,omitempty"`
				PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
				ReadmeUrl         *string   `json:"readme_url,omitempty"`
				RepositoryStorage *string   `json:"repository_storage,omitempty"`
				SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
				StarCount         *int32    `json:"star_count,omitempty"`
				TagList           *[]string `json:"tag_list,omitempty"`
				Topics            *[]string `json:"topics,omitempty"`
				WebUrl            *string   `json:"web_url,omitempty"`
			} `json:"projects,omitempty"`
			Revision    *string                         `json:"revision,omitempty"`
			RunUntagged *string                         `json:"run_untagged,omitempty"`
			RunnerType  *GetApiV4RunnersId200RunnerType `json:"runner_type,omitempty"`
			Status      *string                         `json:"status,omitempty"`
			TagList     *string                         `json:"tag_list,omitempty"`
			Version     *string                         `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4RunnersIdResponse(rsp *http.Response) (*PutApiV4RunnersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4RunnersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AccessLevel  *string    `json:"access_level,omitempty"`
			Active       *bool      `json:"active,omitempty"`
			Architecture *string    `json:"architecture,omitempty"`
			ContactedAt  *string    `json:"contacted_at,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`

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
			Description *string `json:"description,omitempty"`

			// Groups API_Entities_BasicGroupDetails model
			Groups *struct {
				Id     *string `json:"id,omitempty"`
				Name   *string `json:"name,omitempty"`
				WebUrl *string `json:"web_url,omitempty"`
			} `json:"groups,omitempty"`
			Id              *int32  `json:"id,omitempty"`
			IpAddress       *string `json:"ip_address,omitempty"`
			IsShared        *bool   `json:"is_shared,omitempty"`
			Locked          *string `json:"locked,omitempty"`
			MaintenanceNote *string `json:"maintenance_note,omitempty"`
			MaximumTimeout  *string `json:"maximum_timeout,omitempty"`
			Name            *string `json:"name,omitempty"`
			Online          *bool   `json:"online,omitempty"`
			Paused          *bool   `json:"paused,omitempty"`
			Platform        *string `json:"platform,omitempty"`

			// Projects API_Entities_BasicProjectDetails model
			Projects *struct {
				AvatarUrl *string    `json:"avatar_url,omitempty"`
				CreatedAt *time.Time `json:"created_at,omitempty"`

				// CustomAttributes API_Entities_CustomAttribute model
				CustomAttributes *struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				DefaultBranch  *string    `json:"default_branch,omitempty"`
				Description    *string    `json:"description,omitempty"`
				ForksCount     *int32     `json:"forks_count,omitempty"`
				HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
				Id             *int32     `json:"id,omitempty"`
				LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
				License        *struct {
					HtmlUrl   *string `json:"html_url,omitempty"`
					Key       *string `json:"key,omitempty"`
					Name      *string `json:"name,omitempty"`
					Nickname  *string `json:"nickname,omitempty"`
					SourceUrl *string `json:"source_url,omitempty"`
				} `json:"license,omitempty"`
				LicenseUrl        *string `json:"license_url,omitempty"`
				Name              *string `json:"name,omitempty"`
				NameWithNamespace *string `json:"name_with_namespace,omitempty"`
				Namespace         *struct {
					AvatarUrl *string `json:"avatar_url,omitempty"`
					FullPath  *string `json:"full_path,omitempty"`
					Id        *int32  `json:"id,omitempty"`
					Kind      *string `json:"kind,omitempty"`
					Name      *string `json:"name,omitempty"`
					ParentId  *int32  `json:"parent_id,omitempty"`
					Path      *string `json:"path,omitempty"`
					WebUrl    *string `json:"web_url,omitempty"`
				} `json:"namespace,omitempty"`
				Path              *string   `json:"path,omitempty"`
				PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
				ReadmeUrl         *string   `json:"readme_url,omitempty"`
				RepositoryStorage *string   `json:"repository_storage,omitempty"`
				SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
				StarCount         *int32    `json:"star_count,omitempty"`
				TagList           *[]string `json:"tag_list,omitempty"`
				Topics            *[]string `json:"topics,omitempty"`
				WebUrl            *string   `json:"web_url,omitempty"`
			} `json:"projects,omitempty"`
			Revision    *string                         `json:"revision,omitempty"`
			RunUntagged *string                         `json:"run_untagged,omitempty"`
			RunnerType  *PutApiV4RunnersId200RunnerType `json:"runner_type,omitempty"`
			Status      *string                         `json:"status,omitempty"`
			TagList     *string                         `json:"tag_list,omitempty"`
			Version     *string                         `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4RunnersIdJobsResponse(rsp *http.Response) (*GetApiV4RunnersIdJobsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RunnersIdJobsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AllowFailure *bool `json:"allow_failure,omitempty"`

			// Commit API_Entities_Commit model
			Commit *struct {
				AuthorEmail      *string                 `json:"author_email,omitempty"`
				AuthorName       *string                 `json:"author_name,omitempty"`
				AuthoredDate     *time.Time              `json:"authored_date,omitempty"`
				CommittedDate    *time.Time              `json:"committed_date,omitempty"`
				CommitterEmail   *string                 `json:"committer_email,omitempty"`
				CommitterName    *string                 `json:"committer_name,omitempty"`
				CreatedAt        *time.Time              `json:"created_at,omitempty"`
				ExtendedTrailers *map[string]interface{} `json:"extended_trailers,omitempty"`
				Id               *string                 `json:"id,omitempty"`
				Message          *string                 `json:"message,omitempty"`
				ParentIds        *[]string               `json:"parent_ids,omitempty"`
				ShortId          *string                 `json:"short_id,omitempty"`
				Title            *string                 `json:"title,omitempty"`
				Trailers         *map[string]interface{} `json:"trailers,omitempty"`
				WebUrl           *string                 `json:"web_url,omitempty"`
			} `json:"commit,omitempty"`
			Coverage  *float32   `json:"coverage,omitempty"`
			CreatedAt *time.Time `json:"created_at,omitempty"`

			// Duration Time spent running
			Duration      *float32   `json:"duration,omitempty"`
			ErasedAt      *time.Time `json:"erased_at,omitempty"`
			FailureReason *string    `json:"failure_reason,omitempty"`
			FinishedAt    *time.Time `json:"finished_at,omitempty"`
			Id            *int32     `json:"id,omitempty"`
			Name          *string    `json:"name,omitempty"`

			// Pipeline API_Entities_Ci_PipelineBasic model
			Pipeline *struct {
				CreatedAt *time.Time `json:"created_at,omitempty"`
				Id        *int32     `json:"id,omitempty"`
				Iid       *int32     `json:"iid,omitempty"`
				ProjectId *int32     `json:"project_id,omitempty"`
				Ref       *string    `json:"ref,omitempty"`
				Sha       *string    `json:"sha,omitempty"`
				Source    *string    `json:"source,omitempty"`
				Status    *string    `json:"status,omitempty"`
				UpdatedAt *time.Time `json:"updated_at,omitempty"`
				WebUrl    *string    `json:"web_url,omitempty"`
			} `json:"pipeline,omitempty"`
			Project *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"project,omitempty"`

			// QueuedDuration Time spent enqueued
			QueuedDuration *float32   `json:"queued_duration,omitempty"`
			Ref            *string    `json:"ref,omitempty"`
			Stage          *string    `json:"stage,omitempty"`
			StartedAt      *time.Time `json:"started_at,omitempty"`
			Status         *string    `json:"status,omitempty"`
			Tag            *bool      `json:"tag,omitempty"`
			User           *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				Bio              *string `json:"bio,omitempty"`
				Bot              *string `json:"bot,omitempty"`
				CreatedAt        *string `json:"created_at,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Discord         *string `json:"discord,omitempty"`
				Followers       *string `json:"followers,omitempty"`
				Following       *string `json:"following,omitempty"`
				Github          *string `json:"github,omitempty"`
				Id              *int32  `json:"id,omitempty"`
				IsFollowed      *string `json:"is_followed,omitempty"`
				JobTitle        *string `json:"job_title,omitempty"`
				Linkedin        *string `json:"linkedin,omitempty"`
				LocalTime       *string `json:"local_time,omitempty"`
				Location        *string `json:"location,omitempty"`
				Locked          *bool   `json:"locked,omitempty"`
				Name            *string `json:"name,omitempty"`
				Organization    *string `json:"organization,omitempty"`
				Pronouns        *string `json:"pronouns,omitempty"`
				PublicEmail     *string `json:"public_email,omitempty"`
				Skype           *string `json:"skype,omitempty"`
				State           *string `json:"state,omitempty"`
				Twitter         *string `json:"twitter,omitempty"`
				Username        *string `json:"username,omitempty"`
				WebUrl          *string `json:"web_url,omitempty"`
				WebsiteUrl      *string `json:"website_url,omitempty"`
				WorkInformation *string `json:"work_information,omitempty"`
			} `json:"user,omitempty"`
			WebUrl *string `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4RunnersIdManagersResponse(rsp *http.Response) (*GetApiV4RunnersIdManagersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4RunnersIdManagersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			Architecture *string `json:"architecture,omitempty"`
			ContactedAt  *string `json:"contacted_at,omitempty"`
			CreatedAt    *string `json:"created_at,omitempty"`
			Id           *int32  `json:"id,omitempty"`
			IpAddress    *string `json:"ip_address,omitempty"`
			Platform     *string `json:"platform,omitempty"`
			Revision     *string `json:"revision,omitempty"`
			Status       *string `json:"status,omitempty"`
			SystemId     *string `json:"system_id,omitempty"`
			Version      *string `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4RunnersIdResetAuthenticationTokenResponse(rsp *http.Response) (*PostApiV4RunnersIdResetAuthenticationTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4RunnersIdResetAuthenticationTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Token          *string `json:"token,omitempty"`
			TokenExpiresAt *string `json:"token_expires_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
