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
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for PostApiV4UserRunnersJSONBodyAccessLevel.
const (
	NotProtected PostApiV4UserRunnersJSONBodyAccessLevel = "not_protected"
	RefProtected PostApiV4UserRunnersJSONBodyAccessLevel = "ref_protected"
)

// Defines values for PostApiV4UserRunnersJSONBodyRunnerType.
const (
	PostApiV4UserRunnersJSONBodyRunnerTypeGroupType    PostApiV4UserRunnersJSONBodyRunnerType = "group_type"
	PostApiV4UserRunnersJSONBodyRunnerTypeInstanceType PostApiV4UserRunnersJSONBodyRunnerType = "instance_type"
	PostApiV4UserRunnersJSONBodyRunnerTypeProjectType  PostApiV4UserRunnersJSONBodyRunnerType = "project_type"
)

// Defines values for GetApiV4UsersIdEventsParamsTargetType.
const (
	Design       GetApiV4UsersIdEventsParamsTargetType = "design"
	Issue        GetApiV4UsersIdEventsParamsTargetType = "issue"
	MergeRequest GetApiV4UsersIdEventsParamsTargetType = "merge_request"
	Milestone    GetApiV4UsersIdEventsParamsTargetType = "milestone"
	Note         GetApiV4UsersIdEventsParamsTargetType = "note"
	Project      GetApiV4UsersIdEventsParamsTargetType = "project"
	Snippet      GetApiV4UsersIdEventsParamsTargetType = "snippet"
	User         GetApiV4UsersIdEventsParamsTargetType = "user"
	Wiki         GetApiV4UsersIdEventsParamsTargetType = "wiki"
)

// Defines values for GetApiV4UsersIdEventsParamsSort.
const (
	GetApiV4UsersIdEventsParamsSortAsc  GetApiV4UsersIdEventsParamsSort = "asc"
	GetApiV4UsersIdEventsParamsSortDesc GetApiV4UsersIdEventsParamsSort = "desc"
)

// Defines values for GetApiV4UsersUserIdContributedProjectsParamsOrderBy.
const (
	GetApiV4UsersUserIdContributedProjectsParamsOrderByCreatedAt      GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "created_at"
	GetApiV4UsersUserIdContributedProjectsParamsOrderById             GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "id"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByLastActivityAt GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "last_activity_at"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByName           GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "name"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByPackagesSize   GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "packages_size"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByPath           GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "path"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByRepositorySize GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "repository_size"
	GetApiV4UsersUserIdContributedProjectsParamsOrderBySimilarity     GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "similarity"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByStarCount      GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "star_count"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByStorageSize    GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "storage_size"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByUpdatedAt      GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "updated_at"
	GetApiV4UsersUserIdContributedProjectsParamsOrderByWikiSize       GetApiV4UsersUserIdContributedProjectsParamsOrderBy = "wiki_size"
)

// Defines values for GetApiV4UsersUserIdContributedProjectsParamsSort.
const (
	GetApiV4UsersUserIdContributedProjectsParamsSortAsc  GetApiV4UsersUserIdContributedProjectsParamsSort = "asc"
	GetApiV4UsersUserIdContributedProjectsParamsSortDesc GetApiV4UsersUserIdContributedProjectsParamsSort = "desc"
)

