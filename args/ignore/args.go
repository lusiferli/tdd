package ignore

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

type Args struct {
	Schema  string
	Command string

	schemaMap  map[string]string
	commandMap map[string]string

	once sync.Once
}

const (
	DEFAULT_INT_VAL  = 0
	DEFAULT_DIR_VAL  = "."
	DEFAULT_BOOL_VAL = false
)

func (args *Args) init() {
	args.once.Do(
		func() {
			args.resolveSchema()

			args.resolveCommand()

		})
}

func (args *Args) resolveCommand() {
	args.commandMap = map[string]string{}
	commandItemStrArr := strings.Split(args.Command, "-")
	for _, commandItemStr := range commandItemStrArr {
		args.resolveCommandItem(commandItemStr)
	}
}

func (args *Args) resolveCommandItem(commandItemStr string) {
	commandItemStr = strings.TrimSpace(commandItemStr)
	if commandItemStr != "" {
		commandItemArr := strings.Split(commandItemStr, " ")
		flag := commandItemArr[0]
		val := ""
		if len(commandItemArr) > 1 {
			val = commandItemArr[1]
		}
		args.commandMap[flag] = val
	}
}

func (args *Args) resolveSchema() {
	args.schemaMap = map[string]string{}
	schemaItemsStrArr := strings.Split(args.Schema, ",")
	for _, schemaItemStr := range schemaItemsStrArr {
		schemaItemArr := strings.Split(schemaItemStr, ":")
		flag := schemaItemArr[0]
		flagType := schemaItemArr[1]
		args.schemaMap[flag] = flagType
	}
}

func (args *Args) get(flag string) (arg interface{}) {
	args.init()

	flagType, ok := args.schemaMap[flag]
	if ok {

		switch flagType {
		case "int":
			arg = args.getIntArg(flag)
		case "string":
			arg = args.getStringArg(flag)
		case "bool":
			arg = args.getBoolArg(flag)

		}
		return arg
	}
	return errors.New("UnSupportError")
}

func (args *Args) getIntArg(flag string) interface{} {
	commandVal := args.commandMap[flag]
	if commandVal == "" {
		return DEFAULT_INT_VAL
	}
	arg, err := strconv.Atoi(commandVal)
	if err != nil {
		return errors.New("UnSupportError")
	}
	return arg
}

func (args *Args) getStringArg(flag string) interface{} {
	commandVal := args.commandMap[flag]
	if commandVal == "" {
		return DEFAULT_DIR_VAL
	}
	return commandVal
}

func (args *Args) getBoolArg(flag string) interface{} {
	commandVal := args.commandMap[flag]
	if commandVal == "" {
		return DEFAULT_BOOL_VAL
	}
	commandValUp := strings.ToUpper(commandVal)
	if commandValUp == "TRUE" {
		return true
	}
	if commandValUp == "FALSE" {
		return false
	}
	return errors.New("UnSupportError")
}
