package main

import (
	"./ascii"
)

func main() {

	stringliteral := ascii.GetASCIIStringLiteral()
	ascii.IterateOverASCIIStringLiteral(stringliteral)
	ascii.GreetingASCII()
}
