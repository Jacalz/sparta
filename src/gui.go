package main

import (
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

// InitGUI will start up our graphical user interface.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	//app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Initialize the login form that we are to be using.
	userName := widget.NewEntry()
	userName.SetPlaceHolder("Username")

	// Initialize the password input box that we are to be using.
	userPassword := widget.NewPasswordEntry()
	userPassword.SetPlaceHolder("Password")

	// Create the login button that will calculate the 32bit long sha256 hash.
	loginButton := widget.NewButton("Login", func() {
		// Check that a username and password was provided. Without it we show an informative dialog and return.
		if userName.Text == "" || userPassword.Text == "" {
			dialog.ShowInformation("Missing username/password", "Please provide both username and password.", window)
			return
		}

		// Adapt the window to a good size.
		window.Resize(fyne.NewSize(600, 400))

		// Calculate the sha256 hash of the username and password.
		PasswordKey = encrypt.EncryptionKey(userName.Text, userPassword.Text)

		// Create a channel for sending activity data through. Let's us avoid reading the file every time we add a new activity.
		newAddedExercise := make(chan string)

		// Check for the file where we store the data.
		XMLData, empty := file.Check()

		// The button for creating a new exercise.
		newExercise := widget.NewButtonWithIcon("Add new exercise", theme.ContentAddIcon(), func() {
			// Form information for the first row.
			dateEntry := widget.NewEntry()
			dateEntry.SetPlaceHolder("YYYY-MM-DD")

			// Second row of form information.
			clockEntry := widget.NewEntry()
			clockEntry.SetPlaceHolder("HH:MM")

			// Third row of the form.
			activityEntry := widget.NewEntry()

			// Forth row of form data.
			distanceEntry := widget.NewEntry()
			distanceEntry.SetPlaceHolder("kilometers")

			// Fifth row in the form.
			timeEntry := widget.NewEntry()
			timeEntry.SetPlaceHolder("minutes")

			// Create the form for displaying.
			form := &widget.Form{
				OnSubmit: func() {
					// Append new values to a new index.
					XMLData.Exercise = append(XMLData.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: file.ParseFloat(distanceEntry.Text), Time: file.ParseFloat(timeEntry.Text)})

					// Write the data to the file.
					XMLData.Write()

					// Send the formated string from the highest index of the Exercise slice.
					newAddedExercise <- XMLData.Format(len(XMLData.Exercise) - 1)
				},
			}

			// Append all the rows separately in to the form.
			form.Append("Date", dateEntry)
			form.Append("Start time", clockEntry)
			form.Append("Activity", activityEntry)
			form.Append("Distance", distanceEntry)
			form.Append("Time", timeEntry)

			// Show the popup dialog.
			dialog.ShowCustom("Add activity", "Done", form, window)
		})

		// Create a label for displaing some info for the user. Default to showing nothing.
		label := widget.NewLabel("")

		// Start up procedure if the data field is empty.
		if !empty {
			// Add a new label for each exercise and so in a new goroutine to not block the current one.
			go func() {
				// First we loop through the imported file and add the formated info before the previous info (new information comes out on top).
				for i := range XMLData.Exercise {
					label.SetText(XMLData.Format(i) + label.Text)
				}

				// We then block the channel while waiting for an update on the channel.
				for {
					label.SetText(<-newAddedExercise + label.Text)
				}
			}()
		} else {
			// Inform about no exercies being avaliable.
			label.SetText("No exercieses have been created yet.")
		}

		// Set the content to show and do so in a scroll container for the exercieses to show correctly.
		window.SetContent(widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewVBox(newExercise), label)))
	})

	// Set the content to be displayed. It is the userName, userPassword fields and the login button inside a layout.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1), userName, userPassword, loginButton))

	// Set a sane default for the window size on login.
	window.Resize(fyne.NewSize(400, 150))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
