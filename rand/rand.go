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

// RandInt returns a random integer from [0, max).
//
// Keep max < 2^64.
func RandInt(max int) int {
	big_i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	// Maybe undefined if returned value is too large,
	// but that shouldn't happen since max is int64.
	return int(big_i.Int64())
}

// RandChunk generates an array of rand(5,10) random bytes.
//
// Panics if it can't read from getrandom or urandom.
func RandChunk() []byte {
	i := RandInt(6)
	bs := make([]byte, 5+i)
	_, err := rand.Read(bs)
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
