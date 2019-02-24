package cipher

// Cipher Block Chaining (CBC) mode.

import (
	"crypto/cipher"

	"github.com/eenblam/cryptgopals/xor"
)

// Underlying data structure for Encryptor and Decryptor
type cbc struct {
	b         cipher.Block
	blockSize int
	lastBlock []byte
}

func newCBC(b cipher.Block, iv []byte) *cbc {
	if b.BlockSize() != len(iv) {
		panic("cryptgopals/cipher: iv does not match block size")
	}
	out := &cbc{
		b:         b,
		blockSize: b.BlockSize(),
		lastBlock: make([]byte, b.BlockSize()),
	}
	copy(out.lastBlock, iv)
	return out
}

type cbcEncrypter cbc

// NewCBCEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return (*cbcEncrypter)(newCBC(b, iv))
}

func (x *cbcEncrypter) BlockSize() int { return x.blockSize }

func (x *cbcEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("cryptgopals/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cryptgopals/cipher: output smaller than input")
	}
	for len(src) > 0 {
		block := src[:x.blockSize]
		newBlock, err := xor.XORn(block, x.lastBlock)
		if err != nil {
			panic("cryptgopals/cipher: unexpected length mismatch")
		}
		x.b.Encrypt(dst, newBlock)
		copy(x.lastBlock, dst[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type cbcDecrypter cbc

// NewCBCDecrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return (*cbcDecrypter)(newCBC(b, iv))
}

func (x *cbcDecrypter) BlockSize() int { return x.blockSize }

func (x *cbcDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("cryptgopals/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("cryptgopals/cipher: output smaller than input")
	}
	preXOR := make([]byte, x.blockSize)
	for len(src) > 0 {
		x.b.Decrypt(preXOR, src[:x.blockSize])
		plaintext, err := xor.XORn(preXOR, x.lastBlock)
		if err != nil {
			panic("cryptgopals/cipher: unexpected length mismatch")
		}
		copy(dst, plaintext)
		copy(x.lastBlock, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
