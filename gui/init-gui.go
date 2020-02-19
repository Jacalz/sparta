package gui

import (
	"sparta/assets"

	"fyne.io/fyne/app"
)

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	app.SetIcon(assets.AppIcon)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Show the login page and all content after that.
	ShowLoginPage(app, window)
}
