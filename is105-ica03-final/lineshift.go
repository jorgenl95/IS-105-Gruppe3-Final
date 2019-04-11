package main

import (
	"os"
	"github.com/jorgenl95/IS-105-Gruppe3-Final/is105-ica03-final/lineshift"
)

func main() {
	filename := os.Args[1]
	lineshift.SearchForLineshift(filename)
}