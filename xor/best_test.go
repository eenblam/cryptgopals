package xor

import (
	"encoding/hex"
	"reflect"
	"testing"
)

// 1.3
func TestBestXORByte(t *testing.T) {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedString := "Cooking MC's like a pound of bacon"
	expected := []byte(expectedString)
	data, _ := hex.DecodeString(s)
	result, err := BestXORByte(data)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !reflect.DeepEqual(result.PlainText, expected) {
		gotString := string(result.PlainText)
		t.Errorf("Expected %s, got %s", expectedString, gotString)
	}
}
