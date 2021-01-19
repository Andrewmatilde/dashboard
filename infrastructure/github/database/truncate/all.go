package truncate

import (
	"database/sql"
	"log"
)

func ClearTableByName(db *sql.DB, name string) {
	_, err := db.Exec(`truncate table ` + name + `;`)
	if err != nil {
		log.Printf("truncate table "+name+" fail: %v", err)
	}
}

func ClearTablesByName(db *sql.DB, tableNameList []string) {
	for _, name := range tableNameList {
		ClearTableByName(db, name)
	}
}

func AllClear(db *sql.DB) {
	tablesNmaeList := []string{
		"comment",
		"coverage_timeline",
		"issue",
		"issue_label",
		"issue_team",
		"issue_version_affected",
		"issue_version_fixed",
		"label",
		"label_severity_weight",
		"label_sig",
		"pull_request",
		"pull_request_comment",
		"pull_request_label",
		"pull_request_team",
		"repository",
		"tag",
		"team",
		"team_bug_jail",
		"team_issue",
		"timeline",
		"timeline_repository",
		"user",
		"user_bug_jail",
		"user_issue",
		"user_pull_request",
		"version",
		"week_line",
	}
	ClearTablesByName(db, tablesNmaeList)
}
