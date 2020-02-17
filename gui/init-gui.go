package gui

import (
	"sparta/assets"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	app.SetIcon(assets.AppIcon)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Check that we are using the right theme.
	switch app.Preferences().StringWithFallback("Theme", "Dark") {
	case "Dark":
		app.Settings().SetTheme(theme.DarkTheme())
	case "Light":
		app.Settings().SetTheme(theme.LightTheme())
	}

	// Show the login page and all content after that.
	ShowLoginPage(app, window)
}
