package main

import (
	"crypto/sha256"
	"fmt"
	"sparta/file"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// PasswordKey contains the key taken from the username and password.
var PasswordKey [32]byte

// ParseFloat is a wrapper around strconv.ParseFloat to avoid the error when used inline.
func ParseFloat(input string) float64 {
	output, err := strconv.ParseFloat(input, 32)
	if err != nil {
		fmt.Print(err)
	}

	return output
}

// Init will start up our graphical user interface.
func Init(appName string) {
	// Initialize our new fyne interface application.
	app := app.New()

	// Set the application icon for our program.
	//app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow(appName)

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
			message := dialog.NewInformation("Missing username/password", "Please type both username and password.", window)
			message.Show()
			return
		}

		// Adapt the window to a good size.
		window.Resize(fyne.NewSize(600, 400))

		// Calculate the sha256 hash of the username and password.
		PasswordKey = sha256.Sum256([]byte(userName.Text + userPassword.Text))

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
			activityEntry.SetPlaceHolder("Anything")

			// Forth row of form data.
			distanceEntry := widget.NewEntry()
			distanceEntry.SetPlaceHolder("kilometers")

			// Fifth row in the form.
			timeEntry := widget.NewEntry()
			timeEntry.SetPlaceHolder("minutes")

			// Create the form for displaying.
			form := &widget.Form{
				OnSubmit: func() {
					XMLData.Exercise = append(XMLData.Exercise, file.Exercise{Date: dateEntry.Text, Clock: clockEntry.Text, Activity: activityEntry.Text, Distance: ParseFloat(distanceEntry.Text), Time: ParseFloat(timeEntry.Text)})
					file.Write(XMLData)
				},
			}

			// Append all the rows separately in to the form.
			form.Append("Date", dateEntry)
			form.Append("Clock", clockEntry)
			form.Append("Activity", activityEntry)
			form.Append("Distance", distanceEntry)
			form.Append("Time", timeEntry)

			// Show the popup dialog.
			dialog.ShowCustom("Add activity", "Done", form, window)
		})

		// Create a label for displaing some info for the user. Default to showing nothing.
		label := widget.NewLabel("")

		// Append the button and label initially to a vertical box.
		vbox := widget.NewVBox(newExercise, label)

		// Start up procedure if the data fiel is empty.
		if !empty {
			// Add a new label for each exercise and so in a new goroutine to not block the current one.
			go func() {
				for i := range XMLData.Exercise { // Note to self: make this range reversed so new entries come on top.
					vbox.Append(widget.NewLabel(fmt.Sprintf("At %s on %s, you trained %s. The distance was %v kilometers and the exercise lasted for %v minutes, resulting in an average speed of %.3f km/min.",
						XMLData.Exercise[i].Clock, XMLData.Exercise[i].Date, XMLData.Exercise[i].Activity, XMLData.Exercise[i].Distance,
						XMLData.Exercise[i].Time, XMLData.Exercise[i].Distance/XMLData.Exercise[i].Time)))
				}
			}()
		} else {
			// Inform about no exercies being avaliable.
			label.SetText("No exercieses have been created yet.")
		}

		// Set the content to show and do so in a scroll container for the exercieses to show correctly.
		window.SetContent(widget.NewScrollContainer(vbox))
	})

	// Make a container that houses all of our widgets in a one wide grid.
	loginScreenContainer := fyne.NewContainerWithLayout(layout.NewGridLayout(1), userName, userPassword, loginButton)

	// Set the conatiner as what is being displayed.
	window.SetContent(loginScreenContainer)

	// Set a sane default for the window size on login.
	window.Resize(fyne.NewSize(400, 150))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
