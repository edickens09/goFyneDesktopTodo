package services

import (
	"time"

	"goFyneDesktopTodo/internal/db"
	"goFyneDesktopTodo/internal/models"

	"fyne.io/fyne/v2/data/binding"
)

type TodosTree struct {
	binding.UntypedTree
	Dbase db.IDb
}

func newTodosTree(db db.IDb, todos []models.Todo) TodosTree {
	t := TodosTree {
		binding.NewUntypedTree(),
		db,
	}

	for _, td := range todos {
		t.addTree(&td)
	}

	return t
}

func (t *TodosTree) addTree(todo *models.Todo) {

	var dt *time.Time
	if todo.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
		dt, _ = t.Dbase.InsertTodo(todo)
		todo.CreatedAt = *dt
	}

	//t.prependTree(todo)

}

// need to be created to finish tree service
func prependTree(
	tree binding.UntypedTree,
	parentId string,
	itemId string,
	item binding.DataItem,
) {
}
