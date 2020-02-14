package parse

import (
	"fmt"
	"strconv"
)

// Float is a wrapper around strconv.ParseFloat that handles the error to make the function usable inline.
func Float(input string) float64 {
	if input == "" {
		return 0
	}

	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Print(err)
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
		fmt.Print(err)
	}

	return uint(output)
}
