package widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// AdvancedEntry is used to make an entry that reacts to key presses.
type AdvancedEntry struct {
	widget.Entry

	// PressAction for pressing a button on return.
	*PressAction

	// Fields related to switching entry with button.
	*MoveAction
}

// PressAction handles the Button press action.
type PressAction struct {
	Button widget.Button
}

// MoveAction handles focusing a different entry on a specific arrow key press.
type MoveAction struct {
	// Entry widgets to focus on up and down arrow keys respectively.
	UpEntry   *AdvancedEntry
	DownEntry *AdvancedEntry

	// Bools to turn  up and down of in case they are not needed.
	Up, Down bool

	// Window used for focus calls.
	Window fyne.Window
}

// TypedKey handles the key presses inside our UsernameEntry and uses Action to press the linked button.
func (a *AdvancedEntry) TypedKey(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn:
		a.PressAction.Button.OnTapped()
	case fyne.KeyUp:
		if a.Up {
			a.Window.Canvas().Focus(a.MoveAction.UpEntry)
		}
	case fyne.KeyDown:
		if a.Down {
			a.Window.Canvas().Focus(a.MoveAction.DownEntry)
		}
	default:
		a.Entry.TypedKey(ev)
	}
}

// NewAdvancedEntry creates an ExtendedEntry button.
func NewAdvancedEntry(placeholder string, password bool) *AdvancedEntry {
	entry := &AdvancedEntry{}

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
func (a *AdvancedEntry) InitExtend(press widget.Button, move MoveAction) {
	a.PressAction = &PressAction{Button: press}
	a.MoveAction = &move
}

// NewEntryWithPlaceholder makes it easy to create entry widgets with placeholders.
func NewEntryWithPlaceholder(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(text)

	return entry
}
