package gui

import (
	"sort"
	"sparta/file"
	"sparta/file/parse"

	"fmt"
	"regexp"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

// ActivityView shows the opoup for adding a new activity.
func (u *user) ActivityView(window fyne.Window) fyne.CanvasObject {
	// Variables for the entry variables used in the form.
	dateEntry := NewEntryWithPlaceholder("YYYY-MM-DD")
	clockEntry := NewEntryWithPlaceholder("HH:MM")
	activityEntry := NewEntryWithPlaceholder("Name of activity")
	distanceEntry := NewEntryWithPlaceholder("Kilometers")
	durationEntry := NewEntryWithPlaceholder("Minutes")
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
			durationEntry.SetText("")
			setsEntry.SetText("")
			repsEntry.SetText("")
			commentEntry.SetText("")
		},
	}

	// Compile regular expressions for checking numeric input with optional decimals..
	numericFloat, err := regexp.Compile(`^$|(\d*\.)?\d+$`)
	if err != nil {
		fmt.Print(err)
	}

	// Compile regular expressions for checking numeric input without decimals..
	numericUint, err := regexp.Compile(`^$|^[0-9]*$`)
	if err != nil {
		fmt.Print(err)
	}

	// Compile regular expressions for checking date input.
	date, err := regexp.Compile(`^(\d{1,4})-(\d{1,2})-(\d{1,2})$`)
	if err != nil {
		fmt.Print(err)
	}

	// Compile regular expressions for checking clock input.
	clock, err := regexp.Compile(`^(\d{1,2}):(\d{1,2})$`)
	if err != nil {
		fmt.Print(err)
	}

	// Create the form for displaying.
	form.OnSubmit = func() {
		// Bool variable for checking non numeric inputs (default to false).
		nonNumericInput := false

		// Check that input is numeric in given fields.
		switch {
		case !numericFloat.Match([]byte(distanceEntry.Text)):
			nonNumericInput = true
		case !numericFloat.Match([]byte(durationEntry.Text)):
			nonNumericInput = true
		case !numericUint.Match([]byte(setsEntry.Text)):
			nonNumericInput = true
		case !numericUint.Match([]byte(repsEntry.Text)):
			nonNumericInput = true
		}

		// Show and error if any fields does not match the correct input patterns.
		if nonNumericInput || activityEntry.Text == "" || !clock.Match([]byte(clockEntry.Text)) || !date.Match([]byte(dateEntry.Text)) {
			dialog.ShowInformation("Non numeric input or invald formats in fields", "The date and the start time need correct data formating and the activity can not be empty.\nDistance, time, sets and reps can all be empty, however they do need to contain numeric data if non empty.", window)
		} else {
			go func() {
				// Defer the entry fields to be cleaned out last.
				defer form.OnCancel()

				// Create variables for parsing time from date and clock.
				var year, month, day, hour, min int

				// Parse the date and time from strings.
				fmt.Sscanf(dateEntry.Text, "%v-%v-%v", &year, &month, &day)
				fmt.Sscanf(clockEntry.Text, "%v:%v", &hour, &min)

				// Create the time.Time value for the imputed data.
				timeOfExercise := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

				// Append new values to a new index.
				u.Data.Exercise = append(u.Data.Exercise, file.Exercise{Time: timeOfExercise, Clock: clockEntry.Text, Date: dateEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Duration: parse.Float(durationEntry.Text), Reps: parse.Uint(repsEntry.Text), Sets: parse.Uint(setsEntry.Text), Comment: commentEntry.Text})

				// Encrypt and write the data to the configuration file. Do it on another goroutine.
				go u.Data.Write(&u.EncryptionKey)

				// If the data was empty before, we send it as a first exercise and say that it isn't empty anymore.
				if file.Empty() {
					u.FirstExercise <- u.Data.Format(len(u.Data.Exercise) - 1)
					file.SetNonEmpty()
				} else {
					// Check the length of the newly appended slice.
					length := len(u.Data.Exercise)

					// Check if the newest added exercise was after the exercise before that. It means that we won't have to sort the slice.
					if u.Data.Exercise[length-2].Time.Before(u.Data.Exercise[length-1].Time) {
						// Send the formated string from the highest index of the Exercise slice.
						u.NewExercise <- u.Data.Format(len(u.Data.Exercise) - 1)
					} else {
						// Sort all old and new data to make sure that new exercises come first.
						sort.Slice(u.Data.Exercise, func(i, j int) bool {
							return u.Data.Exercise[i].Time.Before(u.Data.Exercise[j].Time)
						})

						// Indicate that the whole slice needs to be redisplayed.
						u.ReorderExercises <- true
					}
				}
			}()
		}
	}

	// Append all the rows separately in to the form.
	form.Append("Date", dateEntry)
	form.Append("Start time", clockEntry)
	form.Append("Activity", activityEntry)
	form.Append("Distance", distanceEntry)
	form.Append("Duration", durationEntry)
	form.Append("Sets", setsEntry)
	form.Append("Reps", repsEntry)
	form.Append("Comment", commentEntry)

	return form
}
