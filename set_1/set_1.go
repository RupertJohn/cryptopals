package set1

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/bits"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// HexToBase64 takes a hex string and converts it to base 64
func HexToBase64(hexString string) (string, error) {
	hexDecoded, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	encoded := b64.StdEncoding.EncodeToString([]byte(hexDecoded))
	return encoded, nil
}

// Xor will take two byte slices and XOR each byte against the other, looping over the key as needed
func Xor(stringBytes, keyBytes []byte) []byte {
	xoredSlice := make([]byte, len(stringBytes))

	for i := range stringBytes {
		j := i % len(keyBytes)
		xoredSlice[i] = stringBytes[i] ^ keyBytes[j]
	}

	return xoredSlice
}

// HexXor will taske two hex strings and xor them against eachother
func HexXor(hexString1, hexString2 string) string {
	hexDecoded1, _ := hex.DecodeString(hexString1)
	hexDecoded2, _ := hex.DecodeString(hexString2)
	xored := hex.EncodeToString(Xor(hexDecoded1, hexDecoded2))
	return xored
}

// FreqAnalysis will count the number of English letters in a phrase
func FreqAnalysis(decStr string) (score float64) {
	count := 0
	for _, v := range decStr {
		if strings.Contains(alphabet, string(v)) {
			count++
		}
	}

	score = float64(count) / float64(len(decStr))
	return score
}

// DecryptSingleByteCipher will xor a hex string against a single byte and determine most 'English' word
func DecryptSingleByteCipher(hexString string) (decStr string, score float64) {
	hexDecoded, _ := hex.DecodeString(hexString)

	for i := 1; i < 256; i++ {
		xored := ""
		for _, v := range hexDecoded {
			xored += string(v ^ byte(i))
		}

		xoredScore := FreqAnalysis(xored)

		if xoredScore > score {
			decStr, score = xored, xoredScore
		}
	}

	return decStr, score
}

// FindEncryptedLineInFile will open a file and return the most English decrypted string
func FindEncryptedLineInFile(filename string) (decStr string, score float64) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Challenge 4, Set 1: %v\n", err)
		return decStr, score
	}
	for _, line := range strings.Split(string(data), "\n") {
		dec, tempScore := DecryptSingleByteCipher(line)
		if tempScore > score {
			decStr, score = dec, tempScore
		}
	}
	return decStr, score
}

// XorString will use a repeating key to encrypt a string
func XorString(text string, key string) string {
	return hex.EncodeToString(Xor([]byte(text), []byte(key)))
}

// HammingDistance calculates the bit difference between two strings
func HammingDistance(string1 string, string2 string) int {
	xoredString := Xor([]byte(string1), []byte(string2))
	bitSum := 0
	for _, xoredValue := range xoredString {
		bitSum += bits.OnesCount(uint(xoredValue))
	}
	return bitSum
}
