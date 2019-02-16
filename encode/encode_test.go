package encode

import (
	"testing"
)

// 1.1
func TestHexToBase64(t *testing.T) {
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	output, err := HexToBase64(inputHex)
	if err != nil {
		t.Error(err)
	}
	if output != expectedBase64 {
		t.Error("Output did not match expected base64")
		t.Fail()
	}
}

// Test that bad hex forces an error.
//
// I don't think there's a good way to force that write error,
// without somehow preventing the process from allocating memory.
func TestHexToBase64Failures(t *testing.T) {
	inputHex := "AAA"
	_, err := HexToBase64(inputHex)
	if err == nil {
		t.Error("Expected error but got none")
	}
}
