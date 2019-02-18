package analysis

import (
	"fmt"
)

// TransposeBlocks splits input data into n chunks, where the i-th
// chunk contains elements (in-order) whose index is congruent to i
// modulo n.
//
// Here, n is a candidate key size.
func TransposeBlocks(s []byte, n int) ([][]byte, error) {
	length := len(s)
	// Perhaps this is a bad assumption... assuming "yes" because padding?
	// Probably worth trying both ways.
	if (length % n) != 0 {
		return nil, fmt.Errorf("Key size %d doesn't divide data size %d",
			n, length)
	}
	blockSize := length / n
	out := make([][]byte, n)
	for i := 0; i < n; i++ {
		out[i] = make([]byte, blockSize)
	}
	var block, index int
	for i, b := range s {
		block = i % blockSize
		index = i / blockSize
		out[block][index] = b
	}
	return out, nil
}

// FlattenBlocks undoes the transformation of TransposeBlocks.
func FlattenBlocks(bs [][]byte) ([]byte, error) {
	blockSize := len(bs)
	chunkSize := len(bs[0])
	for _, chunk := range bs {
		if chunkSize != len(chunk) {
			return nil, fmt.Errorf("Blocks have inconsistent sizes, e.g. %d and %d",
				chunkSize, len(chunk))
		}
	}
	// Okay, everything is the same length.
	outSize := chunkSize * blockSize
	out := make([]byte, outSize)
	var block, index int
	for i := 0; i < outSize; i++ {
		block = i % blockSize
		index = i / blockSize
		out[i] = bs[block][index]
	}
	return out, nil
}
