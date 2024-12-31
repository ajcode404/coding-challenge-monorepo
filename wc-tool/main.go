package main

import (
	"fmt"
	"os"
	"strings"
	"wc-tool/commands"
)

const HELP_TEXT = `Usage wc [OPTION] [FILENAME]
	-c, --bytes            print the byte counts
	-l, --lines            print the newline counts
	-w, --words            print the word counts
`

func PrintHelpText() {
	fmt.Print(HELP_TEXT)
}

type CInput struct {
	commands []string
	fileName string
}

func ParseCommandLineArgs() (*CInput, error) {
	args := os.Args[1:]
	var commandsList []string
	var file string = ""
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") || strings.HasPrefix(args[i], "--") {
			commandsList = append(commandsList, args[i])
		} else {
			// Supports single file for now
			file = args[i]
		}
	}

	if file == "" || len(commandsList) == 0 {
		return nil, fmt.Errorf("either file or commands are not passed")
	}

	return &CInput{commandsList, file}, nil
}

func main() {
	cInput, err := ParseCommandLineArgs()
	if err != nil {
		PrintHelpText()
		return
	}
	file, err := os.ReadFile(cInput.fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var outputString string = ""
	outputString += commands.CountCommand(&commands.WordCount{Commands: cInput.commands, FileBytes: file}) + " "
	outputString += commands.LineCommand(&commands.WordCount{Commands: cInput.commands, FileBytes: file}) + " "
	outputString += commands.WordCommand(&commands.WordCount{Commands: cInput.commands, FileBytes: file}) + " "
	outputString = strings.TrimRight(outputString, " ")
	fmt.Printf("%s %s\n", outputString, cInput.fileName)
}
