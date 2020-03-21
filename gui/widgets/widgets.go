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
	*ButtonFocus
}

// PressAction handles the Button press action.
type PressAction struct {
	Button widget.Button
}

// ButtonFocus handles focusing a different entry on a specific key press.
type ButtonFocus struct {
	ButtonPress  fyne.KeyName
	EntryToFocus *ExtendedEntry
	Window       fyne.Window
}

// TypedKey handles the key presses inside our UsernameEntry and uses Action to press the linked button.
func (e *ExtendedEntry) TypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn:
		if e.PressAction != nil {
			e.PressAction.Button.OnTapped()
		}
	case e.ButtonPress:
		if e.ButtonFocus != nil {
			e.Window.Canvas().Focus(e.ButtonFocus.EntryToFocus)
		}
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
func (e *ExtendedEntry) InitExtend(press widget.Button, focus ButtonFocus) {
	e.PressAction = &PressAction{Button: press}
	e.ButtonFocus = &focus
}

// NewEntryWithPlaceholder makes it easy to create entry widgets with placeholders.
func NewEntryWithPlaceholder(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(text)

	return entry
}
