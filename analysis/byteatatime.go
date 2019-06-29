package analysis

import (
	"bytes"
	"fmt"

	"github.com/eenblam/cryptgopals/rand"
)

// ByteAtATimeECB solves 2.12.
//
// Right now, there's a weird bug when handling the end of the message.
//
// https://cryptopals.com/sets/2/challenges/12
func ByteAtATimeECB() {
	device := rand.NewRandECB()

	// Detect block size
	bs := []byte{}
	ciphertext, err := device.Encrypt(bs)
	if err != nil {
		panic("Ooooof!")
	}
	firstSize := len(ciphertext)
	var blockSize int
	for {
		bs = append(bs, 'A')
		ciphertext, err := device.Encrypt(bs)
		if err != nil {
			panic("Ooooof!")
		}
		if len(ciphertext) > firstSize {
			blockSize = len(ciphertext) - firstSize
			break
		}
	}
	numBlocks := firstSize / blockSize
	fmt.Printf("Secret string is %d bytes long, with %d blocks of size %d.\n",
		firstSize, numBlocks, blockSize)

	// Ensure ECB
	plaintext := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	ciphertext, err = device.Encrypt(plaintext)
	if err != nil {
		fmt.Printf("Oof! %s\n", err)
	}
	isECB, err := GuessECB(ciphertext, blockSize)
	if err != nil {
		fmt.Printf("Oof! %s\n", err)
	}
	fmt.Printf("Is ECB: %t\n", isECB)
	// Iterate over block size, building maps to find each byte in 256 steps.

	decrypted := make([]byte, 0)
	for currentBlock := 1; currentBlock <= numBlocks-1; currentBlock++ {
		currentIndex := (blockSize * currentBlock) - 1
		for i := 0; i < blockSize; i++ {
			numAs := blockSize - i - 1
			testBuffer, err := ThisManyAsInThisManyBytes(numAs, numAs)
			if err != nil {
				panic(err)
			}
			// Get value of target byte by placing at end of current block
			ciphertext, err = device.Encrypt(testBuffer)
			if err != nil {
				panic(err)
			}
			// Get current block
			expectedCiphertext := ciphertext[(currentBlock-1)*blockSize : currentBlock*blockSize]
			// Create buffer - A's until known, then known bytes, then test byte
			testBuffer, err = ThisManyAsInThisManyBytes(numAs, blockSize*currentBlock)
			if err != nil {
				panic(err)
			}
			// Write known to buffer
			for j, b := range decrypted {
				testBuffer[numAs+j] = b
			}
			// Brute force last byte
			for j := 0; j < 256; j++ {
				testBuffer[currentIndex] = byte(j)
				ciphertext, err = device.Encrypt(testBuffer)
				if err != nil {
					panic(err)
				}
				gotCiphertext := ciphertext[(currentBlock-1)*blockSize : currentBlock*blockSize]
				if bytes.Equal(expectedCiphertext, gotCiphertext) {
					decrypted = append(decrypted, byte(j))
					break
				}
				if j == 255 {
					fmt.Printf("Didn't solve byte %d of block #%d of %d. Result:\n", i, currentBlock, numBlocks)
					for i := 0; i < numBlocks; i++ {
						fmt.Println(string(decrypted[i*blockSize : (i+1)*blockSize]))
					}
					panic("Exiting...")
				}
			}
		}
		fmt.Printf("Completed block %d.\n", currentBlock)
	}
	fmt.Println("RESULT:")
	fmt.Println(string(decrypted))
}

func ThisManyAsInThisManyBytes(numAs, totalSize int) ([]byte, error) {
	out := make([]byte, totalSize)
	if numAs > totalSize {
		return nil, fmt.Errorf("Can't fit %d A's in %d bytes!", numAs, totalSize)
	}
	for i := range out {
		out[i] = 0
	}
	for i := 0; i < numAs; i++ {
		out[i] = 'A'
	}
	return out, nil
}
