package gitignore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Parsed response from GitHub API containing relevant repository details
type RepoResponse struct {
	// Parsed list of relevant file data from repository
	Data []RepoResponseItem

	// HTTP Response object
	HTTPResponse *http.Response
}

// Close the HTTP Response Body
func (r *RepoResponse) Close() {
	r.HTTPResponse.Body.Close()
}

// A parsed, single relevant file item from GitHub's gitignore repository
type RepoResponseItem struct {
	// Name of gitignore file
	TypeName string `json:"name"`
	// URL pointing to a downloadable gitignore file
	DownloadURL string `json:"download_url"`
}

// Name or type of a .gitignore file
type Type struct {
	// Complete .gitignore file name (e.g. "Go.gitignore")
	Name string
}

// Shortened, lowercased version of FileName (e.g. "go" (from Go.gitignore))
func (t *Type) ShortName() string {
	return strings.ToLower(strings.TrimSuffix(t.Name, ".gitignore"))
}

// Requested gitignore type does not exist
type DoesNotExistError struct {
	TypeName string
}

func (e *DoesNotExistError) Error() string {
	return fmt.Sprintf("%s.gitignore does not exist.", e.TypeName)
}

// Make an initial
func Request() (*RepoResponse, error) {
	response, err := http.Get("https://api.github.com/repos/github/gitignore/contents")
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

	result := &RepoResponse{
		HTTPResponse: response,
		Data:         contents,
	}

	return result, nil
}

// Fetch a list of supported gitignore types that can be fetched
func TypesList(repoResponse *RepoResponse) []Type {
	types := make([]Type, 0)
	for _, v := range repoResponse.Data {
		if !strings.HasSuffix(v.TypeName, ".gitignore") {
			continue
		}

		t := Type{
			Name: v.TypeName,
		}
		types = append(types, t)
	}

	return types
}

// Fetch gitignore file content from a URL that points to the file
func FetchFromURL(rawFileURL string) ([]byte, error) {
	response, err := http.Get(rawFileURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Fetch gitignore file content from a short name
func FetchFromShortName(repoResponse *RepoResponse, shortName string) ([]byte, error) {
	url := ""
	for _, v := range repoResponse.Data {
		if strings.ToLower(v.TypeName) == strings.ToLower(shortName+".gitignore") {
			url = v.DownloadURL
			break
		}
	}

	if url == "" {
		return nil, &DoesNotExistError{
			TypeName: shortName,
		}
	}

	return FetchFromURL(url)
}
