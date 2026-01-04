package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	c "goFyneDesktopTodo/internal/context"
)

func GetTrashView(ctx *c.AppContext) *fyne.Container {

	navigateBackBtn := navigateBtn(ctx, theme.NavigateBackIcon(), c.List, "")

	bottomCont := container.NewBorder(nil, nil, navigateBackBtn, nil, nil)

	return container.NewBorder(
		nil,
		bottomCont,
		nil,
		nil,
		nil,
		)
}