// Defines values for GetApiV4UsersUserIdProjectsParamsOrderBy.
const (
	GetApiV4UsersUserIdProjectsParamsOrderByCreatedAt      GetApiV4UsersUserIdProjectsParamsOrderBy = "created_at"
	GetApiV4UsersUserIdProjectsParamsOrderById             GetApiV4UsersUserIdProjectsParamsOrderBy = "id"
	GetApiV4UsersUserIdProjectsParamsOrderByLastActivityAt GetApiV4UsersUserIdProjectsParamsOrderBy = "last_activity_at"
	GetApiV4UsersUserIdProjectsParamsOrderByName           GetApiV4UsersUserIdProjectsParamsOrderBy = "name"
	GetApiV4UsersUserIdProjectsParamsOrderByPackagesSize   GetApiV4UsersUserIdProjectsParamsOrderBy = "packages_size"
	GetApiV4UsersUserIdProjectsParamsOrderByPath           GetApiV4UsersUserIdProjectsParamsOrderBy = "path"
	GetApiV4UsersUserIdProjectsParamsOrderByRepositorySize GetApiV4UsersUserIdProjectsParamsOrderBy = "repository_size"
	GetApiV4UsersUserIdProjectsParamsOrderBySimilarity     GetApiV4UsersUserIdProjectsParamsOrderBy = "similarity"
	GetApiV4UsersUserIdProjectsParamsOrderByStarCount      GetApiV4UsersUserIdProjectsParamsOrderBy = "star_count"
	GetApiV4UsersUserIdProjectsParamsOrderByStorageSize    GetApiV4UsersUserIdProjectsParamsOrderBy = "storage_size"
	GetApiV4UsersUserIdProjectsParamsOrderByUpdatedAt      GetApiV4UsersUserIdProjectsParamsOrderBy = "updated_at"
	GetApiV4UsersUserIdProjectsParamsOrderByWikiSize       GetApiV4UsersUserIdProjectsParamsOrderBy = "wiki_size"
)

// Defines values for GetApiV4UsersUserIdProjectsParamsSort.
const (
	GetApiV4UsersUserIdProjectsParamsSortAsc  GetApiV4UsersUserIdProjectsParamsSort = "asc"
	GetApiV4UsersUserIdProjectsParamsSortDesc GetApiV4UsersUserIdProjectsParamsSort = "desc"
)

// Defines values for GetApiV4UsersUserIdProjectsParamsVisibility.
const (
	GetApiV4UsersUserIdProjectsParamsVisibilityInternal GetApiV4UsersUserIdProjectsParamsVisibility = "internal"
	GetApiV4UsersUserIdProjectsParamsVisibilityPrivate  GetApiV4UsersUserIdProjectsParamsVisibility = "private"
	GetApiV4UsersUserIdProjectsParamsVisibilityPublic   GetApiV4UsersUserIdProjectsParamsVisibility = "public"
)

// Defines values for GetApiV4UsersUserIdProjectsParamsMinAccessLevel.
const (
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN10 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 10
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN15 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 15
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN20 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 20
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN30 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 30
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN40 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 40
	GetApiV4UsersUserIdProjectsParamsMinAccessLevelN50 GetApiV4UsersUserIdProjectsParamsMinAccessLevel = 50
)

// Defines values for GetApiV4UsersUserIdStarredProjectsParamsOrderBy.
const (
	CreatedAt      GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "created_at"
	Id             GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "id"
	LastActivityAt GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "last_activity_at"
	Name           GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "name"
	PackagesSize   GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "packages_size"
	Path           GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "path"
	RepositorySize GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "repository_size"
	Similarity     GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "similarity"
	StarCount      GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "star_count"
	StorageSize    GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "storage_size"
	UpdatedAt      GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "updated_at"
	WikiSize       GetApiV4UsersUserIdStarredProjectsParamsOrderBy = "wiki_size"
)

// Defines values for GetApiV4UsersUserIdStarredProjectsParamsSort.
const (
	Asc  GetApiV4UsersUserIdStarredProjectsParamsSort = "asc"
	Desc GetApiV4UsersUserIdStarredProjectsParamsSort = "desc"
)

// Defines values for GetApiV4UsersUserIdStarredProjectsParamsVisibility.
const (
	GetApiV4UsersUserIdStarredProjectsParamsVisibilityInternal GetApiV4UsersUserIdStarredProjectsParamsVisibility = "internal"
	GetApiV4UsersUserIdStarredProjectsParamsVisibilityPrivate  GetApiV4UsersUserIdStarredProjectsParamsVisibility = "private"
	GetApiV4UsersUserIdStarredProjectsParamsVisibilityPublic   GetApiV4UsersUserIdStarredProjectsParamsVisibility = "public"
)

// Defines values for GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel.
const (
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN10 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 10
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN15 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 15
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN20 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 20
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN30 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 30
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN40 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 40
	GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevelN50 GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel = 50
)

