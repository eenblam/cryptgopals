package analysis

// 1.6 Break repeating-key XOR

import (
	"fmt"
	"math"
)

// FindKeySize finds the best key size for the ciphertext
// by evaluating FirstNNormalized scores.
func BestFirstNNormalized(data []byte, minKeySize, maxKeySize int) (int, error) {
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
