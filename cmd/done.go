package cmd

import (
	"fmt"
	"strconv"

	"github.com/jordanpartridge/godo/internal/repository"
)

type DoneCommand struct {
	taskRepo *repository.TaskRepository
}

func NewDoneCommand(taskRepo *repository.TaskRepository) *DoneCommand {
	return &DoneCommand{taskRepo: taskRepo}
}

func (c *DoneCommand) Execute(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: godo done <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	err = c.taskRepo.MarkComplete(id)
	if err != nil {
		return fmt.Errorf("failed to mark task complete: %w", err)
	}

	fmt.Printf("✅ Marked task #%d as complete!\n", id)
	return nil
}