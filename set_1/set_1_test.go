package set1

import "testing"

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
}
