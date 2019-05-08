package set1

import (
	b64 "encoding/base64"
	"encoding/hex"
	"unicode"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var letterScores = map[rune]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702, 'f': 0.02228,
	'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929, 'q': 0.00095, 'r': 0.05987,
	's': 0.06327, 't': 0.09056, 'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074, ' ': 0.2400,
}

// HexToBase64 takes a hex string and converts it to base 64
func HexToBase64(hexString string) (string, error) {
	hexDecoded, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	encoded := b64.StdEncoding.EncodeToString([]byte(hexDecoded))
	return encoded, nil
}

// xor will take two equal length byte arrays and xor the each byte against the other
func xor(array1, array2 []byte) []byte {
	xoredSlice := make([]byte, len(array1))

	for i := range array1 {
		j := i % len(array2)
		xoredSlice[i] = array1[i] ^ array2[j]
	}

	return xoredSlice
}

// HexXor will taske two hex strings and xor them against eachother
func HexXor(hexString1, hexString2 string) string {
	hexDecoded1, _ := hex.DecodeString(hexString1)
	hexDecoded2, _ := hex.DecodeString(hexString2)
	xored := hex.EncodeToString(xor(hexDecoded1, hexDecoded2))
	return xored
}

// ScoreWord will provide a score of the English-ness of a word
func ScoreWord(decStr string) float64 {
	score := 0.0
	for _, v := range decStr {
		score += letterScores[unicode.ToLower(v)]
	}
	return score
}

// DecryptSingleByteCipher will xor a hex string against a single byte and determine most 'English' word
func DecryptSingleByteCipher(hexString string) (decStr string, score float64) {
	hexDecoded, _ := hex.DecodeString(hexString)

	for _, letter := range alphabet {
		xored := ""
		for _, v := range hexDecoded {
			xored += string(v ^ byte(letter))
		}

		xoredScore := ScoreWord(xored)

		if xoredScore > score {
			decStr = xored
			score = xoredScore
		}
	}

	return decStr, score
}
