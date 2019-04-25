package sum

import (
	"fmt"
	"math"
)

func SumInt8(a, b int) int {
	if ((math.MinInt8 < a) && (a < math.MaxInt8)){
		if ((math.MinInt8 < b) && (b < math.MaxInt8)){
			return a + b
		}
		fmt.Println("Error: Second value is outside int8 boundaries")
		return 0
	}
	fmt.Println("Error: One of the values is outside int8 boundaries")
	return 0
}

func SumFloat64(a, b float64) float64 {
	if ((math.SmallestNonzeroFloat64 < a) && (a < math.MaxFloat64)) {
		if ((math.SmallestNonzeroFloat64 < b) && (b < math.MaxFloat64)) {
			return a + b
		}
		fmt.Println("Error: Second value is outside float64 boundaries")
		return 0.0
	}
	fmt.Println("Error: One of the values is outside float64 boundaries")
	return 0.0
}

func SumUint32(a, b uint32) uint32 {
	return a + b
}

func SumInt32(a, b int32) int32 {
	return a + b
}

func SumInt64(a, b int64) int64 {
	return a + b
}