package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// InitialDisplay loops though and produces the text for all existing data.
func (u *user) InitialDisplay() (text string) {
	// We loop through the imported file and add the formated info before the previous info (new information comes out on top).
	for i := range u.Data.Exercise {
		text = u.Data.Format(i) + text
	}

	return text
}

// ExerciseDisplayer runs in the background and updated the label.
func (u *user) ExerciseDisplayer(label *widget.Label) {
	// Handle an empty data file.
	if len(u.Data.Exercise) == 0 {
		// Start by informing  the user that no data is available.
		label.SetText("No exercieses have been created yet.")
	} else {
		// Refresh the widget to show the updated text.
		label.SetText(u.InitialDisplay())
	}

	// We then block the channel while waiting for updates on the channel.
	for {
		select {
		case exercise := <-u.NewExercise:
			label.SetText(exercise + label.Text)
		case exercise := <-u.FirstExercise:
			label.SetText(exercise)
		case <-u.ReorderExercises:
			label.SetText(u.InitialDisplay())
		case <-u.EmptyExercises:
			label.SetText("No exercieses have been created yet.")
		}

	}
}

// ShowMainDataView shows the main view after we are logged in.
func ShowMainDataView(window fyne.Window, app fyne.App, user *user) {
	// Create a label for displaing some info for the user. Default to showing nothing.
	dataLabel := widget.NewLabel("")

	// Start up the function to handle adding exercises in the background.
	go user.ExerciseDisplayer(dataLabel)

	// Tab data for the main window.
	dataPage := widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), dataLabel))

	// Create tabs with data.
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Exercises", theme.HomeIcon(), dataPage),
		widget.NewTabItemWithIcon("Add Exercise", theme.ContentAddIcon(), user.ExerciseView(window)),
		widget.NewTabItemWithIcon("Sync", theme.MailSendIcon(), user.SyncView(window)),
		widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), user.SettingsView(window, app)),
		// TODO: Add an about page with logo, name and version number.
	)

	// Set the tabs to be on top of the page.
	tabs.SetTabLocation(widget.TabLocationTop)

	// Adapt the window to a good size and make it resizable again.
	window.SetFixedSize(false)
	window.Resize(fyne.NewSize(800, 500))

	// Set the content to show and do so in a scroll container for the exercieses to show correctly.
	window.SetContent(tabs)
}
