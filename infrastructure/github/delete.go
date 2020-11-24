package github

import (
	"database/sql"
	"fmt"

	"dashboard/infrastructure/github/crawler/model"
)

// deleteIssueData delete data from issue
func deleteIssueData(db *sql.Tx, issue *model.Issue) {
	_, err := db.Exec(
		`DELETE from ISSUE where ISSUE.ID = ?;`, issue.DatabaseID)
	if err != nil {
		fmt.Println("Delete fail while DELETE from ISSUE where NUMBER = ?;", err)
	}
}
