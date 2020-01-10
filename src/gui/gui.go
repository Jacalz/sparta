package gui

import (
	"sparta/src/bundled"
	"sparta/src/file"
	"sparta/src/file/encrypt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// PasswordKey contains the key taken from the username and password.
var PasswordKey [32]byte

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	app.SetIcon(bundled.MainIcon)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Initialize the login form that we are to be using.
	userName := NewExtendedEntry(false)
	userName.SetPlaceHolder("Username")

	// Initialize the password input box that we are to be using.
	userPassword := NewExtendedEntry(true)
	userPassword.SetPlaceHolder("Password")

	// Create the login button that will calculate the 32bit long sha256 hash.
	loginButton := widget.NewButton("Login", func() {
		// Check that a username and password was provided. Without it we show an informative dialog and return.
		if userName.Text == "" || userPassword.Text == "" {
			dialog.ShowInformation("Missing username/password", "Please provide both username and password.", window)
			return
		} else if userName.Text == userPassword.Text {
			dialog.ShowInformation("Identical username and password", "Please do not use identical username and password.", window)
			return
		} else if len(userPassword.Text) < 8 {
			dialog.ShowInformation("Too short password", "The password should be eight characters or longer.", window)
			return
		}

		// Adapt the window to a good size and make it resizable again.
		window.SetFixedSize(false)
		window.Resize(fyne.NewSize(800, 500))

		// Calculate the sha256 hash of the username and password.
		PasswordKey = encrypt.EncryptionKey(userName.Text, userPassword.Text)

		// Create a channel for sending activity data through. Let's us avoid reading the file every time we add a new activity.
		newAddedExercise := make(chan string)

		// Check for the file where we store the data.
		XMLData, empty := file.Check(&PasswordKey)

		// The button for creating a new exercise.
		newExercise := widget.NewButtonWithIcon("Add new exercise", theme.ContentAddIcon(), func() {

			// Variables for the entry variables used in the form.
			dateEntry := NewEntryWithPlaceholder("YYYY-MM-DD")
			clockEntry := NewEntryWithPlaceholder("HH:MM")
			activityEntry := NewEntryWithPlaceholder("Name of activity")
			distanceEntry := NewEntryWithPlaceholder("Kilometers")
			timeEntry := NewEntryWithPlaceholder("Minutes")
			repsEntry := NewEntryWithPlaceholder("Number of reps")
			setsEntry := NewEntryWithPlaceholder("Number of sets")
			commentEntry := widget.NewMultiLineEntry()

			// Create the form for displaying.
			form := &widget.Form{
				OnSubmit: func() {
					go func() {
						// Append new values to a new index.
						XMLData.Exercise = append(XMLData.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: file.ParseFloat(distanceEntry.Text), Time: file.ParseFloat(timeEntry.Text), Reps: file.ParseInt(repsEntry.Text), Sets: file.ParseInt(setsEntry.Text), Comment: commentEntry.Text})

						// Encrypt and write the data to the configuration file. Do so on another goroutine.
						go XMLData.Write(&PasswordKey)

						// Send the formated string from the highest index of the Exercise slice.
						newAddedExercise <- XMLData.Format(len(XMLData.Exercise) - 1)
					}()
				},
			}

			// Append all the rows separately in to the form.
			form.Append("Date", dateEntry)
			form.Append("Start time", clockEntry)
			form.Append("Activity", activityEntry)
			form.Append("Distance", distanceEntry)
			form.Append("Time", timeEntry)
			form.Append("Reps", repsEntry)
			form.Append("Sets", setsEntry)
			form.Append("Comment", commentEntry)

			// Show the popup dialog.
			dialog.ShowCustom("Add activity", "Done", form, window)
		})

		// Create a label for displaing some info for the user. Default to showing nothing.
		label := widget.NewLabel("")

		go func() {

			// Handle an empty data file.
			if empty {
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

		// Set the content to show and do so in a scroll container for the exercieses to show correctly.
		window.SetContent(widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), newExercise, label)))
	})

	// Add the Action component to make actions work inside the struct. This is used to press the loginButton on pressing enter/return ton the keyboard.
	userName.Action, userPassword.Action = &Action{*loginButton}, &Action{*loginButton}

	// Set the content to be displayed. It is the userName, userPassword fields and the login button inside a layout.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), userName, userPassword, loginButton))

	// Set a sane default for the window size on login and make sure that it isn't resizable.
	window.Resize(fyne.NewSize(400, 100))
	window.SetFixedSize(true)

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
