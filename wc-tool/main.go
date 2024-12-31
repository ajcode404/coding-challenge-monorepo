package main

import (
	"fmt"
	"os"
	"strings"
)

const HELP_TEXT = `Usage wc [OPTION] [FILENAME]
	-c, --bytes            print the byte counts
	-l, --lines            print the newline counts
`

func PrintHelpText() {

	fmt.Print(HELP_TEXT)
}

type CountCommand struct {
	input string
}

func (c CountCommand) execute() {
	bytes, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d %s\n", len(bytes), c.input)
}

type LineCommand struct {
	input string
}

func (c LineCommand) execute() {
	bytes, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := strings.Split(string(bytes), "\n")
	// -1 due to we want to count \n and not the number of lines.
	fmt.Printf("%d %s\n", len(text)-1, c.input)
}

func isCountCommand(command string) bool {
	return command == "-c" || command == "--bytes"
}

func isLineCommand(command string) bool {
	return command == "-l" || command == "--lines"
}

func parseCommandLineArgs() {
	args := os.Args[1:]
	if len(args) != 2 {
		PrintHelpText()
		return
	}
	if isCountCommand(args[0]) {
		countCmd := &CountCommand{
			input: args[1],
		}
		countCmd.execute()
	}

	if isLineCommand(args[0]) {
		lineCmd := &LineCommand{
			input: args[1],
		}
		lineCmd.execute()
	}
}

func main() {
	parseCommandLineArgs()
}
