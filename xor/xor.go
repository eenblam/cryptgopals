package xor

import (
	"errors"
)

// XORByte encrypts arr into dst with rolling-key XOR using the byte b.
func XORByte(dst []byte, arr []byte, b byte, n int) error {
	if (len(dst) != n) || (len(arr) != n) {
		return errors.New("Length mismatch")
	}
	for i := 0; i < n; i++ {
		dst[i] = arr[i] ^ b
	}
	return nil
}

// XORn computes the bitwise XOR of a and b into dst.
//
// 1.2 Fixed XOR
func XORn(dst []byte, a []byte, b []byte, n int) error {
	ldst := len(dst)
	la := len(a)
	lb := len(b)
	if (ldst != n) || (la != n) || (lb != n) {
		return errors.New("Length mismatch")
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return nil
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
