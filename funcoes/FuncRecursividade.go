package main

import "fmt"

func main() {
	fmt.Println(fatorial(4))
	fmt.Println(fatorial(1))
	fmt.Println(loops(4))
}

// Uma função que chama ela mesma
func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial((x - 1))
}

//Utilizar loops ao inves de recursividade
func loops(x int) int {
	total := 1
	for x > 1 {

		total *= x
		x--
		//fmt.Println(x)

	}

	return total

}
