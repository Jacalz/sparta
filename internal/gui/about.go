package gui

import (
	"github.com/Jacalz/sparta/internal/assets"
	"github.com/Jacalz/sparta/internal/file/parse"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

const version string = "v0.8.0"

// AboutView displays the logo and a version link for application information.
func aboutView() fyne.CanvasObject {
	logo := canvas.NewImageFromResource(assets.AppIcon)
	logo.SetMinSize(fyne.NewSize(300, 300))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHBox(
			layout.NewSpacer(),
			widget.NewLabelWithStyle("Sparta", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("-", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewHyperlinkWithStyle(version, parse.URL("https://github.com/Jacalz/sparta/releases/tag/"+version), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}
