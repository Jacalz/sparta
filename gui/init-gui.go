package gui

import (
	"sparta/assets"
	"sparta/file"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

// User holds the data about the user that is currently logged in.
type user struct {
	Username      string
	Password      string
	EncryptionKey [32]byte
	Data          file.Data

	// Data channels for exercises.
	NewExercise      chan string
	FirstExercise    chan string
	EmptyExercises   chan bool
	ReorderExercises chan bool

	// Channels for errors and signaling.
	Errors       chan error
	FinishedSync chan bool
}

func newUser() *user {
	return &user{NewExercise: make(chan string),
		FirstExercise:    make(chan string),
		EmptyExercises:   make(chan bool),
		ReorderExercises: make(chan bool),
		Errors:           make(chan error),
		FinishedSync:     make(chan bool),
	}
}

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	app.SetIcon(assets.AppIcon256)

	// Create the window for our user interface.
	window := app.NewWindow("Sparta")

	// Create the user struct type for later use.
	user := newUser()

	// Check that we are using the right theme.
	switch app.Preferences().StringWithFallback("Theme", "Light") {
	case "Dark":
		app.Settings().SetTheme(theme.DarkTheme())
	case "Light":
		app.Settings().SetTheme(theme.LightTheme())
	}

	// Show the login page and all content after that.
	ShowLoginPage(app, window, user)
}
