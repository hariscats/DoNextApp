package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task

func addTask(description string) {
	task := Task{ID: len(tasks) + 1, Description: description, Completed: false}
	tasks = append(tasks, task)
	fmt.Println("Added task:", description)
}

func viewTasks() {
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, status)
	}
}

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

func main() {
	var command string
	for {
		fmt.Println("\nEnter command (add, view, delete, exit):")
		fmt.Scanln(&command)

		switch command {
		case "add":
			fmt.Println("Enter task description:")
			var description string
			fmt.Scanln(&description)
			addTask(description)
		case "view":
			viewTasks()
		case "delete":
			fmt.Println("Enter task ID to delete:")
			var id int
			fmt.Scanln(&id)
			deleteTask(id)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
