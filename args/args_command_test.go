package main

import (
	"testing"
)

func Test_parse_command(t *testing.T) {
	commandStr := "-f 12"
	var command = CreateCommand(commandStr)
	if command == (Command{}) {
		t.Fail()
	}

	if command.Flag != "-f" {
		t.Fail()
	}
}
