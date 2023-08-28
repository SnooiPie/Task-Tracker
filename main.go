package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	Title       string
	Description string
	Priority    int
	DueDate     time.Time
	Completed   bool
}

var tasks []Task

func addTask() {
	var task Task
	fmt.Print("Task Title: ")
	fmt.Scanln(&task.Title)

	fmt.Print("Task Description: ")
	fmt.Scanln(&task.Description)

	fmt.Print("Priority (1-5, where 1 is highest): ")
	fmt.Scanln(&task.Priority)

	fmt.Print("Due Date (YYYY-MM-DD): ")
	dateStr := ""
	fmt.Scanln(&dateStr)
	task.DueDate, _ = time.Parse("2006-01-02", dateStr)

	task.Completed = false
	tasks = append(tasks, task)
	fmt.Println("Task added:", task.Title)
}

func listTasks() {
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. [%s] %s - Priority: %d, Due: %s, Status: %s\n", i+1, getStatusSymbol(task.Completed), task.Title, task.Priority, task.DueDate.Format("2006-01-02"), status)
	}
}

func markAsCompleted(index int) {
	if index >= 0 && index < len(tasks) {
		tasks[index].Completed = true
		fmt.Println("Task marked as completed:", tasks[index].Title)
	} else {
		fmt.Println("Invalid task index")
	}
}

func getStatusSymbol(completed bool) string {
	if completed {
		return "X"
	}
	return " "
}

func main() {
	for {
		fmt.Println("Task Tracker CLI App")
		fmt.Println("--------------------")

		fmt.Println("Choose an action:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()

		case 2:
			listTasks()

		case 3:
			listTasks()
			var index int
			fmt.Print("Enter task index to mark as completed: ")
			fmt.Scanln(&index)
			markAsCompleted(index - 1)

		case 4:
			fmt.Println("Exiting Task Tracker. Goodbye!")
			os.Exit(0)
		}

		fmt.Println()
	}
}
