package main

import (
	"fmt"

	slice "github.com/stianbar/ica02/Oppgave-5/slice"
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
	fmt.Println("copyslice test: ")
	copySlice := slice.CopySlice(byteslice1)
	fmt.Println(&copySlice[50])

}
