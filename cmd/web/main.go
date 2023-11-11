package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jasonwashburn/go-todo/internal/models"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type Application struct {
	todos *models.TodoModel
}

func main() {
	var dbUrl = "file:./db.sqlite"
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	todoModel := models.NewTodoModel(db)
	app := &Application{
		todos: todoModel,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Get("/todos", app.getTodosHandler)
	r.Post("/todos", app.postTodoHandler)

	http.ListenAndServe(":8080", r)
}
