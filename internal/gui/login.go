package gui

import (
	"github.com/Jacalz/sparta/internal/crypto"
	"github.com/Jacalz/sparta/internal/file"
	"github.com/Jacalz/sparta/internal/gui/widgets"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (u *user) loginTabContainer(a fyne.App, w fyne.Window, t *widget.TabContainer) *widget.TabItem {
	// Create the username entry.
	usernameEntry := widgets.NewAdvancedEntry("Username", false)

	// Create the password entry.
	passwordEntry := widgets.NewAdvancedEntry("Password", true)

	// Create the login button that will login the user.
	loginButton := widget.NewButtonWithIcon("Login", theme.ConfirmIcon(), nil)

	// newUserButton holds the button widget for creating a new user.
	newUserButton := widget.NewButtonWithIcon("Create New User", theme.ContentAddIcon(), func() {
		if !crypto.ValidInput(usernameEntry.Text, passwordEntry.Text, w) {
			return
		}

		// Create the file for the user.
		if err := file.CreateNewUser(usernameEntry.Text); err != nil {
			dialog.ShowError(err, w)
			return
		}

		// Store the password verification.
		if _, err := crypto.SaveNewPasswordHash(passwordEntry.Text, usernameEntry.Text, a); err != nil {
			dialog.ShowError(err, w)
			return
		}

		// Inform the user of the success and show the login button again.
		dialog.ShowInformation("A new user was created", "The new user was created without issues. You can now log in with it.", w)
		loginButton.Show()
	})

	loginButton.OnTapped = func() {
		if loginButton.Hidden {
			newUserButton.OnTapped()
			return
		}

		// Define err here so we can add to u.encryptionKey and u.Data directly.
		var err error

		u.encryptionKey, err = crypto.CorrectCredentials(usernameEntry.Text, passwordEntry.Text, a, w)
		if err != nil {
			dialog.ShowInformation("Wrong username and/or password", "The login credentials are incorrect, please try again.", w)
			return
		}

		// Check files and try to log in using the computed password hash.
		u.data, err = file.ReadData(&u.encryptionKey, usernameEntry.Text)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			// Store the username to the user struct and clear data in widgets.
			u.username, u.password = usernameEntry.Text, passwordEntry.Text
			usernameEntry.Text, passwordEntry.Text = "", ""

			// Run it all in a new goroutine to avoid stalling the main one.
			go func() {
				// Add all the content tabs to the interface.
				t.Append(widget.NewTabItemWithIcon("Exercises", theme.HomeIcon(), u.exercisesView(w, a)))
				t.Append(widget.NewTabItemWithIcon("Add Exercise", theme.ContentAddIcon(), u.addExerciseView(w)))
				t.Append(widget.NewTabItemWithIcon("Sync", theme.MailSendIcon(), u.syncView(w)))
				t.Append(widget.NewTabItemWithIcon("Settings", theme.SettingsIcon(), u.settingsView(w, a)))
				t.Append(widget.NewTabItemWithIcon("About", theme.InfoIcon(), aboutView()))

				// Remove the login tab now that we are logged in.
				t.RemoveIndex(0)

				// Select the tab index to avoid confusing fyne.
				t.SelectTabIndex(0)
			}()
		}

	}

	// No need to show the loginButton if no users exist.
	if file.NoExistingUsers() {
		loginButton.Hide()
	}

	// Extend the AdvancedEntry widgets with extra key press supports.
	usernameEntry.InitExtend(*loginButton, widgets.MoveAction{Down: true, DownEntry: passwordEntry, Window: w})
	passwordEntry.InitExtend(*loginButton, widgets.MoveAction{Up: true, UpEntry: usernameEntry, Window: w})

	if fyne.Device.IsMobile(fyne.CurrentDevice()) {
		return widget.NewTabItem("Login", fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewVBoxLayout(), usernameEntry, passwordEntry, loginButton, newUserButton),
			layout.NewSpacer(),
		))
	}

	return widget.NewTabItem("Login", fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewGridLayout(3), layout.NewSpacer(), widget.NewVBox(usernameEntry, passwordEntry, loginButton, newUserButton), layout.NewSpacer()),
		layout.NewSpacer(),
	))
}
