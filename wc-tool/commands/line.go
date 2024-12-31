package commands

import (
	"fmt"
	"os"
	"strings"
)

type LineCommand struct {
	FileName string
}

func (l LineCommand) Execute() (string, error) {
	bytes, err := os.ReadFile(l.FileName)
	if err != nil {
		return "", err
	}
	newLineCount := len(strings.Split(string(bytes), "\n")) - 1
	return fmt.Sprintf("%d", newLineCount), nil
}
