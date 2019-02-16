package freq

import (
	"math"
	"strings"
)

// Reference: http://www.practicalcryptography.com/cryptanalysis/text-characterisation/chi-squared-statistic/

// A, B, ..., Z
// Expected frequency distribution of English characters
// Using this one instead of WithSpace results in
// "Cooking MC's like a pound of bacon" and
// "cOOKINGmcSLIKEAPOUNDOFBACON       " having the same score.
// (obtained by xor'ing with x58 and x78, respectively.)
var englishFreqs = map[string]float64{
	"A": 0.08167, "B": 0.01492, "C": 0.02782, "D": 0.04253, "E": 0.12702,
	"F": 0.02228, "G": 0.02015, "H": 0.06094, "I": 0.06966, "J": 0.00153,
	"K": 0.00772, "L": 0.04025, "M": 0.02406, "N": 0.06749, "O": 0.07507,
	"P": 0.01929, "Q": 0.00095, "R": 0.05987, "S": 0.06327, "T": 0.09056,
	"U": 0.02758, "V": 0.00978, "W": 0.02360, "X": 0.00150, "Y": 0.01974,
	"Z": 0.00074,
}

// Ripped from @max-b
var englishFreqsWithSpace = map[string]float64{
	"A": 0.0651738, "B": 0.0124248, "C": 0.0217339, "D": 0.0349835,
	"E": 0.1041442, "F": 0.0197881, "G": 0.0158610, "H": 0.0492888,
	"I": 0.0558094, "J": 0.0009033, "K": 0.0050529, "L": 0.0331490,
	"M": 0.0202124, "N": 0.0564513, "O": 0.0596302, "P": 0.0137645,
	"Q": 0.0008606, "R": 0.0497563, "S": 0.0515760, "T": 0.0729357,
	"U": 0.0225134, "V": 0.0082903, "W": 0.0171272, "X": 0.0013692,
	"Y": 0.0145984, "Z": 0.0007836, " ": 0.1918182,
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alphabetWithSpace = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "

// ScoreEnglish returns the ChiSquared score for the given string
// against a more general distribution of English characters.
//
// Assumes text is ASCII English.
//
// 1.3
func ChiSquared(bytes []byte) float64 {
	length := float64(len(bytes))
	// Initialize counters
	counts := make(map[string]uint)
	for _, c := range alphabetWithSpace {
		counts[string(c)] = 0
	}
	// For each byte, coerce to char then upper to disregard case.
	// Otherwise, count it
	for _, b := range bytes {
		upper := strings.ToUpper(string(b))
		counts[upper]++
	}
	// Use counts and length to compute Chi-squared statistic
	// http://www.practicalcryptography.com/cryptanalysis/text-characterisation/chi-squared-statistic/
	// http://practicalcryptography.com/cryptanalysis/letter-frequencies-various-languages/english-letter-frequencies/
	chiSquared := float64(0.0)
	var expected float64
	for c, count := range counts {
		// Penalize punctuation and unprintables.
		relativeExpected, found := englishFreqsWithSpace[c]
		if found {
			expected = length * relativeExpected
		} else {
			// Thanks @max-b!
			expected = 0.0008
		}
		//TODO It would be good to some testing here
		if expected > 0.0 {
			chiSquared += chiTerm(float64(count), expected)
		}
	}
	return chiSquared
}

// Compute one term in the Chi-squared summation
func chiTerm(observed, expected float64) float64 {
	return math.Pow(observed-expected, 2.0) / expected
}
