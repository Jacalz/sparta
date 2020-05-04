package gui

import (
	"github.com/Jacalz/sparta/internal/assets"
	"github.com/Jacalz/sparta/internal/crypto"
	"github.com/Jacalz/sparta/internal/file"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// User holds the data about the user that is currently logged in.
type user struct {
	// User sensitive data.
	username      string
	password      string
	passwordHash  string
	encryptionKey []byte
	data          file.Data

	// Data channels for exercises.
	newExercise      chan string
	firstExercise    chan string
	emptyExercises   chan bool
	reorderExercises chan bool

	// Channel for sending the sync code.
	SyncCode chan string
}

func newUser() *user {
	return &user{newExercise: make(chan string),
		firstExercise:    make(chan string),
		emptyExercises:   make(chan bool),
		reorderExercises: make(chan bool),
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
	t.Append(u.loginTabContainer(a, w, t))

	// Defer creation and storage of a new password hash with random salt.
	defer func() {
		if key, err := crypto.SaveNewPasswordHash(u.password, u.username, a); err != nil {
			fyne.LogError("Error on generating password hash", err)
		} else {
			u.data.Write(&key, u.username)
		}
	}()

	// Set the window to a good size, add the content and lastly run the application.
	w.SetContent(t)
	w.Resize(fyne.NewSize(800, 550))
	w.ShowAndRun()
}
