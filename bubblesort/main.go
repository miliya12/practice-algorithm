package main

import "fmt"

func main() {
	var array = []int{5, 9, 3, 1, 2, 8, 4, 7, 6}
	bubbleSort(array)
	fmt.Println(array)
}

func bubbleSort(ary []int) {
	if len(ary) == 1 {
		return
	}
	for i := len(ary) - 1; i > 0; i-- {
		// swap
		if ary[i] < ary[i-1] {
			tmp := ary[i]
			ary[i] = ary[i-1]
			ary[i-1] = tmp
		}
	}

	bubbleSort(ary[1:])
}
