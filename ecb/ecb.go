package ecb

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
	"crypto/cipher"
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
