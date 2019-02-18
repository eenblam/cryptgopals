package xor

import (
	"testing"
)

func TestBitSum(t *testing.T) {
	cases := []struct {
		Input    byte
		Expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{255, 8},
	}
	for _, test := range cases {
		got := BitSum(test.Input)
		if got != test.Expected {
			t.Errorf("Expected %d, got %d for %s",
				test.Expected, got, []byte{test.Input})
		}
	}
}
