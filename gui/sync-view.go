package gui

import (
	"sparta/gui/widgets"
	"sparta/sync"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// SyncView displays the tab page for syncing data between devices.
func (u *user) SyncView(window fyne.Window) fyne.CanvasObject {

	// Create the channel that we will use to share the code needed for receiving the data.
	syncCodeChan := make(chan string)

	// aboutGroup is a small group to display information about file sharing.
	aboutGroup := widget.NewGroup("About Exercise Synchronization", widget.NewLabelWithStyle("The share support in Sparta enables the user to syncronize their exercises.\nThe file is served encrypted and will be automatically decrypted on a second device using the same login credentials.\nIt is as simple as starting the share on one device and then starting receiving on another device on the same network.\nThe receiving computer will then get all the new activities added locally.", fyne.TextAlignCenter, fyne.TextStyle{}))

	// sendDataButton starts the network server and shares the file over the local internet.
	startSendingDataButton := widget.NewButtonWithIcon("Start Syncing Exercises", theme.ViewRefreshIcon(), nil)

	// recieveCode makes it possible to type in receive code.
	recieveCodeEntry := widgets.NewAdvancedEntry("Receive Code", false)

	// recieveDataButton starts looking for shared data on the local network.
	recieveDataButton := widget.NewButtonWithIcon("Start Receiving Exercises", theme.MailComposeIcon(), nil)

	startSendingDataButton.OnTapped = func() {
		if len(u.Data.Exercise) == 0 {
			dialog.ShowInformation("No exercises to syncronize", "You need to add exercises to be able to sync between devices.", window)
			return
		}

		// Disable the receive button to hinder people from receiving on same computer.
		recieveDataButton.Disable()

		// Start the sharing.
		go sync.StartSync(syncCodeChan, u.Errors, u.FinishedSync)
	}

	recieveDataButton.OnTapped = func() {
		// Disable the button to make sure that users can't do anything bad.
		startSendingDataButton.Disable()

		// Start retrieving data.
		go sync.Retrieve(&u.Data, u.ReorderExercises, u.FirstExercise, u.Errors, u.FinishedSync, &u.EncryptionKey, recieveCodeEntry.Text)
	}

	// shareGroup is a group containing all the options for sharing data.
	shareGroup := widget.NewGroup("Syncronizing Data", startSendingDataButton)

	// Listen in on the channels sent from receive and send and errors.
	go func() {
		for {
			select {
			case code := <-syncCodeChan:
				recieveCodeEntry.SetText(code)
			case <-u.FinishedSync:
				dialog.ShowInformation("Synchronization successful", "The synchronization of exercises finsished successfully.", window)

				// Clean up and make buttons usable again.
				recieveCodeEntry.SetText("")
				recieveDataButton.Enable()
				startSendingDataButton.Enable()
			case err := <-u.Errors:
				dialog.ShowError(err, window)

				// Clean up and make buttons usable again.
				recieveDataButton.Enable()
				startSendingDataButton.Enable()
			}
		}
	}()

	// Extend the recieveCodeEntry to add the option for receiving on pressing enter.
	recieveCodeEntry.InitExtend(*recieveDataButton, widgets.MoveAction{})

	// recieveGroup is a group containing all options related to receiving.
	recieveGroup := widget.NewGroup("Receiving Data", fyne.NewContainerWithLayout(layout.NewGridLayout(2), recieveCodeEntry, recieveDataButton))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), aboutGroup, layout.NewSpacer(), shareGroup, layout.NewSpacer(), recieveGroup)
}
