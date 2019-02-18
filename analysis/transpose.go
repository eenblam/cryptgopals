package analysis

import (
	"fmt"
)

// TransposeBlocks
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
