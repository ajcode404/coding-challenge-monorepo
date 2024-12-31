package main

import (
	"os"
	"testing"
)

func TestParseCommandLineArgs(t *testing.T) {
	// Set up the arguments
	os.Args = []string{"cmd", "-c", "file.txt"}

	wc, err := ParseCommandLineArgs()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if wc.fileName != "file.txt" {
		t.Errorf("Expected file.txt, got %s", wc.fileName)
	}

	if len(wc.commands) != 1 || wc.commands[0] != "-c" {
		t.Errorf("Expected [-c], got %v", wc.commands)
	}
}
