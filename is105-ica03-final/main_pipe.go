package main

import (
	"./pipes"
	"fmt"
)

func main() {
	testString := "This is a test string."
	encodedString := pipes.ReturnGZip(testString)
	fmt.Println(encodedString)
}