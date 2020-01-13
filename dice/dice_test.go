package dice

import (
	"testing"
)

func TestRoll(t *testing.T) {
	valid := []string{"d4", "d6", "d8", "d10", "d12", "d20", "d100", "d42", "1", "4", "8", "15", "16", "23", "42", "425"}
	invalid := []string{"", "d", "d0", "d06", "a6", " d6", "d6 "}

	for _, dx := range valid {
		if _, err := Roll(dx); err != nil {
			t.Errorf("Expected roll to succeed with \"%v\", but it errored with \"%v\"", dx, err)
		}
	}
	for _, dx := range invalid {
		if _, err := Roll(dx); err == nil {
			t.Errorf(`Expected roll to fail with "%s", but it succeeded`, dx)
		}
	}
}
