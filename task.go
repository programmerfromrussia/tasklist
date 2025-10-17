package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Status  string    `json:"status"` // "todo", "done"
	Created time.Time `json:"created"`
}

var globalId = 1

func SaveTask(title string) error {
	if err := loadTasksFromJson(); err != nil {
		return err
	}
	if len(title) == 0 {
		return fmt.Errorf("title can't be empty")
	}
	task := &Task{
		ID:      globalId,
		Title:   title,
		Status:  "todo",
		Created: time.Now(),
	}
	tasks[task.ID] = task
	globalId++
	return saveTaskToJson()
}

func DeleteTask(id int) error {
	if err := loadTasksFromJson(); err != nil {
		return err
	}
	if _, ok := tasks[id]; !ok {
		return fmt.Errorf("task not found")
	}
	delete(tasks, id)
	return saveTaskToJson()
}

func ListTasks() ([]Task, error) {
	if err := loadTasksFromJson(); err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no tasks found")
	}

	var list []Task

	for _, task := range tasks {
		list = append(list, *task)
	}
	return list, nil
}

func DoneTask(id int) error {
	if err := loadTasksFromJson(); err != nil {
		return err
	}
	task, ok := tasks[id]
	if !ok {
		return fmt.Errorf("task not found")
	}
	task.Status = "done"
	return saveTaskToJson()
}
