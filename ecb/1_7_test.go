package ecb

import (
	"bytes"
	"crypto/aes"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

// 1.7 - AES in ECB mode
// Decrypt the file with AES-128-ECB and the provided key.
func Test_1_7(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	filename := util.DataPath("data_1_7.txt")
	data, err := encode.LoadBase64Rows(filename)
	if err != nil {
		t.Errorf("Couldn't load data: %s", err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		t.Errorf("Couldn't create cipher: %s", err)
	}

	plaintext := make([]byte, len(data))
	decrypter := NewECBDecrypter(c)
	decrypter.CryptBlocks(plaintext, data)
	spoilers := []byte{73, 39, 109, 32, 98, 97, 99, 107, 32, 97, 110, 100, 32,
		73, 39, 109, 32, 114, 105, 110, 103, 105, 110, 39, 32, 116, 104, 101,
		32, 98, 101, 108, 108,
	}
	if !bytes.HasPrefix(plaintext, spoilers) {
		t.Error("1.7 failed.")
	}
}
