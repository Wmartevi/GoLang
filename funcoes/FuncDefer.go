package main

import "fmt"

func main() {
	defer fmt.Println("Primeiro")
	defer fmt.Println("Segundo")
	fmt.Println("Terceiro")
}
