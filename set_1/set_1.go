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
