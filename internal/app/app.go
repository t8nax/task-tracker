package app

import (
	"fmt"
	"strconv"

	"github.com/t8nax/task-tracker/internal/models"
	tasksrv "github.com/t8nax/task-tracker/internal/services"
	jsonstorage "github.com/t8nax/task-tracker/internal/storage/json"
)

func Run(args []string) {
	if len(args) == 1 {
		fmt.Println("Usage: task-tracker <commands>")
		return
	}

	command := Command(args[1])
	storage := &jsonstorage.JsonStorage{}
	service := tasksrv.NewTaskService(storage)

	switch command {
	case commandList:
		tasks, err := service.GetAllTasks()
		if err != nil {
			fmt.Println(err)
		}

		for _, task := range tasks {
			fmt.Printf("ID: %d Description: %s Status: %s\n", task.ID, task.Description, task.Status)
		}
	case commandAdd:
		description := args[2]
		task, err := service.AddTask(description)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	case commandMarkDone, commandMarkInProgress:
		ID, err := strconv.ParseUint(args[2], 10, 64)

		if err != nil {
			fmt.Println("Unable to parse task ID")
			return
		}

		if command == commandMarkInProgress {
			_, err = service.MarkTask(ID, models.StatusInProgress)
		} else {
			_, err = service.MarkTask(ID, models.StatusDone)
		}

		if err != nil {
			fmt.Println(err)
		}
	case commandUpdate:
		ID, err := strconv.ParseUint(args[2], 10, 64)

		if err != nil {
			fmt.Println("Unable to parse task ID")
			return
		}

		description := args[3]

		_, err = service.UpdateDescription(ID, description)

		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
