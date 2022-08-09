package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}

	slice3 := append(slice1, slice2...)

	fmt.Println(slice3)
}
