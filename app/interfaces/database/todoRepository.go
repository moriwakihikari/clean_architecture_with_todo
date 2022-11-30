package database

import (
	"github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"
)

type TodoRepository struct {
	Sqlhandler
}

func (repo *TodoRepository) Store(t model.Todo) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"INSERT INTO todos (title, note, day_time, userid) VALUES (? ,? ,? ,? ) RETURNING id;", t.Title, t.Note, t.DayTime, t.UserID,
	)

	if err != nil {
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *TodoRepository) Update(t model.Todo) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"UPDATE todos SET title=?, note=?, day_time=?, userid=? WHERE id=? RETURNING id;",
		t.Title, t.Note, t.DayTime, t.UserID, t.ID,
	)

	if err != nil {
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *TodoRepository) Delete(id int) (err error) {
	_, err = repo.Sqlhandler.Query("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindById(id int) (todo model.Todo, err error) {
	row, err := repo.Sqlhandler.Query("SELECT id, title, note, day_time FROM todos WHERE id = ?;", id)
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DayTime)); err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindByUserId(userID int) (todoList model.Todos, err error) {
	query := `SELECT DISTINCT todos.id, title, note, day_time FROM todos INNER JOIN users ON todos.userid=users.id AND todos.userid=?;`
	rows, err := repo.Sqlhandler.Query(query, userID)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DayTime)); err != nil {
			continue
		}
		todoList = append(todoList, todo)
	}
	return
}

func (repo *TodoRepository) FindAll() (todoList model.Todos, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT id, title, note, day_time FROM todos;")
	if err != nil {
		return
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DayTime)); err != nil {
			continue
		}
		todoList = append(todoList, todo)
	}
	return
}
