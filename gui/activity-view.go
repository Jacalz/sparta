package gui

import (
	"fmt"
	"path/filepath"
	"sparta/file"
	"sparta/file/parse"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

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

	// Create the initial form with a cancel button so it can be used last on submit.
	form := &widget.Form{
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

	// Create the form for displaying.
	form.OnSubmit = func() {
		// Bool variable for checking non numeric inputs (default to false).
		nonNumericInput := false

		// Check that input is numeric in given fields.
		switch {
		case distanceEntry.Text != "" && !parse.IsNumeric(distanceEntry.Text):
			nonNumericInput = true
			fallthrough
		case timeEntry.Text != "" && !parse.IsNumeric(timeEntry.Text):
			nonNumericInput = true
			fallthrough
		case setsEntry.Text != "" && !parse.IsNumeric(setsEntry.Text):
			nonNumericInput = true
			fallthrough
		case repsEntry.Text != "" && !parse.IsNumeric(repsEntry.Text):
			nonNumericInput = true
		}

		// Check that clock input matches the given format.
		correctClockFormat, err := filepath.Match("*:*", clockEntry.Text)
		if err != nil {
			fmt.Print(err)
		}

		// Check that date input matches the given format.
		correctDateFormat, err := filepath.Match("*-*-*", dateEntry.Text)
		if err != nil {
			fmt.Print(err)
		}

		// Show and error if any fields does not match the correct input patterns.
		if nonNumericInput || !correctClockFormat || !correctDateFormat {
			dialog.ShowInformation("Non numeric input or invald formats in fields", "Please make sure that inputed date and start time use the correct data formating.\nPlease also make sure that distance, time, sets and reps all contain numeric data if non empty.", window)
		} else {
			go func() {
				// Defer the entry fields to be cleaned out last.
				defer form.OnCancel()

				// Append new values to a new index.
				exercises.Exercise = append(exercises.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Time: parse.Float(timeEntry.Text), Reps: parse.Int(repsEntry.Text), Sets: parse.Int(setsEntry.Text), Comment: commentEntry.Text})

				// Encrypt and write the data to the configuration file. Do it on another goroutine.
				go exercises.Write(&PasswordKey)

				// Workaround bug that happens after creating a new activity after removing the file. Set the file to be non empty also.
				if file.Empty() {
					dataLabel.Text = ""
					file.SetNonEmpty()
				}

				// Send the formated string from the highest index of the Exercise slice.
				newAddedExercise <- exercises.Format(len(exercises.Exercise) - 1)
			}()
		}
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
