package main

import (
	"os"
	"path/filepath"
)

const SubDir = ".mirror_files"

var (
	ReservoirIndexDirectory = filepath.Join(filepath.Dir(os.Args[0]), SubDir, "reservoir-index")
	GithubRepoPoolDirectory = filepath.Join(filepath.Dir(os.Args[0]), SubDir, "github_repo_pool")
)
