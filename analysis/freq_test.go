package analysis

import (
	"math"
	"testing"
)

func TestChiSquared(t *testing.T) {
	expected := 1291.103216
	tolerance := 0.01
	got := ChiSquared([]byte("Cooking MC's like a pound of bacon"))
	gotDiff := math.Abs(expected - got)
	if gotDiff > tolerance {
		t.Errorf("%f was not within %f of expected score %f",
			got, tolerance, expected)
	}
}
