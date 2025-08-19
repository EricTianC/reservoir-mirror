package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

const ReservoirIndexGithubUrl = "https://github.com/leanprover/reservoir-index.git"

func syncReservoirIndex() {
	_ = SyncGitRepo(ReservoirIndexDirectory, ReservoirIndexGithubUrl)
}

func ReservoirIndexServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	slog.Info("Syncing reservoir index")
	syncReservoirIndex()

	srv := &http.Server{
		Addr:    ":80",
		Handler: ReservoirServerHandler{},
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Info("Serving reservoir index on port 80")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	for {
		select {
		case <-ticker.C:
			slog.Info("Syncing reservoir index")
			syncReservoirIndex()
		case <-ctx.Done():
			slog.Info("Canceling reservoir index mirror server")
			_ = srv.Shutdown(ctx)
			return
		}
	}
}

type ReservoirServerHandler struct{}

func (s ReservoirServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	packageHandler(w, r)
}

func packageHandler(w http.ResponseWriter, r *http.Request) {
	ownerPkg := strings.TrimPrefix(r.URL.Path, "/packages/")
	parts := strings.Split(ownerPkg, "/")
	if len(parts) < 2 {
		slog.Error("Unexpected request", "url", r.URL.Path)
		http.Error(w, "Invalid package name", http.StatusBadRequest)
		return
	}

	owner := parts[0]
	name := parts[1]

	switch {
	case strings.HasSuffix(ownerPkg, "/builds"):
		fetchPackageJson(w, owner, name, "builds")
	case strings.HasSuffix(ownerPkg, "/versions"):
		fetchPackageJson(w, owner, name, "versions")
	default:
		fetchPackageMetaData(w, owner, name)
	}
}

func fetchPackageMetaData(w http.ResponseWriter, owner, name string) {
	fileUrl := path.Join(ReservoirIndexDirectory, owner, name, "metadata.json")
	slog.Info("Fetching package", "fileUrl", fileUrl)

	data, err := os.ReadFile(fileUrl)
	if err != nil {
		slog.Error("Failed to fetch package", "err", err)
		http.Error(w, "Failed to fetch package", http.StatusInternalServerError)
		return
	}

	var metadata MetaData
	if err := json.Unmarshal(data, &metadata); err != nil {
		slog.Error("Failed to fetch package", "err", err)
		return
	}

	metadata.Sources[0].GitUrl = strings.Replace(metadata.Sources[0].GitUrl, "https://", "https://gitclone.com/", 1)
	bMetadata, _ := json.MarshalIndent(metadata, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = fmt.Fprintf(w, string(bMetadata))
	if err != nil {
		return
	}
}

func fetchPackageJson(w http.ResponseWriter, owner, name, sort string) {
	fileUrl := path.Join(ReservoirIndexDirectory, owner, name, fmt.Sprintf("%s.json", sort))
	slog.Info("Fetching package", "fileUrl", fileUrl)

	data, err := os.ReadFile(fileUrl)
	if err != nil {
		slog.Error("Failed to fetch package", "err", err)
		http.Error(w, "Failed to fetch package", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = fmt.Fprintf(w, string(data))
	if err != nil {
		return
	}
}
