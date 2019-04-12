package main

import iso "./isopack"

func main() {

	stringliteral := iso.GetExtASCIIStringLitral()
	iso.IterateOverExtendedASCIIStringLiteral(stringliteral)
	iso.GreetingExtendedASCII()

}
