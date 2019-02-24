package cipher

import (
	"bytes"
	"crypto/aes"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

// 2.10 - Implement CBC mode
func Test_2_10(t *testing.T) {
	filename := util.DataPath("data_2_10.txt")
	ciphertext, err := encode.LoadBase64Rows(filename)
	if err != nil {
		t.Errorf("Couldn't load data: %s", err)
	}

	zeros := []byte{0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0}
	c, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	decrypter := NewCBCDecrypter(c, zeros)
	plaintext := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(plaintext, ciphertext)
	spoilers := []byte{73, 39, 109, 32, 98, 97, 99, 107, 32, 97, 110, 100, 32,
		73, 39, 109, 32, 114, 105, 110, 103, 105, 110, 39, 32, 116, 104, 101,
		32, 98, 101, 108, 108,
	}
	if !bytes.HasPrefix(plaintext, spoilers) {
		t.Error("2.10 failed.")
	}
}
