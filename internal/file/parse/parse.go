package parse

import (
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
)

// Float is a wrapper around strconv.ParseFloat that handles the error to make the function usable inline.
func Float(input string) float64 {
	if input == "" {
		return 0
	}

	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fyne.LogError("Could not parse the Float value", err)
		return 0
	}

	return output
}

// Uint is just a wrapper around strconv.Atoi() returning a uint and handling the error.
func Uint(input string) uint32 {
	if input == "" {
		return 0
	}

	output, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fyne.LogError("Could not parse the Uint value", err)
		return 0
	}

	return uint32(output)
}

// URL parses an url string into the *url.Url type.
func URL(input string) *url.URL {
	link, err := url.Parse(input)
	if err != nil {
		fyne.LogError("Could not parse URL string", err)
	}

	return link
}
