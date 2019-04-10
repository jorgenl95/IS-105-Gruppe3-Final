package main

import (
	"fmt"

	"./bfrequence"
	"./frequence"
)

func main() {

	fmt.Println("frequence test: ")

	frequence.LinjeTeller()
	frequence.RuneTeller()

	fmt.Println("bfrequence test: ")

	bfrequence.LinjeTeller()
	bfrequence.RuneTeller()
}
