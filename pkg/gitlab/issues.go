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
)

type GetApiV4IssuesParams struct {
	// AssigneeId Return issues assigned to the given user id. Mutually exclusive with assignee_username. None returns unassigned issues. Any returns issues with an assignee.
	AssigneeId *int `form:"assignee_id,omitempty" json:"assignee_id,omitempty"`

	// AssigneeUsername Return issues assigned to the given username. Similar to assignee_id and mutually exclusive with assignee_id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
	AssigneeUsername *[]string `form:"assignee_username,omitempty" json:"assignee_username,omitempty"`

	// AuthorId Return issues created by the given user id. Mutually exclusive with author_username. Combine with scope=all or scope=assigned_to_me.
	AuthorId *int `form:"author_id,omitempty" json:"author_id,omitempty"`

	// AuthorUsername Return issues created by the given username. Similar to author_id and mutually exclusive with author_id.
	AuthorUsername *string `form:"author_username,omitempty" json:"author_username,omitempty"`

	// Confidential Filter confidential or public issues.
	Confidential *bool `form:"confidential,omitempty" json:"confidential,omitempty"`

	// CreatedAfter Return issues created on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedAfter *string `form:"created_after,omitempty" json:"created_after,omitempty"`

	// CreatedBefore Return issues created on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	CreatedBefore *string `form:"created_before,omitempty" json:"created_before,omitempty"`

	// DueDate Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), any, today, tomorrow, overdue, week, month, next_month_and_previous_two_weeks.
	DueDate *string `form:"due_date,omitempty" json:"due_date,omitempty"`

	// EpicId Return issues associated with the given epic ID. None returns issues that are not associated with an epic. Any returns issues that are associated with an epic. Premium and Ultimate only.
	EpicId *int `form:"epic_id,omitempty" json:"epic_id,omitempty"`

	// HealthStatus Return issues with the specified health_status. (Introduced in GitLab 15.4). In GitLab 15.5 and later, None returns issues with no health status assigned, and Any returns issues with a health status assigned. Ultimate only.
	HealthStatus *string `form:"health_status,omitempty" json:"health_status,omitempty"`

	// Iids Return only the issues having the given iid.
	Iids *[]int `form:"iids[],omitempty" json:"iids[],omitempty"`

	// In Modify the scope of the search attribute. title, description, or a string joining them with comma. Default is title,description.
	In *string `form:"in,omitempty" json:"in,omitempty"`

	// IssueType Filter to a given type of issue. One of issue, incident, test_case or task.
	IssueType *string `form:"issue_type,omitempty" json:"issue_type,omitempty"`

	// IterationId Return issues assigned to the given iteration ID. None returns issues that do not belong to an iteration. Any returns issues that belong to an iteration. Mutually exclusive with iteration_title. Premium and Ultimate only.
	IterationId *int `form:"iteration_id,omitempty" json:"iteration_id,omitempty"`

	// IterationTitle Return issues assigned to the iteration with the given title. Similar to iteration_id and mutually exclusive with iteration_id. Premium and Ultimate only.
	IterationTitle *string `form:"iteration_title,omitempty" json:"iteration_title,omitempty"`

	// Labels Comma-separated list of label names, issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. No+Label (Deprecated) lists all issues with no labels. Predefined names are case-insensitive.
	Labels *string `form:"labels,omitempty" json:"labels,omitempty"`

	// MilestoneId Returns issues assigned to milestones with a given timebox value (None, Any, Upcoming, and Started). None lists all issues with no milestone. Any lists all issues that have an assigned milestone. Upcoming lists all issues assigned to milestones due in the future. Started lists all issues assigned to open, started milestones. The logic for Upcoming and Started differs from the logic used in the GraphQL API. milestone and milestone_id are mutually exclusive.
	MilestoneId *string `form:"milestone_id,omitempty" json:"milestone_id,omitempty"`

	// Milestone The milestone title. None lists all issues with no milestone. Any lists all issues that have an assigned milestone. Using None or Any will be deprecated in the future. Use milestone_id attribute instead. milestone and milestone_id are mutually exclusive.
	Milestone *string `form:"milestone,omitempty" json:"milestone,omitempty"`

	// MyReactionEmoji Return issues reacted by the authenticated user by the given emoji. None returns issues not given a reaction. Any returns issues given at least one reaction.
	MyReactionEmoji *string `form:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`

	// NonArchived Return issues only from non-archived projects. If false, the response returns issues from both archived and non-archived projects. Default is true.
	NonArchived *bool `form:"non_archived,omitempty" json:"non_archived,omitempty"`

	// Not Return issues that do not match the parameters supplied. Accepts: assignee_id, assignee_username, author_id, author_username, iids, iteration_id, iteration_title, labels, milestone, milestone_id and weight.
	Not *string `form:"not,omitempty" json:"not,omitempty"`

	// OrderBy Return issues ordered by created_at, due_date, label_priority, milestone_due, popularity, priority, relative_position, title, updated_at, or weight fields. Default is created_at.
	OrderBy *string `form:"order_by,omitempty" json:"order_by,omitempty"`

	// Scope Return issues for the given scope: created_by_me, assigned_to_me or all. Defaults to created_by_me.
	Scope *string `form:"scope,omitempty" json:"scope,omitempty"`

	// Search Search issues against their title and description.
	Search *string `form:"search,omitempty" json:"search,omitempty"`

	// Sort Return issues sorted in asc or desc order. Default is desc.
	Sort *string `form:"sort,omitempty" json:"sort,omitempty"`

	// State Return all issues or just those that are opened or closed.
	State *string `form:"state,omitempty" json:"state,omitempty"`

	// UpdatedAfter Return issues updated on or after the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedAfter *string `form:"updated_after,omitempty" json:"updated_after,omitempty"`

	// UpdatedBefore Return issues updated on or before the given time. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	UpdatedBefore *string `form:"updated_before,omitempty" json:"updated_before,omitempty"`

	// Weight Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned. Premium and Ultimate only.
	Weight *int `form:"weight,omitempty" json:"weight,omitempty"`

	// WithLabelsDetails If true, the response returns more details for each label in labels field: :name, :color, :description, :description_html, :text_color. Default is false.
	WithLabelsDetails *bool `form:"with_labels_details,omitempty" json:"with_labels_details,omitempty"`
}
type GetApiV4IssuesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *[]struct {
		Links *struct {
			AwardEmoji          *string `json:"award_emoji,omitempty"`
			ClosedAsDuplicateOf *string `json:"closed_as_duplicate_of,omitempty"`
			Notes               *string `json:"notes,omitempty"`
			Project             *string `json:"project,omitempty"`
			Self                *string `json:"self,omitempty"`
		} `json:"_links,omitempty"`
		Assignee *struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"assignee,omitempty"`
		Assignees *[]struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"assignees,omitempty"`
		Author *struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		ClosedAt           *string   `json:"closed_at,omitempty"`
		ClosedBy           *string   `json:"closed_by,omitempty"`
		Confidential       *bool     `json:"confidential,omitempty"`
		CreatedAt          *string   `json:"created_at,omitempty"`
		Description        *string   `json:"description,omitempty"`
		DiscussionLocked   *bool     `json:"discussion_locked,omitempty"`
		Downvotes          *int      `json:"downvotes,omitempty"`
		DueDate            *string   `json:"due_date,omitempty"`
		HasTasks           *bool     `json:"has_tasks,omitempty"`
		Id                 *int      `json:"id,omitempty"`
		Iid                *int      `json:"iid,omitempty"`
		Imported           *bool     `json:"imported,omitempty"`
		ImportedFrom       *string   `json:"imported_from,omitempty"`
		IssueType          *string   `json:"issue_type,omitempty"`
		Labels             *[]string `json:"labels,omitempty"`
		MergeRequestsCount *int      `json:"merge_requests_count,omitempty"`
		Milestone          *struct {
			CreatedAt   *string `json:"created_at,omitempty"`
			Description *string `json:"description,omitempty"`
			DueDate     *string `json:"due_date,omitempty"`
			Id          *int    `json:"id,omitempty"`
			Iid         *int    `json:"iid,omitempty"`
			ProjectId   *int    `json:"project_id,omitempty"`
			State       *string `json:"state,omitempty"`
			Title       *string `json:"title,omitempty"`
			UpdatedAt   *string `json:"updated_at,omitempty"`
		} `json:"milestone,omitempty"`
		MovedToId  *string `json:"moved_to_id,omitempty"`
		ProjectId  *int    `json:"project_id,omitempty"`
		References *struct {
			Full     *string `json:"full,omitempty"`
			Relative *string `json:"relative,omitempty"`
			Short    *string `json:"short,omitempty"`
		} `json:"references,omitempty"`
		Severity             *string `json:"severity,omitempty"`
		State                *string `json:"state,omitempty"`
		TaskCompletionStatus *struct {
			CompletedCount *int `json:"completed_count,omitempty"`
			Count          *int `json:"count,omitempty"`
		} `json:"task_completion_status,omitempty"`
		TaskStatus *string `json:"task_status,omitempty"`
		TimeStats  *struct {
			HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
			HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
			TimeEstimate        *int    `json:"time_estimate,omitempty"`
			TotalTimeSpent      *int    `json:"total_time_spent,omitempty"`
		} `json:"time_stats,omitempty"`
		Title          *string `json:"title,omitempty"`
		Type           *string `json:"type,omitempty"`
		UpdatedAt      *string `json:"updated_at,omitempty"`
		Upvotes        *int    `json:"upvotes,omitempty"`
		UserNotesCount *int    `json:"user_notes_count,omitempty"`
		WebUrl         *string `json:"web_url,omitempty"`
	}
}
type GetApiV4IssuesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON2XX      *struct {
		Links *struct {
			AwardEmoji          *string `json:"award_emoji,omitempty"`
			ClosedAsDuplicateOf *string `json:"closed_as_duplicate_of,omitempty"`
			Notes               *string `json:"notes,omitempty"`
			Project             *string `json:"project,omitempty"`
			Self                *string `json:"self,omitempty"`
		} `json:"_links,omitempty"`
		Assignee *struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"assignee,omitempty"`
		Assignees *[]struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"assignees,omitempty"`
		Author *struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			Id        *int    `json:"id,omitempty"`
			Name      *string `json:"name,omitempty"`
			State     *string `json:"state,omitempty"`
			Username  *string `json:"username,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"author,omitempty"`
		ClosedAt           *string   `json:"closed_at,omitempty"`
		ClosedBy           *string   `json:"closed_by,omitempty"`
		Confidential       *bool     `json:"confidential,omitempty"`
		CreatedAt          *string   `json:"created_at,omitempty"`
		Description        *string   `json:"description,omitempty"`
		DiscussionLocked   *bool     `json:"discussion_locked,omitempty"`
		Downvotes          *int      `json:"downvotes,omitempty"`
		DueDate            *string   `json:"due_date,omitempty"`
		HasTasks           *bool     `json:"has_tasks,omitempty"`
		Id                 *int      `json:"id,omitempty"`
		Iid                *int      `json:"iid,omitempty"`
		Imported           *bool     `json:"imported,omitempty"`
		ImportedFrom       *string   `json:"imported_from,omitempty"`
		IssueType          *string   `json:"issue_type,omitempty"`
		Labels             *[]string `json:"labels,omitempty"`
		MergeRequestsCount *int      `json:"merge_requests_count,omitempty"`
		Milestone          *struct {
			ClosedAt    *string `json:"closed_at,omitempty"`
			CreatedAt   *string `json:"created_at,omitempty"`
			Description *string `json:"description,omitempty"`
			DueDate     *string `json:"due_date,omitempty"`
			Id          *int    `json:"id,omitempty"`
			Iid         *int    `json:"iid,omitempty"`
			ProjectId   *int    `json:"project_id,omitempty"`
			State       *string `json:"state,omitempty"`
			Title       *string `json:"title,omitempty"`
			UpdatedAt   *string `json:"updated_at,omitempty"`
		} `json:"milestone,omitempty"`
		MovedToId  *string `json:"moved_to_id,omitempty"`
		References *struct {
			Full     *string `json:"full,omitempty"`
			Relative *string `json:"relative,omitempty"`
			Short    *string `json:"short,omitempty"`
		} `json:"references,omitempty"`
		ServiceDeskReplyTo   *string `json:"service_desk_reply_to,omitempty"`
		Severity             *string `json:"severity,omitempty"`
		State                *string `json:"state,omitempty"`
		Subscribed           *bool   `json:"subscribed,omitempty"`
		TaskCompletionStatus *struct {
			CompletedCount *int `json:"completed_count,omitempty"`
			Count          *int `json:"count,omitempty"`
		} `json:"task_completion_status,omitempty"`
		TimeStats *struct {
			HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
			HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
			TimeEstimate        *int    `json:"time_estimate,omitempty"`
			TotalTimeSpent      *int    `json:"total_time_spent,omitempty"`
		} `json:"time_stats,omitempty"`
		Title          *string `json:"title,omitempty"`
		Type           *string `json:"type,omitempty"`
		UpdatedAt      *string `json:"updated_at,omitempty"`
		Upvotes        *int    `json:"upvotes,omitempty"`
		UserNotesCount *int    `json:"user_notes_count,omitempty"`
		WebUrl         *string `json:"web_url,omitempty"`
		Weight         *string `json:"weight,omitempty"`
	}
}

