package gitignore

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type RepoResponseItem struct {
	Name string `json:"name"`
	DownloadURL string `json:"download_url"`
}

const (
	RawFileBaseURL = "https://raw.githubusercontent.com/github/gitignore/master/"
	RepoAPIBaseURL = "https://api.github.com/repos/github/gitignore/contents"
)

func ResponseJSON() ([]RepoResponseItem, error) {
	response, err := http.Get(RepoAPIBaseURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var json []RepoResponseItem
	err := json.Unmarshal(responseBody, &contents)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func NamesList(responseJSON []RepoResponseItem) []string {
	names := make([]string, 0)
	for _, v := range contents {
		if !strings.HasSuffix(v.Name, ".gitignore") {
			continue
		}

		name := strings.ToLower(strings.TrimSuffix(v.Name, "gitignore"))
		names = append(names, name)
	}

	return names
}

func Gitignore(name string) ([]byte, error) {
	url := RawFileBaseURL + name + ".gitignore"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
