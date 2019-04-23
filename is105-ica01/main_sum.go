package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/jorgenl95/is105-ica01/sum"
)

func main() {

	argsNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error occured")
	}

	argsNumber2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("Error occured")
	}

	var tempVar = sum.SumInt32(int32(argsNumber), int32(argsNumber2))

	fmt.Println(tempVar)

}