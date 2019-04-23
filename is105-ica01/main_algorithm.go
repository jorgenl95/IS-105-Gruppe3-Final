package main

import (
	"github.com/jorgenl95/is105-ica01/algorithms"
	"fmt"
)

var list1 = []int{2, 3, 5, 7, 11, 13, 20, 5, 21, 19}

func main() {
	algorithms.Bubble_sort(list1)
	fmt.Printf("%v", list1)
	algorithms.Bubble_sort_modified(list1)
	fmt.Println()
	fmt.Printf("%v", list1)
}
