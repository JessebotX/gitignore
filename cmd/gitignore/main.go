package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jessebotx/gitignore"
)

var ProgramName = "gitignore"

const Version = "1.0.0"
const Usage = `USAGE
=====
gitignore <type...>
    write one or more .gitignore types (e.g. "go", "python", "c++", etc.) into a .gitignore file.
gitignore --print|-p <type...>
    print one or more .gitignore types (e.g. "go", "python", "c++", etc.) contents into stdout.
gitignore --names|--types
    print list of available gitignore types. Source: <https://github.com/github/gitignore>.
gitignore --help|-h
    print this usage information.
gitignore --version
    print program's version.

EXAMPLES
========
Get list of all gitignore types that are supported

   gitignore --types

Create a .gitignore file for a Go project

   gitignore go

Create a .gitignore file for a Go and a Node.js project

   gitignore go node

Print out concatenated .gitignore for a Go and a Node.js project into stdout

   gitignore --print go node
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "[ERROR] Missing arguments")
		os.Exit(1)
	}

	ProgramName = os.Args[0]

	// parse cli args and flags
	printNamesFlag := false
	printHelpFlag := false
	printContentsFlag := false
	printVersionFlag := false

	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			printHelpFlag = true
			break
		}

		if arg == "--version" {
			printVersionFlag = true
			break
		}

		if arg == "--names" || arg == "--types" {
			printNamesFlag = true
			break
		}

		if arg == "-p" || arg == "--print" {
			printContentsFlag = true
			break
		}
	}

	if printHelpFlag {
		fmt.Println(Usage)
		return
	}

	if printVersionFlag {
		fmt.Printf("%s v%s\n", ProgramName, Version)
		return
	}

	if printNamesFlag {
		printNames()
		return
	}

	if printContentsFlag {
		writeContents(os.Stdout, os.Args[1:])
		return
	}

	// default write to <current working dir>/.gitignore file
	f, err := os.Create(".gitignore")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] Cannot create .gitignore in current directory")
		os.Exit(1)
	}
	defer f.Close()

	writeContents(f, os.Args[1:])
}

func writeContents(w io.Writer, args []string) {
	response, err := gitignore.RequestJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}

		bytes, err := gitignore.Gitignore(response, arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[WARNING] failed to get %s.gitignore\n", arg)
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Fprintf(w, string(bytes) + "\n")
	}
}

func printNames() {
	response, err := gitignore.RequestJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	list := gitignore.NamesList(response)
	for _, v := range list {
		fmt.Println(v)
	}
}
