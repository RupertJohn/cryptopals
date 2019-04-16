package set1

import (
  b64 "encoding/base64"
  "encoding/hex"
)

func HexToBase64(hexString string) (string, error) {
  hexDecoded, err := hex.DecodeString(hexString)
  if err != nil {
    return "", err
  }
  encoded := b64.StdEncoding.EncodeToString([]byte(hexDecoded))
  return encoded, nil
}

func Xor(hexString1, hexString2 string) string {
  hexDecoded1, _ := hex.DecodeString(hexString1)
  hexDecoded2, _ := hex.DecodeString(hexString2)
  xoredSlice := make([]byte, len(hexDecoded1))

  for i := range hexDecoded1 {
    xoredSlice[i] = hexDecoded1[i] ^ hexDecoded2[i]
  }

  xored := hex.EncodeToString(xoredSlice)

  return xored
}
