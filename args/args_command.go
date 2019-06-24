package main

type Command struct {
	Flag  string
	Value string
}

func CreateCommand(commandStr string) Command {
	command := Command{}
	command.Flag = ""
	command.Value = ""
	return command
}
