package command

import "errors"

var ErrInvalidArguments = errors.New("invalid arguments")
var ErrUnableToParseTaskId = errors.New("unable to parse task ID")

type CommandHandler interface {
	Execute(args []string) ([]string, error)
}
