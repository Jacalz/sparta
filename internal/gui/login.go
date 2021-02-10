package gui

import (
	"github.com/Jacalz/sparta/internal/crypto"
	"github.com/Jacalz/sparta/internal/crypto/validate"
	"github.com/Jacalz/sparta/internal/file"
	"github.com/Jacalz/sparta/internal/gui/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (u *user) loginTabContainer(a fyne.App, w fyne.Window, t *container.AppTabs) *container.TabItem {
	// Create the username entry.
	usernameEntry := widgets.NewAdvancedEntry("Username", false)
	usernameEntry.Validator = validate.Username

	// Create the password entry.
	passwordEntry := widgets.NewAdvancedEntry("Password", true)
	passwordEntry.Validator = validate.Password

	// Create the login button that will login the user.
	loginButton := widget.NewButtonWithIcon("Login", theme.ConfirmIcon(), nil)

	// newUserButton holds the button widget for creating a new user.
	newUserButton := widget.NewButtonWithIcon("Create New User", theme.ContentAddIcon(), func() {
		if err := passwordEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
			return
		} else if err := usernameEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
			return
		}

		if file.ExistingUser(usernameEntry.Text) {
			dialog.ShowInformation("User already exists", "The requested user already exists. Try a different username.", w)
			return
		}

		// Create the file for the user and the directory if it doesn't exist.
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
		} else if err := passwordEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
			return
		}

		// Define err here so we can add to u.encryptionKey and u.Data directly.
		var err error

		u.encryptionKey, err = validate.Credentials(usernameEntry.Text, passwordEntry.Text, a, w)
		if err != nil {
			dialog.ShowError(err, w)
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
				t.Append(container.NewTabItemWithIcon("Exercises", theme.HomeIcon(), u.exercisesView(w, a)))
				t.Append(container.NewTabItemWithIcon("Add Exercise", theme.ContentAddIcon(), u.addExerciseView(w)))
				t.Append(container.NewTabItemWithIcon("Sync", theme.MailSendIcon(), u.syncView(w)))
				t.Append(container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), u.settingsView(w, a)))
				t.Append(container.NewTabItemWithIcon("About", theme.InfoIcon(), aboutView()))

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
	usernameEntry.OnSubmitted, passwordEntry.OnSubmitted = func(_ string) { loginButton.OnTapped() }, func(_ string) { loginButton.OnTapped() }

	if fyne.Device.IsMobile(fyne.CurrentDevice()) {
		return container.NewTabItem("Login", container.NewVBox(
			layout.NewSpacer(),
			container.NewVBox(usernameEntry, passwordEntry, loginButton, newUserButton),
			layout.NewSpacer(),
		))
	}

	return container.NewTabItem("Login", container.NewGridWithColumns(1,
		layout.NewSpacer(),
		container.NewGridWithColumns(3, layout.NewSpacer(), container.NewVBox(usernameEntry, passwordEntry, loginButton, newUserButton), layout.NewSpacer()),
		layout.NewSpacer(),
	))
}
