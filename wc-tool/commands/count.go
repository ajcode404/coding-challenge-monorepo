package commands

import (
	"os"
)

type CountCommand struct {
	FileName string
}

func (c CountCommand) Execute() (int, error) {
	bytes, err := os.ReadFile(c.FileName)
	if err != nil {
		return -1, err
	}
	return len(bytes), nil
}
