package main

import "fmt"

func main() {
	x := 10

	y := func(x int) int {
		//fmt.Println(x, "vezes 876 é: ", x*876)

		return x * 12

	}
	fmt.Println(x, "vezes 876 é: ", y(x))
	//y(x) //Invoca a função
}
