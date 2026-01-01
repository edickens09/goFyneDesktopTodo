package ui

import (
/*	"fmt"
	"image/color"
	"net/url"
	
	"goFyneDesktopTodo/internal/services"
*/
	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
//	"fyne.io/fyne/v2/widget"

	c "goFyneDesktopTodo/internal/context"
)

func GetTodayView(ctx *c.AppContext) *fyne.Container {
	// create the widgets andn initial containers for views
	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")

	left := container.NewBorder(nil, navigateBackBtn, nil, nil)

	return container.NewBorder(
		nil,
		left,
		nil,
		nil,
		nil,
		)
}
