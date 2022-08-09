package main

import "fmt"

type rango int

var x rango

func main() {
	s := fmt.Sprintf("%v\n%T\n", x, x)
	fmt.Print(s)
	x = 42
	fmt.Println(x)
}
