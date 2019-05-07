package set1

import (
	b64 "encoding/base64"
	"encoding/hex"
)

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
