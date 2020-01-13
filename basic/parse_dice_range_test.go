package basic

import "testing"

func TestParseDiceRange(t *testing.T) {
	tests := []struct {
		input string
		min   uint
		max   uint
		err   bool
	}{
		{"", 0, 0, true},
		{"abc", 0, 0, true},
		{"1", 1, 1, false},
		{"23", 23, 23, false},
		{"0x10", 0, 0, true},
		{"1-4", 1, 4, false},
		{"5-23", 5, 23, false},
		{"1-2-3", 0, 0, true},
	}

	for _, test := range tests {
		min, max, err := parseDiceRange(test.input)
		if err == nil && test.err {
			t.Errorf(`("%s"): Expected error, but no error occurred`, test.input)
		}
		if err != nil && !test.err {
			t.Errorf(`("%s"): Expected no error, but got error "%v"`, test.input, err)
		}
		if min != test.min || max != test.max {
			t.Errorf(`("%s"): Expected (min: %v, max: %v), but got (%v, %v)`, test.input, test.min, test.max, min, max)
		}
	}
}
