package main

import iso "github.com/TeamSolutionsUiA/ica02/is105-ica02-master/iso/isopack"

func main() {

	stringliteral := iso.GetExtASCIIStringLitral()
	iso.IterateOverExtendedASCIIStringLiteral(stringliteral)
	iso.GreetingExtendedASCII()

}
