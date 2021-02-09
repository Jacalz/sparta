package gui

import (
	"github.com/Jacalz/sparta/internal/assets"
	"github.com/Jacalz/sparta/internal/file/parse"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const version = "v0.9.0"
const url = "https://github.com/Jacalz/sparta/releases/tag/" + version

// AboutView displays the logo and a version link for application information.
func aboutView() fyne.CanvasObject {
	logo := canvas.NewImageFromResource(assets.AppIcon)
	logo.SetMinSize(fyne.NewSize(256, 256))

	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabelWithStyle("Sparta", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("-", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewHyperlinkWithStyle(version, parse.URL(url), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}
