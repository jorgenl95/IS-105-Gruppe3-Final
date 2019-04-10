package slice

// Oppgave 5.a
// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte

func AllocateVar(b int) []byte {
	var newSlice []byte
	newSlice = make([]byte, b, b)
	return newSlice
}

// Oppgave 5.a
// AllocateMake tar lengde og kapasitet som b og lager en ny slice
func AllocateMake(b int) []byte {
	newSlice := make([]byte, b, b)
	return newSlice
}

// Oppgave 5.b
// Reslice tar en slice og "slicer" den opp på nytt.
// Skal bruke en av "allocate metodene" for å allokere plass i minnet.
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slc1 := AllocateMake(120)
	slc1 = slc[lidx:uidx]
	return slc1

}

// Oppg. 5.b
// Kopierer en slice fra et innparameter til en ny slice. Den nye slicen printes ut.
func CopySlice(copy1 []byte) []byte {
	copyarr := make([]byte, 60, 110)
	copyarr1 := copy1

	copy(copyarr1, copyarr)

	return copyarr
}
