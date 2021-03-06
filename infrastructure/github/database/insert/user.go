package insert

import (
	"database/sql"
	"fmt"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
)

func User(db *sql.DB, user *model.User) {
	_, err := db.Exec(`
insert into user (id,login,email) values (?,?,?) on duplicate key update login=?;`,
		user.DatabaseID, user.Login, user.Email, user.Login)
	if err != nil {
		fmt.Println("Insert fail while insert into insert into user (id,login,email) ", err)
	}
}

func UserIssue(db *sql.DB, issue *model.Issue, user *model.User) {
	_, err := db.Exec(`
insert into user_issue (user_id, issue_id)
select user.id,?
from user where user.login = ?;`,
		issue.DatabaseID, user.Login)
	if err != nil {
		fmt.Println("insert into user_issue (user_id, issue_id)", err)
	}
}

func UserPullRequest(db *sql.DB, pull_request *model.PullRequest, user *model.User) {
	_, err := db.Exec(`
insert into user_pull_request (user_id, pull_request_id)
select user.id,?
from user where user.login = ?;`,
		pull_request.DatabaseID, user.Login)
	if err != nil {
		fmt.Println("insert into user_pull_request (user_id, pull_request_id)", err)
	}
}

func ActorPullRequest(db *sql.DB, pull_request *model.PullRequest, actor *model.Actor) {
	if (actor.Login == "") || (actor.DatabaseID == nil) {
		fmt.Printf("%v has not enough info\n", actor)
		return
	}
	_, err := db.Exec(`
insert into user (id,login,email) values (?,?,?) on duplicate key update login=?;`,
		actor.DatabaseID, actor.Login, actor.Email, actor.Login)
	if err != nil {
		fmt.Println("Insert fail while insert into insert into user (id,login,email) ", err)
	}

	_, err = db.Exec(`
insert into user_pull_request (user_id, pull_request_id)
select user.id,?
from user where user.login = ?;`,
		pull_request.DatabaseID, actor.Login)
	if err != nil {
		fmt.Println("insert into user_pull_request (user_id, pull_request_id)", err)
	}
}
