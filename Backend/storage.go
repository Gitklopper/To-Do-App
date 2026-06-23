package main

import (
	"time"
	"fmt"
)

var todos = []ToDo{}

var ErrInvalidID = fmt.Errorf("ID cant be lower than 1")
var ErrTodoNotFound = fmt.Errorf("task could'nt be found")

func addTodo(title string) ToDo {
	var task ToDo
	task.ID = len(todos) + 1
	task.Title = title
	task.Status = ProgressionOpen
	task.Role = PriorityNeither
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	todos = append(todos, task)
	return task
}

func findTodoByID(id int) (*ToDo, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	for index := range todos {
		if id == todos[index].ID { 
			return &todos[index], nil
		}
	}
	return nil, ErrTodoNotFound
}

func setStatus(id int, newStatus Progression) error {
	task, err := findTodoByID(id)
	if err != nil {
		return err
	}
	task.Status = newStatus
	task.UpdatedAt = time.Now()
	return nil
}

func setPriority(id int, newPriority Priority) error {
	task, err := findTodoByID(id)
	if err != nil {
		return err
	}
	task.Role = newPriority
	task.UpdatedAt = time.Now()
	return nil
}
