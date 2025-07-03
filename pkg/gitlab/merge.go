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

type GetApiV4MergeRequestsParams struct {
	// AuthorId Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`.
	AuthorId *int32 `form:"author_id,omitempty" json:"author_id,omitempty"`

	// AuthorUsername Returns merge requests created by the given `username`. Mutually exclusive with `author_id`.
	AuthorUsername *string `form:"author_username,omitempty" json:"author_username,omitempty"`

	// AssigneeId Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee.
	AssigneeId *int32 `form:"assignee_id,omitempty" json:"assignee_id,omitempty"`

	// AssigneeUsername Returns merge requests created by the given `username`. Mutually exclusive with `author_id`.
	AssigneeUsername *[]string `form:"assignee_username,omitempty" json:"assignee_username,omitempty"`

	// ReviewerUsername Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8.
	ReviewerUsername *string `form:"reviewer_username,omitempty" json:"reviewer_username,omitempty"`

	// Labels Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive.
	Labels *[]string `form:"labels,omitempty" json:"labels,omitempty"`

	// Milestone Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone.
	Milestone *string `form:"milestone,omitempty" json:"milestone,omitempty"`

	// MyReactionEmoji Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction.
	MyReactionEmoji *string `form:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`

	// ReviewerId Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`.
	ReviewerId *int32 `form:"reviewer_id,omitempty" json:"reviewer_id,omitempty"`

	// State Returns `all` merge requests or just those that are `opened`, `closed`, `locked`, or `merged`.
	State *string `form:"state,omitempty" json:"state,omitempty"`

	// OrderBy Returns merge requests ordered by `created_at`, `label_priority`, `milestone_due`, `popularity`, `priority`, `title`, `updated_at` or `merged_at` fields. Introduced in GitLab 14.8.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Sort Returns merge requests sorted in `asc` or `desc` order.
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`

	// WithLabelsDetails If `true`, response returns more details for each label in labels field: `:name`,`:color`, `:description`, `:description_html`, `:text_color`
	WithLabelsDetails *bool `form:"with_labels_details,omitempty" json:"with_labels_details,omitempty"`

	// WithMergeStatusRecheck If `true`, this projection requests (but does not guarantee) that the `merge_status` field be recalculated asynchronously. Introduced in GitLab 13.0.
	WithMergeStatusRecheck *bool `form:"with_merge_status_recheck,omitempty" json:"with_merge_status_recheck,omitempty"`

	// CreatedAfter Returns merge requests created on or after the given time. Expected in ISO 8601 format.
	CreatedAfter *time.Time `form:"created_after,omitempty" json:"created_after,omitempty"`

	// CreatedBefore Returns merge requests created on or before the given time. Expected in ISO 8601 format.
	CreatedBefore *time.Time `form:"created_before,omitempty" json:"created_before,omitempty"`

	// UpdatedAfter Returns merge requests updated on or after the given time. Expected in ISO 8601 format.
	UpdatedAfter *time.Time `form:"updated_after,omitempty" json:"updated_after,omitempty"`

	// UpdatedBefore Returns merge requests updated on or before the given time. Expected in ISO 8601 format.
	UpdatedBefore *time.Time `form:"updated_before,omitempty" json:"updated_before,omitempty"`

	// View If simple, returns the `iid`, URL, title, description, and basic state of merge request
	View *string `form:"view,omitempty" json:"view,omitempty"`

	// Scope Returns merge requests for the given scope: `created_by_me`, `assigned_to_me` or `all`
	Scope *string `form:"scope,omitempty" json:"scope,omitempty"`

	// SourceBranch Returns merge requests with the given source branch
	SourceBranch *string `form:"source_branch,omitempty" json:"source_branch,omitempty"`

	// SourceProjectId Returns merge requests with the given source project id
	SourceProjectId *int32 `form:"source_project_id,omitempty" json:"source_project_id,omitempty"`

	// TargetBranch Returns merge requests with the given target branch
	TargetBranch *string `form:"target_branch,omitempty" json:"target_branch,omitempty"`

	// Search Search merge requests against their `title` and `description`.
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// In Modify the scope of the search attribute. `title`, `description`, or a string joining them with comma.
	In *string `form:"in,omitempty" json:"in,omitempty"`

	// Wip Filter merge requests against their `wip` status. `yes` to return only draft merge requests, `no` to return non-draft merge requests.
	Wip *string `form:"wip,omitempty" json:"wip,omitempty"`

	// NotAuthorId `<Negated>` Returns merge requests created by the given user `id`. Mutually exclusive with `author_username`. Combine with `scope=all` or `scope=assigned_to_me`.
	NotAuthorId *int32 `form:"not[author_id],omitempty" json:"not[author_id],omitempty"`

	// NotAuthorUsername `<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`.
	NotAuthorUsername *string `form:"not[author_username],omitempty" json:"not[author_username],omitempty"`

	// NotAssigneeId `<Negated>` Returns merge requests assigned to the given user `id`. `None` returns unassigned merge requests. `Any` returns merge requests with an assignee.
	NotAssigneeId *int32 `form:"not[assignee_id],omitempty" json:"not[assignee_id],omitempty"`

	// NotAssigneeUsername `<Negated>` Returns merge requests created by the given `username`. Mutually exclusive with `author_id`.
	NotAssigneeUsername *[]string `form:"not[assignee_username],omitempty" json:"not[assignee_username],omitempty"`

	// NotReviewerUsername `<Negated>` Returns merge requests which have the user as a reviewer with the given `username`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_id`. Introduced in GitLab 13.8.
	NotReviewerUsername *string `form:"not[reviewer_username],omitempty" json:"not[reviewer_username],omitempty"`

	// NotLabels `<Negated>` Returns merge requests matching a comma-separated list of labels. `None` lists all merge requests with no labels. `Any` lists all merge requests with at least one label. Predefined names are case-insensitive.
	NotLabels *[]string `form:"not[labels],omitempty" json:"not[labels],omitempty"`

	// NotMilestone `<Negated>` Returns merge requests for a specific milestone. `None` returns merge requests with no milestone. `Any` returns merge requests that have an assigned milestone.
	NotMilestone *string `form:"not[milestone],omitempty" json:"not[milestone],omitempty"`

	// NotMyReactionEmoji `<Negated>` Returns merge requests reacted by the authenticated user by the given `emoji`. `None` returns issues not given a reaction. `Any` returns issues given at least one reaction.
	NotMyReactionEmoji *string `form:"not[my_reaction_emoji],omitempty" json:"not[my_reaction_emoji],omitempty"`

	// NotReviewerId `<Negated>` Returns merge requests which have the user as a reviewer with the given user `id`. `None` returns merge requests with no reviewers. `Any` returns merge requests with any reviewer. Mutually exclusive with `reviewer_username`.
	NotReviewerId *int32 `form:"not[reviewer_id],omitempty" json:"not[reviewer_id],omitempty"`

	// DeployedBefore Returns merge requests deployed before the given date/time. Expected in ISO 8601 format.
	DeployedBefore *string `form:"deployed_before,omitempty" json:"deployed_before,omitempty"`

	// DeployedAfter Returns merge requests deployed after the given date/time. Expected in ISO 8601 format
	DeployedAfter *string `form:"deployed_after,omitempty" json:"deployed_after,omitempty"`

	// Environment Returns merge requests deployed to the given environment
	Environment *string `form:"environment,omitempty" json:"environment,omitempty"`

	// Approved Filters merge requests by their `approved` status. `yes` returns only approved merge requests. `no` returns only non-approved merge requests.
	Approved *string `form:"approved,omitempty" json:"approved,omitempty"`

	// MergeUserId Returns merge requests which have been merged by the user with the given user `id`. Mutually exclusive with `merge_user_username`.
	MergeUserId *int32 `form:"merge_user_id,omitempty" json:"merge_user_id,omitempty"`

	// MergeUserUsername Returns merge requests which have been merged by the user with the given `username`. Mutually exclusive with `merge_user_id`.
	MergeUserUsername *string `form:"merge_user_username,omitempty" json:"merge_user_username,omitempty"`

	// ApproverIds Return merge requests which have specified the users with the given IDs as an individual approver
	ApproverIds *string `form:"approver_ids,omitempty" json:"approver_ids,omitempty"`

	// ApprovedByIds Return merge requests which have been approved by the specified users with the given IDs
	ApprovedByIds *string `form:"approved_by_ids,omitempty" json:"approved_by_ids,omitempty"`

	// ApprovedByUsernames Return merge requests which have been approved by the specified users with the given
	//
	//
	//
	//
	//             usernames
	ApprovedByUsernames *string `form:"approved_by_usernames,omitempty" json:"approved_by_usernames,omitempty"`

	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type GetApiV4MergeRequestsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		AllowCollaboration    *string `json:"allow_collaboration,omitempty"`
		AllowMaintainerToPush *string `json:"allow_maintainer_to_push,omitempty"`
		ApprovalsBeforeMerge  *string `json:"approvals_before_merge,omitempty"`

		// Assignee API_Entities_UserBasic model
		Assignee *struct {
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
		} `json:"assignee,omitempty"`

		// Assignees API_Entities_UserBasic model
		Assignees *struct {
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
		} `json:"assignees,omitempty"`

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
		BlockingDiscussionsResolved *string `json:"blocking_discussions_resolved,omitempty"`
		ClosedAt                    *string `json:"closed_at,omitempty"`

		// ClosedBy API_Entities_UserBasic model
		ClosedBy *struct {
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
		} `json:"closed_by,omitempty"`
		CreatedAt               *time.Time `json:"created_at,omitempty"`
		Description             *string    `json:"description,omitempty"`
		DescriptionHtml         *string    `json:"description_html,omitempty"`
		DetailedMergeStatus     *string    `json:"detailed_merge_status,omitempty"`
		DiscussionLocked        *string    `json:"discussion_locked,omitempty"`
		Downvotes               *string    `json:"downvotes,omitempty"`
		Draft                   *string    `json:"draft,omitempty"`
		ForceRemoveSourceBranch *string    `json:"force_remove_source_branch,omitempty"`
		HasConflicts            *string    `json:"has_conflicts,omitempty"`
		Id                      *int32     `json:"id,omitempty"`
		Iid                     *int32     `json:"iid,omitempty"`
		Imported                *string    `json:"imported,omitempty"`
		ImportedFrom            *string    `json:"imported_from,omitempty"`
		Labels                  *string    `json:"labels,omitempty"`
		MergeAfter              *string    `json:"merge_after,omitempty"`
		MergeCommitSha          *string    `json:"merge_commit_sha,omitempty"`
		MergeStatus             *string    `json:"merge_status,omitempty"`

		// MergeUser API_Entities_UserBasic model
		MergeUser *struct {
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
		} `json:"merge_user,omitempty"`
		MergeWhenPipelineSucceeds *string `json:"merge_when_pipeline_succeeds,omitempty"`
		MergedAt                  *string `json:"merged_at,omitempty"`

		// MergedBy API_Entities_UserBasic model
		MergedBy *struct {
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
		} `json:"merged_by,omitempty"`
		Milestone *struct {
			CreatedAt   *string `json:"created_at,omitempty"`
			Description *string `json:"description,omitempty"`
			DueDate     *string `json:"due_date,omitempty"`
			Expired     *string `json:"expired,omitempty"`
			GroupId     *string `json:"group_id,omitempty"`
			Id          *string `json:"id,omitempty"`
			Iid         *string `json:"iid,omitempty"`
			ProjectId   *string `json:"project_id,omitempty"`
			StartDate   *string `json:"start_date,omitempty"`
			State       *string `json:"state,omitempty"`
			Title       *string `json:"title,omitempty"`
			UpdatedAt   *string `json:"updated_at,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"milestone,omitempty"`
		PreparedAt *string `json:"prepared_at,omitempty"`
		ProjectId  *int32  `json:"project_id,omitempty"`
		Reference  *string `json:"reference,omitempty"`
		References *struct {
			Full     *string `json:"full,omitempty"`
			Relative *string `json:"relative,omitempty"`
			Short    *string `json:"short,omitempty"`
		} `json:"references,omitempty"`

		// Reviewers API_Entities_UserBasic model
		Reviewers *struct {
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
		} `json:"reviewers,omitempty"`
		Sha                      *string `json:"sha,omitempty"`
		ShouldRemoveSourceBranch *string `json:"should_remove_source_branch,omitempty"`
		SourceBranch             *string `json:"source_branch,omitempty"`
		SourceProjectId          *string `json:"source_project_id,omitempty"`
		Squash                   *string `json:"squash,omitempty"`
		SquashCommitSha          *string `json:"squash_commit_sha,omitempty"`
		SquashOnMerge            *string `json:"squash_on_merge,omitempty"`
		State                    *string `json:"state,omitempty"`
		TargetBranch             *string `json:"target_branch,omitempty"`
		TargetProjectId          *string `json:"target_project_id,omitempty"`
		TaskCompletionStatus     *string `json:"task_completion_status,omitempty"`

		// TimeStats API_Entities_IssuableTimeStats model
		TimeStats *struct {
			HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
			HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
			TimeEstimate        *int32  `json:"time_estimate,omitempty"`
			TotalTimeSpent      *int32  `json:"total_time_spent,omitempty"`
		} `json:"time_stats,omitempty"`
		Title          *string    `json:"title,omitempty"`
		TitleHtml      *string    `json:"title_html,omitempty"`
		UpdatedAt      *time.Time `json:"updated_at,omitempty"`
		Upvotes        *string    `json:"upvotes,omitempty"`
		UserNotesCount *string    `json:"user_notes_count,omitempty"`
		WebUrl         *string    `json:"web_url,omitempty"`
		WorkInProgress *string    `json:"work_in_progress,omitempty"`
	}
}

