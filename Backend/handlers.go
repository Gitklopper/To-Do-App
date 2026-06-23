package main

import (
	"net/http"
	"encoding/json"
	"errors"
)

type CreateTodoRequest struct {
	Title string	`json:"title"`
}

type UpdateTodoRequest struct {
	ID int				`json:"id"`
	Status Progression	`json:"status"`
	Role Priority		`json:"priority"`
}

func getTodos() []ToDo {
	return todos
}

func whichError(err error, w http.ResponseWriter) {
	if errors.Is(err, ErrTodoNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else if errors.Is(err, ErrInvalidID) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var request UpdateTodoRequest
	json.NewDecoder(r.Body).Decode(&request)

	if request.Status != "" {
		err := setStatus(request.ID, request.Status)
		if err != nil {
			whichError(err, w)
			return
		}
	}
	if request.Role != "" {
		err := setPriority(request.ID, request.Role)
		if err != nil {
			whichError(err, w)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getTodos())
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var request CreateTodoRequest
		json.NewDecoder(r.Body).Decode(&request)
		task := addTodo(request.Title)
		json.NewEncoder(w).Encode(task)
	} else {
		json.NewEncoder(w).Encode(getTodos())
	}
}
