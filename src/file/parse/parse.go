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

	output, err := strconv.ParseFloat(input, 32)
	if err != nil {
		fmt.Print(err)
	}

	return output
}

// Int is just a wrapper around strconv.Atoi().
func Int(input string) int {
	if input == "" {
		return 0
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print(err)
	}

	return output
}
