package main

import (
	"encoding/json"
	"io"
	"os"
)

var dataFile = "tasks.json"

func loadTasksFromJson() error {
	file, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(bytes) == 0 {
		tasks = make(map[int]*Task)
		globalId = 1
		return nil
	}

	var list []Task
	if err := json.Unmarshal(bytes, &list); err != nil {

		return err

	}

	tasks = make(map[int]*Task)
	maxID := 0
	for i := range list {
		tasks[list[i].ID] = &list[i]
		if list[i].ID > maxID {
			maxID = list[i].ID
		}
	}
	globalId = maxID + 1
	return nil
}

func saveTaskToJson() error {
	var list []Task
	for _, task := range tasks {
		list = append(list, *task)
	}
	bytes, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, bytes, 0644)
}
