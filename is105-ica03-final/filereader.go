package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"bytes"
	"os"
	"io"
)

func main() {

	file1, err3 := os.Open("./files/text1.txt")
	if err3 != nil {
		log.Fatal(err3)
	}
	defer file1.Close()

	file2, err4 := os.Open("./files/text2.txt")
	if err4 != nil {
		log.Fatal(err4)
	}
	defer file2.Close()

	size, sizeerr := file1.Stat()
	if sizeerr != nil {
		log.Fatal(sizeerr)
	}

	size2, size2err := file2.Stat()
	if size2err != nil {
		log.Fatal(size2err)
	}

	filesize := size.Size()
	filesize2 := size2.Size()

	byteSlice1 := make([]byte, filesize)
	byteslice2 := make([]byte, filesize2)

	_, err5 := io.ReadFull(file1, byteSlice1)
	if err5 != nil {
		log.Fatal(err5)
	}

	_, err6 := io.ReadFull(file2, byteslice2)
	if err6 != nil {
		log.Fatal(err6)
	}

	content1, err := ioutil.ReadFile("./files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	content2, err2 := ioutil.ReadFile("./files/text2.txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("Contents of text1: \n\n%s", content1)
	fmt.Println("\nThe size of text1.txt in bytes is: ", len(byteSlice1))
	fmt.Printf("Contents of text1 as hex: \n\n% X", content1)
	fmt.Printf("\n\nContents of text2: \n\n%s", content2)
	fmt.Println("\nThe size of text2.txt in bytes is: ", len(byteslice2))
	fmt.Printf("\nContents of text2 as hex: \n\n% X", content2)

	if(bytes.Equal(byteSlice1, byteslice2)) {
		fmt.Println("\n\nThe content of the two files is equal.")
	} else {
		fmt.Println("\n\nThe content of the two files is NOT equal.")
	}
}
