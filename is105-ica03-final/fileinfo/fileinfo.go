package main

import (
	"fmt"
	"os"
	"flag"
	"log"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {

	filePtr := flag.String("f", "No File Declared", "Enter the file to be analyzed.")

	flag.Parse()

	fileInfo, err = os.Stat(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	fileSizeKibibytes := float64(fileSize)*0.00098
	fileSizeMebibytes := fileSizeKibibytes*0.00098
	fileSizeGibibytes := fileSizeMebibytes*0.00098

	fileMode := fileInfo.Mode()

	fmt.Println("Information about a file", *filePtr)
	fmt.Println("Size:", fileSize, "in bytes,", fileSizeKibibytes, "in kibibytes,", fileSizeMebibytes, "in mebibytes,", fileSizeGibibytes, "in gibibytes.")
	if fileInfo.IsDir() {
		fmt.Println("Is a directory.")
	} else {
		fmt.Println("Is not a directory.")
	}
	if fileMode.IsRegular() {
		fmt.Println("Is a regular file.")
	} else {
		fmt.Println("Is not a regular file.")
	}
	fmt.Println("Has Unix permission bits:", fileInfo.Mode())

}