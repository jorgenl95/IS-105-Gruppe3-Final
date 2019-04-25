package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string
	Email string
}

// Skriver en JSON-struktur til serveren.
func handler (w http.ResponseWriter, r *http.Request) {
	var n1 = &Person{Name: "Markus",
		Email: "m.sveggen@gmail.com",}

	// js = json
	js, err := json.Marshal(n1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("JSON-struct received")
	w.Write(js)

}

// Starter en HTTP-server.
func main () {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:5000", nil) // "localhost:5000" må endres til "ip:port".
	if err != nil {
		panic(err)
	}
}

