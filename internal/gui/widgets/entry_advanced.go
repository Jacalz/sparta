package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// AdvancedEntry is used to make an entry that reacts to key presses.
type AdvancedEntry struct {
	widget.Entry
}

// TypedKey handles the key presses inside our UsernameEntry and uses Action to press the linked button.
func (a *AdvancedEntry) TypedKey(ev *fyne.KeyEvent) {
	canvas := fyne.CurrentApp().Driver().AllWindows()[0].Canvas()
	switch ev.Name {
	case fyne.KeyUp:
		canvas.FocusPrevious()
	case fyne.KeyDown:
		canvas.FocusNext()
	default:
		a.Entry.TypedKey(ev)
	}
}

// NewAdvancedEntry creates an ExtendedEntry button.
func NewAdvancedEntry(placeholder string, password bool) *AdvancedEntry {
	entry := &AdvancedEntry{}
	entry.ExtendBaseWidget(entry)
	entry.PlaceHolder = placeholder
	entry.Password = password
	return entry
}
