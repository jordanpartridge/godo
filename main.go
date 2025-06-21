package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "done":
		handleDone()
	case "delete":
		handleDelete()
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

func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: godo add <task>")
		return
	}
	
	// Join all args after "add" into one task
	task := strings.Join(os.Args[2:], " ")
	
	// TODO: Save to file
	fmt.Printf("Added task: %s\n", task)
}

func handleList() {
	fmt.Println("List command - TODO")
}

func handleDone() {
	fmt.Println("Done command - TODO")
}

func handleDelete() {
	fmt.Println("Delete command - TODO")
}
