package main

import (
	"fmt"
	"log"
	"io"
	"net/http"
	"encoding/json"
	"regexp"
	"strings"
)

type Response struct {
	Name string `json:"name"`
	DownloadURL string `json:"download_url"`
}

func main() {
	response, err := http.Get("https://api.github.com/repos/github/gitignore/contents")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result []Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err);
	}

	re := regexp.MustCompile("\\.gitignore$")
	for _, v := range result {
		runThis := false // DEBUG set to true to print all types
		if runThis && re.Match([]byte(v.Name)) {
			s := re.ReplaceAll([]byte(v.Name), []byte(""))
			name := strings.ToLower(string(s))
			fmt.Println(name)
		}

		if v.Name == "C.gitignore" {
			fileResponse, err := http.Get(v.DownloadURL)
			if err != nil {
				log.Fatal(err)
			}

			fileBody, err := io.ReadAll(fileResponse.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(fileBody))
		}
	}

	// TODO Steps
	// - fetch from https://api.github.com/repos/github/gitignore/contents
	// - find the correct type from the array by looking at <index>/name
	// - also get download_url from at <index>
	//
	// e.g. GET https://api.github.com/repos/github/gitignore/contents[3].name = Ada.gitignore AND .download_url = raw.github.com/...
}
