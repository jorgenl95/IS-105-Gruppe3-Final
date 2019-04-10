package frequence

import (
	"fmt"
	"io/ioutil"
	"sort"
	"unicode"
)

const textFil = "../files/pg100.txt"

func LinjeTeller() {

	// Åpner og leser linjer fra fil.
	fil, err := ioutil.ReadFile(textFil)
	if err != nil {
		fmt.Println(err)

	}
	// Linjeteller som begynner på 1 og teller for hvert linjeskift.
	linjeTeller := 1

	for i := 0; i < len(fil); i++ {
		currentByte := fmt.Sprintf("%X", fil[i])
		if currentByte == "A" {
			linjeTeller++
		}

	}
	fmt.Println("Filen inneholder", linjeTeller, "linjer")
}

func RuneTeller() {
	// Leser fra fil
	bs, err := ioutil.ReadFile(textFil)
	if err != nil {
		fmt.Println(err)
		return

	}
	// Oppretter en map og legger inn alle runer fra teksten. Plusser på intverdien hver
	// gang en rune blir lagt inn i "map" på samme nøkkel og øker denne hver gang samme rune kommer igjen .
	m := make(map[rune]int)
	for _, r := range string(bs) {
		m[r]++
	}
	// Svaret er nå i m.  Sortering og formattering av output ved å lage en liste.
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
		if unicode.IsGraphic(lf.rune) {
			fmt.Printf("   %c    %7d\n", lf.rune, lf.freq)

		} else {
			fmt.Printf("%U  %7d\n", lf.rune, lf.freq)
		}
		teller++
	}
}

// Lager en struct for en rune med en nummerverdi.
type letterFreq struct {
	rune
	freq int
}

// Lager en liste basert på typen "letterFreq"
type lfList []*letterFreq

// Funksjon for å returnere lengde.
func (lfs lfList) Len() int { return len(lfs) }

// Sammenligne størrelsen på to runer.
func (lfs lfList) Less(i, j int) bool {
	switch fd := lfs[i].freq - lfs[j].freq; {
	case fd < 0:
		return false
	case fd > 0:
		return true
	}
	return lfs[i].rune < lfs[j].rune
}

// Bytter rundt på rekkefølgen i listen.
func (lfs lfList) Swap(i, j int) {
	lfs[i], lfs[j] = lfs[j], lfs[i]
}
