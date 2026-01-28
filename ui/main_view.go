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
	"fyne.io/fyne/v2/layout"

	c "goFyneDesktopTodo/internal/context"
)

// TODO create renderTreeSubItem so they can be different
func renderListItem() fyne.CanvasObject {

	return container.NewHBox(
		widget.NewCheck("", nil),
		widget.NewLabel(""),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), nil),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
		)
}

// This is where we take the base List item and then Bind the data to that singular item
func bindDataToList(todos *services.Todos, w fyne.Window,
) func(di binding.DataItem, co fyne.CanvasObject) {

	return func(di binding.DataItem, co fyne.CanvasObject) {
		//creating the binding variables to be used later in the function
		labelData := binding.NewString()
		checkData := binding.NewBool()

		t := models.NewTodoFromDataItem(di)
		ctr, _ := co.(*fyne.Container)

		c := ctr.Objects[0].(*widget.Check)
		l := ctr.Objects[1].(*widget.Label)
		ctr.Objects[4].(*widget.Button).OnTapped = func() {
			msg := fmt.Sprintf("Are you sure you want to delete the task with Description %q?", t.Description)
			dialog.ShowConfirm("Confirmation", msg, func(b bool) {
				if !b {
					return
				}
				//todos.Remove is what removes it from the displayed list. other moves it to trash
				todos.Remove(t)
				t.Trash = true
				todos.Dbase.UpdateTrash(t)

				if configs.EnableLogger {
					fmt.Printf("The ToDo %q has been successfully moved to trash!\n", t.Description)
				}
			}, w)
		}

		ctr.Objects[3].(*widget.Button).OnTapped = func() {
			t.Today = true
			todos.Dbase.UpdateToday(t)

			if configs.EnableLogger {
				fmt.Printf("The ToDo %q has been successfully added to today\n", t.Description)
			}
		}

		//binding the label to the todo item descriptin
		if (len(t.Description) > 60) {
			labelData.Set(t.Description[:60] + "...")
			l.Bind(labelData)	
		}else {
			labelData.Set(t.Description)
			l.Bind(labelData)
		}

		//binding the bool to the check box todo item selection
		checkData.Set(t.Selected)
		c.Bind(checkData)
		
		
		//this is where the todo item show is being updated based on whether the button is clicked
		c.OnChanged = func(b bool) {
			t.Selected = b
			todos.Dbase.UpdateTodo(t)
		}
	}
}
// Function created so that OnSubmit and Add button add information in the same way.
func AddToList(todos *services.Todos, input *widget.Entry) {
	t := models.NewTodo(input.Text)
	todos.Add(&t)
	input.SetText("")
}

func GetMainView(ctx *c.AppContext) *fyne.Container {
	// Get data from the DB and bind it to an UntypedList
	todos := services.NewTodosFromDb(ctx.Db)

	// Setup Widgets
	input := widget.NewEntry()
	input.PlaceHolder = "New Task...."
	// here is where the enter key function exists
	input.OnSubmitted = func(s string) {
		if len(s) > 2 {
			AddToList(&todos, input)
		}
	}
	addBtn := widget.NewButtonWithIcon(
		"Add", theme.DocumentCreateIcon(), func() {
			AddToList(&todos, input)

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

	// Implementing my navigation buttons here
	todayBtn := navigateBtn(ctx, theme.ListIcon(), c.Today, "Today Tasks")
	trashBtn := navigateBtn(ctx, theme.DeleteIcon(), c.Trash, "Trash")

	settingsBtn := navigateBtn(ctx, theme.SettingsIcon(), c.Settings, "")

	bottomCont := container.NewBorder(nil, nil, settingsBtn, trashBtn, todayBtn)
	//creating tree to be generated moving from list for sub items
	//tree := widget.NewTreeWithData(
	list := widget.NewListWithData(
		// the binding.List type
		todos,
		// func that returns the component structure of the List Item
		// exactly the same as the Simple List
		renderListItem,
		// func that is called for each item in the list and allows
		// but this time we get the actual DataItem we need to cast
		bindDataToList(&todos, ctx.GetWindow()),
	)
	// need to have tree.OnSelected
	list.OnSelected = func(id widget.ListItemID) {
		if configs.EnableLogger {
			fmt.Printf("Selected item: %d\n", id)
		}
	}

	return container.NewBorder(
		nil, // TOP of the container
		// this will be a the BOTTOM of the container
		container.NewBorder(
			nil,//container.NewBorder (nil, nil, nil, todayBtn, nil),    //TOP
			bottomCont,  // BOTTOM
			nil,         // LEFT
			addBtn,      // RIGHT
			input,       // take the rest of the space
		),
		nil,  // Left
		nil,  // Right
		//adding tree here for when change over happens
		//tree,
		list, // the rest will take all the rest of the space
	)
}
