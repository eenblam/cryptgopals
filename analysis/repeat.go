package analysis

// 1.6 Break repeating-key XOR

import (
	"fmt"
	"math"
)

// FindKeySize finds the best key size for the ciphertext
// by evaluating FirstNNormalized scores.
func FindKeySize(data []byte, minKeySize, maxKeySize int) (int, error) {
	bestKeySize := minKeySize
	bestDistance := math.Inf(1)
	for i := minKeySize; i <= maxKeySize; i++ {
		distance, err := FirstNNormalized(data, i)
		if err != nil {
			return 0, fmt.Errorf("Couldn't get key size: %s", err)
		}
		if distance < bestDistance {
			bestKeySize = i
			bestDistance = distance
		}
	}
	return bestKeySize, nil
}

// FindKeySizeFourBlock works like FindKeySize, only it uses four blocks
// of the candidate key size, averaging the Hamming distances between them.
func FindKeySizeFourBlock(data []byte, minKeySize, maxKeySize int) (int, error) {
	bestKeySize := minKeySize
	bestDistance := math.Inf(1)
	for i := minKeySize; i <= maxKeySize; i++ {
		distance, err := FourBlockHamming(data, i)
		if err != nil {
			return 0, fmt.Errorf("Couldn't get key size: %s", err)
		}
		if distance < bestDistance {
			bestKeySize = i
			bestDistance = distance
		}
	}
	return bestKeySize, nil
}

// BreakRepeatingKeyXOR solves 1.6.
//
// Initially attempted with FindKeySize, TransposeBlocks, and FlattenBlocks.
// Had to switch, but maybe I'll need to write a version with these later.
func BreakRepeatingKeyXOR(data []byte) ([]byte, error) {
	//bestKeySize, err := FindKeySize(data, 2, 40)
	bestKeySize, err := FindKeySizeFourBlock(data, 2, 40)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get best key size: %s", err)
	}

	//blocks, err := TransposeBlocks(data, bestKeySize)
	blocks, err := TransposeLoose(data, bestKeySize)
	if err != nil {
		return nil, fmt.Errorf("Couldn't transpose data: %s", err)
	}
	results := make([][]byte, bestKeySize)
	// Solve each transposed block
	for i, block := range blocks {
		result := BestXORByte(block)
		results[i] = result.PlainText
	}

	//flat, err := FlattenBlocks(results)
	flat, err := FlattenLoose(results)
	if err != nil {
		return nil, fmt.Errorf("Couldn't flatten blocks: %s", err)
	}
	return flat, nil
}
