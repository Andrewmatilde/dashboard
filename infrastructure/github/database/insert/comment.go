package insert

import (
	"database/sql"
	"fmt"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
)

func Comment(db *sql.DB, issue *model.Issue, comment *model.IssueComment) {
	_, err := db.Exec(`
insert into comment (id, issue_id, body) values (?,?,?);`,
		comment.DatabaseID, issue.DatabaseID, comment.Body)
	if err != nil {
		fmt.Println("Insert fail while insert into comment (id, issue_id, body) ", err)
	}
}

func PrComment(db *sql.DB, pull_request *model.PullRequest, comment *model.IssueComment) {
	_, err := db.Exec(`
insert into pull_request_comment (id, pull_request_id, body) values (?,?,?);`,
		comment.DatabaseID, pull_request.DatabaseID, comment.Body)
	if err != nil {
		fmt.Println("Insert fail while insert into pull_request_comment (id, pull_request_id, body) ", err)
	}
}
