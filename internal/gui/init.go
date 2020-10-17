package gui

import (
	"github.com/Jacalz/sparta/internal/assets"
	"github.com/Jacalz/sparta/internal/crypto"
	"github.com/Jacalz/sparta/internal/file"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
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
	a.SetIcon(assets.AppIcon)

	// Create the window for our user interface.
	w := a.NewWindow("Sparta")
	w.SetMaster()

	// Create the user struct type for later use.
	u := newUser()

	// Check that we are using the right theme.
	checkTheme(a.Preferences().StringWithFallback("Theme", "Adaptive (requires restart)"), a)

	// Create the tab handler for the user interface and set up the login view.
	t := &container.AppTabs{}
	t.Append(u.loginTabContainer(a, w, t))

	// Create and store of a new password hash with random salt on window close.
	w.SetOnClosed(func() {
		if key, err := crypto.SaveNewPasswordHash(u.password, u.username, a); err != nil {
			fyne.LogError("Error on generating password hash", err)
		} else {
			u.data.Write(&key, u.username)
		}
	})

	// Set the window to a good size, add the content and lastly run the application.
	w.SetContent(t)
	w.Resize(fyne.NewSize(800, 550))
	w.ShowAndRun()
}
