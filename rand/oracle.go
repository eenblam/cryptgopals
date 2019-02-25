package rand

// Encryption "oracle" described by https://cryptopals.com/sets/2/challenges/11

import (
	"fmt"

	"github.com/eenblam/cryptgopals/cipher"
)

// EncryptionOracle encrypts the provided plaintext under a random key
// after padding it at front and back with a random number of random bits.
//
// AES-128 is used, but the mode varies. 50% of the time, ECB is used.
// The rest of the time, CBC is used with a random 16 byte IV.
// Following the random byte padding, the plaintext is padded via PKCS#7.
//
// analysis.GuessECB consistently guesses which mode is used by
// EncryptionOracle, solving challenge 2.11: An ECB/CBC detection oracle.
func EncryptionOracle(plaintext []byte) ([]byte, error) {
	key := RandKey()
	paddedPlaintext := RandPad(plaintext)
	if RandInt(2) > 0 {
		fmt.Println("SPOILER: ECB")
		return cipher.ECBAESEncrypt(key, paddedPlaintext)
	} else {
		fmt.Println("SPOILER: CBC")
		iv := RandKey()
		return cipher.CBCAESEncrypt(key, iv, paddedPlaintext)
	}
}
