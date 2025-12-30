//go:build !prod

package configs

import c "github.com/emarifer/go-fyne-desktop-todoapp/internal/context"

const (
	AppId        = "ftodo_main"
	WindowTitle  = "fToDo App - a mini task manager"
	WindowWidth  = 480
	WindowHeight = 600
	WindowFixed  = false
	InitialRoute = c.List
	InitialTheme = c.Dark
	DbName       = "ftodo_data.db"
	EnableLogger = true
	Version      = "DEV"
)
