package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoRepository interface {
	Create(todo Todo) (int, error)
	Update(id int, todo Todo) error
	Delete(id int) error
	Get(id int) (*Todo, error)
	GetAll() ([]Todo, error)
}

type InMemoryTodoRepository struct {
	todos  map[int]Todo
	autoID int
}

func NewInMemoryRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos:  make(map[int]Todo),
		autoID: 0,
	}
}

func (r *InMemoryTodoRepository) Create(todo Todo) (int, error) {
	todo.ID = r.autoID
	r.todos[r.autoID] = todo
	r.autoID++
	return todo.ID, nil
}

func (r *InMemoryTodoRepository) Get(id int) (*Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}

func (r *InMemoryTodoRepository) GetAll() ([]Todo, error) {
	todos := []Todo{}
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *InMemoryTodoRepository) Update(id int, todo Todo) error {
	if _, ok := r.todos[id]; !ok {
		return errors.New("todo not found")
	}
	if id != todo.ID {
		return errors.New("id cannot be updated")
	}
	r.todos[id] = todo
	return nil
}

func (r *InMemoryTodoRepository) Delete(id int) error {
	if _, ok := r.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete(r.todos, id)
	return nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Get("/todos", getTodosHandler)

	http.ListenAndServe(":8080", r)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	repo := NewInMemoryRepository()
	todos, err := repo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
