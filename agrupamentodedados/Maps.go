package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"Alfredo": 1,
		"João":    2,
		"Carlos":  3,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["Alfredo"])

	amigos["Maria"] = 4

	fmt.Println(amigos)

	if exists, ok := amigos["Marcelo"]; !ok {
		fmt.Println("Não existe")
	} else {
		fmt.Println(exists)
	}
}
