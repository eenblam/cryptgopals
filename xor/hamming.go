package xor

import (
	"fmt"
)

// Sum bits in byte b
func BitSum(b byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(b & 1)
		b = b >> 1
	}
	return count
}
