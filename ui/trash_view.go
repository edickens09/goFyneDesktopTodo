package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"

	"goFyneDesktopTodo/internal/services"
	"goFyneDesktopTodo/internal/models"

	c "goFyneDesktopTodo/internal/context"
)

// right now I don't know how to add extra buttons to the area in the container and it not throw an error when returning the CanvasObject
func RenderListItemsTrash() fyne.CanvasObject {

	return container.NewHBox(
		widget.NewCheck("", nil), //left
		widget.NewLabel(""),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), nil),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
		)
}

func BindItemsToListTrash(todos *services.Todos, w fyne.Window,
) func(di binding.DataItem, co fyne.CanvasObject) {

	return func(di binding.DataItem, co fyne.CanvasObject) {

		t := models.NewTodoFromDataItem(di)
		ctr, _ := co.(*fyne.Container)

		c := ctr.Objects[0].(*widget.Check)
		l := ctr.Objects[1].(*widget.Label)
		ctr.Objects[4].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("Are you sure you want to permanently delete the task with Description %q", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool) {

				if !b {
					return
				}

				todos.Remove(t)
				todos.Dbase.DeleteTodo(t)
			
			}, w)
		}
		// putting a button here for future use
		ctr.Objects[3].(*widget.Button).OnTapped = func() { 
			msg:= fmt.Sprintf("This button works %q", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool){
				if !b {
					return
				}
			}, w)
		}
		
		/* I think there is a better way impliment this that is more dynamic but this seems 
		to work on my laptop in half screen mode so that is good enough for now */
		if (len(t.Description) > 60) {

			ellip := t.Description[:60] +"..."

			l.Bind(binding.BindString(&ellip))

		}else {
			l.Bind(binding.BindString(&t.Description))
		}

		c.Bind(binding.BindBool(&t.Selected))

		//l.Truncation = fyne.TextTruncateEllipsis
		c.OnChanged = func(b bool) {
			t.Selected = b
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
