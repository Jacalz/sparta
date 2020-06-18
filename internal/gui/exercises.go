package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// InitialDisplay loops though and produces the text for all existing data.
func (u *user) initialDisplay() (text string) {
	// We loop through the imported file and add the formated info before the previous info (new information comes out on top).
	for i := range u.data.Exercise {
		text = u.data.Format(i) + text
	}

	return text
}

// ExerciseDisplayer runs in the background and updates the label.
func (u *user) exerciseDisplayer(label *widget.Label) {
	// Handle an empty data file.
	if len(u.data.Exercise) == 0 {
		// Start by informing  the user that no data is available.
		label.SetText("No exercises have been created yet.")
	} else {
		// Refresh the widget to show the updated text.
		label.SetText(u.initialDisplay())
	}

	// We then block the channel while waiting for updates on the channel.
	for {
		select {
		case exercise := <-u.newExercise:
			label.SetText(exercise + label.Text)
		case exercise := <-u.firstExercise:
			label.SetText(exercise)
		case <-u.reorderExercises:
			label.SetText(u.initialDisplay())
		case <-u.emptyExercises:
			label.SetText("No exercises have been created yet.")
		}

	}
}

// ExercisesView shows the main view after we are logged in.
func (u *user) exercisesView(w fyne.Window, a fyne.App) fyne.CanvasObject {
	// Create a label for displaing some info for the user. Default to showing nothing.
	dataLabel := widget.NewLabel("")
	dataLabel.Wrapping = fyne.TextWrapWord

	// Start up the function to handle adding exercises in the background.
	go u.exerciseDisplayer(dataLabel)

	return widget.NewScrollContainer(dataLabel)
}
