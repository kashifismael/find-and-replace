package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cliArgs := os.Args[1:]

	textToSearchFor := cliArgs[0]
	textToReplaceWith := cliArgs[1]
	pathToFile := cliArgs[2]

	cmd := exec.Command("rg", textToSearchFor, pathToFile, "--no-heading", "--with-filename", "--line-number")

	out, err := cmd.Output()

	if err != nil {
		log.Fatal("No results")
	}

	rawResults := strings.Split(string(out), "\n")

	var results []LineContext = make([]LineContext, 0)

	for _, line := range rawResults {

		if strings.TrimSpace(line) != "" {
			context := BuildLineContext(line, textToSearchFor, textToReplaceWith, len(results) + 1)
			results = append(results, context)
		}

	}

	for _, context := range results {
		Replace(context, len(results))
	}
}
