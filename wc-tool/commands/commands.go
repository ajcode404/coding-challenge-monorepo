package commands

import (
	"os"
	"strings"
)

func CountCommand(fileName string) (int, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}
	return len(bytes), nil
}

func LineCommand(fileName string) (int, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}
	newLineCount := len(strings.Split(string(bytes), "\n")) - 1
	return newLineCount, nil
}

func WordCommand(fileName string) (int, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}
	return len(strings.Fields(string(bytes))), nil
}
