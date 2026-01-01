package ui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"goFyneDesktopTodo/internal/context"
)

func h1(text string) *canvas.Text {
	txt := canvas.NewText(text, color.RGBA{61, 133, 255, 255})
	txt.TextSize = 20
	txt.TextStyle = fyne.TextStyle{Bold: true}

	return txt
}

func small(text string) *canvas.Text {
	txt := canvas.NewText(text, color.RGBA{61, 133, 255, 255})
	txt.TextSize = 12
	txt.Alignment = fyne.TextAlignCenter
	txt.TextStyle = fyne.TextStyle{Bold: true}

	return txt
}

func centered(obj fyne.CanvasObject) *fyne.Container {

	return container.NewCenter(obj)
}

/* func leftAligned(obj fyne.CanvasObject) *fyne.Container {

	return container.NewBorder(nil, nil, obj, nil)
}

func rightAligned(obj fyne.CanvasObject) *fyne.Container {

	return container.NewBorder(nil, nil, nil, obj)
} */

func successMessage(msg string, textItem *canvas.Text) {
	flasMessage(
		msg,
		textItem,
		time.Second*3,
		color.RGBA{0, 255, 50, 255},
	)
}

func errorMessage(msg string, textItem *canvas.Text) {
	flasMessage(
		msg,
		textItem,
		time.Second*3,
		color.RGBA{255, 0, 50, 255},
	)
}

func flasMessage(
	msg string,
	textItem *canvas.Text,
	duration time.Duration,
	color color.Color,
) {
	textItem.Color = color
	textItem.Text = msg

	go func() {
		time.Sleep(duration)
		textItem.Text = ""
	}()
}

//this is the button used for navigation, can I adapt this so that it works for the today route?
func navigateBtn(
	ctx *context.AppContext, icon fyne.Resource, route context.AppRoute, name string,
) *widget.Button {

	return widget.NewButtonWithIcon(name, icon, func() {
		ctx.NavigateTo(route)
	})
}

func toggleThemeBtn(
	ctx *context.AppContext, icon fyne.Resource,
) *widget.Button {

	return widget.NewButtonWithIcon("Toggle Theme", icon, func() {
		switch ctx.CurrentTheme() {
		case context.Dark:
			ctx.ChangeThemeTo(context.Light)
		case context.Light:
			ctx.ChangeThemeTo(context.Dark)
		}
	})
}

func newFlashTxtPlaceholder() *canvas.Text {
	return canvas.NewText("", theme.Color(theme.ColorNameForeground))
}
