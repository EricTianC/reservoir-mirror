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
	proxyOpt := transport.ProxyOptions{
		URL: proxyAddress,
	}

	repo, err := git.PlainOpen(path)
	if err == nil {
		remotes, _ := repo.Remotes()
		remoteUrl := remotes[0].Config().URLs[0]

		if remoteUrl != url {
			slog.Debug("Repo already exists, but the remoteUrl doesn't match.", "path", path, "url", url, "remote", remoteUrl)
			err = os.RemoveAll(path)
			return err
		}
		slog.Debug("Repo already exists, updating.", "path", path, "url", url, "remote", remoteUrl)
		wt, err := repo.Worktree()
		if err != nil {
			return err
		}
		err = wt.Pull(&git.PullOptions{Progress: progress, ProxyOptions: proxyOpt})
		return err
	}
	slog.Debug("Cloning repo.", "path", path, "url", url, "proxy", proxyAddress)

	repo, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:          url,
		ProxyOptions: proxyOpt,
		Progress:     progress,
	})

	return err
}
