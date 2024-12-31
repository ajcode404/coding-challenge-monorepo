package commands

import (
	"os"
	"strings"
)

type LineCommand struct {
	FileName string
}

func (l LineCommand) Execute() (int, error) {
	bytes, err := os.ReadFile(l.FileName)
	if err != nil {
		return -1, err
	}
	newLineCount := len(strings.Split(string(bytes), "\n")) - 1
	return newLineCount, nil
}
