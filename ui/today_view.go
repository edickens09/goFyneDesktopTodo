package ui

import (
	"fmt"
	
	"goFyneDesktopTodo/internal/services"
	"goFyneDesktopTodo/internal/models"
	"goFyneDesktopTodo/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"

	c "goFyneDesktopTodo/internal/context"
)

func RenderListItemsToday() fyne.CanvasObject {

	return container.NewHBox(
		widget.NewCheck("", nil),
		widget.NewLabel(""),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
		)
}

func BindItemsToListToday(todos *services.Todos, w fyne.Window,
) func (di binding.DataItem, co fyne.CanvasObject) {
	return func (di binding.DataItem, co fyne.CanvasObject) {

		t := models.NewTodoFromDataItem(di)
		ctr, _ := co.(*fyne.Container)

		c := ctr.Objects[0].(*widget.Check)
		l := ctr.Objects[1].(*widget.Label)

		ctr.Objects[3].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("This is the edit button, tapping works now")
			dialog.ShowInformation("Confirm Here", msg, w)
		}

		ctr.Objects[4].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("Are you sure you want to Delete %q", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool) {
				if !b {
					return
				}
				todos.Remove(t)
				t.Trash = true
				todos.Dbase.UpdateTrash(t)

				if configs.EnableLogger {
					fmt.Printf("%q has been moved to the trash", t.Description)
				}
			}, w)
		}

		l.Bind(binding.BindString(&t.Description))
		c.Bind(binding.BindBool(&t.Selected))
	}
}

func GetTodayView(ctx *c.AppContext) *fyne.Container {

	todos := services.TodayTodosFromDb(ctx.Db)

	list := widget.NewListWithData(
		todos,
		RenderListItemsToday,
		BindItemsToListToday(&todos, ctx.GetWindow()),
		)
	//navigate buttons should have a variable that is tracked that is the last area that they were in and that's where they travel to
	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")
	settingsBtn := navigateBtn(ctx, theme.SettingsIcon(), c.Settings, "")

	bottomCont := container.NewBorder(nil, nil, navigateBackBtn, settingsBtn, nil)

	return container.NewBorder(
		nil,
		container.NewBorder(nil, bottomCont, nil, nil),
		nil,
		nil,
		list,
		)
}
