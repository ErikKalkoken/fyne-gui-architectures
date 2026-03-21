package model

import (
	"database/sql"
	"fmt"
)

type Model struct {
	db          *sql.DB
	isErrorMode bool
}

func NewModel(db *sql.DB, isErrorMode bool) *Model {
	m := &Model{
		db:          db,
		isErrorMode: isErrorMode,
	}
	return m
}

func (m *Model) Init() error {
	sqlStmt := "CREATE TABLE IF NOT EXISTS tasks (title TEXT NOT NULL);"
	_, err := m.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) AddTask(title string) error {
	if m.isErrorMode {
		return fmt.Errorf("adding task error")
	}
	sqlStmt := "INSERT INTO tasks (title) VALUES (?)"
	_, err := m.db.Exec(sqlStmt, title)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) DeleteTask(title string) error {
	sqlStmt := "DELETE FROM tasks WHERE title = ?"
	_, err := m.db.Exec(sqlStmt, title)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) ListTasks() ([]string, error) {
	queryStmt := "SELECT title FROM tasks ORDER BY title"
	rows, err := m.db.Query(queryStmt)
	if err != nil {
		return nil, err
	}

	var tasks []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, title)
	}
	return tasks, nil
}
