package gui

import (
	"sparta/file"
	"sparta/file/parse"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// TODO: Handle invalid inputs and empty fields.

// ActivityView shows the opoup for adding a new activity.
func ActivityView(window fyne.Window, exercises *file.Data, dataLabel *widget.Label, newAddedExercise chan string) fyne.CanvasObject {
	// Variables for the entry variables used in the form.
	dateEntry := NewEntryWithPlaceholder("YYYY-MM-DD")
	clockEntry := NewEntryWithPlaceholder("HH:MM")
	activityEntry := NewEntryWithPlaceholder("Name of activity")
	distanceEntry := NewEntryWithPlaceholder("Kilometers")
	timeEntry := NewEntryWithPlaceholder("Minutes")
	setsEntry := NewEntryWithPlaceholder("Number of sets")
	repsEntry := NewEntryWithPlaceholder("Number of reps")
	commentEntry := widget.NewMultiLineEntry()
	commentEntry.SetPlaceHolder("Type your comment here")

	// Create the form for displaying.
	form := &widget.Form{
		OnSubmit: func() {
			go func() {
				// Append new values to a new index.
				exercises.Exercise = append(exercises.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Time: parse.Float(timeEntry.Text), Reps: parse.Int(repsEntry.Text), Sets: parse.Int(setsEntry.Text), Comment: commentEntry.Text})

				// Make sure to clean out the text for all the entry widgets.
				dateEntry.SetText("")
				clockEntry.SetText("")
				activityEntry.SetText("")
				distanceEntry.SetText("")
				timeEntry.SetText("")
				setsEntry.SetText("")
				repsEntry.SetText("")
				commentEntry.SetText("")

				// Encrypt and write the data to the configuration file. Do so on another goroutine.
				go exercises.Write(&PasswordKey)

				// Workaround bug that happens after creating a new activity after removing the file.
				if file.Empty() {
					dataLabel.Text = ""
				}

				// Send the formated string from the highest index of the Exercise slice.
				newAddedExercise <- exercises.Format(len(exercises.Exercise) - 1)

				// Now set the status to not be empty.
				file.SetNonEmpty()
			}()
		},
		OnCancel: func() {
			// Make sure to clean out the text for all the entry widgets.
			dateEntry.SetText("")
			clockEntry.SetText("")
			activityEntry.SetText("")
			distanceEntry.SetText("")
			timeEntry.SetText("")
			setsEntry.SetText("")
			repsEntry.SetText("")
			commentEntry.SetText("")
		},
	}

	// Append all the rows separately in to the form.
	form.Append("Date", dateEntry)
	form.Append("Start time", clockEntry)
	form.Append("Activity", activityEntry)
	form.Append("Distance", distanceEntry)
	form.Append("Time", timeEntry)
	form.Append("Sets", setsEntry)
	form.Append("Reps", repsEntry)
	form.Append("Comment", commentEntry)

	return form
}
