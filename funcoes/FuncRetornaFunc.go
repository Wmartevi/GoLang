package main

import "fmt"

func main() {
	x := retornaumafuncao()
	y := x(3)

	z := retornaumafuncao()(4)

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(retornaumafuncao()(5))

}

func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
