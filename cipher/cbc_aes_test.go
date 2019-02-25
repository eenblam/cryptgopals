package cipher

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestCBCAES(t *testing.T) {
	plaintext := []byte("0123456789abcdefFEDCBA9876543210")
	zeros := []byte{0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0}
	c, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	encrypter := NewCBCEncrypter(c, zeros)
	ciphertext := make([]byte, len(plaintext))
	encrypter.CryptBlocks(ciphertext, plaintext)

	c2, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	decrypter := NewCBCDecrypter(c2, zeros)
	got := make([]byte, len(plaintext))
	decrypter.CryptBlocks(got, ciphertext)

	if !bytes.Equal(plaintext, got) {
		t.Errorf("Expected %s, got %s", plaintext, got)
	}
}

func TestCBCAES_EncryptDecrypt(t *testing.T) {
	plaintext := []byte("0123456789abcdefFEDCBA9876543210")
	iv := []byte("abcdefghijklmnop")
	key := []byte("YELLOW SUBMARINE")
	ciphertext, err := CBCAESEncrypt(key, iv, plaintext)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	got, err := CBCAESDecrypt(key, iv, ciphertext)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(plaintext, got) {
		t.Errorf("Expected %s, got %s", string(plaintext), string(got))
	}
}
