package analysis

import (
	"fmt"

	"github.com/eenblam/cryptgopals/xor"
)

// BitSum sums the bits in byte b.
func BitSum(b byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(b & 1)
		b = b >> 1
	}
	return count
}

// BitSumBytes sums the bits in bytes bs.
func BitSumBytes(bs []byte) int {
	count := 0
	for _, b := range bs {
		count += BitSum(b)
	}
	return count
}

// Hamming computes the Hamming distance (edit distance) between a and b.
func Hamming(a, b []byte) (int, error) {
	// XORn handles length matching
	xord, err := xor.XORn(a, b)
	if err != nil {
		return 0, err
	}
	return BitSumBytes(xord), nil
}

// HammingString computes the edit distance between the byte representations
// of strings a and b.
func HammingString(a, b string) (int, error) {
	return Hamming([]byte(a), []byte(b))
}

// FirstNHamming computes the Hamming distances of s[:n] and s[n:2n].
func FirstNHamming(s []byte, n int) (int, error) {
	if 2*n > len(s) {
		msg := "Key size (%d) can't be more than half the size of the input array (%d)"
		return 0, fmt.Errorf(msg, n, len(s))
	}
	left, right := s[:n], s[n:2*n]
	return Hamming(left, right)
}
