package analysis

import (
	"bytes"
	"testing"
)

func TestTransposeBlocks(t *testing.T) {
	data := []byte("123456789")
	expected := []byte("147258369")
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
		t.Error("Expected error, but got none")
	}
}

func TestFlattenBlocks(t *testing.T) {
	data := []byte("123456789")
	transposed, err := TransposeBlocks(data, 3)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	flat, err := FlattenBlocks(transposed)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(flat, data) {
		t.Errorf("Flatten(Transpose(\"%s\")) returned %s", string(data), string(flat))
	}

	badBlocks := [][]byte{[]byte("147"), []byte("25"), []byte("369")}
	_, err = FlattenBlocks(badBlocks)
	if err == nil {
		t.Error("Expected error, but got none")
	}
}
