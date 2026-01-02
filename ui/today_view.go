package ui

import (
/*	"fmt"
	"image/color"
	
	"goFyneDesktopTodo/internal/services"
*/
	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	//"fyne.io/fyne/v2/widget"

	c "goFyneDesktopTodo/internal/context"
)

func GetTodayView(ctx *c.AppContext) *fyne.Container {
	// create the widgets andn initial containers for views
	//navigate buttons should have a variable that is tracked that is the last area that they were in and that's where they travel to
	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")
	settingsBtn := navigateBtn(ctx, theme.SettingsIcon(), c.Settings, "")
	//trashBtn := navigateBtn(ctx theme.DeleteIcon(), c.Trash, "")

	bottomCont := container.NewBorder(nil, nil, navigateBackBtn, settingsBtn, nil)

	return container.NewBorder(
		nil,
		container.NewBorder(nil, bottomCont, nil, nil),
		nil,
		nil,
		nil,
		)
}