type PostApiV4UserRunnersJSONBody struct {
	// AccessLevel The access level of the runner
	AccessLevel *PostApiV4UserRunnersJSONBodyAccessLevel `json:"access_level,omitempty"`

	// Description Description of the runner
	Description *string `json:"description,omitempty"`

	// GroupId The ID of the group that the runner is created in
	GroupId int32 `json:"group_id"`

	// Locked Specifies if the runner should be locked for the current project (defaults to false)
	Locked *bool `json:"locked,omitempty"`

	// MaintenanceNote Free-form maintenance notes for the runner (1024 characters)
	MaintenanceNote *string `json:"maintenance_note,omitempty"`

	// MaximumTimeout Maximum timeout that limits the amount of time (in seconds) that runners can run jobs
	MaximumTimeout *int32 `json:"maximum_timeout,omitempty"`

	// Paused Specifies if the runner should ignore new jobs (defaults to false)
	Paused *bool `json:"paused,omitempty"`

	// ProjectId The ID of the project that the runner is created in
	ProjectId int32 `json:"project_id"`

	// RunUntagged Specifies if the runner should handle untagged jobs  (defaults to true)
	RunUntagged *bool `json:"run_untagged,omitempty"`

	// RunnerType Specifies the scope of the runner
	RunnerType PostApiV4UserRunnersJSONBodyRunnerType `json:"runner_type"`

	// TagList A list of runner tags
	TagList *[]string `json:"tag_list,omitempty"`
}
type PostApiV4UserRunnersJSONBodyAccessLevel string
type PostApiV4UserRunnersJSONBodyRunnerType string
type GetApiV4UsersIdEventsParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Action Event action to filter on
	Action *string `form:"action,omitempty" json:"action,omitempty"`

	// TargetType Event target type to filter on
	TargetType *GetApiV4UsersIdEventsParamsTargetType `form:"target_type,omitempty" json:"target_type,omitempty"`

	// Before Include only events created before this date
	Before *openapi_types.Date `form:"before,omitempty" json:"before,omitempty"`

	// After Include only events created after this date
	After *openapi_types.Date `form:"after,omitempty" json:"after,omitempty"`

	// Sort Return events sorted in ascending and descending order
	Sort *GetApiV4UsersIdEventsParamsSort `form:"sort,omitempty" json:"sort,omitempty"`
}
type GetApiV4UsersIdEventsParamsTargetType string
type GetApiV4UsersIdEventsParamsSort string
type GetApiV4UsersUserIdContributedProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *GetApiV4UsersUserIdContributedProjectsParamsOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Sort Return projects sorted in ascending and descending order
	Sort *GetApiV4UsersUserIdContributedProjectsParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty"`
}
type GetApiV4UsersUserIdContributedProjectsParamsOrderBy string
type GetApiV4UsersUserIdContributedProjectsParamsSort string
type GetApiV4UsersUserIdProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *GetApiV4UsersUserIdProjectsParamsOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Sort Return projects sorted in ascending and descending order
	Sort *GetApiV4UsersUserIdProjectsParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Archived Limit by archived status
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`

	// Visibility Limit by visibility
	Visibility *GetApiV4UsersUserIdProjectsParamsVisibility `form:"visibility,omitempty" json:"visibility,omitempty"`

	// Search Return list of projects matching the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// SearchNamespaces Include ancestor namespaces when matching search criteria
	SearchNamespaces *bool `form:"search_namespaces,omitempty" json:"search_namespaces,omitempty"`

	// Owned Limit by owned by authenticated user
	Owned *bool `form:"owned,omitempty" json:"owned,omitempty"`

	// Starred Limit by starred status
	Starred *bool `form:"starred,omitempty" json:"starred,omitempty"`

	// Imported Limit by imported by authenticated user
	Imported *bool `form:"imported,omitempty" json:"imported,omitempty"`

	// Membership Limit by projects that the current user is a member of
	Membership *bool `form:"membership,omitempty" json:"membership,omitempty"`

	// WithIssuesEnabled Limit by enabled issues feature
	WithIssuesEnabled *bool `form:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty"`

	// WithMergeRequestsEnabled Limit by enabled merge requests feature
	WithMergeRequestsEnabled *bool `form:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty"`

	// WithProgrammingLanguage Limit to repositories which use the given programming language
	WithProgrammingLanguage *string `form:"with_programming_language,omitempty" json:"with_programming_language,omitempty"`

	// MinAccessLevel Limit by minimum access level of authenticated user
	MinAccessLevel *GetApiV4UsersUserIdProjectsParamsMinAccessLevel `form:"min_access_level,omitempty" json:"min_access_level,omitempty"`

	// IdAfter Limit results to projects with IDs greater than the specified ID
	IdAfter *int32 `form:"id_after,omitempty" json:"id_after,omitempty"`

	// IdBefore Limit results to projects with IDs less than the specified ID
	IdBefore *int32 `form:"id_before,omitempty" json:"id_before,omitempty"`

	// LastActivityAfter Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityAfter *time.Time `form:"last_activity_after,omitempty" json:"last_activity_after,omitempty"`

	// LastActivityBefore Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityBefore *time.Time `form:"last_activity_before,omitempty" json:"last_activity_before,omitempty"`

	// RepositoryStorage Which storage shard the repository is on. Available only to admins
	RepositoryStorage *string `form:"repository_storage,omitempty" json:"repository_storage,omitempty"`

	// Topic Comma-separated list of topics. Limit results to projects having all topics
	Topic *[]string `form:"topic,omitempty" json:"topic,omitempty"`

	// TopicId Limit results to projects with the assigned topic given by the topic ID
	TopicId *int32 `form:"topic_id,omitempty" json:"topic_id,omitempty"`

	// UpdatedBefore Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedBefore *time.Time `form:"updated_before,omitempty" json:"updated_before,omitempty"`

	// UpdatedAfter Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedAfter *time.Time `form:"updated_after,omitempty" json:"updated_after,omitempty"`

	// IncludePendingDelete Include projects in pending delete state. Can only be set by admins
	IncludePendingDelete *bool `form:"include_pending_delete,omitempty" json:"include_pending_delete,omitempty"`

	// MarkedForDeletionOn Date when the project was marked for deletion
	MarkedForDeletionOn *openapi_types.Date `form:"marked_for_deletion_on,omitempty" json:"marked_for_deletion_on,omitempty"`

	// Active Limit by projects that are not archived and not marked for deletion
	Active *bool `form:"active,omitempty" json:"active,omitempty"`

	// WikiChecksumFailed Limit by projects where wiki checksum is failed
	WikiChecksumFailed *bool `form:"wiki_checksum_failed,omitempty" json:"wiki_checksum_failed,omitempty"`

	// RepositoryChecksumFailed Limit by projects where repository checksum is failed
	RepositoryChecksumFailed *bool `form:"repository_checksum_failed,omitempty" json:"repository_checksum_failed,omitempty"`

	// IncludeHidden Include hidden projects. Can only be set by admins
	IncludeHidden *bool `form:"include_hidden,omitempty" json:"include_hidden,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty"`

	// Statistics Include project statistics
	Statistics *bool `form:"statistics,omitempty" json:"statistics,omitempty"`

	// WithCustomAttributes Include custom attributes in the response
	WithCustomAttributes *bool `form:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty"`
}
type GetApiV4UsersUserIdProjectsParamsOrderBy string
type GetApiV4UsersUserIdProjectsParamsSort string
type GetApiV4UsersUserIdProjectsParamsVisibility string
type GetApiV4UsersUserIdProjectsParamsMinAccessLevel int32
type GetApiV4UsersUserIdStarredProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *GetApiV4UsersUserIdStarredProjectsParamsOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Sort Return projects sorted in ascending and descending order
	Sort *GetApiV4UsersUserIdStarredProjectsParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Archived Limit by archived status
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`

	// Visibility Limit by visibility
	Visibility *GetApiV4UsersUserIdStarredProjectsParamsVisibility `form:"visibility,omitempty" json:"visibility,omitempty"`

	// Search Return list of projects matching the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// SearchNamespaces Include ancestor namespaces when matching search criteria
	SearchNamespaces *bool `form:"search_namespaces,omitempty" json:"search_namespaces,omitempty"`

	// Owned Limit by owned by authenticated user
	Owned *bool `form:"owned,omitempty" json:"owned,omitempty"`

	// Starred Limit by starred status
	Starred *bool `form:"starred,omitempty" json:"starred,omitempty"`

	// Imported Limit by imported by authenticated user
	Imported *bool `form:"imported,omitempty" json:"imported,omitempty"`

	// Membership Limit by projects that the current user is a member of
	Membership *bool `form:"membership,omitempty" json:"membership,omitempty"`

	// WithIssuesEnabled Limit by enabled issues feature
	WithIssuesEnabled *bool `form:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty"`

	// WithMergeRequestsEnabled Limit by enabled merge requests feature
	WithMergeRequestsEnabled *bool `form:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty"`

	// WithProgrammingLanguage Limit to repositories which use the given programming language
	WithProgrammingLanguage *string `form:"with_programming_language,omitempty" json:"with_programming_language,omitempty"`

	// MinAccessLevel Limit by minimum access level of authenticated user
	MinAccessLevel *GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel `form:"min_access_level,omitempty" json:"min_access_level,omitempty"`

	// IdAfter Limit results to projects with IDs greater than the specified ID
	IdAfter *int32 `form:"id_after,omitempty" json:"id_after,omitempty"`

	// IdBefore Limit results to projects with IDs less than the specified ID
	IdBefore *int32 `form:"id_before,omitempty" json:"id_before,omitempty"`

	// LastActivityAfter Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityAfter *time.Time `form:"last_activity_after,omitempty" json:"last_activity_after,omitempty"`

	// LastActivityBefore Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityBefore *time.Time `form:"last_activity_before,omitempty" json:"last_activity_before,omitempty"`

	// RepositoryStorage Which storage shard the repository is on. Available only to admins
	RepositoryStorage *string `form:"repository_storage,omitempty" json:"repository_storage,omitempty"`

	// Topic Comma-separated list of topics. Limit results to projects having all topics
	Topic *[]string `form:"topic,omitempty" json:"topic,omitempty"`

	// TopicId Limit results to projects with the assigned topic given by the topic ID
	TopicId *int32 `form:"topic_id,omitempty" json:"topic_id,omitempty"`

	// UpdatedBefore Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedBefore *time.Time `form:"updated_before,omitempty" json:"updated_before,omitempty"`

	// UpdatedAfter Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedAfter *time.Time `form:"updated_after,omitempty" json:"updated_after,omitempty"`

	// IncludePendingDelete Include projects in pending delete state. Can only be set by admins
	IncludePendingDelete *bool `form:"include_pending_delete,omitempty" json:"include_pending_delete,omitempty"`

	// MarkedForDeletionOn Date when the project was marked for deletion
	MarkedForDeletionOn *openapi_types.Date `form:"marked_for_deletion_on,omitempty" json:"marked_for_deletion_on,omitempty"`

	// Active Limit by projects that are not archived and not marked for deletion
	Active *bool `form:"active,omitempty" json:"active,omitempty"`

	// WikiChecksumFailed Limit by projects where wiki checksum is failed
	WikiChecksumFailed *bool `form:"wiki_checksum_failed,omitempty" json:"wiki_checksum_failed,omitempty"`

	// RepositoryChecksumFailed Limit by projects where repository checksum is failed
	RepositoryChecksumFailed *bool `form:"repository_checksum_failed,omitempty" json:"repository_checksum_failed,omitempty"`

	// IncludeHidden Include hidden projects. Can only be set by admins
	IncludeHidden *bool `form:"include_hidden,omitempty" json:"include_hidden,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty"`

	// Statistics Include project statistics
	Statistics *bool `form:"statistics,omitempty" json:"statistics,omitempty"`
}
type GetApiV4UsersUserIdStarredProjectsParamsOrderBy string
type GetApiV4UsersUserIdStarredProjectsParamsSort string
type GetApiV4UsersUserIdStarredProjectsParamsVisibility string
type GetApiV4UsersUserIdStarredProjectsParamsMinAccessLevel int32
type PostApiV4UserRunnersJSONRequestBody PostApiV4UserRunnersJSONBody
type PostApiV4UserRunnersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Id             *string `json:"id,omitempty"`
		Token          *string `json:"token,omitempty"`
		TokenExpiresAt *string `json:"token_expires_at,omitempty"`
	}
}
type GetApiV4UserCountsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AssignedIssues               *int32 `json:"assigned_issues,omitempty"`
		AssignedMergeRequests        *int32 `json:"assigned_merge_requests,omitempty"`
		MergeRequests                *int32 `json:"merge_requests,omitempty"`
		ReviewRequestedMergeRequests *int32 `json:"review_requested_merge_requests,omitempty"`
		Todos                        *int32 `json:"todos,omitempty"`
	}
}
type GetApiV4UsersIdEventsResponse struct {
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
type GetApiV4UsersUserIdContributedProjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
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
	}
}
type GetApiV4UsersUserIdProjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
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
	}
}
type GetApiV4UsersUserIdStarredProjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
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
	}
}

