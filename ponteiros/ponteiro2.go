package main

import "fmt"

func main() {
	x := 10
	//recebeovalor(x) //passa o valor
	recebeoponteiro(&x) // passa o endereço de memória

	fmt.Println(x)

}

func recebeovalor(x int) {
	x++
	fmt.Println(x)
}

func recebeoponteiro(x *int) {
	*x++
	fmt.Println(*x)
}