func (c *Client) GetApiV4Issues(ctx context.Context, params *GetApiV4IssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4IssuesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4IssuesId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4IssuesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4IssuesRequest(server string, params *GetApiV4IssuesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/issues")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "assignee_username", runtime.ParamLocationQuery, *params.AssigneeUsername); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

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

		if params.Confidential != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "confidential", runtime.ParamLocationQuery, *params.Confidential); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

		if params.DueDate != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "due_date", runtime.ParamLocationQuery, *params.DueDate); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.EpicId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "epic_id", runtime.ParamLocationQuery, *params.EpicId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.HealthStatus != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "health_status", runtime.ParamLocationQuery, *params.HealthStatus); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Iids != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "iids[]", runtime.ParamLocationQuery, *params.Iids); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

		if params.IssueType != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "issue_type", runtime.ParamLocationQuery, *params.IssueType); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.IterationId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "iteration_id", runtime.ParamLocationQuery, *params.IterationId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.IterationTitle != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "iteration_title", runtime.ParamLocationQuery, *params.IterationTitle); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "labels", runtime.ParamLocationQuery, *params.Labels); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.MilestoneId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "milestone_id", runtime.ParamLocationQuery, *params.MilestoneId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

		if params.NonArchived != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "non_archived", runtime.ParamLocationQuery, *params.NonArchived); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Not != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "not", runtime.ParamLocationQuery, *params.Not); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

		if params.Weight != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "weight", runtime.ParamLocationQuery, *params.Weight); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4IssuesIdRequest(server string, id int) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/issues/%s", pathParam0)
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
func (r GetApiV4IssuesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4IssuesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4IssuesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4IssuesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4IssuesWithResponse(ctx context.Context, params *GetApiV4IssuesParams, reqEditors ...RequestEditorFn) (*GetApiV4IssuesResponse, error) {
	rsp, err := c.GetApiV4Issues(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4IssuesResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4IssuesIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetApiV4IssuesIdResponse, error) {
	rsp, err := c.GetApiV4IssuesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4IssuesIdResponse(rsp)
}
func ParseGetApiV4IssuesResponse(rsp *http.Response) (*GetApiV4IssuesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4IssuesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest []struct {
			Links *struct {
				AwardEmoji          *string `json:"award_emoji,omitempty"`
				ClosedAsDuplicateOf *string `json:"closed_as_duplicate_of,omitempty"`
				Notes               *string `json:"notes,omitempty"`
				Project             *string `json:"project,omitempty"`
				Self                *string `json:"self,omitempty"`
			} `json:"_links,omitempty"`
			Assignee *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"assignee,omitempty"`
			Assignees *[]struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"assignees,omitempty"`
			Author *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			ClosedAt           *string   `json:"closed_at,omitempty"`
			ClosedBy           *string   `json:"closed_by,omitempty"`
			Confidential       *bool     `json:"confidential,omitempty"`
			CreatedAt          *string   `json:"created_at,omitempty"`
			Description        *string   `json:"description,omitempty"`
			DiscussionLocked   *bool     `json:"discussion_locked,omitempty"`
			Downvotes          *int      `json:"downvotes,omitempty"`
			DueDate            *string   `json:"due_date,omitempty"`
			HasTasks           *bool     `json:"has_tasks,omitempty"`
			Id                 *int      `json:"id,omitempty"`
			Iid                *int      `json:"iid,omitempty"`
			Imported           *bool     `json:"imported,omitempty"`
			ImportedFrom       *string   `json:"imported_from,omitempty"`
			IssueType          *string   `json:"issue_type,omitempty"`
			Labels             *[]string `json:"labels,omitempty"`
			MergeRequestsCount *int      `json:"merge_requests_count,omitempty"`
			Milestone          *struct {
				CreatedAt   *string `json:"created_at,omitempty"`
				Description *string `json:"description,omitempty"`
				DueDate     *string `json:"due_date,omitempty"`
				Id          *int    `json:"id,omitempty"`
				Iid         *int    `json:"iid,omitempty"`
				ProjectId   *int    `json:"project_id,omitempty"`
				State       *string `json:"state,omitempty"`
				Title       *string `json:"title,omitempty"`
				UpdatedAt   *string `json:"updated_at,omitempty"`
			} `json:"milestone,omitempty"`
			MovedToId  *string `json:"moved_to_id,omitempty"`
			ProjectId  *int    `json:"project_id,omitempty"`
			References *struct {
				Full     *string `json:"full,omitempty"`
				Relative *string `json:"relative,omitempty"`
				Short    *string `json:"short,omitempty"`
			} `json:"references,omitempty"`
			Severity             *string `json:"severity,omitempty"`
			State                *string `json:"state,omitempty"`
			TaskCompletionStatus *struct {
				CompletedCount *int `json:"completed_count,omitempty"`
				Count          *int `json:"count,omitempty"`
			} `json:"task_completion_status,omitempty"`
			TaskStatus *string `json:"task_status,omitempty"`
			TimeStats  *struct {
				HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
				HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
				TimeEstimate        *int    `json:"time_estimate,omitempty"`
				TotalTimeSpent      *int    `json:"total_time_spent,omitempty"`
			} `json:"time_stats,omitempty"`
			Title          *string `json:"title,omitempty"`
			Type           *string `json:"type,omitempty"`
			UpdatedAt      *string `json:"updated_at,omitempty"`
			Upvotes        *int    `json:"upvotes,omitempty"`
			UserNotesCount *int    `json:"user_notes_count,omitempty"`
			WebUrl         *string `json:"web_url,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}
func ParseGetApiV4IssuesIdResponse(rsp *http.Response) (*GetApiV4IssuesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4IssuesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2:
		var dest struct {
			Links *struct {
				AwardEmoji          *string `json:"award_emoji,omitempty"`
				ClosedAsDuplicateOf *string `json:"closed_as_duplicate_of,omitempty"`
				Notes               *string `json:"notes,omitempty"`
				Project             *string `json:"project,omitempty"`
				Self                *string `json:"self,omitempty"`
			} `json:"_links,omitempty"`
			Assignee *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"assignee,omitempty"`
			Assignees *[]struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"assignees,omitempty"`
			Author *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				Id        *int    `json:"id,omitempty"`
				Name      *string `json:"name,omitempty"`
				State     *string `json:"state,omitempty"`
				Username  *string `json:"username,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"author,omitempty"`
			ClosedAt           *string   `json:"closed_at,omitempty"`
			ClosedBy           *string   `json:"closed_by,omitempty"`
			Confidential       *bool     `json:"confidential,omitempty"`
			CreatedAt          *string   `json:"created_at,omitempty"`
			Description        *string   `json:"description,omitempty"`
			DiscussionLocked   *bool     `json:"discussion_locked,omitempty"`
			Downvotes          *int      `json:"downvotes,omitempty"`
			DueDate            *string   `json:"due_date,omitempty"`
			HasTasks           *bool     `json:"has_tasks,omitempty"`
			Id                 *int      `json:"id,omitempty"`
			Iid                *int      `json:"iid,omitempty"`
			Imported           *bool     `json:"imported,omitempty"`
			ImportedFrom       *string   `json:"imported_from,omitempty"`
			IssueType          *string   `json:"issue_type,omitempty"`
			Labels             *[]string `json:"labels,omitempty"`
			MergeRequestsCount *int      `json:"merge_requests_count,omitempty"`
			Milestone          *struct {
				ClosedAt    *string `json:"closed_at,omitempty"`
				CreatedAt   *string `json:"created_at,omitempty"`
				Description *string `json:"description,omitempty"`
				DueDate     *string `json:"due_date,omitempty"`
				Id          *int    `json:"id,omitempty"`
				Iid         *int    `json:"iid,omitempty"`
				ProjectId   *int    `json:"project_id,omitempty"`
				State       *string `json:"state,omitempty"`
				Title       *string `json:"title,omitempty"`
				UpdatedAt   *string `json:"updated_at,omitempty"`
			} `json:"milestone,omitempty"`
			MovedToId  *string `json:"moved_to_id,omitempty"`
			References *struct {
				Full     *string `json:"full,omitempty"`
				Relative *string `json:"relative,omitempty"`
				Short    *string `json:"short,omitempty"`
			} `json:"references,omitempty"`
			ServiceDeskReplyTo   *string `json:"service_desk_reply_to,omitempty"`
			Severity             *string `json:"severity,omitempty"`
			State                *string `json:"state,omitempty"`
			Subscribed           *bool   `json:"subscribed,omitempty"`
			TaskCompletionStatus *struct {
				CompletedCount *int `json:"completed_count,omitempty"`
				Count          *int `json:"count,omitempty"`
			} `json:"task_completion_status,omitempty"`
			TimeStats *struct {
				HumanTimeEstimate   *string `json:"human_time_estimate,omitempty"`
				HumanTotalTimeSpent *string `json:"human_total_time_spent,omitempty"`
				TimeEstimate        *int    `json:"time_estimate,omitempty"`
				TotalTimeSpent      *int    `json:"total_time_spent,omitempty"`
			} `json:"time_stats,omitempty"`
			Title          *string `json:"title,omitempty"`
			Type           *string `json:"type,omitempty"`
			UpdatedAt      *string `json:"updated_at,omitempty"`
			Upvotes        *int    `json:"upvotes,omitempty"`
			UserNotesCount *int    `json:"user_notes_count,omitempty"`
			WebUrl         *string `json:"web_url,omitempty"`
			Weight         *string `json:"weight,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON2XX = &dest

	}

	return response, nil
}
