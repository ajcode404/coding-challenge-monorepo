package main

import (
	"fmt"
	"os"
)

const HELP_TEXT = `Usage wc [OPTION] [FILENAME]
	-c, --bytes            print the byte counts
`

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

func parseCommandLineArgs() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Print(HELP_TEXT)
		return
	}
	if args[0] != "-c" && args[0] != "--bytes" {
		fmt.Print(HELP_TEXT)
		return
	}
	countCmd := &CountCommand{
		input: args[1],
	}
	countCmd.execute()
}

func main() {
	parseCommandLineArgs()
}
