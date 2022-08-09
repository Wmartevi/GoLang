package main

import "fmt"

//fora do package level a declaração deve ser feita com "var"
var y = "Olá bom dia!!!"

func main() {
	//Declaração
	x := 10 + 10
	//y := "olá bom dia"

	fmt.Printf("y : %v, %T\n", y, y)
	fmt.Printf("x : %v, %T\n", x, x)

	//Atribuição
	x, y = 20, "Olá tarde"
	fmt.Printf("y : %v, %T\n", y, y)
	fmt.Printf("x : %v, %T\n", y, y)

	qualquercoisa(x)
}

func qualquercoisa(x int) {
	fmt.Print(x, y)
}
