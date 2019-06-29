package rand

import (
	"github.com/eenblam/cryptgopals/cipher"
)

type ECBKey []byte

func (k *ECBKey) Encrypt(plaintext []byte) ([]byte, error) {
	return cipher.ECBAESEncrypt([]byte(*k), plaintext)
}

func (k *ECBKey) Decrypt(ciphertext []byte) ([]byte, error) {
	return cipher.ECBAESDecrypt([]byte(*k), ciphertext)
}

func RandECBKey() ECBKey {
	return ECBKey(RandKey())
}
