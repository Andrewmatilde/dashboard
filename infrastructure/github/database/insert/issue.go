package insert

import (
	"database/sql"
	"fmt"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/util"
	util2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/util"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions"
	model2 "github.com/PingCAP-QE/dashboard/infrastructure/github/processing/versions/model"
)

// insert.Issue insert data into table ISSUE
func Issue(db *sql.DB, repo *model.Repository, issue *model.Issue) {
	closeAt := util.GetTimeOrNull(issue.Closed, issue.ClosedAt)
	closeAtDate := util.GetDateOrNull(issue.Closed, issue.ClosedAt)
	createAtDate := util2.ParseDate(issue.CreatedAt)
	_, err := db.Exec(`
insert into issue 
(id,number, repository_id, closed, closed_at, 
 closed_week, created_at, created_week, title, url) 
values (?,?,?,?,?,
date_add(?, interval 7 - weekday(?) day),?,date_add(?, interval 7 - weekday(?) day),?,?);`,
		issue.DatabaseID, issue.Number, repo.DatabaseID, issue.Closed,
		closeAt, closeAtDate, closeAtDate, issue.CreatedAt, createAtDate, createAtDate, issue.Title, issue.Url)
	if err != nil {
		fmt.Println("Insert fail while insert into issue (id,number, repository_id, closed, closed_at, created_at, title, url) ", err)
	}
}

func IssueLabel(db *sql.DB, repo *model.Repository, issue *model.Issue, label *model.Label) {
	_, err := db.Exec(`
insert into issue_label (issue_id,label_id)
select ?, label.id
from label where label.name = ? and
                 repository_id = ?;`,
		issue.DatabaseID, label.Name, repo.DatabaseID)
	if err != nil {
		fmt.Println("Insert fail while insert into issue_label (issue_id,label_id)", err)
	}
}

func IssueVersion(db *sql.DB, issue *model.Issue, body *string) {
	affectedVersions, fixedVersions, err := versions.GetVersions(body)
	if err != nil {
		return
	}

	for _, version := range affectedVersions {
		err := issueVersionAffected(db, issue, &version)
		if err != nil {
			fmt.Println(err)
			fmt.Println("insert fail issue Number = ", issue.Number)
			fmt.Println("###########################")
		}
	}
	for _, version := range fixedVersions {
		err := issueVersionFixed(db, issue, &version)
		if err != nil {
			fmt.Println(err)
			fmt.Println("insert fail issue Number = ", issue.Number)
			fmt.Println("###########################")
		}
	}
}

func issueVersionAffected(db *sql.DB, issue *model.Issue, version *model2.Version) error {
	versionDatabaseID, err := util.GenIDFromVersion(*version)
	if err != nil {
		return nil
	}
	_, err = db.Exec(`
insert into issue_version_affected 
(issue_id, version_id)
values (?,?)`,
		issue.DatabaseID, versionDatabaseID)
	if err != nil {
		return fmt.Errorf(" insert fail while insert into issue_version_affected (issue_id, version_id) , %v", err)
	}
	return nil
}

func issueVersionFixed(db *sql.DB, issue *model.Issue, version *model2.Version) error {
	versionDatabaseID, err := util.GenIDFromVersion(*version)
	if err != nil {
		return nil
	}
	_, err = db.Exec(`
insert into issue_version_fixed
(issue_id, version_id)
values (?,?);`,
		issue.DatabaseID, versionDatabaseID)
	if err != nil {
		return fmt.Errorf(" insert fail while insert into issue_version_fixed (issue_id, version_id), %v", err)
	}
	return nil
}
