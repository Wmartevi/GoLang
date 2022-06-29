package main

import "fmt"

func main() {
	x := 387

	func(x int) {
		fmt.Println(x, "vezes 876 é: ", x*876)

	}(x) //Invoca a função
}
