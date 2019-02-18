package analysis

import (
	"fmt"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

// 1.7 - AES in CBC mode
func TestAESCBC(t *testing.T) {
	filename := util.DataPath("data_1_7.txt")
	_, err := encode.LoadBase64Rows(filename)
	if err != nil {
		t.Errorf("Couldn't load data: %s", err)
	}
}
