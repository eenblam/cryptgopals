package xor

import (
	"encoding/hex"
	"fmt"
	"math"
	"sort"

	"github.com/eenblam/cryptgopals/freq"
)

type XORResult struct {
	Byte      byte
	PlainText []byte
	// Chi-squared score
	Score float64
}

// Allow sorting of many scores
type ByScore []*XORResult

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

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

// BestXORByteOfMany computes all possible single byte XORs for each input,
// then returns the XORResult with the lowest Chi-squared score against
// our expected distribution of English characters.
//
// This could be broken up a bit to return the top n sorted results.
func BestXORByteOfMany(dataRows [][]byte) (*XORResult, error) {
	results := make(ByScore, 0)
	for _, row := range dataRows {
		result, err := BestXORByte(row)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	sort.Sort(results)
	return results[0], nil
}
