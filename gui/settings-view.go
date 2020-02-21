package gui

import (
	"sparta/file/encrypt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// TODO: Multi user support by labling the data file exercises-user.xml

// SettingsView contains the gui information for the settings screen.
func SettingsView(window fyne.Window, app fyne.App, dataLabel *widget.Label, user *user) fyne.CanvasObject {

	// TODO: Add setting for changing language.

	// Make it possible for the user to switch themes.
	themeSwitcher := widget.NewSelect([]string{"Dark", "Light"}, func(selected string) {
		switch selected {
		case "Dark":
			app.Settings().SetTheme(theme.DarkTheme())
		case "Light":
			app.Settings().SetTheme(theme.LightTheme())
		}

		// Set the theme to the selected one and save it using the preferences api in fyne.
		app.Preferences().SetString("Theme", selected)
	})

	// Default theme is light and thus we set the placeholder to that and then refresh it (without a refresh, it doesn't show until hovering on to widget).
	themeSwitcher.PlaceHolder = app.Preferences().StringWithFallback("Theme", "Light")
	themeSwitcher.Refresh()

	// Add the theme switcher next to a label.
	themeChanger := fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("Application Theme"), themeSwitcher)

	// An entry for typing the new username.
	usernameEntry := NewEntryWithPlaceholder("New Username")

	// Create the button used for changing the username.
	usernameButton := widget.NewButtonWithIcon("Change Username", theme.ConfirmIcon(), func() {
		// Check that the username is valid.
		if usernameEntry.Text == user.Password || usernameEntry.Text == "" {
			dialog.ShowInformation("Please enter a valid username", "Usernames need to not be empty and not the same as the password.", window)
		} else {
			// Ask the user to confirm what we are about to do.
			dialog.ShowConfirm("Are you sure that you want to continue?", "The action will permanently change your username.", func(change bool) {
				if change {
					// Calculate the new PasswordKey.
					user.EncryptionKey = encrypt.EncryptionKey(usernameEntry.Text, user.Password)

					// Clear out the text inside the entry.
					usernameEntry.SetText("")

					// Write the data encrypted using the new key and do so concurrently.
					go user.ExerciseData.Write(&user.EncryptionKey)
				}
			}, window)
		}

	})

	usernameChanger := fyne.NewContainerWithLayout(layout.NewGridLayout(2), usernameEntry, usernameButton)

	// Create the entry for updating the password.
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("New Password")

	// Create the button used for changing the password.
	passwordButton := widget.NewButtonWithIcon("Change Password", theme.ConfirmIcon(), func() {
		// Check that the password is valid.
		if len(passwordEntry.Text) < 8 || passwordEntry.Text == usernameEntry.Text {
			dialog.ShowInformation("Please enter a valid password", "Passwords need to be at least eight characters long.", window)
		} else {
			// Ask the user to confirm what we are about to do.
			dialog.ShowConfirm("Are you sure that you want to continue?", "The action will permanently change your password.", func(change bool) {
				if change {
					// Calculate the new PasswordKey.
					user.EncryptionKey = encrypt.EncryptionKey(user.Username, passwordEntry.Text)

					// Clear out the text inside the entry.
					passwordEntry.SetText("")

					// Write the data encrypted using the new key and do so concurrently.
					go user.ExerciseData.Write(&user.EncryptionKey)
				}
			}, window)
		}
	})

	// passwordChanger holds the widgets for the password changer.
	passwordChanger := fyne.NewContainerWithLayout(layout.NewGridLayout(2), passwordEntry, passwordButton)

	// revertToDefaultSettings reverts all settings to their default values.
	revertToDefaultSettings := widget.NewButtonWithIcon("Reset settings to default values", theme.ViewRefreshIcon(), func() {
		// Make sure to set the themeSwitcher to Dark if it isn't already.
		if app.Preferences().String("Theme") != "Dark" {
			// Set the widget placeholder to dark.
			themeSwitcher.PlaceHolder = "Dark"
			themeSwitcher.Refresh()

			// Set the visible theme to dark.
			app.Settings().SetTheme(theme.DarkTheme())

			// Set the saved theme to dark.
			app.Preferences().SetString("Theme", "Dark")
		}
	})
	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Delete all saved activities", theme.DeleteIcon(), func() {

		// Ask the user to confirm what we are about to do.
		dialog.ShowConfirm("Are you sure that you want to continue?", "Deleting your data will remove all of your exercises and activities.", func(remove bool) {
			if remove {
				// Run the delete function and do it concurrently to avoid stalling the thread with file io.
				go user.ExerciseData.Delete()

				// Clear all the data inside the data label so we don't display the old data.
				dataLabel.SetText("No exercieses have been created yet.")
			}
		}, window)
	})

	// userInterfaceSettings is a group holding widgets related to user interface settings such as theme.
	userInterfaceSettings := widget.NewGroup("User Interface Settings", themeChanger)

	// accountPasswordSettings groups together all settings related to usernames and passwords.
	accountPasswordSettings := widget.NewGroup("Account and Password Settings", usernameChanger, passwordChanger)

	// advancedSettings is a group holding widgets related to advanced settings.
	advancedSettings := widget.NewGroup("Advanced Settings", revertToDefaultSettings, widget.NewLabel(""), deleteButton)

	// settingsContentView holds all widget groups and content for the settings page.
	settingsContentView := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), userInterfaceSettings, layout.NewSpacer(), accountPasswordSettings, layout.NewSpacer(), advancedSettings)

	return widget.NewScrollContainer(settingsContentView)
}
