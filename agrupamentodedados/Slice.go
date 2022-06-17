package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice2 := append(slice, 16)
	fmt.Println(slice2)

	for indice, valor := range slice2 {

		fmt.Printf("No índice: %v  temos o valor: %v\n", indice, valor)

		if valor == 16 {
			slice2 = append(slice2, 17)
		}

	}

	for _, valor := range slice2 {

		fmt.Printf("Temos o valor: %v\n", valor)

	}

}
