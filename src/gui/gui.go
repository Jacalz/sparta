package gui

import (
	"crypto/sha256"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
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
	userName.SetPlaceHolder("Användarnamn")

	// Initialize the password input box that we are to be using.
	userPassword := widget.NewPasswordEntry()
	userPassword.SetPlaceHolder("Lösenord")

	// Create the login button that will calculate the 32bit long sha256 hash.
	loginButton := widget.NewButton("Login", func() {
		PasswordKey = sha256.Sum256([]byte(userName.Text + userPassword.Text))
	})

	// Make a container that houses all of our widgets in a one wide grid.
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		userName,
		userPassword,
		loginButton,
	)

	// Set the conatiner as what is being displayed.
	window.SetContent(container)

	// Set a sane default for the window size.
	window.Resize(fyne.NewSize(200, 100))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
