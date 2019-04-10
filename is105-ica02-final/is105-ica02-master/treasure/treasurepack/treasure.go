package treasure

import (
	"bytes"
	"fmt"
)

// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	treasureByte := []byte(treasureString)
	treasureByte = bytes.Replace(treasureByte, []byte("\xf8"), []byte("\xc3\xb8"), 1)
	treasureByte = bytes.Replace(treasureByte, []byte("\xe6"), []byte("\xc3\xa6"), 1)
	treasureByte = bytes.Replace(treasureByte, []byte("\xe5"), []byte("\xc3\xa5"), 1)
	fmt.Printf("%s", treasureByte)
}
