package rand

import (
	"crypto/rand"
	"math/big"
)

// RandKey returns a random 16 byte key.
//
// Panics if it can't read from getrandom or urandom.
//
// Going to just assume 16 bytes for now until I need something else.
func RandKey() []byte {
	bs := make([]byte, 16)
	_, err := rand.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs
}

// RandChunk generates an array of rand(5,10) random bytes.
//
// Panics if it can't read from getrandom or urandom.
func RandChunk() []byte {
	// 0-5; rand.Int is right-exclusive.
	big_i, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		panic(err)
	}
	// Shouldn't fail, since it's at most 5.
	i := big_i.Int64()
	bs := make([]byte, 5+i)
	_, err = rand.Read(bs)
	if err != nil {
		panic(err)
	}
	return bs
}

// RandPad appends random bytes to the front and back of provided bytes.
func RandPad(bs []byte) []byte {
	front := RandChunk()
	back := RandChunk()
	lf, lbs := len(front), len(bs)
	out := make([]byte, lf+lbs+lf)
	copy(out[:lf], front)
	copy(out[lf:lf+lbs], bs)
	copy(out[lf+lbs:], back)
	return out
}
