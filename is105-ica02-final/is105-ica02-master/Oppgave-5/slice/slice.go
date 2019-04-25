package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere 
// Returnerer en slice av type []byte
// 
func AllocateVar(b int) []byte {
		var newSlice []byte
		newSlice = make([]byte, b, b)
		return newSlice
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	newSlice := make([]byte, b, b)
	return newSlice
}

// Reslice takes a slice and reslices it
func Reslice(slice []byte, lidx int, uidx int) []byte {
	slice1 := AllocateMake(120)
	slice1 = slice[lidx:uidx]
	return slice1
}

// CopySlice ???
func CopySlice(slice []byte) []byte {
	sliceCopy := make([]byte, 60, 110)
	sliceCopy1 := slice

	copy(sliceCopy1, sliceCopy)

	return sliceCopy
}
