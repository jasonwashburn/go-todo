package models

import (
	"database/sql"
	"log"
)

type TodoModel struct {
	DB *sql.DB
}

func NewTodoModel(db *sql.DB) *TodoModel {
	stmt := `CREATE TABLE IF NOT EXISTS todos(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		completed BOOLEAN
	)`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatalf("Error creating todos table: %v", err)
	}
	return &TodoModel{DB: db}
}

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (m *TodoModel) Insert(title string) (int, error) {
	stmt := `INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id`
	result, err := m.DB.Exec(stmt, title, false)
	if err != nil {
		log.Printf("Error inserting new todo: %v", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v", err)
		return 0, err
	}
	return int(id), nil
}

func (m *TodoModel) GetById(id int) (*Todo, error) {
	stmt := `SELECT id, title, completed FROM todos WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)
	t := &Todo{}
	err := row.Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		log.Printf("Error getting todo by id: %v", err)
		return nil, err
	}
	return t, nil
}

func (m *TodoModel) GetAll() ([]*Todo, error) {
	stmt := `SELECT id, title, completed FROM todos`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Printf("Error getting all todos: %v", err)
		return nil, err
	}
	defer rows.Close()
	todos := []*Todo{}
	for rows.Next() {
		t := &Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			log.Printf("Error getting todo by id: %v", err)
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}
