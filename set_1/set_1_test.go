package set1

import "testing"

type hexPair struct {
  hexString     string
  want          string
}

var hexCases = []hexPair{
  {
    "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
    "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
  },
}

func TestSet1(t *testing.T) {
  for _, str := range hexCases {
    b64String, _ := HexToBase64(str.hexString)
    if b64String != str.want {
      t.Error(
        "For", str.hexString,
        "expected", str.want,
        "got", b64String,
      )
    }
  }

}
