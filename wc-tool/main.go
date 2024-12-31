package main

import (
	"fmt"
	"os"
	"strings"

	"ajcode404.github.io/m/commands"
)

const HELP_TEXT = `Usage wc [OPTION] [FILENAME]
	-c, --bytes            print the byte counts
	-l, --lines            print the newline counts
	-w, --words            print the word counts
`

func PrintHelpText() {
	fmt.Print(HELP_TEXT)
}

func isCountCommand(commands []string) bool {
	for _, el := range commands {
		if el == "-c" || el == "--bytes" {
			return true
		}
	}
	return false
}

func isLineCommand(commands []string) bool {
	for _, el := range commands {
		if el == "-l" || el == "--lines" {
			return true
		}
	}
	return false
}

func isWordCommand(commands []string) bool {
	for _, el := range commands {
		if el == "-w" || el == "--words" {
			return true
		}
	}
	return false
}

// Cannot think of any good name TBH
type WordCount struct {
	commands []string
	fileName string
}

func parseCommandLineArgs() (*WordCount, error) {
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
		PrintHelpText()
		return nil, fmt.Errorf("either file or commands are not passed")
	}

	return &WordCount{
		commands: commandsList,
		fileName: file,
	}, nil
}

func main() {
	wordCount, err := parseCommandLineArgs()
	if err != nil {
		fmt.Println(err)
	}

	var outputString string = ""

	if isCountCommand(wordCount.commands) {
		o, err := commands.CountCommand(wordCount.fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputString += fmt.Sprintf("c:%d", o) + " "
	}
	if isLineCommand(wordCount.commands) {
		o, err := commands.LineCommand(wordCount.fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputString += fmt.Sprintf("l:%d", o) + " "
	}

	if isWordCommand(wordCount.commands) {
		o, err := commands.WordCommand(wordCount.fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputString += fmt.Sprintf("w:%d", o) + " "
	}
	outputString = strings.TrimRight(outputString, " ")
	fmt.Printf("%s %s\n", outputString, wordCount.fileName)
}
