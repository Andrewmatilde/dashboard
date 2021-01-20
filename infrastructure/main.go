package main

import (
	"fmt"
	"time"

	"github.com/PingCAP-QE/dashboard/infrastructure/codecov"
	"github.com/PingCAP-QE/dashboard/infrastructure/config"
	"github.com/PingCAP-QE/dashboard/infrastructure/github"
)

func main() {
	c := config.GetConfig("./config.toml")
	github.RunInfrastructure(c.GithubConfig)
	codecov.RunInfrastructure(c)
	fmt.Printf(
		`
	###########################################################################################
	init db ok %v
	###########################################################################################
	`, time.Now())
	for {
		time.Sleep(time.Hour)
		c = config.GetConfig("./config.toml")
		github.RunInfrastructure(c.GithubConfig)
		codecov.RunInfrastructure(c)
		fmt.Printf(
			`
	###########################################################################################
	update database %v
	###########################################################################################
	`, time.Now())
	}
}
