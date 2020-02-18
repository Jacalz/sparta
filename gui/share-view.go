package gui

import (
	"sparta/file"
	"sparta/share"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// ShareView displays the tab page for syncing data between devices.
func ShareView(exercises *file.Data, newAddedExercise chan string, key *[32]byte) fyne.CanvasObject {

	// Create the channel that we will use to share the code needed for recieving the data.
	shareCodeChan := make(chan string)

	// aboutGroup is a small group to display information about file sharing.
	aboutGroup := widget.NewGroup("About Exercise Syncronization", widget.NewLabelWithStyle("The share support in Sparta enables the user to syncronize their exercises.\nThe file is served encrypted and will be automatically decrypted on a second device using the same login credentials.\nIt is as simple as starting the share on one device and then starting recieving on another device on the same network.\nThe recieving computer will then get all the new activities added locally.", fyne.TextAlignCenter, fyne.TextStyle{}))

	// recieveCode makes it possible to type in recieve code.
	recieveCodeEntry := NewEntryWithPlaceholder("Recieve Code")

	// recieveDataButton starts looking for shared data on the local network.
	recieveDataButton := widget.NewButtonWithIcon("Start recieving exercises", theme.MailComposeIcon(), func() {
		go share.Retrieve(exercises, newAddedExercise, key, recieveCodeEntry.Text)
	})

	// sendDataButton starts the network server and shares the file over the local internet.
	startSendingDataButton := widget.NewButtonWithIcon("Start sharing exercises", theme.ViewRefreshIcon(), func() {
		// Create a channel for comunicating when we are done.
		finished := make(chan struct{})

		// Disable the recieve button to hinder people from recieving on same computer.
		recieveDataButton.Disable()

		// Start the sharing.
		go share.StartSharing(shareCodeChan, finished)

		// Listen in on the channels and clear stuff up when we are finished.
		go func() {
			for {
				select {
				case code := <-shareCodeChan:
					recieveCodeEntry.SetText(code)
				case <-finished:
					recieveDataButton.Enable()
					return
				}
			}
		}()
	})

	// shareGroup is a group containing all the options for sharing data.
	shareGroup := widget.NewGroup("Sharing Data", startSendingDataButton)

	// recieveGroup is a group containing all options related to recieving.
	recieveGroup := widget.NewGroup("Recieving Data", fyne.NewContainerWithLayout(layout.NewGridLayout(2), recieveCodeEntry, recieveDataButton))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), aboutGroup, layout.NewSpacer(), shareGroup, layout.NewSpacer(), recieveGroup)
}
