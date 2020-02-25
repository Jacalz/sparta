package gui

import (
	"sparta/assets"
	"sparta/file"

	"fyne.io/fyne/app"
)

// User holds the data about the user that is currently logged in.
type user struct {
	Username         string
	Password         string
	EncryptionKey    [32]byte
	Data             file.Data
	NewExercise      chan string
	FirstExercise    chan string
	EmptyExercises   chan bool
	ReorderExercises chan bool
	Errors           chan error
}

func newUser() *user {
	return &user{NewExercise: make(chan string),
		FirstExercise:    make(chan string),
		EmptyExercises:   make(chan bool),
		ReorderExercises: make(chan bool),
		Errors:           make(chan error),
	}
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
	user := newUser()

	// Show the login page and all content after that.
	ShowLoginPage(app, window, user)
}
