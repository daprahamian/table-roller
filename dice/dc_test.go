package dice

import "testing"

func TestDCRoll(t *testing.T) {
	tests := []uint{0, 1, 2, 3, 10, 42, 5000}

	for _, val := range tests {
		actual := DC(val).Roll()
		if val != actual {
			t.Errorf(`Expected DC(%v) = %v, but got %v`, val, val, actual)
		}
	}
}
