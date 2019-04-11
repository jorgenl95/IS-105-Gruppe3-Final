package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"bytes"
	"os"
	//"io"
)

func main() {

	file1, err := os.Open("./files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Open("./files/text2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	size, err := file1.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size2, err := file2.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := size.Size()
	filesize2 := size2.Size()

	byteSlice1 := make([]byte, filesize)
	byteslice2 := make([]byte, filesize2)
/*
	_, err := io.ReadFull(file1, byteSlice1)
	if err != nil {
		log.Fatal(err)
	}

	_, err := io.ReadFull(file2, byteslice2)
	if err != nil {
		log.Fatal(err)
	}
*/
	content1, err := ioutil.ReadFile("./files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	content2, err := ioutil.ReadFile("./files/text2.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contents of text1: \n\n%s", content1)
	fmt.Println("\nThe size of text1.txt in bytes is: ", len(byteSlice1))
	fmt.Printf("\nContents of text1 as hex: \n\n% X", content1)
	fmt.Printf("\n\nContents of text2: \n\n%s", content2)
	fmt.Println("\nThe size of text2.txt in bytes is: ", len(byteslice2))
	fmt.Printf("\nContents of text2 as hex: \n\n% X", content2)

	if(bytes.Equal(byteSlice1, byteslice2)) {
		fmt.Println("\n\nThe content of the two files is equal.")
	} else {
		fmt.Println("\n\nThe content of the two files is NOT equal.")
	}
}
