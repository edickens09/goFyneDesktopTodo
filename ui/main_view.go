package ui

import (
	"fmt"

	"goFyneDesktopTodo/configs"
	"goFyneDesktopTodo/internal/models"
	"goFyneDesktopTodo/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	c "goFyneDesktopTodo/internal/context"
)

type tappableEntry struct {
	widget.Entry
}

//this is where the info lives on the app. Change info? Get rid of? Unsure what to do with
//Don't forget there is an area down below that sets this back to display when it gets cleared
func newTappableEntry() *tappableEntry {
	e := &tappableEntry{
		widget.Entry{
			PlaceHolder: "What should I do with this?",
			TextStyle:   fyne.TextStyle{Monospace: true},
		},
	}
	e.ExtendBaseWidget(e)

	return e
}

//I'm not sure what this does yet
func (e *tappableEntry) Tapped(_ *fyne.PointEvent) {
	e.Disable()
}

func renderListItem() fyne.CanvasObject {
	return container.NewBorder(
		nil, nil, // Top & bottom
		// ↓ left of the border ↓
		widget.NewCheck("", nil), // func(b bool) {}
		// ↓ right of the border ↓
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
		// take the rest of the space ↓
		widget.NewLabel(""),
	)
}

//I believe this is where the work is being done for the individual list items
func bindDataToList(
	displayText *tappableEntry, todos *services.Todos, w fyne.Window,
) func(di binding.DataItem, co fyne.CanvasObject) {
	return func(di binding.DataItem, co fyne.CanvasObject) {
		t := models.NewTodoFromDataItem(di)
		ctr, _ := co.(*fyne.Container)

		l := ctr.Objects[0].(*widget.Label)
		c := ctr.Objects[1].(*widget.Check)
		ctr.Objects[2].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("Are you sure you want to delete the task with Description %q?", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool) {
				if !b {
					return
				}
				todos.Remove(t)
				todos.Dbase.DeleteTodo(t)

				if configs.EnableLogger {
					fmt.Printf("The ToDo with description %q has been successfully removed!\n", t.Description)
				}
				displayText.SetText(fmt.Sprintf("%q has been successfully removed!", t.Description))
			}, w)
		}
		
		//I'm not entirely sure what is happening here, but it is somehow binding the list object item to the database object
		l.Bind(binding.BindString(&t.Description))
		c.Bind(binding.BindBool(&t.Done))
		
		//this is where the todo item show is being updated based on whether the button is clicked
		l.Truncation = fyne.TextTruncateEllipsis
		c.OnChanged = func(b bool) {
			t.Done = b
			todos.Dbase.UpdateTodo(t)
		}
	}
}

func GetMainView(ctx *c.AppContext) *fyne.Container {
	// Get data from the DB and bind it to an UntypedList
	todos := services.NewTodosFromDb(ctx.Db)

	// Setup Widgets
	input := widget.NewEntry()
	input.PlaceHolder = "New Task...."
	input.OnSubmitted = func(s string) {
		if len(s) > 2 {
			t := models.NewTodo(input.Text)
			todos.Add(&t)
			input.SetText("")
		}
	}
	addBtn := widget.NewButtonWithIcon(
		"Add", theme.DocumentCreateIcon(), func() {
			t := models.NewTodo(input.Text)
			todos.Add(&t)
			input.SetText("")
		},
	)
	addBtn.Disable()
	input.OnChanged = func(s string) {
		// ↓ so that if we delete characters it will be disabled again ↓
		addBtn.Disable()
		if len(s) > 2 {
			addBtn.Enable()
		}
	}

	// today button, needs to be implimented only a placeholder right now
	// how can I make it say text?
	todayBtn := navigateBtn(ctx, theme.ListIcon(), c.Today, "Today")

	displayText := newTappableEntry()

	deleteBtn := widget.NewButtonWithIcon(
		"Reset", theme.ViewRefreshIcon(), func() {
			dialog.ShowConfirm(
				"Confirmation",
				"Are you sure you want to delete all the data you have saved? This action is irreversible!!",
				func(b bool) {
					if !b {
						return
					}

					todos.Drop()

					displayText.SetText("Display")
				}, ctx.GetWindow(),
			)
		},
	)

	settingsBtn := navigateBtn(ctx, theme.SettingsIcon(), c.Settings, "")

	bottomCont := container.NewBorder(nil, nil, nil, settingsBtn, deleteBtn)

	list := widget.NewListWithData(
		// the binding.List type
		todos,
		// func that returns the component structure of the List Item
		// exactly the same as the Simple List
		renderListItem,
		// func that is called for each item in the list and allows
		// but this time we get the actual DataItem we need to cast
		bindDataToList(displayText, &todos, ctx.GetWindow()),
	)
	list.OnSelected = func(id widget.ListItemID) {
		t := todos.All()
		displayText.SetText(t[id].String())
		displayText.Enable()
		if configs.EnableLogger {
			fmt.Printf("Selected item: %d\n", id)
		}
	}
	// the directions are not important the order is important
	return container.NewBorder(
		nil, // TOP of the container
		// this will be a the BOTTOM of the container
		container.NewBorder(
			container.NewBorder (nil, nil, nil, todayBtn, displayText),    //TOP
			bottomCont,  // BOTTOM
			nil,         // LEFT
			addBtn,      // RIGHT
			input,       // take the rest of the space
		),
		nil,  // Left
		nil,  // Right
		list, // the rest will take all the rest of the space
	)
}
