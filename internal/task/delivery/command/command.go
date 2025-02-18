package command

type Command string

const (
	СommandList           Command = "list"
	СommandAdd            Command = "add"
	СommandMarkDone       Command = "mark-done"
	СommandMarkInProgress Command = "mark-in-progress"
	СommandUpdate         Command = "update"
)
