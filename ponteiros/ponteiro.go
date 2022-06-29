package main

import "fmt"

func main() {
	x := 10
	y := &x
	fmt.Println(&x) // endereço de memória
	fmt.Println(*y) // conteúdo do endereço de memória
	fmt.Println(&y) // endereço de memória
}
