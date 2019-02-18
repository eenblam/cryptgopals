package analysis

import (
	"bytes"
	"testing"
)

func TestTransposeBlocks(t *testing.T) {
	data := []byte("123456789")
	expected := []byte("147258369")
	//expected := [][]byte([]byte("147"), []byte("258"), []byte("369"))
	got, err := TransposeBlocks(data, 3)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	gotFlat := bytes.Join(got, []byte{})
	if !bytes.Equal(expected, gotFlat) {
		t.Errorf("Expected %s, got %s", string(expected), string(gotFlat))
	}

	_, shouldBeErr := TransposeBlocks(data, 5)
	if shouldBeErr == nil {
		t.Errorf("Expected error, but got none")
	}
}
