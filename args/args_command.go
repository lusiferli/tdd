package main

import (
	"errors"
	"strconv"
	"strings"
)

type Command struct {
	Flag  string
	Value string
}

type Schema struct {
	Flag         string
	Type         string
	DefaultValue interface{}
}

func CreateCommand(commandStr string) Command {
	items := strings.Split(commandStr, " ")
	var flag string
	var value string
	command := Command{}
	if len(items) > 2 || len(items) == 0 {
		return command
	}
	flag = items[0]
	flag = parseFlag(flag)
	if len(items) == 2 {
		value = items[1]
	}
	command.Flag = flag
	command.Value = value
	return command
}

func parseFlag(flagStr string) string {
	return strings.Replace(flagStr, "-", "", 1)
}

func ParseCommands(commandStr string) []Command {
	var commands []Command
	commandItemStrArr := strings.Split(commandStr, "-")
	for _, commandItemStr := range commandItemStrArr {
		if commandItemStr == "" {
			continue
		}
		commandItemStr = strings.TrimSpace(commandItemStr)
		commandItemStr = "-" + commandItemStr
		commands = append(commands, CreateCommand(commandItemStr))
	}
	return commands
}

func ParseCommandsToMap(schemaStr string) map[string]Command {
	var commands = ParseCommands(schemaStr)
	commandMap := make(map[string]Command)
	for _, command := range commands {
		commandMap[command.Flag] = command
	}
	return commandMap
}

func (c *Command) GetValueWithSchema(schema Schema) interface{} {
	if schema.Type == "int" {
		return getIntValue(schema, c.Value)
	} else if schema.Type == "bool" {
		return getBoolValue(schema, c.Value)
	} else if schema.Type == "string" {
		return c.Value
	}
	return nil
}

func getIntValue(schema Schema, commandValue string) interface{} {
	if commandValue == "" {
		return schema.DefaultValue
	}
	value, err := strconv.Atoi(commandValue)
	if err != nil {
		return errors.New("FormatException")
	}
	return value
}

func getBoolValue(schema Schema, commandValue string) interface{} {
	if commandValue == "" {
		return schema.DefaultValue
	}
	commandValUp := strings.ToUpper(commandValue)
	boolValue := false
	if commandValUp == "TRUE" {
		boolValue = true
	} else if commandValUp == "FALSE" {
		boolValue = false
	} else {
		return errors.New("FormatException")
	}
	return boolValue
}
