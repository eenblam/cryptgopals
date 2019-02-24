package encode

import (
	"fmt"
)

// PadBytesTo pads bs to a length that is a multiple of blockSize.
//
// An error is returned for an invalid block size.
//
// This padding is done in accordance with RFC 5652 Section 6.3.
// https://tools.ietf.org/html/rfc5652#section-6.3
func PadBytesTo(bs []byte, blockSize int) ([]byte, error) {
	if blockSize >= 256 {
		return nil, fmt.Errorf("Block size greater than 256: %d", blockSize)
	}
	lbs := len(bs)
	var missing int
	remainder := lbs % blockSize
	if remainder == 0 {
		missing = blockSize
	} else {
		missing = blockSize - remainder
	}
	out := make([]byte, lbs+missing)
	copy(out, bs)
	for i := 0; i < missing; i++ {
		out[lbs+i] = byte(missing)
	}
	return out, nil
}

// UnPadBytesBy removes PKCS#7 padding from bs assuming blockSize.
//
// Errors are returned for malformed padding or an invalid block size.
//
// https://tools.ietf.org/html/rfc5652#section-6.3
func UnpadBytesBy(bs []byte, blockSize int) ([]byte, error) {
	if blockSize >= 256 {
		return nil, fmt.Errorf("Block size greater than 256: %d", blockSize)
	}
	lbs := len(bs)
	if lbs%blockSize != 0 {
		return nil, fmt.Errorf("Block size %d does not divide length of bytes %d",
			blockSize, lbs)
	}
	// Validate padding
	last := bs[lbs-1]
	last_i := int(last)
	for i := 1; i <= last_i; i++ {
		j := lbs - i
		if bs[j] != last {
			return nil, fmt.Errorf("%b != %b at index %d", bs[j], last, j)
		}
	}
	return bs[:lbs-last_i], nil
}
