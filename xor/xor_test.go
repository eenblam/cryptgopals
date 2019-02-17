package xor

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

// 1.2
func TestXORn(t *testing.T) {
	left, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	right, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	n := len(left)
	result := make([]byte, n)
	err := XORn(result, left, right, n)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if !bytes.Equal(result, expected) {
		fmt.Println("Output bytes did not match expected bytes")
		t.Fail()
	}
}
func TestXORnErrors(t *testing.T) {
	a := []byte{0xFF, 0x00, 0x0F}
	b := []byte{0x00, 0xFF}
	c := make([]byte, 3)
	xorError := XORn(c, a, b, 3)
	if xorError == nil {
		fmt.Println("No error on length mismatch")
	}
}

// Related to 1.3
func TestXORByte(t *testing.T) {
	a := []byte{0xFF, 0x00, 0x0F}
	b := byte(0x00)
	expected := []byte{0xFF, 0x00, 0x0F}
	results := make([]byte, 3)
	xorError := XORByte(results, a, b, 3)
	if xorError != nil {
		t.Error(xorError)
	}
	if !bytes.Equal(results, expected) {
		t.Error("Output bytes do not equal expected bytes")
	}
}

// 1.4
func TestXORRepeat(t *testing.T) {
	data := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	ciphertext := XORRepeat(data, key)
	cipherhex := hex.EncodeToString(ciphertext)
	expectedhex := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if !reflect.DeepEqual(cipherhex, expectedhex) {
		t.Errorf("Expected\n%s\nGot\n%s\n", expectedhex, cipherhex)
	}
}
