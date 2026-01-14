package ui

import (
	"fmt"
//	"image/color"
	
	"goFyneDesktopTodo/internal/services"
	"goFyneDesktopTodo/internal/models"

	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
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
	}
}

func GetTodayView(ctx *c.AppContext) *fyne.Container {
	// create the widgets andn initial containers for views
	//navigate buttons should have a variable that is tracked that is the last area that they were in and that's where they travel to
	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")
	settingsBtn := navigateBtn(ctx, theme.SettingsIcon(), c.Settings, "")

	bottomCont := container.NewBorder(nil, nil, navigateBackBtn, settingsBtn, nil)

	return container.NewBorder(
		nil,
		container.NewBorder(nil, bottomCont, nil, nil),
		nil,
		nil,
		nil,
		)
}
