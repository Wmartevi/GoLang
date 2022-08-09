package main

import (
	"fmt"
)

func main() {
	sliceofslice := [][]int{
		[]int{1, 2, 3},
		[]int{4, 6, 7},
		[]int{8, 9, 10},
		[]int{11, 12, 13},
	}
	fmt.Println(sliceofslice[2][2])

	sliceofslice2 := [][][]int{
		[][]int{{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9}},
		[][]int{{11, 12, 13},
			[]int{14, 15, 16},
			[]int{17, 18, 19}},
	}
	fmt.Println(sliceofslice2[0][1][2])
}
