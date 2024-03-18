package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessebotx/gitignore"
)

var ProgramName = "gitignore"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing arguments")
		os.Exit(1)
	}

	ProgramName = os.Args[0]

	// parse cli args and flags
	printNamesFlag := false
	printHelpFlag := false
	printContentsFlag := false

	for _, arg := range os.Args {
		if arg == "--names" || arg == "--type" {
			printNamesFlag = true
			break
		}

		if arg == "--help" || arg == "-h" {
			printHelpFlag = true
			break
		}

		if arg == "-p" || arg == "--print" {
			printContentsFlag = true
			break
		}
	}

	if printNamesFlag {
		printNames()
		return
	}

	if printHelpFlag {
		printHelp()
		return
	}

	if printContentsFlag {
		printContents(os.Args[1:])
		return
	}
}

func printContents(args []string) {
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

		fmt.Println(string(bytes))
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

func printHelp() {
	fmt.Printf("USAGE\n")
}
