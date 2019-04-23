package main

import (
	"encoding/json"
	"net"
)

type Person struct {
	Name  string
	Email string
}

// Skriver en JSON-struktur til serveren og oppretter en goroutine.
func handler(c net.Conn) {
	var n = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",}

	// js = json
	js, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	c.Write([]byte(js))
	c.Close()
}

// Starter en tcp-server.
func main() {
	l, err := net.Listen("tcp", ":5000") // ":5000" må endres til din "ip:port". l = listen
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
