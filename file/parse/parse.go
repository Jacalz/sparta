package parse

import (
	"strconv"

	"fyne.io/fyne"
)

// Float is a wrapper around strconv.ParseFloat that handles the error to make the function usable inline.
func Float(input string) float64 {
	if input == "" {
		return 0
	}

	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fyne.LogError("Error on parsing float", err)
		return 0
	}

	return output
}

// Uint is just a wrapper around strconv.Atoi() returning a uint and handling the error.
func Uint(input string) uint {
	if input == "" {
		return 0
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		fyne.LogError("Error on parsing int", err)
		return 0
	}

	return uint(output)
}
