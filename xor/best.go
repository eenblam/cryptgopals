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

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
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
		if !isASCII(string(xord)) {
			continue
		}
		score := freq.ChiSquared(xord)
		//fmt.Printf("%f, %s\n", score, xord)
		//if (score < minScore) && isASCII(string(xord)) {
		if score < minScore {
			minScore = score
			copy(bestXord, xord)
			bestByte = b
		}
		//fmt.Printf("%s, %s, %f\n", bestByte, bestXord, minScore)
	}
	result := &XORResult{bestByte, bestXord, minScore}
	return result, nil
}
