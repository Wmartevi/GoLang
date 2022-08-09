package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Println(slice)
	fmt.Println(slice, len(slice), cap(slice))
}
