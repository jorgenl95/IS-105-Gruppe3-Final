package pipes

import (
	"encoding/base64"
	"fmt"
)

func ReturnBase64(s string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(s))
	return fmt.Sprintf(b64)
}