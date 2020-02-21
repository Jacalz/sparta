package gui

import (
	"sparta/assets"
	"sparta/file"

	"fyne.io/fyne/app"
)

// User holds the data about the user that is currently logged in.
type user struct {
	Username      string
	Password      string
	EncryptionKey [32]byte
	NewExercise   chan string
	ExerciseData  file.Data
}

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	app.SetIcon(assets.AppIcon)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Create the user struct type for later use.
	user := &user{}

	// Show the login page and all content after that.
	ShowLoginPage(app, window, user)
}
