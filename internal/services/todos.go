package services

import (
	"time"

	"goFyneDesktopTodo/internal/db"
	"goFyneDesktopTodo/internal/models"

	"fyne.io/fyne/v2/data/binding"
)

// TODO change from list to Tree for subitems
// TODO have seperate struct for an untypedList and Tree
type Todos struct {
	binding.UntypedList // composition
//	binding.UntypedTree
	Dbase               db.IDb
}

func NewTodosFromDb(db db.IDb) Todos {
	todoList := db.GetAllTodos()

	return newTodos(db, todoList)
}

func TrashTodosFromDb(db db.IDb) Todos {
	trashList := db.GetAllTrash()

	return newTodos(db, trashList)
}

//might need a different on so that it shows as a list rather than a tree
func TodayTodosFromDb(db db.IDb) Todos {
	todayList := db.GetAllToday()

	return newTodos(db, todayList)
}

func newTodos(db db.IDb, todos []models.Todo) Todos {
	t := Todos{
		binding.NewUntypedList(),
	//	binding.NewUntypedTree(),
		db,
	}

	for _, td := range todos {
		t.Add(&td)
	}

	return t
}

func (t *Todos) Add(todo *models.Todo) {
	// If created_at is the value 'zero' of time.Time,
	// we insert the data into the DB
	var dt *time.Time
	if todo.CreatedAt.String() == "0001-01-01 00:00:00 +0000 UTC" {
		dt, _ = t.Dbase.InsertTodo(todo)
		todo.CreatedAt = *dt
	}

	t.Prepend(todo)
	
}

// This function doesn't seem to be used currently, I will leave it
// for the time being

/*func (t *Todos) All() []*models.Todo {
	result := []*models.Todo{}
	for i := 0; i < t.Length(); i++ {
		di, err := t.GetItem(i)
		if err != nil {
			break
		}
		result = append(result, models.NewTodoFromDataItem(di))
	}

	return result
}*/

func (t *Todos) Drop() {
	t.Dbase.Drop()

	// list, _ := t.Get()
	// list = list[:0]
	t.Set([]any{})
}
