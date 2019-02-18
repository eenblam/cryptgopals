package analysis

import (
	"bytes"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

// 1.6 - Break repeating-key XOR
func TestBreakRepeatingKeyXOR(t *testing.T) {
	filename := util.DataPath("data_1_6.txt")
	data, err := encode.LoadBase64Rows(filename)
	if err != nil {
		t.Errorf("Couldn't load data: %s", err)
	}
	result, err := BreakRepeatingKeyXOR(data)
	if err != nil {
		t.Errorf("Couldn't solve 1.6: %s", err)
	}

	spoilers := []byte{73, 39, 109, 32, 98, 97, 99, 107, 32, 97, 110, 100, 32,
		73, 39, 109, 32, 114, 105, 110, 103, 105, 110, 39, 32, 116, 104, 101,
		32, 98, 101, 108, 108,
	}
	if !bytes.HasPrefix(result, spoilers) {
		t.Error("1.6 failed.")
	}
}
