package xor

import (
	"fmt"
)

// XORByte encrypts arr with rolling-key XOR using the byte b.
func XORByte(s []byte, b byte) []byte {
	n := len(s)
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = s[i] ^ b
	}
	return dst
}

// XORn computes the bitwise XOR of a and b, returning an error if
// their lengths don't match.
//
// 1.2 Fixed XOR
func XORn(a, b []byte) ([]byte, error) {
	la, lb := len(a), len(b)
	if la != lb {
		return nil, fmt.Errorf("Length mismatch: %d and %d", la, lb)
	}
	dst := make([]byte, la)
	for i := 0; i < la; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst, nil
}

// XORRepeat computes the
//
// 1.5 Implement repeating-key XOR
func XORRepeat(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	keyLength := len(key)
	for i, b := range plaintext {
		keyIndex := i % keyLength
		ciphertext[i] = b ^ key[keyIndex]
	}
	return ciphertext
}
