package app

import (
	"fmt"
	"github.com/t8nax/task-tracker/internal/services"
	jsonstorage "github.com/t8nax/task-tracker/internal/storage/json"
)

func Run(args []string) {
	if len(args) == 1 {
		fmt.Println("Usage: task-tracker <commands>")
		return
	}

	command := Command(args[1])
	storage := &jsonstorage.JsonStorage{}
	service := services.NewTaskService(storage)

	switch command {
	case commandList:
		tasks, err := service.GetAllTasks()
		if err != nil {
			fmt.Println(err)
		}

		for _, task := range tasks {
			fmt.Printf("ID: %d Description: %s Status: %s", task.Id, task.Description, task.Status)
		}
	default:
		fmt.Printf("Unknown command: %s.", command)
	}
}
