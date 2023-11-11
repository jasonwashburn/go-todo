package main

import "testing"

func TestInMemoryTodoRepositoryCreate(t *testing.T) {
	repo := NewInMemoryRepository()
	todo := Todo{
		Title:     "Test",
		Completed: false,
	}
	id, err := repo.Create(todo)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	if id != 0 {
		t.Fatalf("expected id to be 0, got %d", id)
	}
	want := Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}
	if repo.todos[0] != want {
		t.Fatalf("got %v; want: %v", repo.todos[0], want)
	}
}

func TestInMemoryTodoRepositoryGet(t *testing.T) {
	repo := NewInMemoryRepository()
	want := Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}
	repo.todos[0] = want

	todo, err := repo.Get(0)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if *todo != want {
		t.Fatalf("got %v; want: %v", todo, want)
	}
}

func TestInMemoryTodoRepositoryGetNotFound(t *testing.T) {
	repo := NewInMemoryRepository()
	_, err := repo.Get(0)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInMemoryTodoRepositoryUpdate(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.todos[0] = Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}

	want := Todo{
		ID:        0,
		Title:     "Test",
		Completed: true,
	}

	err := repo.Update(0, want)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	if repo.todos[0] != want {
		t.Fatalf("got %v; want: %v", repo.todos[0], want)
	}
}

func TestInMemoryTodoRepositoryUpdateNotFound(t *testing.T) {
	repo := NewInMemoryRepository()
	err := repo.Update(0, Todo{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInMemoryTodoRepositoryUpdateID(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.todos[0] = Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}

	err := repo.Update(0, Todo{
		ID:        1,
		Title:     "Test",
		Completed: true,
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInMemoryTodoRepositoryDelete(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.todos[0] = Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}

	err := repo.Delete(0)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if _, ok := repo.todos[0]; ok {
		t.Fatalf("expected todo to be deleted")
	}
}

func TestInMemoryTodoRepositoryDeleteNonExisting(t *testing.T) {
	repo := NewInMemoryRepository()
	err := repo.Delete(0)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestInMemoryTodoRepositoryGetAll(t *testing.T) {
	repo := NewInMemoryRepository()
	repo.todos[0] = Todo{
		ID:        0,
		Title:     "Test",
		Completed: false,
	}
	repo.todos[1] = Todo{
		ID:        1,
		Title:     "Test 2",
		Completed: true,
	}

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	for index, todo := range todos {
		if todo != repo.todos[index] {
			t.Fatalf("got %v; want: %v", todo, repo.todos[index])
		}
	}
}
