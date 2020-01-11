package gui

import (
	"sparta/src/file"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// SettingsView contains the gui information for the settings screen.
func SettingsView(XMLData *file.Data, dataLabel *widget.Label) fyne.CanvasObject {

	// TODO: Add setting for changing language.

	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Remove all activities", theme.DeleteIcon(), func() {

		// TODO: Add a dialog to confirm what we are doing.

		// Run the delete function.
		XMLData.Delete()

		// Clear all the data inside the data label.
		dataLabel.SetText("No exercieses have been created yet.")
	})

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), deleteButton)
}
