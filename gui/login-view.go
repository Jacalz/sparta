package gui

import (
	"sparta/crypto"
	"sparta/file"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// CheckValidInput checks if the inputed username and passwords are valid and creates a message if they are not.
func CheckValidInput(username, password string, window fyne.Window) (valid bool) {
	if username == "" || password == "" {
		dialog.ShowInformation("Missing username/password", "Please provide both username and password.", window)
	} else if username == password {
		dialog.ShowInformation("Identical username and password", "Use separate usernames and passwords.", window)
	} else if len(password) < 8 {
		dialog.ShowInformation("Too short password", "The password should be eight characters or longer.", window)
	} else {
		valid = true
	}

	return valid
}

// TODO: Logout support and fix weirdly cut dialog boxes.

// ShowLoginPage shows the login page that handles the inertaface for logging in.
func ShowLoginPage(app fyne.App, window fyne.Window, user *user) {
	// Initialize the login form that we are to be using.
	username := NewExtendedEntry("Username", false)

	// Initialize the password input box that we are to be using.
	password := NewExtendedEntry("Password", true)

	// Create the login button that will calculate the 32bit long sha256 hash.
	loginButton := widget.NewButton("Login", func() {
		// Check the inputed data to handle invalid inputs.
		valid := CheckValidInput(username.Text, password.Text, window)
		if !valid {
			return
		}

		// Calculate the sha256 hash of the username and password.
		user.EncryptionKey = crypto.Hash(username.Text, password.Text)

		// Store the username and password to user structs and clear data in widgets.
		user.Username, username.Text = username.Text, ""
		user.Password, password.Text = password.Text, ""

		// Check for the file where we store the data. The user inputed the wrong password if we get an error.
		exercises, err := file.Check(&user.EncryptionKey)
		if err != nil {
			dialog.ShowInformation("Wrong username or password", "The login credentials are incorrect, please try again.", window)
			return
		}

		// Store the exercises inside the user struct.
		user.Data = exercises

		ShowMainDataView(window, app, user)
	})

	// Add the Action component to make actions work inside the struct. This is used to press the loginButton on pressing enter/return ton the keyboard.
	username.Action, password.Action = &Action{*loginButton}, &Action{*loginButton}

	// Set the content to be displayed. It is the userName, userPassword fields and the login button inside a layout.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), username, password, loginButton))

	// Set a sane default for the window size on login and make sure that it isn't resizable.
	window.Resize(fyne.NewSize(400, 100))
	window.SetFixedSize(true)

	// Update widgets if it is a first run.
	if file.FirstRun() {
		username.SetPlaceHolder("New Username")
		password.SetPlaceHolder("New Password")
		loginButton.SetText("Create User and Login")
	}

	// Check that we are using the right theme.
	switch app.Preferences().StringWithFallback("Theme", "Light") {
	case "Dark":
		app.Settings().SetTheme(theme.DarkTheme())
	case "Light":
		app.Settings().SetTheme(theme.LightTheme())
	}

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}
