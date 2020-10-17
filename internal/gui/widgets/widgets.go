package widgets

import (
	"errors"
	"regexp"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// AdvancedEntry is used to make an entry that reacts to key presses.
type AdvancedEntry struct {
	widget.Entry

	// PressFuncfor running a function on return.
	PressFunc func()

	// Fields related to switching entry with button.
	*MoveAction
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
		if a.PressFunc != nil {
			a.PressFunc()
		}
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
func (a *AdvancedEntry) InitExtend(pressFunc func(), move MoveAction) {
	a.PressFunc = pressFunc
	a.MoveAction = &move
}

// NewEntryWithPlaceholder makes it easy to create entry widgets with placeholders.
func NewEntryWithPlaceholder(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(text)

	return entry
}

// NewFormEntry returns a new validated entry with placeholder
func NewFormEntry(placeholder, reason string, validation *regexp.Regexp, multiline bool) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	entry.Validator = func(input string) error {
		if validation != nil && !validation.MatchString(input) {
			return errors.New(reason)
		}

		return nil // Nothing to validate with, same as having no validator.
	}

	entry.MultiLine = multiline
	if entry.MultiLine {
		entry.Wrapping = fyne.TextWrapWord
	}

	return entry
}