func (c *Client) PostApiV4UserRunnersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UserRunnersRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4UserRunners(ctx context.Context, body PostApiV4UserRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4UserRunnersRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UserCounts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UserCountsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsersIdEvents(ctx context.Context, id string, params *GetApiV4UsersIdEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsersIdEventsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsersUserIdContributedProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdContributedProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsersUserIdContributedProjectsRequest(c.Server, userId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsersUserIdProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsersUserIdProjectsRequest(c.Server, userId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4UsersUserIdStarredProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdStarredProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4UsersUserIdStarredProjectsRequest(c.Server, userId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4UserRunnersRequest(server string, body PostApiV4UserRunnersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4UserRunnersRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4UserRunnersRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/user/runners")
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
func NewGetApiV4UserCountsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/user_counts")
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
func NewGetApiV4UsersIdEventsRequest(server string, id string, params *GetApiV4UsersIdEventsParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/users/%s/events", pathParam0)
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
func NewGetApiV4UsersUserIdContributedProjectsRequest(server string, userId string, params *GetApiV4UsersUserIdContributedProjectsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/users/%s/contributed_projects", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

		if params.Simple != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "simple", runtime.ParamLocationQuery, *params.Simple); err != nil {
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
func NewGetApiV4UsersUserIdProjectsRequest(server string, userId string, params *GetApiV4UsersUserIdProjectsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/users/%s/projects", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

		if params.Archived != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
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

		if params.Visibility != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "visibility", runtime.ParamLocationQuery, *params.Visibility); err != nil {
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

		if params.SearchNamespaces != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "search_namespaces", runtime.ParamLocationQuery, *params.SearchNamespaces); err != nil {
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

		if params.Owned != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "owned", runtime.ParamLocationQuery, *params.Owned); err != nil {
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

		if params.Starred != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "starred", runtime.ParamLocationQuery, *params.Starred); err != nil {
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

		if params.Imported != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "imported", runtime.ParamLocationQuery, *params.Imported); err != nil {
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

		if params.Membership != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "membership", runtime.ParamLocationQuery, *params.Membership); err != nil {
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

		if params.WithIssuesEnabled != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_issues_enabled", runtime.ParamLocationQuery, *params.WithIssuesEnabled); err != nil {
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

		if params.WithMergeRequestsEnabled != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_merge_requests_enabled", runtime.ParamLocationQuery, *params.WithMergeRequestsEnabled); err != nil {
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

		if params.WithProgrammingLanguage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_programming_language", runtime.ParamLocationQuery, *params.WithProgrammingLanguage); err != nil {
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

		if params.MinAccessLevel != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "min_access_level", runtime.ParamLocationQuery, *params.MinAccessLevel); err != nil {
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

		if params.IdAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_after", runtime.ParamLocationQuery, *params.IdAfter); err != nil {
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

		if params.IdBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_before", runtime.ParamLocationQuery, *params.IdBefore); err != nil {
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

		if params.LastActivityAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_activity_after", runtime.ParamLocationQuery, *params.LastActivityAfter); err != nil {
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

		if params.LastActivityBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_activity_before", runtime.ParamLocationQuery, *params.LastActivityBefore); err != nil {
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

		if params.Topic != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "topic", runtime.ParamLocationQuery, *params.Topic); err != nil {
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

		if params.TopicId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "topic_id", runtime.ParamLocationQuery, *params.TopicId); err != nil {
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

		if params.UpdatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "updated_before", runtime.ParamLocationQuery, *params.UpdatedBefore); err != nil {
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

		if params.UpdatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "updated_after", runtime.ParamLocationQuery, *params.UpdatedAfter); err != nil {
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

		if params.IncludePendingDelete != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "include_pending_delete", runtime.ParamLocationQuery, *params.IncludePendingDelete); err != nil {
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

		if params.MarkedForDeletionOn != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marked_for_deletion_on", runtime.ParamLocationQuery, *params.MarkedForDeletionOn); err != nil {
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

		if params.Active != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "active", runtime.ParamLocationQuery, *params.Active); err != nil {
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

		if params.WikiChecksumFailed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "wiki_checksum_failed", runtime.ParamLocationQuery, *params.WikiChecksumFailed); err != nil {
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

		if params.RepositoryChecksumFailed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "repository_checksum_failed", runtime.ParamLocationQuery, *params.RepositoryChecksumFailed); err != nil {
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

		if params.IncludeHidden != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "include_hidden", runtime.ParamLocationQuery, *params.IncludeHidden); err != nil {
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

		if params.Simple != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "simple", runtime.ParamLocationQuery, *params.Simple); err != nil {
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

		if params.Statistics != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "statistics", runtime.ParamLocationQuery, *params.Statistics); err != nil {
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

		if params.WithCustomAttributes != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_custom_attributes", runtime.ParamLocationQuery, *params.WithCustomAttributes); err != nil {
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
func NewGetApiV4UsersUserIdStarredProjectsRequest(server string, userId string, params *GetApiV4UsersUserIdStarredProjectsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/users/%s/starred_projects", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

		if params.Archived != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
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

		if params.Visibility != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "visibility", runtime.ParamLocationQuery, *params.Visibility); err != nil {
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

		if params.SearchNamespaces != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "search_namespaces", runtime.ParamLocationQuery, *params.SearchNamespaces); err != nil {
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

		if params.Owned != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "owned", runtime.ParamLocationQuery, *params.Owned); err != nil {
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

		if params.Starred != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "starred", runtime.ParamLocationQuery, *params.Starred); err != nil {
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

		if params.Imported != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "imported", runtime.ParamLocationQuery, *params.Imported); err != nil {
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

		if params.Membership != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "membership", runtime.ParamLocationQuery, *params.Membership); err != nil {
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

		if params.WithIssuesEnabled != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_issues_enabled", runtime.ParamLocationQuery, *params.WithIssuesEnabled); err != nil {
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

		if params.WithMergeRequestsEnabled != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_merge_requests_enabled", runtime.ParamLocationQuery, *params.WithMergeRequestsEnabled); err != nil {
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

		if params.WithProgrammingLanguage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_programming_language", runtime.ParamLocationQuery, *params.WithProgrammingLanguage); err != nil {
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

		if params.MinAccessLevel != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "min_access_level", runtime.ParamLocationQuery, *params.MinAccessLevel); err != nil {
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

		if params.IdAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_after", runtime.ParamLocationQuery, *params.IdAfter); err != nil {
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

		if params.IdBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id_before", runtime.ParamLocationQuery, *params.IdBefore); err != nil {
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

		if params.LastActivityAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_activity_after", runtime.ParamLocationQuery, *params.LastActivityAfter); err != nil {
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

		if params.LastActivityBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "last_activity_before", runtime.ParamLocationQuery, *params.LastActivityBefore); err != nil {
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

		if params.Topic != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "topic", runtime.ParamLocationQuery, *params.Topic); err != nil {
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

		if params.TopicId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "topic_id", runtime.ParamLocationQuery, *params.TopicId); err != nil {
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

		if params.UpdatedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "updated_before", runtime.ParamLocationQuery, *params.UpdatedBefore); err != nil {
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

		if params.UpdatedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "updated_after", runtime.ParamLocationQuery, *params.UpdatedAfter); err != nil {
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

		if params.IncludePendingDelete != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "include_pending_delete", runtime.ParamLocationQuery, *params.IncludePendingDelete); err != nil {
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

		if params.MarkedForDeletionOn != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marked_for_deletion_on", runtime.ParamLocationQuery, *params.MarkedForDeletionOn); err != nil {
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

		if params.Active != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "active", runtime.ParamLocationQuery, *params.Active); err != nil {
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

		if params.WikiChecksumFailed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "wiki_checksum_failed", runtime.ParamLocationQuery, *params.WikiChecksumFailed); err != nil {
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

		if params.RepositoryChecksumFailed != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "repository_checksum_failed", runtime.ParamLocationQuery, *params.RepositoryChecksumFailed); err != nil {
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

		if params.IncludeHidden != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "include_hidden", runtime.ParamLocationQuery, *params.IncludeHidden); err != nil {
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

		if params.Simple != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "simple", runtime.ParamLocationQuery, *params.Simple); err != nil {
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

		if params.Statistics != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "statistics", runtime.ParamLocationQuery, *params.Statistics); err != nil {
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
func (r PostApiV4UserRunnersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4UserRunnersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UserCountsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UserCountsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsersIdEventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsersIdEventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsersUserIdContributedProjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsersUserIdContributedProjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsersUserIdProjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsersUserIdProjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4UsersUserIdStarredProjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4UsersUserIdStarredProjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4UserRunnersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UserRunnersResponse, error) {
	rsp, err := c.PostApiV4UserRunnersWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UserRunnersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4UserRunnersWithResponse(ctx context.Context, body PostApiV4UserRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UserRunnersResponse, error) {
	rsp, err := c.PostApiV4UserRunners(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4UserRunnersResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UserCountsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UserCountsResponse, error) {
	rsp, err := c.GetApiV4UserCounts(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UserCountsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsersIdEventsWithResponse(ctx context.Context, id string, params *GetApiV4UsersIdEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersIdEventsResponse, error) {
	rsp, err := c.GetApiV4UsersIdEvents(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsersIdEventsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsersUserIdContributedProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdContributedProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdContributedProjectsResponse, error) {
	rsp, err := c.GetApiV4UsersUserIdContributedProjects(ctx, userId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsersUserIdContributedProjectsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsersUserIdProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdProjectsResponse, error) {
	rsp, err := c.GetApiV4UsersUserIdProjects(ctx, userId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsersUserIdProjectsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4UsersUserIdStarredProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdStarredProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdStarredProjectsResponse, error) {
	rsp, err := c.GetApiV4UsersUserIdStarredProjects(ctx, userId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4UsersUserIdStarredProjectsResponse(rsp)
}
func ParsePostApiV4UserRunnersResponse(rsp *http.Response) (*PostApiV4UserRunnersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4UserRunnersResponse{
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
func ParseGetApiV4UserCountsResponse(rsp *http.Response) (*GetApiV4UserCountsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UserCountsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AssignedIssues               *int32 `json:"assigned_issues,omitempty"`
			AssignedMergeRequests        *int32 `json:"assigned_merge_requests,omitempty"`
			MergeRequests                *int32 `json:"merge_requests,omitempty"`
			ReviewRequestedMergeRequests *int32 `json:"review_requested_merge_requests,omitempty"`
			Todos                        *int32 `json:"todos,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4UsersIdEventsResponse(rsp *http.Response) (*GetApiV4UsersIdEventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsersIdEventsResponse{
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
func ParseGetApiV4UsersUserIdContributedProjectsResponse(rsp *http.Response) (*GetApiV4UsersUserIdContributedProjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsersUserIdContributedProjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
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
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4UsersUserIdProjectsResponse(rsp *http.Response) (*GetApiV4UsersUserIdProjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsersUserIdProjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
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
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4UsersUserIdStarredProjectsResponse(rsp *http.Response) (*GetApiV4UsersUserIdStarredProjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4UsersUserIdStarredProjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
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
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