func (c *Client) GetApiV4MergeRequests(ctx context.Context, params *GetApiV4MergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4MergeRequestsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4MergeRequestsRequest(server string, params *GetApiV4MergeRequestsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/merge_requests")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.AuthorId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "author_id", runtime.ParamLocationQuery, *params.AuthorId); err != nil {
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

		if params.AuthorUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "author_username", runtime.ParamLocationQuery, *params.AuthorUsername); err != nil {
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

		if params.AssigneeId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "assignee_id", runtime.ParamLocationQuery, *params.AssigneeId); err != nil {
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

		if params.AssigneeUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "assignee_username", runtime.ParamLocationQuery, *params.AssigneeUsername); err != nil {
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

		if params.ReviewerUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "reviewer_username", runtime.ParamLocationQuery, *params.ReviewerUsername); err != nil {
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

		if params.Labels != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "labels", runtime.ParamLocationQuery, *params.Labels); err != nil {
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

		if params.Milestone != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "milestone", runtime.ParamLocationQuery, *params.Milestone); err != nil {
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

		if params.MyReactionEmoji != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "my_reaction_emoji", runtime.ParamLocationQuery, *params.MyReactionEmoji); err != nil {
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

		if params.ReviewerId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "reviewer_id", runtime.ParamLocationQuery, *params.ReviewerId); err != nil {
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

		if params.State != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "state", runtime.ParamLocationQuery, *params.State); err != nil {
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

		if params.WithLabelsDetails != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_labels_details", runtime.ParamLocationQuery, *params.WithLabelsDetails); err != nil {
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

		if params.WithMergeStatusRecheck != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "with_merge_status_recheck", runtime.ParamLocationQuery, *params.WithMergeStatusRecheck); err != nil {
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

		if params.View != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "view", runtime.ParamLocationQuery, *params.View); err != nil {
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

		if params.SourceBranch != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "source_branch", runtime.ParamLocationQuery, *params.SourceBranch); err != nil {
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

		if params.SourceProjectId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "source_project_id", runtime.ParamLocationQuery, *params.SourceProjectId); err != nil {
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

		if params.TargetBranch != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "target_branch", runtime.ParamLocationQuery, *params.TargetBranch); err != nil {
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

		if params.In != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "in", runtime.ParamLocationQuery, *params.In); err != nil {
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

		if params.Wip != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "wip", runtime.ParamLocationQuery, *params.Wip); err != nil {
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

		if params.NotAuthorId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[author_id]", runtime.ParamLocationQuery, *params.NotAuthorId); err != nil {
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

		if params.NotAuthorUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[author_username]", runtime.ParamLocationQuery, *params.NotAuthorUsername); err != nil {
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

		if params.NotAssigneeId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[assignee_id]", runtime.ParamLocationQuery, *params.NotAssigneeId); err != nil {
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

		if params.NotAssigneeUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "not[assignee_username]", runtime.ParamLocationQuery, *params.NotAssigneeUsername); err != nil {
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

		if params.NotReviewerUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[reviewer_username]", runtime.ParamLocationQuery, *params.NotReviewerUsername); err != nil {
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

		if params.NotLabels != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "not[labels]", runtime.ParamLocationQuery, *params.NotLabels); err != nil {
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

		if params.NotMilestone != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[milestone]", runtime.ParamLocationQuery, *params.NotMilestone); err != nil {
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

		if params.NotMyReactionEmoji != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[my_reaction_emoji]", runtime.ParamLocationQuery, *params.NotMyReactionEmoji); err != nil {
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

		if params.NotReviewerId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not[reviewer_id]", runtime.ParamLocationQuery, *params.NotReviewerId); err != nil {
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

		if params.DeployedBefore != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "deployed_before", runtime.ParamLocationQuery, *params.DeployedBefore); err != nil {
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

		if params.DeployedAfter != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "deployed_after", runtime.ParamLocationQuery, *params.DeployedAfter); err != nil {
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

		if params.Environment != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "environment", runtime.ParamLocationQuery, *params.Environment); err != nil {
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

		if params.Approved != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "approved", runtime.ParamLocationQuery, *params.Approved); err != nil {
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

		if params.MergeUserId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "merge_user_id", runtime.ParamLocationQuery, *params.MergeUserId); err != nil {
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

		if params.MergeUserUsername != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "merge_user_username", runtime.ParamLocationQuery, *params.MergeUserUsername); err != nil {
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

		if params.ApproverIds != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "approver_ids", runtime.ParamLocationQuery, *params.ApproverIds); err != nil {
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

		if params.ApprovedByIds != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "approved_by_ids", runtime.ParamLocationQuery, *params.ApprovedByIds); err != nil {
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

		if params.ApprovedByUsernames != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "approved_by_usernames", runtime.ParamLocationQuery, *params.ApprovedByUsernames); err != nil {
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
func (r GetApiV4MergeRequestsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4MergeRequestsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4MergeRequestsWithResponse(ctx context.Context, params *GetApiV4MergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4MergeRequestsResponse, error) {
	rsp, err := c.GetApiV4MergeRequests(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4MergeRequestsResponse(rsp)
}
func ParseGetApiV4MergeRequestsResponse(rsp *http.Response) (*GetApiV4MergeRequestsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4MergeRequestsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			AllowCollaboration    *string `json:"allow_collaboration,omitempty"`
			AllowMaintainerToPush *string `json:"allow_maintainer_to_push,omitempty"`
			ApprovalsBeforeMerge  *string `json:"approvals_before_merge,omitempty"`

			// Assignee API_Entities_UserBasic model
			Assignee *struct {
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
			} `json:"assignee,omitempty"`

			// Assignees API_Entities_UserBasic model
			Assignees *struct {
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
			} `json:"assignees,omitempty"`

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
			BlockingDiscussionsResolved *string `json:"blocking_discussions_resolved,omitempty"`
			ClosedAt                    *string `json:"closed_at,omitempty"`

			// ClosedBy API_Entities_UserBasic model
			ClosedBy *struct {
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
			} `json:"closed_by,omitempty"`
			CreatedAt               *time.Time `json:"created_at,omitempty"`
			Description             *string    `json:"description,omitempty"`
			DescriptionHtml         *string    `json:"description_html,omitempty"`
			DetailedMergeStatus     *string    `json:"detailed_merge_status,omitempty"`
			DiscussionLocked        *string    `json:"discussion_locked,omitempty"`
			Downvotes               *string    `json:"downvotes,omitempty"`
			Draft                   *string    `json:"draft,omitempty"`
			ForceRemoveSourceBranch *string    `json:"force_remove_source_branch,omitempty"`
			HasConflicts            *string    `json:"has_conflicts,omitempty"`
			Id                      *int32     `json:"id,omitempty"`
			Iid                     *int32     `json:"iid,omitempty"`
			Imported                *string    `json:"imported,omitempty"`
			ImportedFrom            *string    `json:"imported_from,omitempty"`
			Labels                  *string    `json:"labels,omitempty"`
			MergeAfter              *string    `json:"merge_after,omitempty"`
			MergeCommitSha          *string    `json:"merge_commit_sha,omitempty"`
			MergeStatus             *string    `json:"merge_status,omitempty"`

			// MergeUser API_Entities_UserBasic model
			MergeUser *struct {
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
			} `json:"merge_user,omitempty"`
			MergeWhenPipelineSucceeds *string `json:"merge_when_pipeline_succeeds,omitempty"`
			MergedAt                  *string `json:"merged_at,omitempty"`

			// MergedBy API_Entities_UserBasic model
			MergedBy *struct {
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
			} `json:"merged_by,omitempty"`
			Milestone *struct {
				CreatedAt   *string `json:"created_at,omitempty"`
				Description *string `json:"description,omitempty"`
				DueDate     *string `json:"due_date,omitempty"`
				Expired     *string `json:"expired,omitempty"`
				GroupId     *string `json:"group_id,omitempty"`
				Id          *string `json:"id,omitempty"`
				Iid         *string `json:"iid,omitempty"`
				ProjectId   *string `json:"project_id,omitempty"`
				StartDate   *string `json:"start_date,omitempty"`
				State       *string `json:"state,omitempty"`
				Title       *string `json:"title,omitempty"`
				UpdatedAt   *string `json:"updated_at,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"milestone,omitempty"`
			PreparedAt *string `json:"prepared_at,omitempty"`
			ProjectId  *int32  `json:"project_id,omitempty"`
			Reference  *string `json:"reference,omitempty"`
			References *struct {
				Full     *string `json:"full,omitempty"`
				Relative *string `json:"relative,omitempty"`
				Short    *string `json:"short,omitempty"`
			} `json:"references,omitempty"`

			// Reviewers API_Entities_UserBasic model
			Reviewers *struct {
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
			} `json:"reviewers,omitempty"`
			Sha                      *string `json:"sha,omitempty"`
			ShouldRemoveSourceBranch *string `json:"should_remove_source_branch,omitempty"`
			SourceBranch             *string `json:"source_branch,omitempty"`
			SourceProjectId          *string `json:"source_project_id,omitempty"`
			Squash                   *string `json:"squash,omitempty"`
			SquashCommitSha          *string `json:"squash_commit_sha,omitempty"`
			SquashOnMerge            *string `json:"squash_on_merge,omitempty"`
			State                    *string `json:"state,omitempty"`
			TargetBranch             *string `json:"target_branch,omitempty"`
			TargetProjectId          *string `json:"target_project_id,omitempty"`
			TaskCompletionStatus     *string `json:"task_completion_status,omitempty"`

			// TimeStats API_Entities_IssuableTimeStats model
			TimeStats *struct {
				HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
				HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
				TimeEstimate        *int32  `json:"time_estimate,omitempty"`
				TotalTimeSpent      *int32  `json:"total_time_spent,omitempty"`
			} `json:"time_stats,omitempty"`
			Title          *string    `json:"title,omitempty"`
			TitleHtml      *string    `json:"title_html,omitempty"`
			UpdatedAt      *time.Time `json:"updated_at,omitempty"`
			Upvotes        *string    `json:"upvotes,omitempty"`
			UserNotesCount *string    `json:"user_notes_count,omitempty"`
			WebUrl         *string    `json:"web_url,omitempty"`
			WorkInProgress *string    `json:"work_in_progress,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
