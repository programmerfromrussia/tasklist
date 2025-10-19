package main

import (
	"fmt"
	"os"
	"strconv"
)

var tasks = make(map[int]*Task)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no tasks specified")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: add <task name> <task description>")
			os.Exit(1)
		}
		title := os.Args[2]
		if err := SaveTask(title); err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Saved task:", title)
	case "list":
		taskList, err := ListTasks()
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		for _, task := range taskList {
			fmt.Printf("%d. [%s] %s (%s)\n", task.ID, task.Status, task.Title, task.Created.Format("2006-01-02"))
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("usage: done <task name> <task description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid task id")
			os.Exit(1)
		}
		if err := DoneTask(id); err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Done task:", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("usage: delete <task name> <task description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid task id")
			os.Exit(1)
		}

		if err := DeleteTask(id); err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}
		fmt.Println("Deleted task:", id)
	default:
		fmt.Println("error: invalid command")
	}
}
