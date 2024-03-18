package gitignore

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"fmt"
)

const (
	RepoAPIBaseURL = "https://api.github.com/repos/github/gitignore/contents"
)

type RepoResponseItem struct {
	Name string `json:"name"`
	DownloadURL string `json:"download_url"`
}

type DoesNotExistError struct {
	Name string
}

func (e *DoesNotExistError) Error() string {
	return fmt.Sprintf("%s.gitignore does not exist.", e.Name)
}


func RequestJSON() ([]RepoResponseItem, error) {
	response, err := http.Get(RepoAPIBaseURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var contents []RepoResponseItem
	err = json.Unmarshal(body, &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func NamesList(responseJSON []RepoResponseItem) []string {
	names := make([]string, 0)
	for _, v := range responseJSON {
		if !strings.HasSuffix(v.Name, ".gitignore") {
			continue
		}

		name := strings.ToLower(strings.TrimSuffix(v.Name, "gitignore"))
		names = append(names, name)
	}

	return names
}

func Gitignore(requestJSON []RepoResponseItem, name string) ([]byte, error) {
	url := ""
	for _, v := range requestJSON {
		if strings.ToLower(v.Name) == strings.ToLower(name + ".gitignore") {
			url = v.DownloadURL
			break
		}
	}

	if url == "" {
		return nil, &DoesNotExistError{
			Name: name,
		}
	}

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
