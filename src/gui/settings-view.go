package gui

import (
	"fmt"
	"sparta/src/file"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// SettingsView contains the gui information for the settings screen.
func SettingsView(window fyne.Window, XMLData *file.Data, dataLabel *widget.Label) fyne.CanvasObject {

	// TODO: Add setting for changing language.

	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Remove all activities", theme.DeleteIcon(), func() {

		// Ask the user to confirm what we are about to do.
		dialog.ShowConfirm("Are you sure that you want to continue?", "Deleting your data will remove all of your exercises and activities.", func(remove bool) {
			if remove {
				// Run the delete function.
				XMLData.Delete()

				fmt.Print("We are removing :(")

				// Clear all the data inside the data label.
				dataLabel.SetText("No exercieses have been created yet.")
			}
		}, window)
	})

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), deleteButton)
}
