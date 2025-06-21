package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jordanpartridge/godo/cmd"
	"github.com/jordanpartridge/godo/internal/database"
	"github.com/jordanpartridge/godo/internal/repository"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db.DB)

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		addCmd := cmd.NewAddCommand(taskRepo)
		if err := addCmd.Execute(args); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		listCmd := cmd.NewListCommand(taskRepo)
		if err := listCmd.Execute(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "done":
		doneCmd := cmd.NewDoneCommand(taskRepo)
		if err := doneCmd.Execute(args); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "delete":
		fmt.Println("Delete command - TODO")
	case "help", "-h", "--help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("godo - A simple todo CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  godo add <task>     - Add a new task")
	fmt.Println("  godo list           - List all tasks")
	fmt.Println("  godo done <id>      - Mark task as done")
	fmt.Println("  godo delete <id>    - Delete a task")
	fmt.Println("  godo help           - Show this help")
}
