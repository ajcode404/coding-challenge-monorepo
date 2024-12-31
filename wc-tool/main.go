package main

import (
	"fmt"
	"os"

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

func isCountCommand(command string) bool {
	return command == "-c" || command == "--bytes"
}

func isLineCommand(command string) bool {
	return command == "-l" || command == "--lines"
}

func isWordCommand(command string) bool {
	return command == "-w" || command == "--words"
}

func parseCommandLineArgs() {
	args := os.Args[1:]
	if len(args) != 2 {
		PrintHelpText()
		return
	}

	if isCountCommand(args[0]) {
		countCmd := &commands.CountCommand{
			FileName: args[1],
		}
		o, err := countCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d %s\n", o, countCmd.FileName)
	}

	if isLineCommand(args[0]) {
		lineCmd := &commands.LineCommand{
			FileName: args[1],
		}
		o, err := lineCmd.Execute()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d %s\n", o, lineCmd.FileName)
	}

	if isWordCommand(args[0]) {
		wordCommand := &commands.WordCommand{
			FileName: args[1],
		}
		o, err := wordCommand.Execute()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d %s\n", o, wordCommand.FileName)
	}
}

func main() {
	parseCommandLineArgs()
}
