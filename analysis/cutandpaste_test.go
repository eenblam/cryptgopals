package analysis

import (
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/rand"
)

func TestCutAndPaste(t *testing.T) {
	k := rand.RandECBKey()
	// What does our admin block look like?
	// 0123456789ABCDEF
	// email=foo@bar.io
	// adminXXXXXXXXXXX (where X is byte(11)) <-- Read the second block
	// &uid=10&role=use
	// rYYYYYYYYYYYYYYY (where Y is just the PKCS#7 padding)
	eleven := string([]byte{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11})
	email := "foo@bar.ioadmin" + eleven
	plaintext := encode.ProfileFor(email).Cookie()
	ciphertext, err := k.Encrypt([]byte(plaintext))
	if err != nil {
		t.Error(err)
	}
	// This is our payload
	ourBlock := ciphertext[16:32]

	// Now, isolate the data we want to replace into its own block:
	// 0123456789ABCDEF
	// email=AAAfoo@bar <-- Pad email a bit to align
	// .io&uid=10&role=
	// userYYYYYYYYYYYY (Y is just padding) <-- Replace this block with "ourBlock"
	email = "AAAfoo@bar.io"
	plaintext = encode.ProfileFor(email).Cookie()
	ciphertext, err = k.Encrypt([]byte(plaintext))
	if err != nil {
		t.Error(err)
	}

	// We know that:
	// 1. First two blocks of ciphertext will create an email
	// 2. Last block is only role + padding
	// 3. ourBlock will decrypt with the same key
	// So we just replace the last block with ourBlock
	for i := 0; i < 16; i++ {
		ciphertext[32+i] = ourBlock[i]
	}
	// Now we're done - this ciphertext will decrypt to produce the desired profile.
	decrypted, err := k.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	// Test
	cookie := string(decrypted)
	profile, err := encode.ParseCookie(cookie).GetProfile()
	if err != nil {
		t.Error(err)
	}
	if profile.Role != "admin" {
		t.Errorf("Expected role admin, got %s\n", profile.Role)
	}
}
