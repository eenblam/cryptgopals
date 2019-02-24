package encode

import (
	"bytes"
	"testing"
)

func TestPadBytesTo(t *testing.T) {
	_, err := PadBytesTo([]byte("asdf"), 256)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	// Test that extra block is appended when blocks align
	expected := []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 5, 5, 5, 5, 5}
	got, err := PadBytesTo([]byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 5)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func Test_2_9(t *testing.T) {
	data := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	got, err := PadBytesTo(data, 20)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestUnpadBytesBy(t *testing.T) {
	_, err := UnpadBytesBy([]byte("asdf"), 256)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	// Test that extra block is appended when blocks align
	expected := []byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	got, err := UnpadBytesBy([]byte{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 5, 5, 5, 5, 5}, 5)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
	// Should error when blockSize doesn't divide length
	_, err = UnpadBytesBy([]byte("ICE ICE ICE ICE\x04\x04\x04\x04"), 16)
	if err == nil {
		t.Error("Expected error, but got none")
	}
}

func Test_2_15(t *testing.T) {
	data := []byte("ICE ICE BABY\x04\x04\x04\x04")
	expected := []byte("ICE ICE BABY")
	got, err := UnpadBytesBy(data, 16)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Expected %s, got %s", string(expected), string(got))
	}

	_, err = UnpadBytesBy([]byte("ICE ICE BABY\x05\x05\x05\x05"), 16)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	_, err = UnpadBytesBy([]byte("ICE ICE BABY\x01\x02\x03\x04"), 16)
	if err == nil {
		t.Error("Expected error, but got none")
	}
}
