package codecov

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/PingCAP-QE/dashboard/infrastructure/config"
	"github.com/PingCAP-QE/libs/coverage"
)

var db *sql.DB
var err error

// init Set mysql database connection
func initDB(c config.Config) {
	db, err = sql.Open("mysql", c.GithubConfig.DatabaseConfig.DatabaseUrl)
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

// RunInfrastructure fetch all the data first and then fetch data 10 days before.
func RunInfrastructure(c config.Config) {
	initDB(c)

	fmt.Println("Processing coverage...")
	for _, repositoryArg := range c.GithubConfig.RepositoryArgs {
		err := coverage.ProcessCoverage(db, repositoryArg.Owner, repositoryArg.Name)
		if err != nil {
			fmt.Printf("ProcessCoverage error: %v\n", err)
		}
	}
}
