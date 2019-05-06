package pipes

import (
	"encoding/hex"
	"fmt"
)

func HexReturn(s string) string {
	text := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(text)))
	hex.Encode(dst, text)
	return fmt.Sprintf("%s", dst)
}