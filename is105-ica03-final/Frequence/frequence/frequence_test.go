package frequence

import (
	"testing"
	
)

// https://golang.org/doc/effective_go.html#init


func benchmarkBRuneT100(b *testing.B) {
	benchmarkBRuneT(100, b)
}

func benchmarkBRuneT1000(b *testing.B) {
	benchmarkBRuneT(1000, b)
}

func benchmarkBRuneT10000(b *testing.B) {
	benchmarkBRuneT(10000, b)
}

func benchmarkBRuneT(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		b.StartTimer()
		RuneTeller()
	}
}

// Implementasjon av testfunksjoner for Quicksort algoritmen


func  benchLinjeTeller100(b *testing.B) {
	benchLinjeTeller(100, b)
}

func  benchLinjeTeller1000(b *testing.B) {
	benchLinjeTeller(1000, b)
}

func  benchLinjeTeller10000(b *testing.B) {
	benchLinjeTeller(10000, b)
}

func  benchLinjeTeller(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		b.StartTimer()
		LinjeTeller()
	}
}

