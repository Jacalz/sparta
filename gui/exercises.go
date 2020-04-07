package gui

import (
	"fyne.io/fyne"
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

// ExerciseDisplayer runs in the background and updates the label.
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

// ExercisesView shows the main view after we are logged in.
func (u *user) ExercisesView(w fyne.Window, a fyne.App) fyne.CanvasObject {
	// Create a label for displaing some info for the user. Default to showing nothing.
	dataLabel := widget.NewLabel("")

	// Start up the function to handle adding exercises in the background.
	go u.ExerciseDisplayer(dataLabel)

	return widget.NewScrollContainer(dataLabel)
}
