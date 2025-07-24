package main

import (
	"github.com/asif/mispa/config"
	"github.com/asif/mispa/github"
	"github.com/asif/mispa/utils"
)

func main() {
	cfg := config.LoadConfig()

	client := github.NewGitHubClient(cfg.GitHub.Token)
	files := github.DownloadFiles(client, cfg)

	utils.ReadMarkdownFiles(cfg.GitHub.DownloadDir, files)
}
