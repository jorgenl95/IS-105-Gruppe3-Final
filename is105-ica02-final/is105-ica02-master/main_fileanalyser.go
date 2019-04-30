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

	fmt.Println("Contents of lang01.wl: ")
	fmt.Printf("%c \n", byteSlice1Subset)

	fmt.Println("Contents of lang02.wl: ")
	fmt.Printf("%c \n", byteSlice2Subset)

	fmt.Println("Contents of lang03.wl: ")
	fmt.Printf("%c \n", byteSlice3Subset)
}