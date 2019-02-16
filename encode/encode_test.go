package encode

import (
	"fmt"
	"testing"
)

// 1.1
func TestHexToBase64(t *testing.T) {
	inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	output, err := HexToBase64(inputHex)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if output != expectedBase64 {
		fmt.Println("Output did not match expected base64")
		t.Fail()
	}
}
