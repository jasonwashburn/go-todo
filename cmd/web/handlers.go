package main

import (
	"encoding/json"
	"net/http"

	"github.com/jasonwashburn/go-todo/internal/models"
)

func (app *Application) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.todos.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) postTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := app.todos.Insert(todo.Title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	todo.ID = id
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
