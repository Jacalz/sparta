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

	// Create the channel that we will use to shut down the server.
	shutdown := make(chan bool)

	// aboutGroup is a small group to display information about file sharing.
	aboutGroup := widget.NewGroup("About Exercise Syncronization", widget.NewLabelWithStyle("The share support in Sparta enables the user to syncronize their exercises.\nThe file is served encrypted and will be automatically decrypted on a second device using the same login credentials.\nIt is as simple as starting the share on one device and then starting recieving on another device on the same network.\nThe recieving computer will then get all the new activities added locally.", fyne.TextAlignCenter, fyne.TextStyle{}))

	// sendDataButton starts the network server and shares the file over the local internet.
	startSendingDataButton := widget.NewButtonWithIcon("Start sharing exercises", theme.ViewRefreshIcon(), func() {
		go share.StartServer(shutdown)
	})

	//stopSendingDataButton will stop the data sharing from running in the background.
	stopSendingDataButton := widget.NewButtonWithIcon("Stop sharing exercises", theme.CancelIcon(), func() {
		shutdown <- true
	})

	// shareGroup is a group containing all the options for sharing data.
	shareGroup := widget.NewGroup("Sharing Data", startSendingDataButton, widget.NewLabel(""), stopSendingDataButton)

	// recieveDataButton starts looking for shared data on the local network.
	recieveDataButton := widget.NewButtonWithIcon("Start recieving exercises", theme.MailComposeIcon(), func() {
		go share.Retrieve(exercises, newAddedExercise, key)
	})

	// recieveGroup is a group containing all options related to recieving.
	recieveGroup := widget.NewGroup("Recieving Data", recieveDataButton)

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), aboutGroup, layout.NewSpacer(), shareGroup, layout.NewSpacer(), recieveGroup)
}
