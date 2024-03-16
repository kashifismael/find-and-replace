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

	cmd := exec.Command("rg", textToSearchFor, "-H", pathToFile, "--line-number", "--no-filename")

	out, err := cmd.Output()

	if err != nil {
		log.Fatal("No results")
	}

	results := strings.Split(string(out), "\n")

	for _, line := range results {
		Replace(pathToFile, line, textToSearchFor, textToReplaceWith)
	}
}
