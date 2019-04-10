package bfrequence

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const textFil = "../files/pg100.txt"

func LinjeTeller() {

	// Åpner og leser fra fil.¢
	linjeTeller := 1

	fil, err := os.Open(textFil)
	if err != nil {
		log.Fatal(err)
	}
	//buffer := bufio.NewReader(fil)
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		linjeTeller++

	}
	fmt.Println("Filen inneholder", linjeTeller, "linjer")
}

func RuneTeller() {

	// Hente inn tekstfilen som runer i en array-buffer.
	fil, err := os.Open(textFil)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fil)
	scanner.Split(bufio.ScanRunes)

	m := make(map[string]int)
	for scanner.Scan() {

		r := scanner.Text()

		m[r]++
	}
	// answer is now in m.  sort and format output:
	lfs := make(lfList, 0, len(m))
	for l, f := range m {
		lfs = append(lfs, &letterFreq{l, f})
	}
	sort.Sort(lfs)
	fmt.Println("Presentasjon av de 5 mest brukte runene i teksten:")
	teller := 0
	for _, lf := range lfs {
		if teller >= 5 {
			break
		}

		fmt.Printf("   %+v    %7d\n", lf.string, lf.freq)
		teller++
	}

}

type letterFreq struct {
	string
	freq int
}
type lfList []*letterFreq

func (lfs lfList) Len() int { return len(lfs) }

func (lfs lfList) Less(i, j int) bool {
	switch fd := lfs[i].freq - lfs[j].freq; {
	case fd < 0:
		return false
	case fd > 0:
		return true
	}
	return lfs[i].string < lfs[j].string
}
func (lfs lfList) Swap(i, j int) {
	lfs[i], lfs[j] = lfs[j], lfs[i]
}
