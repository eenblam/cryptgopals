package rand

// Generate a random key, then keep encrypting AES-128-ECB with it.
// https://cryptopals.com/sets/2/challenges/12

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/eenblam/cryptgopals/cipher"
	"github.com/eenblam/cryptgopals/encode"
)

func NewRandECB() *RandECB {
	key := RandKey()
	pad, err := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	if err != nil {
		panic("Bad base64 string! You broke the buillllllld!")
	}
	return &RandECB{key, pad}
}

type RandECB struct {
	key      []byte
	rightPad []byte
}

// Encrypt encrypts the provided plaintext under the RandECB's fixed key.
//
// For 2.12.
func (r *RandECB) Encrypt(plaintext []byte) ([]byte, error) {
	var bs bytes.Buffer
	n, err := bs.Write(plaintext)
	if err != nil {
		return nil, err
	}
	if n != len(plaintext) {
		return nil, fmt.Errorf("Could only write %d of %d bytes", n, len(plaintext))
	}
	n, err = bs.Write(r.rightPad)
	if err != nil {
		return nil, err
	}
	if n != len(r.rightPad) {
		return nil, fmt.Errorf("Could only write %d of %d bytes", n, len(plaintext))
	}
	paddedPlaintext, err := encode.PadBytesTo(bs.Bytes(), 16)
	if err != nil {
		return nil, err
	}
	return cipher.ECBAESEncrypt(r.key, paddedPlaintext)
}
