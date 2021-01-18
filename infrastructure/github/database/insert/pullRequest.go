package insert

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/util"
	util2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/util"
)

// insert.Issue insert data into table ISSUE
func PullRequest(db *sql.DB, repo *model.Repository, pull_request *model.PullRequest) {
	closeAt := util.GetIssueClosedTime(pull_request.Closed, pull_request.ClosedAt)
	closeAtWeek := util.GetIssueClosedWeek(pull_request.Closed, pull_request.ClosedAt)
	createAtWeek := util2.ParseDate(pull_request.CreatedAt)
	_, err := db.Exec(`
insert into pull_request 
(id,number, repository_id, closed, closed_at, 
 closed_week, created_at, created_week, title, url) 
values (?,?,?,?,?,
date_add(?, interval 7 - weekday(?) day),?,date_add(?, interval 7 - weekday(?) day),?,?);`,
		pull_request.DatabaseID, pull_request.Number, repo.DatabaseID, pull_request.Closed,
		closeAt, closeAtWeek, closeAtWeek, pull_request.CreatedAt, createAtWeek, createAtWeek, pull_request.Title, pull_request.Url)
	if err != nil {
		fmt.Println("Insert fail while insert into pull_request (id,number, repository_id, closed, closed_at, created_at, title, url) ", err)
	}
}

func PullRequestLabel(db *sql.DB, repo *model.Repository, pull_request *model.PullRequest, label *model.Label) {
	_, err := db.Exec(`
insert into pull_request_label (pull_request_id,label_id)
select ?, label.id
from label where label.name = ? and
                 repository_id = ?;`,
		pull_request.DatabaseID, label.Name, repo.DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into pull_request_label (pull_request_id,label_id)", err)
	}
}
