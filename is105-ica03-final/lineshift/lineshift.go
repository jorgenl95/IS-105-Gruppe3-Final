package lineshift

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func SearchForLineshift(fileToSearch string) {

	content, err := ioutil.ReadFile(fileToSearch)
	if err != nil {
		log.Fatal(err)
	}

	contentAsString := fmt.Sprintf("% X", content)

	//Removing 0D in text file to test the if statement below
	//contentAsStringTrimmed := strings.ReplaceAll(contentAsString, "0D ", "")

	if strings.Contains(contentAsString, "0D 0A") {
		fmt.Println("The file contains CRLF and is formatted for MS-DOS.")
	} else if (strings.Contains(contentAsString, "0A") && !strings.Contains(contentAsString, "0D")) {
		fmt.Println("The file contains LF and is formatted for POSIX.")
	} else {
		fmt.Println("The file does not contain LineShift.")
	}
}