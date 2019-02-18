package analysis

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/eenblam/cryptgopals/encode"
	"github.com/eenblam/cryptgopals/util"
)

// 1.3
func TestBestXORByte(t *testing.T) {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedString := "Cooking MC's like a pound of bacon"
	expected := []byte(expectedString)
	data, _ := hex.DecodeString(s)
	result := BestXORByte(data)
	if !reflect.DeepEqual(result.PlainText, expected) {
		gotString := string(result.PlainText)
		t.Errorf("Expected %s, got %s", expectedString, gotString)
	}
}

// 1.4
func TestXORByteOfMany(t *testing.T) {
	filepath := util.DataPath("data_1_4.txt")
	rows, err := encode.LoadHexRows(filepath)
	if err != nil {
		t.Errorf("Could not load data from file %s: %s", filepath, err)
	}
	expected := []byte("Now that the party is jumping")
	got := BestXORByteOfMany(rows)
	if reflect.DeepEqual(expected, got.PlainText) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
