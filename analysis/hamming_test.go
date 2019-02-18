package analysis

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

func TestBitSumBytes(t *testing.T) {
	cases := []struct {
		Input    []byte
		Expected int
	}{
		{[]byte{}, 0},
		{[]byte{0}, 0},
		{[]byte{0, 1, 2, 3, 255}, 12},
	}
	for _, test := range cases {
		got := BitSumBytes(test.Input)
		if got != test.Expected {
			t.Errorf("Expected %d, got %d for %s",
				test.Expected, got, test.Input)
		}
	}
}

// Tests both Hamming and HammingString
func TestHamming(t *testing.T) {
	cases := []struct {
		A        string
		B        string
		Expected int
	}{
		{
			"this is a test",
			"wokka wokka!!!",
			37,
		},
		{
			"",
			"",
			0,
		},
	}
	for _, test := range cases {
		got, err := HammingString(test.A, test.B)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if got != test.Expected {
			t.Errorf("Expected %d, got %d", test.Expected, got)
		}
	}
}

// Tests both Hamming and HammingString
func TestHammingError(t *testing.T) {
	_, err := HammingString("short", "long")
	if err == nil {
		t.Error("Expected error, but got none")
	}
}

func TestFirstNHamming(t *testing.T) {
	bs := []byte("this is a testwokka wokka!!!ignore this part")
	expectedDistance := 37
	distance, err := FirstNHamming(bs, 14)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if distance != expectedDistance {
		t.Errorf("Expected %d, got %d", expectedDistance, distance)
	}
	_, shouldBeErr := FirstNHamming(bs, 1+(len(bs)/2))
	if shouldBeErr == nil {
		t.Error("Expected error, got none")
	}
}
