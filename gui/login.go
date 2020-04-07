package gui

import (
	"sparta/crypto"
	"sparta/file"
	"sparta/gui/widgets"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// ValidInput checks if the inputed username and passwords are valid and creates a message if they are not.
func ValidInput(username, password string, window fyne.Window) (valid bool) {
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

func (u *user) LoginTabContainer(a fyne.App, w fyne.Window, t *widget.TabContainer) *widget.TabItem {
	// Create the username entry.
	usernameEntry := widgets.NewAdvancedEntry("Username", false)

	// Create the password entry.
	passwordEntry := widgets.NewAdvancedEntry("Password", true)

	// Create the login button that will login the user.
	loginButton := widget.NewButton("Login", func() {
		if ValidInput(usernameEntry.Text, passwordEntry.Text, w) {
			u.EncryptionKey = crypto.Hash(usernameEntry.Text, passwordEntry.Text)

			// Deine err here so we can add exercises to u.Data directly.
			var err error

			// Check files and try to log in using the computed password hash.
			u.Data, err = file.Check(&u.EncryptionKey)
			if err != nil {
				dialog.ShowInformation("Wrong username or password", "The login credentials are incorrect, please try again.", w)
			} else {
				// Store the username and password to user structs and clear data in widgets.
				u.Username, usernameEntry.Text = usernameEntry.Text, ""
				u.Password, passwordEntry.Text = passwordEntry.Text, ""

				// Run it all in a new goroutine to avoid stalling the main one.
				go func() {
					// Add all the content tabs to the interface.
					t.Append(widget.NewTabItemWithIcon("Exercises", theme.HomeIcon(), u.ExercisesView(w, a)))
					t.Append(widget.NewTabItemWithIcon("Add Exercise", theme.ContentAddIcon(), u.AddExerciseView(w)))
					t.Append(widget.NewTabItemWithIcon("Sync", theme.MailSendIcon(), u.SyncView(w)))
					t.Append(widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), u.SettingsView(w, a)))
					t.Append(widget.NewTabItemWithIcon("About", theme.InfoIcon(), AboutView()))

					// Remove the login tab now that we are logged in.
					t.RemoveIndex(0)

					// Select the tab index to avoid confusing fyne.
					t.SelectTabIndex(0)
				}()
			}
		}
	})

	// Update widgets if it is a first run.
	if file.FirstRun() {
		usernameEntry.SetPlaceHolder("New Username")
		passwordEntry.SetPlaceHolder("New Password")
		loginButton.SetText("Create User and Login")
	}

	// Extend the AdvancedEntry widgets with extra key press supports.
	usernameEntry.InitExtend(*loginButton, widgets.MoveAction{Down: true, DownEntry: passwordEntry, Window: w})
	passwordEntry.InitExtend(*loginButton, widgets.MoveAction{Up: true, UpEntry: usernameEntry, Window: w})

	if fyne.Device.IsMobile(fyne.CurrentDevice()) {
		return widget.NewTabItem("Login", fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), usernameEntry, passwordEntry, loginButton),
			layout.NewSpacer(),
		))
	}

	return widget.NewTabItem("Login", fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewGridLayout(3), layout.NewSpacer(), widget.NewVBox(usernameEntry, passwordEntry, loginButton), layout.NewSpacer()),
		layout.NewSpacer(),
	))
}
