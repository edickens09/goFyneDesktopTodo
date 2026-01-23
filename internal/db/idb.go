package db

import (
	"time"

	"goFyneDesktopTodo/internal/models"
)

type IDb interface {
	Close()
	DeleteTodo(todo *models.Todo) bool
	Drop() bool
	ExportData() bool
	GetAllTodos() []models.Todo
	ImportData() bool
	InsertTodo(todo *models.Todo) (*time.Time, bool)
	UpdateTodo(todo *models.Todo) bool
	UpdateTrash(todo *models.Todo) bool
	GetAllTrash() []models.Todo
	GetAllToday() []models.Todo
	UpdateToday(todo *models.Todo) bool
}
