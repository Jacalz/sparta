package widgets

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/sparta/internal/crypto/validate"
)

// NewFormEntry returns a new validated entry with placeholder
func NewFormEntry(placeholder, reason string, validation *regexp.Regexp, multiline bool) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	entry.Validator = validate.NewRegexp(validation, reason)

	entry.MultiLine = multiline
	if entry.MultiLine {
		entry.Wrapping = fyne.TextWrapWord
	} else {
		entry.Wrapping = fyne.TextWrapOff
	}

	return entry
}
