package commands

import (
	"fmt"
	"strings"
)

// Cannot think of any good name TBH
type WordCount struct {
	Commands  []string
	FileBytes []byte
}

func CountCommand(wc *WordCount) string {
	if !isCountCommand(wc.Commands) {
		return ""
	}
	return fmt.Sprint(len(wc.FileBytes))
}
func LineCommand(wc *WordCount) string {
	if !isLineCommand(wc.Commands) {
		return ""
	}
	newLineCount := len(strings.Split(string(wc.FileBytes), "\n")) - 1
	return fmt.Sprint(newLineCount)
}

func WordCommand(wc *WordCount) string {
	if !isWordCommand(wc.Commands) {
		return ""
	}
	return fmt.Sprint(len(strings.Fields(string(wc.FileBytes))))
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
