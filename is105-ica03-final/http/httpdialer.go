package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
)

type Person1 struct {
	Name  string
	Email string
}

var person1 Person1

//Skriver JSON-strukturen til en txt. fil. Om filen allerede inneholder tekst, blir ny data lagt til på neste linje.
func writeToFile () {
	info := "\nNavn: " + person1.Name + " - Epost: " + person1.Email

	file, err := os.OpenFile("http.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Skriver slicen med bytes til "http.txt".
	_, err = file.Write([]byte(info))
	if err != nil {
		panic(err)
	}

}

//Kobler seg opp mot en server og dekoder JSON-strukturen som blir hentet fra serveren.
func main () {
	httpAddr2  := flag.String("http", "foo", "HTTP address")
	flag.Parse()

	resp, err := http.Get(*httpAddr2) //"http://localhost:5000"
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//Leser data fra server (JSON-struktur) og dekoder den.
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &person1); err != nil {
		panic(err)
	}

	writeToFile()
}
