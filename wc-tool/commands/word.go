package commands

import (
	"os"
	"strings"
)

type WordCommand struct {
	FileName string
}

func (w WordCommand) Execute() (int, error) {
	bytes, err := os.ReadFile(w.FileName)
	if err != nil {
		return -1, err
	}
	return len(strings.Fields(string(bytes))), nil
}
