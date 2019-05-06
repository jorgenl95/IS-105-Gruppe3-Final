package pipes

import (
	"bytes"
	"compress/gzip"
	"log"
	"fmt"
)

func ReturnGZip(s string) string {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	err = zw.Close()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintln(buf)
}