package gui

import (
	"regexp"

	"github.com/Jacalz/sparta/internal/gui/widgets"
	"github.com/Jacalz/sparta/internal/sync"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Regular expression for verifying sync code.
var validCode = regexp.MustCompile(`^\d\d?-\w{2,12}-\w{2,12}$`)

// SyncView displays the tab page for syncing data between devices.
func (u *user) syncView(w fyne.Window) fyne.CanvasObject {

	// aboutGroup is a small group to display information about file sharing.
	aboutLabel := widget.NewLabelWithStyle("The share support in Sparta enables the user to syncronize their exercises. The file is served encrypted and will be automatically decrypted on a second device using the same login credentials. It is as simple as starting the share on one device and then starting receiving on another device on the same network. The receiving computer will then get all the new activities added locally.", fyne.TextAlignCenter, fyne.TextStyle{})
	aboutLabel.Wrapping = fyne.TextWrapWord
	aboutGroup := widget.NewGroup("About Exercise Synchronization", aboutLabel)

	// sendDataButton starts the network server and shares the file over the local internet.
	startSendingDataButton := widget.NewButtonWithIcon("Start Syncing Exercises", theme.ViewRefreshIcon(), nil)

	// recieveCode makes it possible to type in receive code.
	recieveCodeEntry := widgets.NewAdvancedEntry("Receive Code", false)

	// recieveDataButton starts looking for shared data on the local network.
	recieveDataButton := widget.NewButtonWithIcon("Start Receiving Exercises", theme.MailComposeIcon(), nil)

	startSendingDataButton.OnTapped = func() {
		if len(u.data.Exercise) == 0 {
			dialog.ShowInformation("No exercises to syncronize", "You need to add exercises to be able to sync between devices.", w)
			return
		}

		// Disable the receive button to hinder people from receiving on same computer.
		recieveDataButton.Disable()

		go func() {
			// Start sharing our exercises.
			err := sync.StartSync(u.SyncCode, u.username, &u.encryptionKey)
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				dialog.ShowInformation("Synchronization successful", "The synchronization of exercises finsished successfully.", w)
			}

			// Clean up and make buttons usable again.
			recieveCodeEntry.SetText("")
			recieveDataButton.Enable()
		}()

		// Add the sync code to the recieveCodeEntry.
		recieveCodeEntry.SetText(<-u.SyncCode)
	}

	recieveDataButton.OnTapped = func() {
		if validCode.MatchString(recieveCodeEntry.Entry.Text) {
			// Disable the button to make sure that users can't do anything bad.
			startSendingDataButton.Disable()

			code := recieveCodeEntry.Entry.Text

			go func() {
				err := sync.Receive(&u.data, u.reorderExercises, u.firstExercise, &u.encryptionKey, code, u.username)
				if err != nil {
					dialog.ShowError(err, w)
				} else {
					dialog.ShowInformation("Synchronization successful", "The synchronization of exercises finsished successfully.", w)
				}
			}()

			// Clean up and make buttons usable again.
			recieveCodeEntry.SetText("")
			startSendingDataButton.Enable()
		}
	}

	// shareGroup is a group containing all the options for sharing data.
	shareGroup := widget.NewGroup("Syncronizing Data", startSendingDataButton)

	// Extend the recieveCodeEntry to add the option for receiving on pressing enter.
	recieveCodeEntry.InitExtend(*recieveDataButton, widgets.MoveAction{})

	// recieveGroup is a group containing all options related to receiving.
	recieveGroup := widget.NewGroup("Receiving Data", fyne.NewContainerWithLayout(layout.NewGridLayout(2), recieveCodeEntry, recieveDataButton))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), aboutGroup, layout.NewSpacer(), shareGroup, layout.NewSpacer(), recieveGroup)
}
