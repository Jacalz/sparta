package gui

import (
	"regexp"
	"sort"
	"time"

	"github.com/Jacalz/sparta/internal/file"
	"github.com/Jacalz/sparta/internal/file/parse"
	"github.com/Jacalz/sparta/internal/gui/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	floatValidation    = regexp.MustCompile(`^$|(\d+\.)?\d+$`)
	uintValidation     = regexp.MustCompile(`^$|^\d*$`)
	notEmptyValidation = regexp.MustCompile(`(.|\s)*\S(.|\s)*`)
)

// AddExerciseView shows the opoup for adding a new activity.
func (u *user) addExerciseView(w fyne.Window) fyne.CanvasObject {
	// Variables for the entry variables used in the form.
	var (
		dateEntry     = widgets.NewFormEntry("YYYY-MM-DD", "Incorrect date formatting.", nil, false)
		clockEntry    = widgets.NewFormEntry("HH:MM", "Incorrect clock formatting.", nil, false)
		activityEntry = widgets.NewFormEntry("Name of exercise", "The field must be filled in.", notEmptyValidation, false)
		distanceEntry = widgets.NewFormEntry("Kilometers", "Not a valid number.", floatValidation, false)
		durationEntry = widgets.NewFormEntry("Minutes", "Not a valid number.", floatValidation, false)
		setsEntry     = widgets.NewFormEntry("Number of sets", "Not a valid integer number.", uintValidation, false)
		repsEntry     = widgets.NewFormEntry("Number of reps", "Not a valid integer number.", uintValidation, false)
		commentEntry  = widgets.NewFormEntry("A comment about the exercise...", "", nil, true)
	)

	form := widget.NewForm(
		&widget.FormItem{Text: "Date", HintText: "The date of the exercise.", Widget: dateEntry},
		&widget.FormItem{Text: "Start Time", HintText: "The start time.", Widget: clockEntry},
		&widget.FormItem{Text: "Activity", HintText: "The activity name.", Widget: activityEntry},
		&widget.FormItem{Text: "Distance", HintText: "The travelled distance, can be empty.", Widget: distanceEntry},
		&widget.FormItem{Text: "Duration", HintText: "The elapsed duration, can be empty.", Widget: durationEntry},
		&widget.FormItem{Text: "Sets", HintText: "The amount of sets performed, can be empty.", Widget: setsEntry},
		&widget.FormItem{Text: "Reps", HintText: "The amount of reps performed per set, can be empty.", Widget: repsEntry},
		&widget.FormItem{Text: "Comment", Widget: commentEntry},
	)

	form.OnCancel = func() {
		for _, item := range form.Items {
			item.Widget.(*widget.Entry).SetText("")
		}
	}

	form.OnSubmit = func() {
		go func() {
			// Defer the entry fields to be cleaned out last.
			defer form.OnCancel()

			timeOfExercise, err := time.Parse("2006-01-02|15:04", dateEntry.Text+"|"+clockEntry.Text)
			if err != nil {
				fyne.LogError("Error on parsing exercise time", err)
				return
			}

			// Append new values to a new index.
			u.data.Exercise = append(u.data.Exercise, file.Exercise{Time: timeOfExercise, Clock: clockEntry.Text, Date: dateEntry.Text, Activity: activityEntry.Text, Distance: parse.Float(distanceEntry.Text), Duration: parse.Float(durationEntry.Text), Reps: parse.Uint(repsEntry.Text), Sets: parse.Uint(setsEntry.Text), Comment: commentEntry.Text})

			// Encrypt and write the data to the configuration file. Do it on another goroutine.
			go u.data.Write(&u.encryptionKey, u.username)

			dlen := len(u.data.Exercise)

			if dlen-1 == 0 { // First exercise to be added
				u.firstExercise <- u.data.Format(dlen - 1)
				return
			}

			// Check if the newest exercise is later than the exercise before. It means that we won't have to sort the slice.
			if u.data.Exercise[dlen-2].Time.Before(u.data.Exercise[dlen-1].Time) {
				u.newExercise <- u.data.Format(dlen - 1)
			} else {
				// Sort all old and new data to make sure that new exercises come first.
				sort.Slice(u.data.Exercise, func(i, j int) bool {
					return u.data.Exercise[i].Time.Before(u.data.Exercise[j].Time)
				})

				// Indicate that the whole slice needs to be redisplayed.
				u.reorderExercises <- true

			}
		}()

	}

	return container.NewScroll(container.NewPadded(form))
}
