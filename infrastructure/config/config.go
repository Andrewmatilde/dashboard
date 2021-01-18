package config

import (
	githubConfig "github.com/PingCAP-QE/dashboard/infrastructure/github/config"
)

type Config struct {
	GithubConfig githubConfig.Config
}

func GetConfig(path string) Config {
	var c Config
	c.GithubConfig = githubConfig.GetConfig(path)
	return c
}
