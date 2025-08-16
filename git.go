package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"log/slog"
	"os"
)

func SyncGitRepo(path string, url string) error {
	var progress sideband.Progress
	if debugOpt {
		progress = os.Stdout
	}

	repo, err := git.PlainOpen(path)
	if err == nil {
		remotes, _ := repo.Remotes()
		remoteUrl := remotes[0].Config().URLs[0]
		slog.Debug("Repo already exists, skip cloning.", "path", path, "url", url, "remote", remoteUrl)
	} else {
		slog.Debug("Cloning repo.", "path", path, "url", url, "proxy", proxyAddress)

		repo, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: url,
			ProxyOptions: transport.ProxyOptions{
				URL: proxyAddress,
			},
			Progress: progress,
		})
	}
	return err
}
