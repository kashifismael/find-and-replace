package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func Replace(fileName string, lineOfText string, textToReplace string, textToAdd string) {

	lineOfTextArgs := strings.Split(lineOfText, ":")

	if len(lineOfTextArgs) == 1 {
		return
	}

	lineNumber := lineOfTextArgs[0]
	text := lineOfTextArgs[1]

	fmt.Println("")
	fmt.Println(text)
	message := fmt.Sprintf("(k)eep '%v', or (r)eplace with '%v'", textToReplace, textToAdd)
	fmt.Println(message)

	var input string
	fmt.Scan(&input)

	if input != "k" && input != "r" {
		panic("received unrecognised value")
	}

	if input == "k" {
		fmt.Println("skipping...")
		return
	}

	sedArgs := fmt.Sprintf("%vs/%v/%v/", lineNumber, textToReplace, textToAdd)
	cmd := exec.Command("sed", "-i", "", sedArgs, fileName)

	_, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

}
