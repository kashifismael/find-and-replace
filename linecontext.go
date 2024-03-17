package main

import (
	"strings"
)

type LineContext struct {
	fileName     string
	lineNumber   string
	lineContent  string
	wordToRemove string
	wordToAdd    string
	resultIndex  int
}

func BuildLineContext(rawGrepLine string, wordToRemove string, wordToAdd string, resultIndex int) LineContext {
	lineOfTextArgs := strings.Split(rawGrepLine, ":")

	return LineContext{
		lineOfTextArgs[0],
		lineOfTextArgs[1],
		lineOfTextArgs[2],
		wordToRemove,
		wordToAdd,
		resultIndex,
	}
}
