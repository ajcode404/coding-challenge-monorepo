package commands

import (
	"fmt"
	"os"
)

type CountCommand struct {
	FileName string
}

func (c CountCommand) Execute() (string, error) {
	bytes, err := os.ReadFile(c.FileName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", len(bytes)), nil
}
