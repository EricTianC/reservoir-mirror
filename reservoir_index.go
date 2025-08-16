package main

const ReservoirIndexGithubUrl = "https://github.com/leanprover/reservoir-index.git"

func syncReservoirIndex() {
	_ = SyncGitRepo(ReservoirIndexDirectory, ReservoirIndexGithubUrl)
}
