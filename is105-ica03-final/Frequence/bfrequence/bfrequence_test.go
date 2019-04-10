package bfrequence

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


func  benchBufLinjeTeller100(b *testing.B) {
	benchBufLinjeTeller(100, b)
}

func  benchBufLinjeTeller1000(b *testing.B) {
	benchBufLinjeTeller(1000, b)
}

func  benchBufLinjeTeller10000(b *testing.B) {
	benchBufLinjeTeller(10000, b)
}

func  benchBufLinjeTeller(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		b.StartTimer()
		LinjeTeller()
	}
}

