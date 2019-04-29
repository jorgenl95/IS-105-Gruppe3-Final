package main

import iso "./isopack"

func main() {

	stringliteral := iso.GetExtASCIIStringLiteral()
	iso.IterateOverExtendedASCIIStringLiteral(stringliteral)
	iso.GreetingExtendedASCII()

}
