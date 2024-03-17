package main

import (
	"fmt"
	"os/exec"
)

func Replace(lineCtx LineContext, totalNumberOfLines int) {

	textToReplace := lineCtx.wordToRemove
	textToAdd := lineCtx.wordToAdd
	fileName := lineCtx.fileName
	lineNumber := lineCtx.lineNumber

	userIntent := requestUserInput(lineCtx, totalNumberOfLines)

	if userIntent == "KEEP" {
		return
	}

	sedArgs := fmt.Sprintf("%vs/%v/%v/", lineNumber, textToReplace, textToAdd)
	cmd := exec.Command("sed", "-i", "", sedArgs, fileName)

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

}

func requestUserInput(lineCtx LineContext, totalNumberOfLines int) string {
	text := lineCtx.lineContent
	textToReplace := lineCtx.wordToRemove
	textToAdd := lineCtx.wordToAdd
	fileName := lineCtx.fileName
	lineNumber := lineCtx.lineNumber
	resultIndex := lineCtx.resultIndex

	fmt.Println("")

	fileInfo := fmt.Sprintf("%v:%v \n", fileName, lineNumber)
	fmt.Println(Format(YELLOW, fileInfo))

	fmt.Println(text)

	fmt.Println("")
	message := Format(
		PURPLE,
		fmt.Sprintf("(%v/%v) Do you want to replace '%v' with '%v'? (y/n)",
			resultIndex,
			totalNumberOfLines,
			textToReplace,
			textToAdd,
		))
	fmt.Println(message)

	var input string
	fmt.Scan(&input)

	if input != "y" && input != "n" {
		panic("received unrecognised value")
	}

	if input == "n" {
		return "KEEP"
	}

	return "REPLACE"
}
