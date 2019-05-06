package pipes

import (
	"encoding/hex"
)

func ReturnHex(s string) string {
	text := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(text)))
	hex.Encode(dst, text)
	return string(dst)
}