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

func parseCommandLineArgs() {
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
		return
	}

	var outputString string = ""

	if isCountCommand(commandsList) {

		countCmd := &commands.CountCommand{
			FileName: file,
		}
		o, err := countCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		outputString += fmt.Sprintf("%d", o) + " "
	}

	if isLineCommand(commandsList) {
		lineCmd := &commands.LineCommand{
			FileName: file,
		}
		o, err := lineCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		outputString += fmt.Sprintf("%d", o) + " "
	}

	if isWordCommand(commandsList) {
		wordCommand := &commands.WordCommand{
			FileName: file,
		}
		o, err := wordCommand.Execute()
		if err != nil {
			fmt.Println(err)
		}
		outputString += fmt.Sprintf("%d", o) + " "
	}
	outputString = strings.TrimRight(outputString, " ")
	fmt.Printf("%s %s\n", outputString, file)
}

func main() {
	parseCommandLineArgs()
}
