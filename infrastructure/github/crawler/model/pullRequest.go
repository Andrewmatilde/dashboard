package model

import "time"

type PullRequestConnection struct {
	Nodes      []*PullRequest `json:"nodes"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type PullRequest struct {
	Assignees  *UserConnection         `json:"assignees"`
	Body       string                  `json:"body"`
	Closed     bool                    `json:"closed"`
	ClosedAt   *time.Time              `json:"closedAt"`
	CreatedAt  time.Time               `json:"createdAt"`
	Comments   *IssueCommentConnection `json:"comments"`
	DatabaseID int                     `json:"databaseId"`
	Merged     bool                    `json:"merged"`
	MergedAt   *time.Time              `json:"mergedAt"`
	Labels     *LabelConnection        `json:"labels"`
	Number     int                     `json:"number"`
	Title      string                  `json:"title"`
	Url        string                  `json:"url"`
}
