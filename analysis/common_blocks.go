package analysis

import (
	"fmt"
)

// ToBlocks computes a slice of blocks of size blockSize
// from a slice of bytes.
//
// An error is returned if blockSize does not divide the length
// of the input bytes,
func ToBlocks(bs []byte, blockSize int) ([][]byte, error) {
	numBytes := len(bs)
	if numBytes%blockSize != 0 {
		return nil, fmt.Errorf("Found %d bytes; not divisible by block size %d. Padding needing.",
			numBytes, blockSize)
	}
	numBlocks := numBytes / blockSize
	blocks := make([][]byte, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blocks[i] = bs[i*blockSize : (i+1)*blockSize]
	}
	return blocks, nil
}

// GetBlockCounts returns a map of counts for each block in the provided
// block array.
//
// Because Golang refuses to hash a slice of byte, they're first coerced to string.
func GetBlockCounts(blocks [][]byte) map[string]int {
	counts := make(map[string]int)
	for _, block := range blocks {
		s := string(block)
		_, found := counts[s]
		if found {
			counts[s]++
		} else {
			counts[s] = 1
		}
	}
	return counts
}

// HasRedundantBlocks returns true as soon as a block is found to have occurred
// more than once.
func HasRedundantBlocks(counts map[string]int) bool {
	for _, v := range counts {
		if v > 1 {
			return true
		}
	}
	return false
}

// GuessECB returns true if the provided bytes contain redundant blocks.
//
// Consistently guesses the block cipher mode used by rand.EncryptionOracle,
// solving challenge 2.11: An ECB/CBC detection oracle.
func GuessECB(bs []byte, blockSize int) (bool, error) {
	blocks, err := ToBlocks(bs, blockSize)
	if err != nil {
		return false, err
	}
	counts := GetBlockCounts(blocks)
	if HasRedundantBlocks(counts) {
		return true, nil
	}
	return false, nil
}

// DetectFirstECB searches for the first ciphertext with redundant blocks
// of blockSize, returning it if found and an error otherwise.
func DetectFirstECB(lines [][]byte, blockSize int) ([]byte, error) {
	for _, line := range lines {
		isECB, err := GuessECB(line, blockSize)
		if err != nil {
			return nil, err
		}
		if isECB {
			return line, nil
		}
	}
	return nil, fmt.Errorf("No ciphertext with redundant blocks of size %d",
		blockSize)
}

// DetectPossibleECBs searches the input ciphertexts for redundant
// blockSize-byte blocks.
//
// Any ciphertexts with redundant blocks are returned
// as a map of <ciphertext>:<number redundant blocks>.
func DetectPossibleECBs(ciphertexts [][]byte, blockSize int) (map[string]int, error) {
	results := make(map[string]int)
	for _, ciphertext := range ciphertexts {
		blocks, err := ToBlocks(ciphertext, blockSize)
		if err != nil {
			return nil, err
		}
		counts := GetBlockCounts(blocks)
		nonUnique := 0
		for _, v := range counts {
			if v > 1 {
				nonUnique += v
			}
		}
		if nonUnique > 0 {
			results[string(ciphertext)] = nonUnique
		}
	}
	return results, nil
}
