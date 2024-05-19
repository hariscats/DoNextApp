package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task

// Add a new task to the list
func addTask(description string) {
	task := Task{ID: len(tasks) + 1, Description: description, Completed: false}
	tasks = append(tasks, task)
	fmt.Println("Added task:", description)
}

// View all tasks in the list
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, status)
	}
}

// Delete a task by its ID
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Deleted task ID:", id)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

// Mark a task as completed
func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Completed task ID:", id)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nEnter command (add, view, delete, complete, exit):")
		scanner.Scan()
		command := strings.ToLower(scanner.Text())

		switch command {
		case "add":
			fmt.Println("Enter task description:")
			scanner.Scan()
			description := scanner.Text()
			addTask(description)
		case "view":
			viewTasks()
		case "delete":
			fmt.Println("Enter task ID to delete:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID format")
				continue
			}
			deleteTask(id)
		case "complete":
			fmt.Println("Enter task ID to complete:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID format")
				continue
			}
			completeTask(id)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
