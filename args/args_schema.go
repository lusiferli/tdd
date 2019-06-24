package main

import (
	"strings"
)

func CreateSchema(schemaStr string) Schema {
	items := strings.Split(schemaStr, ":")
	if len(items) != 2 {
		return Schema{}
	}
	schema := Schema{}
	schema.Flag = items[0]
	schema.Type = items[1]
	switch items[1] {
	case "int":
		schema.DefaultValue = 0
		break
	case "bool":
		schema.DefaultValue = true
		break
	case "string":
		schema.DefaultValue = ""
		break
	}
	return schema
}
