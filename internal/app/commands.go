package app

type Command string

const (
	commandList           Command = "list"
	commandAdd            Command = "add"
	commandMarkDone       Command = "mark-done"
	commandMarkInProgress Command = "mark-in-progress"
)
