package main

import(
	"fmt"
	"./fileutils"
)

func main() {

	byteSlice1 := fileutils.FileToByteslice("files/lang01.wl")
	byteSlice2 := fileutils.FileToByteslice("files/lang02.wl")
	byteSlice3 := fileutils.FileToByteslice("files/lang03.wl")

	byteSlice1Subset := byteSlice1[0:30]
	byteSlice2Subset := byteSlice2[0:30]
	byteSlice3Subset := byteSlice3[0:30]

	fmt.Printf("%c ", byteSlice1Subset)

	fmt.Printf("\n\n%c ", byteSlice2Subset)

	fmt.Printf("\n\n%c ", byteSlice3Subset)
}