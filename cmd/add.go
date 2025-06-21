package cmd

import (
	"fmt"
	"strings"

	"github.com/jordanpartridge/godo/internal/repository"
)

type AddCommand struct {
	taskRepo *repository.TaskRepository
}

func NewAddCommand(taskRepo *repository.TaskRepository) *AddCommand {
	return &AddCommand{taskRepo: taskRepo}
}

func (c *AddCommand) Execute(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: godo add <task>")
	}

	title := strings.Join(args, " ")

	task, err := c.taskRepo.Create(title, "")
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	fmt.Printf("✅ Added task #%d: %s\n", task.ID, task.Title)
	return nil
}
