package xor

import (
    "bytes"
    "encoding/hex"
    "fmt"
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
func testXORByte(t *testing.T) {
    a := []byte{0xFF, 0x00, 0x0F}
    b := byte(0x00)
    expected := []byte{0xFF, 0x00, 0xF0}
    results := make([]byte, 3)
    xorError := XORByte(results, a, b, 3)
    if xorError != nil {
        fmt.Println(xorError)
        t.Fail()
    }
    if !bytes.Equal(results, expected) {
        fmt.Println("Output bytes do not equal expected bytes")
    }
}
