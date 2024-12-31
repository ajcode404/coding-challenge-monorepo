package commands

import (
	"testing"
)

func TestCountCommand(t *testing.T) {
	wc := &WordCount{Commands: []string{"-c"}, FileBytes: []byte("Hello World")}
	res := CountCommand(wc)
	if res != "11" {
		t.Error("Expected 11 got", res)
	}
}

func TestLineCommand(t *testing.T) {
	wc := &WordCount{Commands: []string{"-l"}, FileBytes: []byte("Hello\nWorld")}
	res := LineCommand(wc)
	if res != "1" {
		t.Error("Expected 1 got", res)
	}
}

func TestWordCommand(t *testing.T) {
	wc := &WordCount{Commands: []string{"-w"}, FileBytes: []byte("Hello World")}
	res := WordCommand(wc)
	if res != "2" {
		t.Error("Expected 2 got", res)
	}
}
