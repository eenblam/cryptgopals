package analysis

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

func TestToBlocks(t *testing.T) {
	bs := []byte("YELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINE")
	expected := [][]byte{
		[]byte("YELLOW SUBMARINE"), []byte("YELLOW SUBMARINE"),
		[]byte("YELLOW SUBMARINE"),
	}
	got, err := ToBlocks(bs, 16)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(got) != len(expected) {
		t.Errorf("Expected %d blocks, got %d blocks", len(expected), len(got))
	}
	if !reflect.DeepEqual(expected, got) {
		t.Error("Received incorrect blocks")
	}
	// Error triggers when expected
	_, shouldBeErr := ToBlocks(bs, 15)
	if shouldBeErr == nil {
		t.Error("Expected error, but got none")
	}
}

func TestGetBlockCounts(t *testing.T) {
	bs := []byte("YELLOW SUBMARINEYELLOW SUBMARINE0123456789ABCDEFYELLOW SUBMARINE")
	blocks, err := ToBlocks(bs, 16)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	counts := GetBlockCounts(blocks)
	expected := map[string]int{
		"YELLOW SUBMARINE": 3,
		"0123456789ABCDEF": 1,
	}
	if !reflect.DeepEqual(expected, counts) {
		t.Error("Received incorrect block counts")
	}
}

func TestDetectFirstECB(t *testing.T) {
}

func TestDetectPossibleECBs(t *testing.T) {
}

func Test_1_8(t *testing.T) {
	expected := []byte{216, 128, 97, 151, 64, 168, 161, 155, 120, 64, 168, 163, 28, 129, 10, 61, 8, 100, 154, 247, 13, 192, 111, 79, 213, 210, 214, 156, 116, 76, 210, 131, 226, 221, 5, 47, 107, 100, 29, 191, 157, 17, 176, 52, 133, 66, 187, 87, 8, 100, 154, 247, 13, 192, 111, 79, 213, 210, 214, 156, 116, 76, 210, 131, 148, 117, 201, 223, 219, 193, 212, 101, 151, 148, 157, 156, 126, 130, 191, 90, 8, 100, 154, 247, 13, 192, 111, 79, 213, 210, 214, 156, 116, 76, 210, 131, 151, 169, 62, 171, 141, 106, 236, 213, 102, 72, 145, 84, 120, 154, 107, 3, 8, 100, 154, 247, 13, 192, 111, 79, 213, 210, 214, 156, 116, 76, 210, 131, 212, 3, 24, 12, 152, 200, 246, 219, 31, 42, 63, 156, 64, 64, 222, 176, 171, 81, 178, 153, 51, 242, 193, 35, 197, 131, 134, 176, 111, 186, 24, 106}
	filepath := util.DataPath("data_1_8.txt")
	data, err := encode.LoadHexRows(filepath)
	if err != nil {
		t.Errorf("Couldn't load hex rows from %s: %s", filepath, err)
	}
	got, err := DetectFirstECB(data, 16)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !bytes.Equal(got, expected) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
