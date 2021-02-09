package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Jacalz/sparta/internal/assets"
	"github.com/Jacalz/sparta/internal/gui"
)

const appID = "com.github.jacalz.sparta"

func main() {
	a := app.NewWithID(appID)
	a.SetIcon(assets.AppIcon)
	w := a.NewWindow("Sparta")

	w.SetContent(gui.Create(a, w))
	w.Resize(fyne.NewSize(800, 550))
	w.ShowAndRun()
}
