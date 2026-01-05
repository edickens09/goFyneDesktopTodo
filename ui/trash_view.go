package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"

	"goFyneDesktopTodo/internal/services"
	"goFyneDesktopTodo/internal/models"

	c "goFyneDesktopTodo/internal/context"
)

func RenderListItemsTrash() fyne.CanvasObject {
	/*
	rightCon := container.NewBorder(
		nil, nil,
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),

		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), nil),

		nil,
		) */
	return container.NewBorder(
		nil, nil,
		widget.NewCheck("", nil),
		
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),

		widget.NewLabel(""),
		)
}

func BindItemsToListTrash(todos *services.Todos, w fyne.Window,
) func(di binding.DataItem, co fyne.CanvasObject) {

	return func(di binding.DataItem, co fyne.CanvasObject) {

		t := models.NewTodoFromDataItem(di)
		ctr, _ := co.(*fyne.Container)

		l := ctr.Objects[0].(*widget.Label)
		c := ctr.Objects[1].(*widget.Check)
		ctr.Objects[2].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("Are you sure you want to permanently delete the task with Description %q", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool) {

				if !b {
					return
				}

				todos.Remove(t)
				todos.Dbase.DeleteTodo(t)
			
			}, w)
		}
		l.Bind(binding.BindString(&t.Description))
		c.Bind(binding.BindBool(&t.Done))

		l.Truncation = fyne.TextTruncateEllipsis
		c.OnChanged = func(b bool) {
			t.Done = b
			todos.Dbase.UpdateTodo(t)
		}

	}

}


func GetTrashView(ctx *c.AppContext) *fyne.Container {
	
	todos := services.TrashTodosFromDb(ctx.Db)

	list := widget.NewListWithData(

		todos,

		RenderListItemsTrash,

		BindItemsToListTrash(&todos, ctx.GetWindow()),

		)

	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")

	bottomCont := container.NewBorder(nil, nil, navigateBackBtn, nil, nil)

	return container.NewBorder(
		nil,
		bottomCont,
		nil,
		nil,
		list,
		)
}
