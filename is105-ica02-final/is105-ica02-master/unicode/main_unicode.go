package main

import (
	"fmt"

	unicode "github.com/TeamSolutionsUiA/ica02/is105-ica02-master/unicode/uniPack"
)

func main() {

	// Test av funksjonen Translate for Islandsk og Japansk oversettelse.
	input := "\xE2\x80\x9C\x6E\x6F\x72\x64\x20\x6F\x67\x20\x73\xC3\xB8\x72\xE2\x80\x9D\xE2\x80\x8B"
	outputIs := unicode.Translate(input, "is")
	outputJp := unicode.Translate(input, "jp")

	fmt.Println(outputIs)
	fmt.Println(outputJp)

	//Test av UnicodeCodePointDemo
	unicode.UnicodeCodePointDemo()
}
