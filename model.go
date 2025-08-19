package main

import "time"

type MetaData struct {
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	FullName    string    `json:"fullName"`
	Description string    `json:"description"`
	Keywords    []string  `json:"keywords"`
	Homepage    *string   `json:"homepage"`
	License     string    `json:"license"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Stars       int       `json:"stars"`
	Sources     []struct {
		Type          string `json:"type"`
		Host          string `json:"host"`
		Id            string `json:"id"`
		FullName      string `json:"fullName"`
		RepoUrl       string `json:"repoUrl"`
		GitUrl        string `json:"gitUrl"`
		DefaultBranch string `json:"defaultBranch"`
	} `json:"sources"`
	SchemaVersion string `json:"schemaVersion"`
}

type Builds struct {
	SchemaVersion string `json:"schemaVersion"`
	Data          []struct {
		Built          bool      `json:"built"`
		Tested         *bool     `json:"tested"`
		Toolchain      string    `json:"toolchain"`
		RequiredUpdate *bool     `json:"requiredUpdate"`
		ArchiveSize    *int      `json:"archiveSize"`
		ArchiveHash    *string   `json:"archiveHash"`
		RunAt          time.Time `json:"runAt"`
		Url            string    `json:"url"`
		Revision       string    `json:"revision"`
	} `json:"data"`
}

type Versions struct {
	SchemaVersion string `json:"schemaVersion"`
	Data          []struct {
		Version             string    `json:"version"`
		Revision            string    `json:"revision"`
		Date                time.Time `json:"date"`
		Tag                 *string   `json:"tag"`
		Toolchain           string    `json:"toolchain"`
		PlatformIndependent *string   `json:"platformIndependent"`
		License             *string   `json:"license"`
		LicenseFiles        []string  `json:"licenseFiles"`
		ReadmeFile          string    `json:"readmeFile"`
		Dependencies        []struct {
			Type       string  `json:"type"`
			Name       string  `json:"name"`
			Scope      *string `json:"scope"`
			Version    string  `json:"version"`
			Transitive bool    `json:"transitive"`
			Rev        string  `json:"rev"`
			InputRev   *string `json:"inputRev"`
			Url        string  `json:"url"`
		} `json:"dependencies"`
	} `json:"data"`
}
