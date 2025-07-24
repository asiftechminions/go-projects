package github

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/asif/mispa/config"
	"github.com/google/go-github/v55/github"
)

func DownloadFiles(client *github.Client, cfg *config.Config) []string {
	ctx := context.Background()
	opt := &github.RepositoryContentGetOptions{Ref: cfg.GitHub.Branch}

	contents, dir, _, err := client.Repositories.GetContents(ctx, cfg.GitHub.Owner, cfg.GitHub.Repo, "", opt)
	if err != nil {
		panic(err)
	}

	os.MkdirAll(cfg.GitHub.DownloadDir, os.ModePerm)

	var savedFiles []string

	if dir != nil {
		for _, file := range dir {
			if file.GetType() == "file" {
				url := file.GetDownloadURL()
				resp, err := http.Get(url)
				if err != nil {
					continue
				}
				defer resp.Body.Close()

				filePath := cfg.GitHub.DownloadDir + "/" + file.GetName()
				out, err := os.Create(filePath)
				if err != nil {
					continue
				}
				defer out.Close()

				io.Copy(out, resp.Body)
				fmt.Printf("Downloaded: %s\n", file.GetName())
				savedFiles = append(savedFiles, filePath)
			}
		}
	} else if contents != nil {
		content, err := contents.GetContent()
		if err != nil {
			panic(err)
		}
		fmt.Println("File contents:", content)
	}

	return savedFiles
}
