package iso

import "fmt"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

const extAscii = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f" +
	"\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf" +
	"\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf" +
	"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
	"\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf" +
	"\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef" +
	"\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"

// Return constant extAscii.
func GetExtASCIIStringLitral() string {

	return extAscii

}

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data
func IterateOverExtendedASCIIStringLiteral(extAscii string) {
	// Kode for Oppgave 2a
	for i := 0; i < len(extAscii); i++ {
		fmt.Printf("%X  ", extAscii[i])
		fmt.Printf("%q  ", extAscii[i])
		fmt.Printf("%b  \n", extAscii[i])
	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	byteArray := []byte{'\x22', '\x53', '\x61', '\x6C', '\x75', '\x74', '\x2C', '\x20', '\xE7', '\x61', '\x20', '\x76',
		'\x61', '\x20', '\xB0', '\x2D', '\x29', '\x20', '\xC7', '\x61', '\x20', '\x63', '\x6F', '\x75', '\x74', '\x65',
		'\x20', '\x80', '\x35', '\x30', '\x22', '\x0A'}
	greetingString := ""
	for i := 0; i < len(byteArray); i++ {
		greetingString += fmt.Sprintf("%c", byteArray[i])
	}
	fmt.Println(greetingString)
	return greetingString
}
