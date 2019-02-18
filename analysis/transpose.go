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
	// Bad assumption; this is how TransposeLoose came about.
	if (length % n) != 0 {
		return nil, fmt.Errorf("Key size %d doesn't divide data size %d",
			n, length)
	}
	blockSize := length / n
	// Initialize data
	out := make([][]byte, n)
	for i := 0; i < n; i++ {
		out[i] = make([]byte, blockSize)
	}
	// Run transposition
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
	numBlocks := len(bs)
	blockSize := len(bs[0])
	for _, block := range bs {
		if blockSize != len(block) {
			return nil, fmt.Errorf("Blocks have inconsistent sizes, e.g. %d and %d",
				blockSize, len(block))
		}
	}
	// Okay, everything is the same length.
	outSize := blockSize * numBlocks
	out := make([]byte, outSize)
	var block, index int
	for i := 0; i < outSize; i++ {
		block = i % numBlocks
		index = i / numBlocks
		out[i] = bs[block][index]
	}
	return out, nil
}

// TransposeLoose is like TransposeBlocks, but it doesn't care if the
// candidate key size divides the ciphertext length cleanly.
func TransposeLoose(s []byte, n int) ([][]byte, error) {
	// Initialize data
	out := make([][]byte, n)
	for i := 0; i < n; i++ {
		out[i] = make([]byte, 0)
	}
	// Run transposition
	var block int
	for i, b := range s {
		block = i % n
		out[block] = append(out[block], b)
	}
	return out, nil
}

// FlattenLoose undoes the transformation of TransposeLoose.
func FlattenLoose(bs [][]byte) ([]byte, error) {
	numBlocks := len(bs)
	blockSize := len(bs[0])
	startSize := blockSize
	seenDrop := false
	for _, block := range bs {
		diff := startSize - len(block)
		if diff < 0 {
			return nil, fmt.Errorf("Block sizes shouldn't grow")
		}
		if diff > 1 {
			return nil, fmt.Errorf("Block sizes aren't close (within 1)")
		}
		if diff == 0 && seenDrop {
			return nil, fmt.Errorf("Block sizes aren't consistently decreasing")
		}
		if diff == 1 && !seenDrop {
			seenDrop = true
		}
		// Everything else should be fine
	}
	outSize := 0
	for _, block := range bs {
		outSize += len(block)
	}
	out := make([]byte, outSize)
	var block, index int
	for i := 0; i < outSize; i++ {
		block = i % numBlocks
		index = i / numBlocks
		out[i] = bs[block][index]
	}
	return out, nil
}
