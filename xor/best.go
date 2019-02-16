package xor

import (
	"encoding/hex"
	"fmt"
	"math"

	"github.com/eenblam/cryptgopals/freq"
)

type XORResult struct {
	Byte      byte
	PlainText []byte
	// Chi-squared score
	Score float64
}

func (r *XORResult) String() string {
	hexByte := make([]byte, 1)
	hexByte[0] = r.Byte
	hexString := hex.EncodeToString(hexByte)
	return fmt.Sprintf("Byte (hex): %s\nPlainText: %s\nScore: %f",
		hexString, string(r.PlainText), r.Score)
}

// isASCII returns false iff a byte has its highest bit set.
func isASCII(bytes []byte) bool {
	for _, b := range bytes {
		if b^128 == 128 {
			return false
		}
	}
	return true
}

// BestXORByte computes every possible single-byte XOR of the input bytes,
// returning the result with lowest ChiSquared statistic.
func BestXORByte(bytes []byte) (*XORResult, error) {
	length := len(bytes)
	xord := make([]byte, length)
	minScore := math.Inf(1)
	bestXord := make([]byte, length)
	bestByte := byte(0)
	for i := 0; i < 256; i++ {
		b := byte(i)
		err := XORByte(xord, bytes, b, length)
		if err != nil {
			return nil, err
		}
		if !isASCII(xord) {
			continue
		}
		score := freq.ChiSquared(xord)
		if score < minScore {
			minScore = score
			copy(bestXord, xord)
			bestByte = b
		}
	}
	result := &XORResult{bestByte, bestXord, minScore}
	return result, nil
}
