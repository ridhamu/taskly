package internal

import (
	"encoding/json"
	"os"
	"time"
)

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func LoadTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename) // reading the file

	if err != nil {
		if os.IsNotExist(err) { // check if the error because file doesn't exist
			return []Task{}, nil // return empty if file doesn't exist
		}
		return nil, err // return other error
	}

	var task []Task
	err = json.Unmarshal(file, &task)
	return task, err
}

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0777)
}
