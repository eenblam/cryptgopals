package cipher

// Adapted from https://codereview.appspot.com/7860047/patch/23001/24001
// which proposed adding ECB to the Go standard library.

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Electronic Code Book (ECB) mode.

// ECB provides confidentiality by assigning a fixed ciphertext block to each
// plaintext block.

// See NIST SP 800-38A, pp 08-09

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/eenblam/cryptgopals/encode"
)

// Underlying data structure for Encryptor and Decryptor
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	//TODO Is panic actually part of the usual API?
	if len(src)%x.blockSize != 0 {
		panic("cryptgopals/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cryptgopals/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		//TODO I assume this just increments pointers into the slice,
		// but would an index not be at least as efficient?
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	//TODO Is panic actually part of the usual API?
	if len(src)%x.blockSize != 0 {
		panic("cryptgopals/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cryptgopals/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		//TODO I assume this just increments pointers into the slice,
		// but would an index not be at least as efficient?
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ECBAESEncrypt encrypts the plaintext (padded with PKCS#7) with AES
// using the provided key.
func ECBAESEncrypt(key, plaintext []byte) ([]byte, error) {
	//TODO require 16 bytes?
	blockSize := len(key)
	// Pad with PKCS 7
	paddedPlaintext, padErr := encode.PadBytesTo(plaintext, blockSize)
	if padErr != nil {
		return nil, padErr
	}
	// Get cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Get blockmode & encrypt
	encrypter := NewECBEncrypter(c)
	ciphertext := make([]byte, len(paddedPlaintext))
	encrypter.CryptBlocks(ciphertext, paddedPlaintext)
	return ciphertext, nil
}

// ECBAESDecrypt decrypts the ciphertext (assuming PKCS#7 padding)
// with AES using the provided key.
func ECBAESDecrypt(key, ciphertext []byte) ([]byte, error) {
	blockSize := len(key)
	if len(ciphertext)%blockSize != 0 {
		return nil, fmt.Errorf("Key size %d does not divide ciphertext length %d",
			blockSize, len(ciphertext))
	}
	// Get cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Get blockmode & decrypt
	decrypter := NewECBDecrypter(c)
	paddedPlaintext := make([]byte, len(ciphertext))
	decrypter.CryptBlocks(paddedPlaintext, ciphertext)
	// Unad with PKCS 7
	plaintext, padErr := encode.UnpadBytesBy(paddedPlaintext, blockSize)
	if padErr != nil {
		return nil, padErr
	}
	return plaintext, nil
}
