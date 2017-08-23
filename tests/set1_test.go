package set1

import (
    "bytes"
    "encoding/hex"
    "fmt"
    "testing"

    "github.com/eenblam/cryptgopals/encode"
    "github.com/eenblam/cryptgopals/xor"
)

func TestHexToBase64(t *testing.T) {
    inputHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
    output, err := encode.HexToBase64(inputHex)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    if output != expectedBase64 {
        fmt.Println("Output did not match expected base64")
        t.Fail()
    }
}
func TestXORn(t *testing.T) {
    left, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
    right, _ := hex.DecodeString("686974207468652062756c6c277320657965")
    expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

    n := len(left)
    result := make([]byte, n)
    err := xor.XORn(result, left, right, n)
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
    xorError := xor.XORn(c, a, b, 3)
    if xorError == nil {
        fmt.Println("No error on length mismatch")
    }
}
