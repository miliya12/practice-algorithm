package main

import "fmt"

func main() {
	var array = []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	selectionSort(array)
	fmt.Println(array)
}

func selectionSort(ary []int) {
	if len(ary) == 1 {
		return
	}
	var minIdx int
	for i := range ary {
		if ary[minIdx] > ary[i] {
			minIdx = i
		}

	}
	if minIdx != 0 {
		// swap
		tmp := ary[minIdx]
		ary[minIdx] = ary[0]
		ary[0] = tmp
	}
	selectionSort(ary[1:])
}
