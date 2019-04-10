package unicode

import (
	"fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	transJp := "\x20\x70\xC3\xA5\x20\x6A\x61\x70\x61\x6E\x73\x6B\x20\x65\x72\x20\xE2\x80\x8B" +
		"\xE2\x80\x9E\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\xE2\x80\x9D"
	transIs := "\x20\x70\xC3\xA5\x20\x69\x73\x6C\x61\x6E\x64\x73\x6B\x20\x65\x72\x20\xE2\x80\x8B" +
		"\xE2\x80\x9C\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\xE2\x80\x9D"

	returnString := expression

	if language == "jp" {
		returnString += transJp
	} else if language == "is" {
		returnString += transIs

	}
	returnString = fmt.Sprintf("%s", returnString)
	return returnString
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {

	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
