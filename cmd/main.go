package main

import (
	"fmt"
	"os"

	"github.com/t8nax/task-tracker/internal/task/delivery/command"
	"github.com/t8nax/task-tracker/internal/task/repository"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: task-tracker <commands>")
		return
	}

	repo := &repository.JsonRepository{}
	uCase := usecase.NewTaskUseCase(repo)
	handler_factory := command.NewCommandHandlerFactory(uCase)

	cmd := command.Command(os.Args[1])

	handler, err := handler_factory.GetHandler(cmd)

	if err != nil {
		fmt.Println(err)
		return
	}

	messages, err := handler.Execute(os.Args)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
