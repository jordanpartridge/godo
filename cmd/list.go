package cmd

import (
	"fmt"

	"github.com/jordanpartridge/godo/internal/repository"
)

type ListCommand struct {
	taskRepo *repository.TaskRepository
}

func NewListCommand(taskRepo *repository.TaskRepository) *ListCommand {
	return &ListCommand{taskRepo: taskRepo}
}

func (c *ListCommand) Execute() error {
	tasks, err := c.taskRepo.All()
	if err != nil {
		return fmt.Errorf("failed to fetch tasks: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return nil
	}

	fmt.Println("Your tasks:")
	for _, task := range tasks {
		status := "☐"
		if task.Completed {
			status = "☑"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}

	return nil
}
