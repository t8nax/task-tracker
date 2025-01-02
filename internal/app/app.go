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
			fmt.Printf("ID: %d Description: %s Status: %s\n", task.Id, task.Description, task.Status)
		}
	case commandAdd:
		description := args[2]

		if description == "" {
			fmt.Println("To add task description must be specified")
			return
		}

		task, err := service.AddTask(description)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", task.Id)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
