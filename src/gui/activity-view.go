package gui

import (
	"sparta/src/file"
	"sparta/src/file/parse"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// ActivityView shows the opoup for adding a new activity.
func ActivityView(window fyne.Window, XMLData *file.Data, newAddedExercise chan string) fyne.CanvasObject {
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
				XMLData.Exercise = append(XMLData.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Time: parse.Float(timeEntry.Text), Reps: parse.Int(repsEntry.Text), Sets: parse.Int(setsEntry.Text), Comment: commentEntry.Text})

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
				go XMLData.Write(&PasswordKey)

				// Send the formated string from the highest index of the Exercise slice.
				newAddedExercise <- XMLData.Format(len(XMLData.Exercise) - 1)
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
