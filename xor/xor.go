package xor

import (
    "errors"
)

func XORn(dst []byte, a []byte, b []byte, n int) (error) {
    ldst := len(dst)
    la := len(a)
    lb := len(b)
    if (ldst != n) || (la != n) || (lb != n) {
        return errors.New("Length mismatch")
    }
    for i := 0; i < n; i++ {
        dst[i] = a[i] ^ b[i]
    }
    return nil
}
