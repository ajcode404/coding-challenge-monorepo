package main

import "os"

type Command struct {
	keyword string
	input   string
}

func parseCommandLineArgs() string {
	args := os.Args[1:]
	if len(args)%2 != 0 {
		panic("Incorrect format of the command")
	}
	for i := 0; i < len(args); i++ {
		keyword, input := args[i], args[i+1]
		if keyword == "-c" {

		}

	}
}

func main() {
	parseCommandLineArgs()
}
