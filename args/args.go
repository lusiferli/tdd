package main

type Args struct {
	schemaStr  string
	commandStr string
	SchemaMap  map[string]Schema
	CommandMap map[string]Command
}

func CreateArgs(schemaStr string, commandStr string) Args {
	commandMap := ParseCommandsToMap(commandStr)
	schemaMap := ParseSchemasToMap(schemaStr)
	args := Args{}
	args.SchemaMap = schemaMap
	args.CommandMap = commandMap
	return args
}
