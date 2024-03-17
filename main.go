package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// TODO Steps
	// - fetch from https://api.github.com/repos/github/gitignore/contents
	// - find the correct type from the array by looking at <index>/name
	// - also get download_url from at <index>
	//
	// e.g. GET https://api.github.com/repos/github/gitignore/contents[3].name = Ada.gitignore AND .download_url = raw.github.com/...
}
