package entity

import "fmt"

type Status string

const (
	StatusNone       Status = ""
	StatusToDo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func GetErrInvalidStatus(s string) error {
	return fmt.Errorf("unable to parse status: %s", s)
}

var validStatuses = map[string]Status{
	"":            StatusNone,
	"todo":        StatusToDo,
	"in-progress": StatusInProgress,
	"done":        StatusDone,
}

func ParseStatus(s string) (Status, error) {
	if status, ok := validStatuses[s]; ok {
		return status, nil
	}
	return StatusNone, GetErrInvalidStatus(s)
}
