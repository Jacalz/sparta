package gui

import (
	"github.com/Jacalz/sparta/internal/crypto"
	"github.com/Jacalz/sparta/internal/file"

	"fyne.io/fyne"
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

	// Preferences
	theme string

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

// Create will create our graphical user interface.
func Create(a fyne.App, w fyne.Window) *container.AppTabs {
	// Create the user struct type for later use.
	u := newUser()

	// Check that we are using the right theme.
	u.theme = checkTheme(a.Preferences().StringWithFallback("Theme", "Adaptive (requires restart)"), a)

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

	return t
}
