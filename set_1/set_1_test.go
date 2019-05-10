package set1

import (
	"fmt"
	"testing"
)

type hexPair struct {
	hexString string
	want      string
}

type hexXor struct {
	hexString1 string
	hexString2 string
	want       string
}

func TestSet1(t *testing.T) {
	hexCase := hexPair{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"}
	b64String, _ := HexToBase64(hexCase.hexString)
	if b64String != hexCase.want {
		t.Error(
			"For", hexCase.hexString,
			"expected", hexCase.want,
			"got", b64String,
		)
	}

	xor := hexXor{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"}
	xored := HexXor(xor.hexString1, xor.hexString2)
	if xored != xor.want {
		t.Error(
			"For", xor.hexString1, xor.hexString2,
			"expected", xor.want,
			"got", xored,
		)
	}

	score := ScoreWord("Hello There")
	if score != 1.04894 {
		t.Error(
			"For 'Hello There'",
			"expected", 1.04894,
			"got", score,
		)
	}

	decStr, score := DecryptSingleByteCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Printf("Decrypted: %v\nScore: %v\n", decStr, score)

	decStr, score = FindEncryptedLineInFile("4.txt")
	fmt.Printf("Decrypted: %v\nScore: %v\n", decStr, score)

	encStr := EncryptString("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE")
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if encStr != expected {
		t.Error(
			"Received: ", encStr,
			"Expected: ", expected,
		)
	}
}
