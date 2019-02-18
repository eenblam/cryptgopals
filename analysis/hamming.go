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

// FirstNNormalized computes the Hamming distances of s[:n] and s[n:2n]
// then normalizes the result by dividing by n.
//
// Here, n is a candidate key size.
func FirstNNormalized(s []byte, n int) (float64, error) {
	if n < 2 {
		return 0, fmt.Errorf("You probably didn't mean to use only %d bytes.", n)
	}
	distance, err := FirstNHamming(s, n)
	if err != nil {
		return 0, err
	}
	return float64(distance) / float64(n), nil
}

//
func FourBlockHamming(s []byte, n int) (float64, error) {
	if n < 2 {
		return 0, fmt.Errorf("You probably didn't mean to use only %d bytes.", n)
	}
	if 4*n > len(s) {
		return 0, fmt.Errorf("Can't select %d bytes from input of size %d",
			4*n, len(s))
	}
	part1 := s[:n]
	part2 := s[n : 2*n]
	part3 := s[2*n : 3*n]
	part4 := s[3*n : 4*n]
	total := float64(0)
	// These are all the same length (n)
	total += UnsafeHamming(part1, part2)
	total += UnsafeHamming(part1, part3)
	total += UnsafeHamming(part1, part4)
	total += UnsafeHamming(part2, part3)
	total += UnsafeHamming(part2, part4)
	total += UnsafeHamming(part3, part4)
	return total / float64(6), nil
}

// UnsafeHamming returns the normalized hamming distance of a and b
// or just 0 if you messed up and used the wrong inputs.
//
// Born out of frustration in order to keep FourBlockHamming concise.
func UnsafeHamming(a, b []byte) float64 {
	distance, err := Hamming(a, b)
	if err != nil {
		// Probably different lengths
		return 0
	}
	return float64(distance) / float64(len(a))
}
