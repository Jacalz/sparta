package gui

import (
	"sparta/src/file"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// ShowMainDataView shows the main view after we are logged in.
func ShowMainDataView(window fyne.Window, XMLData *file.Data, newAddedExercise chan string) {
	// Create a label for displaing some info for the user. Default to showing nothing.
	label := widget.NewLabel("")

	go func() {

		// Handle an empty data file.
		if file.Empty() {
			// Start by inorming  the user that no data is avaliable.
			label.SetText("No exercieses have been created yet.")

			// Then wait for more data to come running down the pipe.
			label.SetText(<-newAddedExercise)
		} else {
			// We loop through the imported file and add the formated info before the previous info (new information comes out on top).
			for i := range XMLData.Exercise {
				label.SetText(XMLData.Format(i) + label.Text)
			}
		}

		// We then block the channel while waiting for an update on the channel.
		for {
			label.SetText(<-newAddedExercise + label.Text)
		}
	}()

	// Tab data for the main window.
	dataPage := widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label))

	// Create tabs with data.
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Activities", theme.HomeIcon(), dataPage),
		widget.NewTabItemWithIcon("Add activity", theme.ContentAddIcon(), ActivityView(XMLData, newAddedExercise)),
		widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), SettingsView()),
	)

	// Set the tabs to be on top of the page.
	tabs.SetTabLocation(widget.TabLocationTop)

	// Set the content to show and do so in a scroll container for the exercieses to show correctly.
	window.SetContent(tabs)
}
