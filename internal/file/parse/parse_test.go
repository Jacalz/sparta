package parse

import "testing"

var floatTests = []struct {
	input    string
	expected float64
}{
	{"55.5555", 55.5555},
	{"000.00145", 0.00145},
	{"50.56", 50.56},
	{"INVALID", 0},
	{"", 0},
}

var uintTests = []struct {
	input    string
	expected uint32
}{
	{"55", 55},
	{"00145", 145},
	{"5067", 5067},
	{"INVALID", 0},
	{"", 0},
}

func TestTableFloat(t *testing.T) {
	for _, test := range floatTests {
		if output := Float(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, received %v", test.input, test.expected, output)
		}
	}
}

func TestTableUint(t *testing.T) {
	for _, test := range uintTests {
		if output := Uint(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, received %v", test.input, test.expected, output)
		}
	}
}
