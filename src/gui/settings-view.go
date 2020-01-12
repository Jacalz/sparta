package gui

import (
	"sparta/src/file"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// TODO: Multi user support by labling the data file exercises-user.xml

// SettingsView contains the gui information for the settings screen.
func SettingsView(window fyne.Window, app fyne.App, XMLData *file.Data, dataLabel *widget.Label) fyne.CanvasObject {

	// TODO: Add setting for changing language.

	// Make it possible for the user to switch themes.
	themeSwitcher := widget.NewSelect([]string{"Dark", "Light"}, func(selected string) {
		switch selected {
		case "Dark":
			app.Settings().SetTheme(theme.DarkTheme())
		case "Light":
			app.Settings().SetTheme(theme.LightTheme())
		}

		// Set the theme to the selected one.
		config.Theme = selected

		// Write the theme update to the file and do it concurrently to not stop the thread.
		go config.Write()
	})

	// Default theme is dark and thus we set the placeholder to that and then refresh it (without a refresh, it doesn't show until hovering on to widget).
	themeSwitcher.PlaceHolder = config.Theme
	themeSwitcher.Refresh()

	// Add the theme switcher next to a label.
	themeSettings := fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("Change theme"), themeSwitcher)

	// TODO: Change password for the encrypted file.

	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Remove all activities", theme.DeleteIcon(), func() {

		// Ask the user to confirm what we are about to do.
		dialog.ShowConfirm("Are you sure that you want to continue?", "Deleting your data will remove all of your exercises and activities.", func(remove bool) {
			if remove {
				// Run the delete function.
				XMLData.Delete()

				// Clear all the data inside the data label.
				dataLabel.SetText("No exercieses have been created yet.")
			}
		}, window)
	})

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), themeSettings, layout.NewSpacer(), deleteButton)
}
