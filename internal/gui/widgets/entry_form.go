package widgets

import (
	"regexp"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/sparta/internal/crypto/validate"
)

func validateDate(input string) error {
	_, err := time.Parse("2006-01-02", input)
	return err
}

func validateTime(input string) error {
	_, err := time.Parse("15:04", input)
	return err
}

// NewFormEntry returns a new validated entry with placeholder
func NewFormEntry(placeholder, reason string, validation *regexp.Regexp, multiline bool) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)

	if placeholder == "YYYY-MM-DD" {
		entry.Validator = validateDate
	} else if placeholder == "HH:MM" {
		entry.Validator = validateTime
	} else {
		entry.Validator = validate.NewRegexp(validation, reason)
	}

	entry.MultiLine = multiline
	if entry.MultiLine {
		entry.Wrapping = fyne.TextWrapWord
	} else {
		entry.Wrapping = fyne.TextWrapOff
	}

	return entry
}
