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

type PostApiV4UserRunnersJSONBody struct {
	// AccessLevel The access level of the runner
	AccessLevel *string `json:"access_level,omitempty" jsonschema:"description=The access level of the runner,enum=not_protected,enum=ref_protected"`

	// Description Description of the runner
	Description *string `json:"description,omitempty" jsonschema:"description=Description of the runner"`

	// GroupId The ID of the group that the runner is created in
	GroupId int32 `json:"group_id" jsonschema:"description=The ID of the group that the runner is created in,format=int32"`

	// Locked Specifies if the runner should be locked for the current project (defaults to false)
	Locked *bool `json:"locked,omitempty" jsonschema:"description=Specifies if the runner should be locked for the current project (defaults to false)"`

	// MaintenanceNote Free-form maintenance notes for the runner (1024 characters)
	MaintenanceNote *string `json:"maintenance_note,omitempty" jsonschema:"description=Free-form maintenance notes for the runner (1024 characters)"`

	// MaximumTimeout Maximum timeout that limits the amount of time (in seconds) that runners can run jobs
	MaximumTimeout *int32 `json:"maximum_timeout,omitempty" jsonschema:"description=Maximum timeout that limits the amount of time (in seconds) that runners can run jobs,format=int32"`

	// Paused Specifies if the runner should ignore new jobs (defaults to false)
	Paused *bool `json:"paused,omitempty" jsonschema:"description=Specifies if the runner should ignore new jobs (defaults to false)"`

	// ProjectId The ID of the project that the runner is created in
	ProjectId int32 `json:"project_id" jsonschema:"description=The ID of the project that the runner is created in,format=int32"`

	// RunUntagged Specifies if the runner should handle untagged jobs  (defaults to true)
	RunUntagged *bool `json:"run_untagged,omitempty" jsonschema:"description=Specifies if the runner should handle untagged jobs (defaults to true)"`

	// RunnerType Specifies the scope of the runner
	RunnerType string `json:"runner_type" jsonschema:"description=Specifies the scope of the runner,enum=instance_type,enum=group_type,enum=project_type"`

	// TagList A list of runner tags
	TagList *[]string `json:"tag_list,omitempty" jsonschema:"description=A list of runner tags"`
}
type GetApiV4UsersIdEventsParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty" jsonschema:"description=Current page number,default=desc,format=date"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty" jsonschema:"description=Number of items per page,default=desc,format=date"`

	// Action Event action to filter on
	Action *string `form:"action,omitempty" json:"action,omitempty" jsonschema:"description=Event action to filter on"`

	// TargetType Event target type to filter on
	TargetType *string `form:"target_type,omitempty" json:"target_type,omitempty" jsonschema:"description=Event target type to filter on,enum=issue,enum=milestone,enum=merge_request,enum=note,enum=project,enum=snippet,enum=user,enum=wiki,enum=design"`

	// Before Include only events created before this date
	Before *openapi_types.Date `form:"before,omitempty" json:"before,omitempty" jsonschema:"description=Include only events created before this date,format=date"`

	// After Include only events created after this date
	After *openapi_types.Date `form:"after,omitempty" json:"after,omitempty" jsonschema:"description=Include only events created after this date,format=date"`

	// Sort Return events sorted in ascending and descending order
	Sort *string `form:"sort,omitempty" json:"sort,omitempty" jsonschema:"description=Return events sorted in ascending and descending order,default=desc,enum=asc,enum=desc"`
}
type GetApiV4UsersUserIdContributedProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty" jsonschema:"description=Return projects ordered by field. storage_size\\, repository_size\\, wiki_size\\, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.,default=false,enum=id,enum=name,enum=path,enum=created_at,enum=updated_at,enum=last_activity_at,enum=similarity,enum=star_count,enum=storage_size,enum=repository_size,enum=wiki_size,enum=packages_size"`

	// Sort Return projects sorted in ascending and descending order
	Sort *string `form:"sort,omitempty" json:"sort,omitempty" jsonschema:"description=Return projects sorted in ascending and descending order,default=false,enum=asc,enum=desc"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty" jsonschema:"description=Current page number,default=false,format=int32"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty" jsonschema:"description=Number of items per page,default=false,format=int32"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty" jsonschema:"description=Return only the ID\\, URL\\, name\\, and path of each project,default=false"`
}
type GetApiV4UsersUserIdProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty" jsonschema:"description=Return projects ordered by field. storage_size\\, repository_size\\, wiki_size\\, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.,default=false,enum=id,enum=name,enum=path,enum=created_at,enum=updated_at,enum=last_activity_at,enum=similarity,enum=star_count,enum=storage_size,enum=repository_size,enum=wiki_size,enum=packages_size"`

	// Sort Return projects sorted in ascending and descending order
	Sort *string `form:"sort,omitempty" json:"sort,omitempty" jsonschema:"description=Return projects sorted in ascending and descending order,default=false,enum=asc,enum=desc"`

	// Archived Limit by archived status
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty" jsonschema:"description=Limit by archived status"`

	// Visibility Limit by visibility
	Visibility *string `form:"visibility,omitempty" json:"visibility,omitempty" jsonschema:"description=Limit by visibility,enum=private,enum=internal,enum=public"`

	// Search Return list of projects matching the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty" jsonschema:"description=Return list of projects matching the search criteria"`

	// SearchNamespaces Include ancestor namespaces when matching search criteria
	SearchNamespaces *bool `form:"search_namespaces,omitempty" json:"search_namespaces,omitempty" jsonschema:"description=Include ancestor namespaces when matching search criteria"`

	// Owned Limit by owned by authenticated user
	Owned *bool `form:"owned,omitempty" json:"owned,omitempty" jsonschema:"description=Limit by owned by authenticated user,default=false"`

	// Starred Limit by starred status
	Starred *bool `form:"starred,omitempty" json:"starred,omitempty" jsonschema:"description=Limit by starred status,default=false"`

	// Imported Limit by imported by authenticated user
	Imported *bool `form:"imported,omitempty" json:"imported,omitempty" jsonschema:"description=Limit by imported by authenticated user,default=false"`

	// Membership Limit by projects that the current user is a member of
	Membership *bool `form:"membership,omitempty" json:"membership,omitempty" jsonschema:"description=Limit by projects that the current user is a member of,default=false"`

	// WithIssuesEnabled Limit by enabled issues feature
	WithIssuesEnabled *bool `form:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty" jsonschema:"description=Limit by enabled issues feature,default=false"`

	// WithMergeRequestsEnabled Limit by enabled merge requests feature
	WithMergeRequestsEnabled *bool `form:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty" jsonschema:"description=Limit by enabled merge requests feature,default=false"`

	// WithProgrammingLanguage Limit to repositories which use the given programming language
	WithProgrammingLanguage *string `form:"with_programming_language,omitempty" json:"with_programming_language,omitempty" jsonschema:"description=Limit to repositories which use the given programming language"`

	// MinAccessLevel Limit by minimum access level of authenticated user
	MinAccessLevel *int32 `form:"min_access_level,omitempty" json:"min_access_level,omitempty" jsonschema:"description=Limit by minimum access level of authenticated user,format=int32,enum=10,enum=15,enum=20,enum=30,enum=40,enum=50"`

	// IdAfter Limit results to projects with IDs greater than the specified ID
	IdAfter *int32 `form:"id_after,omitempty" json:"id_after,omitempty" jsonschema:"description=Limit results to projects with IDs greater than the specified ID,format=int32"`

	// IdBefore Limit results to projects with IDs less than the specified ID
	IdBefore *int32 `form:"id_before,omitempty" json:"id_before,omitempty" jsonschema:"description=Limit results to projects with IDs less than the specified ID,format=int32"`

	// LastActivityAfter Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityAfter *time.Time `form:"last_activity_after,omitempty" json:"last_activity_after,omitempty" jsonschema:"description=Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// LastActivityBefore Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityBefore *time.Time `form:"last_activity_before,omitempty" json:"last_activity_before,omitempty" jsonschema:"description=Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// RepositoryStorage Which storage shard the repository is on. Available only to admins
	RepositoryStorage *string `form:"repository_storage,omitempty" json:"repository_storage,omitempty" jsonschema:"description=Which storage shard the repository is on. Available only to admins"`

	// Topic Comma-separated list of topics. Limit results to projects having all topics
	Topic *[]string `form:"topic,omitempty" json:"topic,omitempty" jsonschema:"description=Comma-separated list of topics. Limit results to projects having all topics"`

	// TopicId Limit results to projects with the assigned topic given by the topic ID
	TopicId *int32 `form:"topic_id,omitempty" json:"topic_id,omitempty" jsonschema:"description=Limit results to projects with the assigned topic given by the topic ID,format=int32"`

	// UpdatedBefore Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedBefore *time.Time `form:"updated_before,omitempty" json:"updated_before,omitempty" jsonschema:"description=Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// UpdatedAfter Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedAfter *time.Time `form:"updated_after,omitempty" json:"updated_after,omitempty" jsonschema:"description=Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// IncludePendingDelete Include projects in pending delete state. Can only be set by admins
	IncludePendingDelete *bool `form:"include_pending_delete,omitempty" json:"include_pending_delete,omitempty" jsonschema:"description=Include projects in pending delete state. Can only be set by admins"`

	// MarkedForDeletionOn Date when the project was marked for deletion
	MarkedForDeletionOn *openapi_types.Date `form:"marked_for_deletion_on,omitempty" json:"marked_for_deletion_on,omitempty" jsonschema:"description=Date when the project was marked for deletion,format=int32"`

	// Active Limit by projects that are not archived and not marked for deletion
	Active *bool `form:"active,omitempty" json:"active,omitempty" jsonschema:"description=Limit by projects that are not archived and not marked for deletion"`

	// WikiChecksumFailed Limit by projects where wiki checksum is failed
	WikiChecksumFailed *bool `form:"wiki_checksum_failed,omitempty" json:"wiki_checksum_failed,omitempty" jsonschema:"description=Limit by projects where wiki checksum is failed,default=false"`

	// RepositoryChecksumFailed Limit by projects where repository checksum is failed
	RepositoryChecksumFailed *bool `form:"repository_checksum_failed,omitempty" json:"repository_checksum_failed,omitempty" jsonschema:"description=Limit by projects where repository checksum is failed,default=false"`

	// IncludeHidden Include hidden projects. Can only be set by admins
	IncludeHidden *bool `form:"include_hidden,omitempty" json:"include_hidden,omitempty" jsonschema:"description=Include hidden projects. Can only be set by admins,default=false"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty" jsonschema:"description=Current page number,default=false,format=int32"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty" jsonschema:"description=Number of items per page,default=false,format=int32"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty" jsonschema:"description=Return only the ID\\, URL\\, name\\, and path of each project,default=false"`

	// Statistics Include project statistics
	Statistics *bool `form:"statistics,omitempty" json:"statistics,omitempty" jsonschema:"description=Include project statistics,default=false"`

	// WithCustomAttributes Include custom attributes in the response
	WithCustomAttributes *bool `form:"with_custom_attributes,omitempty" json:"with_custom_attributes,omitempty" jsonschema:"description=Include custom attributes in the response,default=false"`
}
type GetApiV4UsersUserIdStarredProjectsParams struct {
	// OrderBy Return projects ordered by field. storage_size, repository_size, wiki_size, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty" jsonschema:"description=Return projects ordered by field. storage_size\\, repository_size\\, wiki_size\\, packages_size are only available to admins. Similarity is available when searching and is limited to projects the user has access to.,default=false,enum=id,enum=name,enum=path,enum=created_at,enum=updated_at,enum=last_activity_at,enum=similarity,enum=star_count,enum=storage_size,enum=repository_size,enum=wiki_size,enum=packages_size"`

	// Sort Return projects sorted in ascending and descending order
	Sort *string `form:"sort,omitempty" json:"sort,omitempty" jsonschema:"description=Return projects sorted in ascending and descending order,default=false,enum=asc,enum=desc"`

	// Archived Limit by archived status
	Archived *bool `form:"archived,omitempty" json:"archived,omitempty" jsonschema:"description=Limit by archived status"`

	// Visibility Limit by visibility
	Visibility *string `form:"visibility,omitempty" json:"visibility,omitempty" jsonschema:"description=Limit by visibility,enum=private,enum=internal,enum=public"`

	// Search Return list of projects matching the search criteria
	Search *string `form:"search,omitempty" json:"search,omitempty" jsonschema:"description=Return list of projects matching the search criteria"`

	// SearchNamespaces Include ancestor namespaces when matching search criteria
	SearchNamespaces *bool `form:"search_namespaces,omitempty" json:"search_namespaces,omitempty" jsonschema:"description=Include ancestor namespaces when matching search criteria"`

	// Owned Limit by owned by authenticated user
	Owned *bool `form:"owned,omitempty" json:"owned,omitempty" jsonschema:"description=Limit by owned by authenticated user,default=false"`

	// Starred Limit by starred status
	Starred *bool `form:"starred,omitempty" json:"starred,omitempty" jsonschema:"description=Limit by starred status,default=false"`

	// Imported Limit by imported by authenticated user
	Imported *bool `form:"imported,omitempty" json:"imported,omitempty" jsonschema:"description=Limit by imported by authenticated user,default=false"`

	// Membership Limit by projects that the current user is a member of
	Membership *bool `form:"membership,omitempty" json:"membership,omitempty" jsonschema:"description=Limit by projects that the current user is a member of,default=false"`

	// WithIssuesEnabled Limit by enabled issues feature
	WithIssuesEnabled *bool `form:"with_issues_enabled,omitempty" json:"with_issues_enabled,omitempty" jsonschema:"description=Limit by enabled issues feature,default=false"`

	// WithMergeRequestsEnabled Limit by enabled merge requests feature
	WithMergeRequestsEnabled *bool `form:"with_merge_requests_enabled,omitempty" json:"with_merge_requests_enabled,omitempty" jsonschema:"description=Limit by enabled merge requests feature,default=false"`

	// WithProgrammingLanguage Limit to repositories which use the given programming language
	WithProgrammingLanguage *string `form:"with_programming_language,omitempty" json:"with_programming_language,omitempty" jsonschema:"description=Limit to repositories which use the given programming language"`

	// MinAccessLevel Limit by minimum access level of authenticated user
	MinAccessLevel *int32 `form:"min_access_level,omitempty" json:"min_access_level,omitempty" jsonschema:"description=Limit by minimum access level of authenticated user,format=int32,enum=10,enum=15,enum=20,enum=30,enum=40,enum=50"`

	// IdAfter Limit results to projects with IDs greater than the specified ID
	IdAfter *int32 `form:"id_after,omitempty" json:"id_after,omitempty" jsonschema:"description=Limit results to projects with IDs greater than the specified ID,format=int32"`

	// IdBefore Limit results to projects with IDs less than the specified ID
	IdBefore *int32 `form:"id_before,omitempty" json:"id_before,omitempty" jsonschema:"description=Limit results to projects with IDs less than the specified ID,format=int32"`

	// LastActivityAfter Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityAfter *time.Time `form:"last_activity_after,omitempty" json:"last_activity_after,omitempty" jsonschema:"description=Limit results to projects with last_activity after specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// LastActivityBefore Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	LastActivityBefore *time.Time `form:"last_activity_before,omitempty" json:"last_activity_before,omitempty" jsonschema:"description=Limit results to projects with last_activity before specified time. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// RepositoryStorage Which storage shard the repository is on. Available only to admins
	RepositoryStorage *string `form:"repository_storage,omitempty" json:"repository_storage,omitempty" jsonschema:"description=Which storage shard the repository is on. Available only to admins"`

	// Topic Comma-separated list of topics. Limit results to projects having all topics
	Topic *[]string `form:"topic,omitempty" json:"topic,omitempty" jsonschema:"description=Comma-separated list of topics. Limit results to projects having all topics"`

	// TopicId Limit results to projects with the assigned topic given by the topic ID
	TopicId *int32 `form:"topic_id,omitempty" json:"topic_id,omitempty" jsonschema:"description=Limit results to projects with the assigned topic given by the topic ID,format=int32"`

	// UpdatedBefore Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedBefore *time.Time `form:"updated_before,omitempty" json:"updated_before,omitempty" jsonschema:"description=Return projects updated before the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// UpdatedAfter Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
	UpdatedAfter *time.Time `form:"updated_after,omitempty" json:"updated_after,omitempty" jsonschema:"description=Return projects updated after the specified datetime. Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ,format=int32"`

	// IncludePendingDelete Include projects in pending delete state. Can only be set by admins
	IncludePendingDelete *bool `form:"include_pending_delete,omitempty" json:"include_pending_delete,omitempty" jsonschema:"description=Include projects in pending delete state. Can only be set by admins"`

	// MarkedForDeletionOn Date when the project was marked for deletion
	MarkedForDeletionOn *openapi_types.Date `form:"marked_for_deletion_on,omitempty" json:"marked_for_deletion_on,omitempty" jsonschema:"description=Date when the project was marked for deletion,format=int32"`

	// Active Limit by projects that are not archived and not marked for deletion
	Active *bool `form:"active,omitempty" json:"active,omitempty" jsonschema:"description=Limit by projects that are not archived and not marked for deletion"`

	// WikiChecksumFailed Limit by projects where wiki checksum is failed
	WikiChecksumFailed *bool `form:"wiki_checksum_failed,omitempty" json:"wiki_checksum_failed,omitempty" jsonschema:"description=Limit by projects where wiki checksum is failed,default=false"`

	// RepositoryChecksumFailed Limit by projects where repository checksum is failed
	RepositoryChecksumFailed *bool `form:"repository_checksum_failed,omitempty" json:"repository_checksum_failed,omitempty" jsonschema:"description=Limit by projects where repository checksum is failed,default=false"`

	// IncludeHidden Include hidden projects. Can only be set by admins
	IncludeHidden *bool `form:"include_hidden,omitempty" json:"include_hidden,omitempty" jsonschema:"description=Include hidden projects. Can only be set by admins,default=false"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty" jsonschema:"description=Current page number,default=false,format=int32"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty" jsonschema:"description=Number of items per page,default=false,format=int32"`

	// Simple Return only the ID, URL, name, and path of each project
	Simple *bool `form:"simple,omitempty" json:"simple,omitempty" jsonschema:"description=Return only the ID\\, URL\\, name\\, and path of each project,default=false"`

	// Statistics Include project statistics
	Statistics *bool `form:"statistics,omitempty" json:"statistics,omitempty" jsonschema:"description=Include project statistics,default=false"`
}
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
		AssignedIssues               *int32 `json:"assigned_issues,omitempty" jsonschema:",format=int32"`
		AssignedMergeRequests        *int32 `json:"assigned_merge_requests,omitempty" jsonschema:",format=int32"`
		MergeRequests                *int32 `json:"merge_requests,omitempty" jsonschema:",format=int32"`
		ReviewRequestedMergeRequests *int32 `json:"review_requested_merge_requests,omitempty" jsonschema:",format=int32"`
		Todos                        *int32 `json:"todos,omitempty" jsonschema:",format=int32"`
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
			Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"author,omitempty" jsonschema:"description=API_Entities_UserBasic model"`
		AuthorId       *int32  `json:"author_id,omitempty" jsonschema:",format=int32"`
		AuthorUsername *string `json:"author_username,omitempty"`
		CreatedAt      *string `json:"created_at,omitempty"`
		Id             *int32  `json:"id,omitempty" jsonschema:",format=int32"`
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
				Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty" jsonschema:"description=API_Entities_UserBasic model"`
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
				Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"resolved_by,omitempty" jsonschema:"description=API_Entities_UserBasic model"`

			// Suggestions API_Entities_Suggestion model
			Suggestions *struct {
				Appliable   *string `json:"appliable,omitempty"`
				Applied     *string `json:"applied,omitempty"`
				FromContent *string `json:"from_content,omitempty"`
				FromLine    *string `json:"from_line,omitempty"`
				Id          *string `json:"id,omitempty"`
				ToContent   *string `json:"to_content,omitempty"`
				ToLine      *string `json:"to_line,omitempty"`
			} `json:"suggestions,omitempty" jsonschema:"description=API_Entities_Suggestion model"`
			System    *string `json:"system,omitempty"`
			Type      *string `json:"type,omitempty"`
			UpdatedAt *string `json:"updated_at,omitempty"`
		} `json:"note,omitempty"`
		ProjectId *int32 `json:"project_id,omitempty" jsonschema:",format=int32"`
		PushData  *struct {
			Action      *string `json:"action,omitempty"`
			CommitCount *int32  `json:"commit_count,omitempty" jsonschema:",format=int32"`
			CommitFrom  *string `json:"commit_from,omitempty"`
			CommitTitle *string `json:"commit_title,omitempty"`
			CommitTo    *string `json:"commit_to,omitempty"`
			Ref         *string `json:"ref,omitempty"`
			RefCount    *int32  `json:"ref_count,omitempty" jsonschema:",format=int32"`
			RefType     *string `json:"ref_type,omitempty"`
		} `json:"push_data,omitempty"`
		TargetId    *int32  `json:"target_id,omitempty" jsonschema:",format=int32"`
		TargetIid   *int32  `json:"target_iid,omitempty" jsonschema:",format=int32"`
		TargetTitle *string `json:"target_title,omitempty"`
		TargetType  *string `json:"target_type,omitempty"`

		// WikiPage API_Entities_WikiPageBasic model
		WikiPage *struct {
			Format         *string `json:"format,omitempty"`
			Slug           *string `json:"slug,omitempty"`
			Title          *string `json:"title,omitempty"`
			WikiPageMetaId *int32  `json:"wiki_page_meta_id,omitempty" jsonschema:",format=int32"`
		} `json:"wiki_page,omitempty" jsonschema:"description=API_Entities_WikiPageBasic model"`
	}
}
type GetApiV4UsersUserIdContributedProjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		AvatarUrl *string    `json:"avatar_url,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

		// CustomAttributes API_Entities_CustomAttribute model
		CustomAttributes *struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
		DefaultBranch  *string    `json:"default_branch,omitempty"`
		Description    *string    `json:"description,omitempty"`
		ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
		HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
		Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
		LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
			Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
			Kind      *string `json:"kind,omitempty"`
			Name      *string `json:"name,omitempty"`
			ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
			Path      *string `json:"path,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"namespace,omitempty"`
		Path              *string   `json:"path,omitempty"`
		PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
		ReadmeUrl         *string   `json:"readme_url,omitempty"`
		RepositoryStorage *string   `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
		StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
		CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

		// CustomAttributes API_Entities_CustomAttribute model
		CustomAttributes *struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
		DefaultBranch  *string    `json:"default_branch,omitempty"`
		Description    *string    `json:"description,omitempty"`
		ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
		HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
		Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
		LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
			Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
			Kind      *string `json:"kind,omitempty"`
			Name      *string `json:"name,omitempty"`
			ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
			Path      *string `json:"path,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"namespace,omitempty"`
		Path              *string   `json:"path,omitempty"`
		PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
		ReadmeUrl         *string   `json:"readme_url,omitempty"`
		RepositoryStorage *string   `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
		StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
		CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

		// CustomAttributes API_Entities_CustomAttribute model
		CustomAttributes *struct {
			Key   *string `json:"key,omitempty"`
			Value *string `json:"value,omitempty"`
		} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
		DefaultBranch  *string    `json:"default_branch,omitempty"`
		Description    *string    `json:"description,omitempty"`
		ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
		HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
		Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
		LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
			Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
			Kind      *string `json:"kind,omitempty"`
			Name      *string `json:"name,omitempty"`
			ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
			Path      *string `json:"path,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"namespace,omitempty"`
		Path              *string   `json:"path,omitempty"`
		PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
		ReadmeUrl         *string   `json:"readme_url,omitempty"`
		RepositoryStorage *string   `json:"repository_storage,omitempty"`
		SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
		StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
			AssignedIssues               *int32 `json:"assigned_issues,omitempty" jsonschema:",format=int32"`
			AssignedMergeRequests        *int32 `json:"assigned_merge_requests,omitempty" jsonschema:",format=int32"`
			MergeRequests                *int32 `json:"merge_requests,omitempty" jsonschema:",format=int32"`
			ReviewRequestedMergeRequests *int32 `json:"review_requested_merge_requests,omitempty" jsonschema:",format=int32"`
			Todos                        *int32 `json:"todos,omitempty" jsonschema:",format=int32"`
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
				Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"author,omitempty" jsonschema:"description=API_Entities_UserBasic model"`
			AuthorId       *int32  `json:"author_id,omitempty" jsonschema:",format=int32"`
			AuthorUsername *string `json:"author_username,omitempty"`
			CreatedAt      *string `json:"created_at,omitempty"`
			Id             *int32  `json:"id,omitempty" jsonschema:",format=int32"`
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
					Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
					Locked      *bool   `json:"locked,omitempty"`
					Name        *string `json:"name,omitempty"`
					PublicEmail *string `json:"public_email,omitempty"`
					State       *string `json:"state,omitempty"`
					Username    *string `json:"username,omitempty"`
					WebUrl      *string `json:"web_url,omitempty"`
				} `json:"author,omitempty" jsonschema:"description=API_Entities_UserBasic model"`
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
					Id          *int32  `json:"id,omitempty" jsonschema:",format=int32"`
					Locked      *bool   `json:"locked,omitempty"`
					Name        *string `json:"name,omitempty"`
					PublicEmail *string `json:"public_email,omitempty"`
					State       *string `json:"state,omitempty"`
					Username    *string `json:"username,omitempty"`
					WebUrl      *string `json:"web_url,omitempty"`
				} `json:"resolved_by,omitempty" jsonschema:"description=API_Entities_UserBasic model"`

				// Suggestions API_Entities_Suggestion model
				Suggestions *struct {
					Appliable   *string `json:"appliable,omitempty"`
					Applied     *string `json:"applied,omitempty"`
					FromContent *string `json:"from_content,omitempty"`
					FromLine    *string `json:"from_line,omitempty"`
					Id          *string `json:"id,omitempty"`
					ToContent   *string `json:"to_content,omitempty"`
					ToLine      *string `json:"to_line,omitempty"`
				} `json:"suggestions,omitempty" jsonschema:"description=API_Entities_Suggestion model"`
				System    *string `json:"system,omitempty"`
				Type      *string `json:"type,omitempty"`
				UpdatedAt *string `json:"updated_at,omitempty"`
			} `json:"note,omitempty"`
			ProjectId *int32 `json:"project_id,omitempty" jsonschema:",format=int32"`
			PushData  *struct {
				Action      *string `json:"action,omitempty"`
				CommitCount *int32  `json:"commit_count,omitempty" jsonschema:",format=int32"`
				CommitFrom  *string `json:"commit_from,omitempty"`
				CommitTitle *string `json:"commit_title,omitempty"`
				CommitTo    *string `json:"commit_to,omitempty"`
				Ref         *string `json:"ref,omitempty"`
				RefCount    *int32  `json:"ref_count,omitempty" jsonschema:",format=int32"`
				RefType     *string `json:"ref_type,omitempty"`
			} `json:"push_data,omitempty"`
			TargetId    *int32  `json:"target_id,omitempty" jsonschema:",format=int32"`
			TargetIid   *int32  `json:"target_iid,omitempty" jsonschema:",format=int32"`
			TargetTitle *string `json:"target_title,omitempty"`
			TargetType  *string `json:"target_type,omitempty"`

			// WikiPage API_Entities_WikiPageBasic model
			WikiPage *struct {
				Format         *string `json:"format,omitempty"`
				Slug           *string `json:"slug,omitempty"`
				Title          *string `json:"title,omitempty"`
				WikiPageMetaId *int32  `json:"wiki_page_meta_id,omitempty" jsonschema:",format=int32"`
			} `json:"wiki_page,omitempty" jsonschema:"description=API_Entities_WikiPageBasic model"`
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
			CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

			// CustomAttributes API_Entities_CustomAttribute model
			CustomAttributes *struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
			DefaultBranch  *string    `json:"default_branch,omitempty"`
			Description    *string    `json:"description,omitempty"`
			ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
			HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
			Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
			LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
				Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespace,omitempty"`
			Path              *string   `json:"path,omitempty"`
			PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
			ReadmeUrl         *string   `json:"readme_url,omitempty"`
			RepositoryStorage *string   `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
			StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
			CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

			// CustomAttributes API_Entities_CustomAttribute model
			CustomAttributes *struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
			DefaultBranch  *string    `json:"default_branch,omitempty"`
			Description    *string    `json:"description,omitempty"`
			ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
			HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
			Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
			LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
				Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespace,omitempty"`
			Path              *string   `json:"path,omitempty"`
			PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
			ReadmeUrl         *string   `json:"readme_url,omitempty"`
			RepositoryStorage *string   `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
			StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
			CreatedAt *time.Time `json:"created_at,omitempty" jsonschema:",format=date-time"`

			// CustomAttributes API_Entities_CustomAttribute model
			CustomAttributes *struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty" jsonschema:"description=API_Entities_CustomAttribute model"`
			DefaultBranch  *string    `json:"default_branch,omitempty"`
			Description    *string    `json:"description,omitempty"`
			ForksCount     *int32     `json:"forks_count,omitempty" jsonschema:",format=int32"`
			HttpUrlToRepo  *string    `json:"http_url_to_repo,omitempty"`
			Id             *int32     `json:"id,omitempty" jsonschema:",format=int32"`
			LastActivityAt *time.Time `json:"last_activity_at,omitempty" jsonschema:",format=date-time"`
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
				Id        *int32  `json:"id,omitempty" jsonschema:",format=int32"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty" jsonschema:",format=int32"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespace,omitempty"`
			Path              *string   `json:"path,omitempty"`
			PathWithNamespace *string   `json:"path_with_namespace,omitempty"`
			ReadmeUrl         *string   `json:"readme_url,omitempty"`
			RepositoryStorage *string   `json:"repository_storage,omitempty"`
			SshUrlToRepo      *string   `json:"ssh_url_to_repo,omitempty"`
			StarCount         *int32    `json:"star_count,omitempty" jsonschema:",format=int32"`
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
