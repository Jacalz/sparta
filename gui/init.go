package gui

import (
	"sparta/assets"
	"sparta/file"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
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

	// Channel for sending the sync code.
	SyncCode chan string
}

func newUser() *user {
	return &user{NewExercise: make(chan string),
		FirstExercise:    make(chan string),
		EmptyExercises:   make(chan bool),
		ReorderExercises: make(chan bool),
		SyncCode:         make(chan string),
	}
}

// Init will start up our graphical user interface.
func Init() {
	// Initialize our new fyne interface application.
	a := app.NewWithID("com.github.jacalz.sparta")

	// Set the application icon for our program.
	a.SetIcon(assets.AppIcon256)

	// Create the window for our user interface.
	w := a.NewWindow("Sparta")

	// Create the user struct type for later use.
	u := newUser()

	// Check that we are using the right theme.
	switch a.Preferences().StringWithFallback("Theme", "Light") {
	case "Dark":
		a.Settings().SetTheme(theme.DarkTheme())
	case "Light":
		a.Settings().SetTheme(theme.LightTheme())
	}

	// Create the tab handler for the user interface and set up the login view.
	t := &widget.TabContainer{}
	t.Append(u.LoginTabContainer(a, w, t))

	// Set the window to a good size, add the content and lastly run the application.
	w.SetContent(t)
	w.Resize(fyne.NewSize(800, 550))
	w.ShowAndRun()
}
