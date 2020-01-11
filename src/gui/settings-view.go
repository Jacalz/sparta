package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// SettingsView contains the gui information for the settings screen.
func SettingsView() fyne.CanvasObject {
	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Remove all activities", theme.DeleteIcon(), func() {})

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), deleteButton)
}
