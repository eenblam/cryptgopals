package encode

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func HexToBase64(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	_, writeErr := encoder.Write(decoded)
	if writeErr != nil {
		return "", writeErr
	}
	encoder.Close()
	return buf.String(), nil
}

// LoadHexRows reads a file of form <hexstring>\n<hexstring>\n...
// and returns the data as an array of byte arrays.
func LoadHexRows(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	out := make([][]byte, 0)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ascii := scanner.Text()
		bytes, err := hex.DecodeString(ascii)
		if err != nil {
			return nil, fmt.Errorf("Could not decode string %s: %s", ascii, err)
		}
		out = append(out, bytes)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Encountered scan error: %s", err)
	}
	return out, nil
}

// Read base64 rows from file, decoding them to bytes.
func LoadBase64Rows(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Returns `&decoder{enc: enc, r: &newlineFilteringReader{r}}`
	decoder := base64.NewDecoder(base64.StdEncoding, file)
	decodedBytes, readErr := ioutil.ReadAll(decoder)
	if readErr != nil {
		return nil, fmt.Errorf("Could not decode file as base64: %s", err)
	}
	return decodedBytes, nil
}
