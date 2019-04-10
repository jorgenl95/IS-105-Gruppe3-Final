// mainpakke for å teste pakken "slice".
// Oppgave 5 i ica02.

package main

import (
	"fmt"

	slice "github.com/TeamSolutionsUiA/ica02/is105-ica02-master/slice/slicepack"
)

func main() {
	byteslice1 := make([]byte, 51, 100)

	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteslice1[50])

	// For å teste copyslice funksjonen.
	fmt.Println("test av copyslice: ")
	copySlice := slice.CopySlice(byteslice1)
	fmt.Println(&copySlice[50])

}
