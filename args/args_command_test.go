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

	if command.Flag != "f" {
		t.Fail()
	}
	if command.Value != "12" {
		t.Fail()
	}
}

func Test_parse_type(t *testing.T) {
	commandStr := "-f 12"
	var command = CreateCommand(commandStr)
	schema := Schema{}
	schema.Flag = "f"
	schema.Type = "int"
	schema.DefaultValue = 0
	if command.GetValueWithSchema(schema) != 12 {
		t.Fail()
	}
}
