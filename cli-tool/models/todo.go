package models

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Id          string
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Status      bool      `json:"status"`
}

const filename = "tasks.json"

func SaveTasks(tasks []Todo) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.New("cannot create a file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func LoadTasks() ([]Todo, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Todo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return []Todo{}, nil
	}

	return tasks, nil
}

func UpdateTasks(id string) ([]Todo, error) {
	if id == "" {
		return nil, errors.New("ID is null")
	}

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, errors.New("error occurred in updating")
	}
	defer file.Close()

	var tasks []Todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, errors.New("error decoding tasks")
	}

	taskIndex, err := strconv.Atoi(id)
	if err != nil || taskIndex < 0 || taskIndex >= len(tasks) {
		return nil, errors.New("invalid task ID")
	}

	tasks[taskIndex].Status = !tasks[taskIndex].Status

	if err := SaveTasks(tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
