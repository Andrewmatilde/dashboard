package github

import (
	"database/sql"
	"fmt"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/insert/team"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/process"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/PingCAP-QE/dashboard/infrastructure/github/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/client"
	crawlerConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/crawler/model"
	dbConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/database/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/insert"
	"github.com/PingCAP-QE/dashboard/infrastructure/github/database/truncate"
)

var db *sql.DB
var err error

// init Set mysql database connection
func initDB(c dbConfig.Config) {
	db, err = sql.Open("mysql", c.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		log.Panicf("open database fail: %v", err)
		return
	}
	fmt.Println("connect success")
}

// Fetch fetch all data
func FetchQuery(c crawlerConfig.Config, owner, name string) (model.Query, error) {
	client.InitClient(c)
	request := client.NewClient()
	opt := crawler.FetchOption{
		Owner:    owner,
		RepoName: name,
		IssueFilters: map[string]interface{}{
			"states": []string{"CLOSED", "OPEN"},
		},
	}

	totalData, err := crawler.FetchByRepoSafe(request, opt)
	return totalData, err
}

//insertData insert all the data fetched from database
func InsertQuery(db *sql.DB, totalData model.Query, owner string, c *config.Config) {
	fmt.Printf("Processing %s\n...", totalData.Repository.Name)
	var wg sync.WaitGroup

	fmt.Println("Inserting Repository...")
	insert.Repository(db, totalData.Repository, owner)

	fmt.Println("Inserting Timeline...")
	insert.Timeline(db, c)
	insert.WeekLine(db, c)
	insert.TimelineRepository(db, totalData.Repository)

	fmt.Println("Inserting User...")

	for _, user := range totalData.Repository.AssignableUsers.Nodes {
		wg.Add(1)
		go func(user *model.User) {
			insert.User(db, user)
			defer wg.Done()
		}(user)
	}
	wg.Wait()

	fmt.Println("Inserting Label...")
	for _, label := range totalData.Repository.Labels.Nodes {
		wg.Add(1)
		go func(label *model.Label) {
			insert.Label(db, totalData.Repository, label)
			insert.LabelSig(db, totalData.Repository, label)
			defer wg.Done()
		}(label)
	}
	wg.Wait()
	for _, weight := range c.LabelSeverityWeight {
		wg.Add(1)
		go func(weight config.LabelSeverityWeight) {
			insert.LabelSeverityWeight(db, totalData.Repository, weight)
			defer wg.Done()
		}(weight)
	}
	wg.Wait()

	fmt.Println("Inserting Issue...")
	for _, issue := range totalData.Repository.Issues.Nodes {
		wg.Add(1)
		go func(issue *model.Issue) {
			insert.Issue(db, totalData.Repository, issue)
			defer wg.Done()
		}(issue)
	}
	wg.Wait()

	fmt.Println("Inserting Comment...")
	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, issueComment := range issue.Comments.Nodes {
			wg.Add(1)
			go func(issue *model.Issue, issueComment *model.IssueComment) {
				insert.Comment(db, issue, issueComment)
				defer wg.Done()
			}(issue, issueComment)
		}
	}
	wg.Wait()

	fmt.Println("Inserting PullRequest...")
	for _, pr := range totalData.Repository.PullRequests.Nodes {
		wg.Add(1)
		go func(pr *model.PullRequest) {
			insert.PullRequest(db, totalData.Repository, pr)
			defer wg.Done()
		}(pr)
	}
	wg.Wait()

	fmt.Println("Inserting PrComment...")
	for _, pr := range totalData.Repository.PullRequests.Nodes {
		for _, prComment := range pr.Comments.Nodes {
			wg.Add(1)
			go func(pr *model.PullRequest, prComment *model.IssueComment) {
				insert.PrComment(db, pr, prComment)
				defer wg.Done()
			}(pr, prComment)
		}
	}
	wg.Wait()

	fmt.Println("Inserting Version...")
	for _, ref := range totalData.Repository.Refs.Nodes {
		wg.Add(1)
		go func(ref *model.Ref) {
			insert.Tag(db, totalData.Repository, ref)
			insert.Version(db, ref)
			defer wg.Done()
		}(ref)
	}
	wg.Wait()
	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, issueComment := range issue.Comments.Nodes {
			wg.Add(1)
			go func(issue *model.Issue, body *string) {
				insert.IssueVersion(db, issue, body)
				defer wg.Done()
			}(issue, &issueComment.Body)
		}
	}
	wg.Wait()

	fmt.Println("Inserting IssueLabel...")
	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, label := range issue.Labels.Nodes {
			wg.Add(1)
			go func(issue *model.Issue, label *model.Label) {
				insert.IssueLabel(db, totalData.Repository, issue, label)
				defer wg.Done()
			}(issue, label)
		}
	}
	wg.Wait()

	fmt.Println("Inserting PullRequestsLabel...")
	for _, pr := range totalData.Repository.PullRequests.Nodes {
		for _, label := range pr.Labels.Nodes {
			wg.Add(1)
			go func(pr *model.PullRequest, label *model.Label) {
				insert.PullRequestLabel(db, totalData.Repository, pr, label)
				defer wg.Done()
			}(pr, label)
		}
	}
	wg.Wait()

	fmt.Println("Inserting UserIssue...")
	for _, issue := range totalData.Repository.Issues.Nodes {
		for _, user := range issue.Assignees.Nodes {
			wg.Add(1)
			go func(issue *model.Issue, user *model.User) {
				insert.UserIssue(db, issue, user)
				defer wg.Done()
			}(issue, user)
		}
	}
	wg.Wait()

	fmt.Println("Inserting UserPullRequests...")
	for _, pr := range totalData.Repository.PullRequests.Nodes {
		for _, user := range pr.Assignees.Nodes {
			wg.Add(1)
			go func(pr *model.PullRequest, user *model.User) {
				insert.UserPullRequest(db, pr, user)
				defer wg.Done()
			}(pr, user)
		}
		wg.Add(1)
		go func(pr *model.PullRequest, actor *model.Actor) {
			insert.ActorPullRequest(db, pr, actor)
			defer wg.Done()
		}(pr, pr.Actor)
	}
	wg.Wait()

	fmt.Println("Inserting TeamIssue...")
	for _, issue := range totalData.Repository.Issues.Nodes {
		teams := team.GetTeams(totalData.Repository.Name, issue)
		for _, team := range teams {
			wg.Add(1)
			go func(issue *model.Issue, team string) {
				insert.TeamIssue(db, issue, team)
				defer wg.Done()
			}(issue, team)
		}
	}
	wg.Wait()

}

// RunInfrastructure fetch all the data first and then fetch data 10 days before.
func RunInfrastructure(c config.Config) {

	initDB(c.DatabaseConfig)

	queries := make([]model.Query, len(c.RepositoryArgs))
	for i, repositoryArg := range c.RepositoryArgs {
		retryTimes := 3
		for {
			queries[i], err = FetchQuery(c.CrawlerConfig, repositoryArg.Owner, repositoryArg.Name)
			if err != nil {
				if retryTimes == 0 {
					panic(err)
				} else {
					retryTimes--
				}
			} else {
				break
			}
		}
	}

	truncate.AllClear(db)
	insert.Team(db, &c)
	for i, query := range queries {
		InsertQuery(db, query, c.RepositoryArgs[i].Owner, &c)
	}

	process.ProcessAll(db)

	if err != nil {
		fmt.Println(err)
	}
}
