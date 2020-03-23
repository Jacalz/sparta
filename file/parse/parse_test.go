package parse

import "testing"

func TestTableFloat(t *testing.T) {
	var tests = []struct {
		input    string
		expected float64
	}{
		{"55.5555", 55.5555},
		{"000.00145", 0.00145},
		{"50.56", 50.56},
		{"INVALID", 0},
		{"", 0},
	}

	for _, test := range tests {
		if output := Float(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, received %v", test.input, test.expected, output)
		}
	}
}

func TestTableUint(t *testing.T) {
	var tests = []struct {
		input    string
		expected uint
	}{
		{"55", 55},
		{"00145", 145},
		{"5067", 5067},
		{"INVALID", 0},
		{"", 0},
	}

	for _, test := range tests {
		if output := Uint(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, received %v", test.input, test.expected, output)
		}
	}
}
