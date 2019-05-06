package main

import(
	"fmt"
	"io/ioutil"
	"./pipes"
	"os"
	"log"
)

func main() {
	args := os.Args[1]

	file, err := ioutil.ReadFile(args)
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(file)

	hex := pipes.ReturnHex(fileString)
	base64 := pipes.ReturnBase64(fileString)
	gzip := pipes.ReturnGZip(fileString)

	hexLimited := hex[:20]
	base64Limited := base64[:20]
	gzipLimited := gzip[:20]

	hexLength := len(hex)
	base64Length := len(base64)
	gzipLength := len(gzip)

	fmt.Println("First 20 bytes of hex encoded output: ", hexLimited)
	fmt.Println("\nLength of hex encoded output: ", hexLength)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("First 20 bytes of base64 encoded output: ", base64Limited)
	fmt.Println("\nLength of base64 encoded output: ", base64Length)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("First 20 bytes of gzip encoded output: ", gzipLimited)
	fmt.Println("\nLength of gzip encoded output: ", gzipLength)
}
