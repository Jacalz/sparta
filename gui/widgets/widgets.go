package widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// ExtendedEntry is used to make an entry that reacts to key presses.
type ExtendedEntry struct {
	widget.Entry

	// PressAction for pressing a button on return.
	*PressAction

	// Fields related to switching entry with button.
	ButtonPress  fyne.KeyName
	EntryToFocus *ExtendedEntry
	Window       fyne.Window
}

// PressAction handles the Button press action.
type PressAction struct {
	Button widget.Button
}

// TypedKey handles the key presses inside our UsernameEntry and uses Action to press the linked button.
func (e *ExtendedEntry) TypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn:
		e.PressAction.Button.OnTapped()
	case e.ButtonPress:
		e.Window.Canvas().Focus(e.EntryToFocus)
	default:
		e.Entry.TypedKey(ev)
	}
}

// NewExtendedEntry creates an ExtendedEntry button.
func NewExtendedEntry(placeholder string, password bool) *ExtendedEntry {
	entry := &ExtendedEntry{}

	// Extend the base widget.
	entry.ExtendBaseWidget(entry)

	// Set placeholder for the entry.
	entry.SetPlaceHolder(placeholder)

	// Check if we are creating a password entry.
	if password {
		entry.Password = true
	}

	return entry
}

// InitExtend adds extra data to the extended entry.
func (e *ExtendedEntry) InitExtend(window fyne.Window, key fyne.KeyName, focus *ExtendedEntry, press widget.Button) {
	e.PressAction = &PressAction{Button: press}
	e.EntryToFocus = focus
	e.ButtonPress = key
	e.Window = window
}

// NewEntryWithPlaceholder makes it easy to create entry widgets with placeholders.
func NewEntryWithPlaceholder(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(text)

	return entry
}
