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

func Test_parse_type_bool(t *testing.T) {
	commandStr := "-b true"
	var command = CreateCommand(commandStr)
	schema := Schema{}
	schema.Flag = "b"
	schema.Type = "bool"
	schema.DefaultValue = false
	if command.GetValueWithSchema(schema) != true {
		t.Fail()
	}
}

func Test_parse_type_string(t *testing.T) {
	commandStr := "-s jjjjjj"
	var command = CreateCommand(commandStr)
	schema := Schema{}
	schema.Flag = "s"
	schema.Type = "string"
	schema.DefaultValue = ""
	if command.GetValueWithSchema(schema) != "jjjjjj" {
		t.Fail()
	}
}

func Test_parse_command_str(t *testing.T) {
	commandStr := "-d 12 -f st"
	commands := ParseCommands(commandStr)
	if commands == nil {
		t.Fail()
	}
	schemaString := Schema{}
	schemaString.Flag = "f"
	schemaString.Type = "string"
	schemaString.DefaultValue = ""
	if commands[1].GetValueWithSchema(schemaString) != "st" {
		t.Fail()
	}
	schemaInt := Schema{}
	schemaInt.Flag = "d"
	schemaInt.Type = "int"
	schemaInt.DefaultValue = "0"

	if commands[0].GetValueWithSchema(schemaInt) != 12 {
		t.Fail()
	}
}

func Test_parse_schema_str(t *testing.T) {
	schemaStr := "d:int"
	schema := CreateSchema(schemaStr)
	if schema == (Schema{}) {
		t.Fail()
	}

	if schema.Type != "int" {
		t.Fail()
	}
}

func Test_parse_schema_strs(t *testing.T) {
	schemaStr := "d:int,f:string"
	schemas := ParseSchemas(schemaStr)
	if schemas == nil {
		t.Fail()
	}

	if len(schemas) != 2 {
		t.Fail()
	}

	if schemas[0].Type != "int" && schemas[0].Flag != "d" && schemas[0].DefaultValue != "0" {
		t.Fail()
	}
	if schemas[1].Type != "string" && schemas[1].Flag != "f" && schemas[10].DefaultValue != "" {
		t.Fail()
	}
}
