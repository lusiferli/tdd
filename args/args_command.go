package main

import "strings"

type Command struct {
	Flag  string
	Value string
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
