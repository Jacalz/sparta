package main

import (
	"crypto/sha256"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"sparta/file"
)

// PasswordKey contains the key taken from the username and password.
var PasswordKey [32]byte

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

		// Calculate the sha256 hash of the username and password.
		PasswordKey = sha256.Sum256([]byte(userName.Text + userPassword.Text))

		// Check for the file where we store the data.
		XMLData, empty := file.Check()

		// The button for creating a new exercise.
		newExercise := widget.NewButtonWithIcon("Add new exercise", theme.ContentAddIcon(), func() {
		})

		// Create a label for displaing some info for the user. Default to showing nothing.
		label := widget.NewLabel("")

		// Append the button and label initially to a vertical box.
		vbox := widget.NewVBox(newExercise, label)

		// Start up procedure if the data fiel is empty.
		if !empty {
			// Set a bigger window when we have content to display.
			window.Resize(fyne.NewSize(400, 250))

			// Add a new label for each exercise and so in a new goroutine to not block the current one.
			go func() {
				for i := range XMLData.Exercise {
					vbox.Append(widget.NewLabel(fmt.Sprintf("At %s on %s, you trained %s. The distance was %v and the exercise lasted for %v minutes, resulting in an average speed of %v.",
						XMLData.Exercise[i].Clock, XMLData.Exercise[i].Date, XMLData.Exercise[i].Activity, XMLData.Exercise[i].Distance,
						XMLData.Exercise[i].Length, XMLData.Exercise[i].Distance/XMLData.Exercise[i].Length)))
				}
			}()
		} else {
			// Adapt the window size for empty content.
			window.Resize(fyne.NewSize(400, 150))

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

	// Set a sane default for the window size.
	window.Resize(fyne.NewSize(400, 150))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
