package gui

import (
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/Jacalz/sparta/internal/file"
	"github.com/Jacalz/sparta/internal/file/parse"
	"github.com/Jacalz/sparta/internal/gui/widgets"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

// AddExerciseView shows the opoup for adding a new activity.
func (u *user) addExerciseView(w fyne.Window) fyne.CanvasObject {
	// Variables for the entry variables used in the form.
	dateEntry := widgets.NewEntryWithPlaceholder("YYYY-MM-DD")
	clockEntry := widgets.NewEntryWithPlaceholder("HH:MM")
	activityEntry := widgets.NewEntryWithPlaceholder("Name of exercise")
	distanceEntry := widgets.NewEntryWithPlaceholder("Kilometers")
	durationEntry := widgets.NewEntryWithPlaceholder("Minutes")
	setsEntry := widgets.NewEntryWithPlaceholder("Number of sets")
	repsEntry := widgets.NewEntryWithPlaceholder("Number of reps")
	commentEntry := widget.NewMultiLineEntry()
	commentEntry.SetPlaceHolder("Type your comment here...")

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

	// Compile regular expressions for checking numeric input with optional decimals.
	validFloat := regexp.MustCompile(`^$|(\d+\.)?\d+$`)

	// Compile regular expressions for checking numeric input without decimals.
	validUint := regexp.MustCompile(`^$|^\d*$`)

	// Compile regular expressions for checking date input.
	validDate := regexp.MustCompile(`^([12]\d{3})-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[0-1])$`)

	// Compile regular expressions for checking clock input.
	validClock := regexp.MustCompile(`^([0-1][0-9]|2[0-3]):[0-5][0-9]$`)

	// Create the form for displaying.
	form.OnSubmit = func() {
		// Bool variable for checking non numeric inputs (default to false).
		nonNumericInput := false

		// Check that input is numeric in given fields.
		switch {
		case !validFloat.MatchString(distanceEntry.Text):
			nonNumericInput = true
		case !validFloat.MatchString(durationEntry.Text):
			nonNumericInput = true
		case !validUint.MatchString(setsEntry.Text):
			nonNumericInput = true
		case !validUint.MatchString(repsEntry.Text):
			nonNumericInput = true
		}

		// Show and error if any fields does not match the correct input patterns.
		if nonNumericInput || activityEntry.Text == "" || !validClock.MatchString(clockEntry.Text) || !validDate.MatchString(dateEntry.Text) {
			dialog.ShowInformation("Non numeric input or invald formats in fields", "The date and the start time need correct data formating and the exercise can not be empty.\nDistance, time, sets and reps can all be empty, however they do need to contain numeric data if non empty.", w)
		} else {
			go func() {
				// Defer the entry fields to be cleaned out last.
				defer form.OnCancel()

				// Create variables for parsing time from date and clock.
				var year, month, day, hour, min int

				// Parse the date and time from strings.
				_, err := fmt.Sscanf(dateEntry.Text, "%v-%v-%v", &year, &month, &day)
				if err != nil {
					fmt.Println("Parsing date: ", err)
					return
				}

				_, err = fmt.Sscanf(clockEntry.Text, "%v:%v", &hour, &min)
				if err != nil {
					fmt.Println("Parsing time: ", err)
					return
				}

				// Create the time.Time value for the imputed data.
				timeOfExercise := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

				// Append new values to a new index.
				u.data.Exercise = append(u.data.Exercise, file.Exercise{Time: timeOfExercise, Clock: clockEntry.Text, Date: dateEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Duration: parse.Float(durationEntry.Text), Reps: parse.Uint(repsEntry.Text), Sets: parse.Uint(setsEntry.Text), Comment: commentEntry.Text})

				// Encrypt and write the data to the configuration file. Do it on another goroutine.
				go u.data.Write(&u.encryptionKey, u.username)

				// Check if the length before appending was zero. If it was, the file is empy and sees its first exercise added.
				if len(u.data.Exercise)-1 == 0 {
					u.firstExercise <- u.data.Format(len(u.data.Exercise) - 1)
				} else {
					// Check the length of the newly appended slice.
					length := len(u.data.Exercise)

					// Check if the newest added exercise was after the exercise before that. It means that we won't have to sort the slice.
					if u.data.Exercise[length-2].Time.Before(u.data.Exercise[length-1].Time) {
						// Send the formated string from the highest index of the Exercise slice.
						u.newExercise <- u.data.Format(len(u.data.Exercise) - 1)
					} else {
						// Sort all old and new data to make sure that new exercises come first.
						sort.Slice(u.data.Exercise, func(i, j int) bool {
							return u.data.Exercise[i].Time.Before(u.data.Exercise[j].Time)
						})

						// Indicate that the whole slice needs to be redisplayed.
						u.reorderExercises <- true
					}
				}
			}()
		}
	}

	// Append all the rows separately in to the form.
	form.Append("Date", dateEntry)
	form.Append("Start Time", clockEntry)
	form.Append("Activity", activityEntry)
	form.Append("Distance", distanceEntry)
	form.Append("Duration", durationEntry)
	form.Append("Sets", setsEntry)
	form.Append("Reps", repsEntry)
	form.Append("Comment", commentEntry)

	return form
}
