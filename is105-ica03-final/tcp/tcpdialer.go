package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net"
	"os"
)

type Person1 struct {
	Name  string
	Email string
}

var person1 Person1

// Skriver JSON-strukturen til en txt. fil. Om filen allerede inneholder tekst, blir ny data lagt til på neste linje.
func WriteToFile () {
	info := "\nNavn: " + person1.Name + " - Epost: " + person1.Email

	file, err := os.OpenFile("tcp.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Skriver slicen med bytes til "tcp.txt".
	_, err = file.Write([]byte(info))
	if err != nil {
		panic(err)
	}
}

// Kobler seg opp mot en server og dekoder JSON-strukturen som blir hentet fra serveren.
func main () {
	tcpAddr1 := flag.String("tcp", "foo", "a string")
	flag.Parse()

	conn, err := net.Dial("tcp", *tcpAddr1) //":5000"
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// Leser data fra server (JSON-struktur) og dekoder den. bs = bytes
	bs, _ := ioutil.ReadAll(conn)
	if err := json.Unmarshal(bs, &person1); err != nil {
		panic(err)
	}

	WriteToFile()
}
