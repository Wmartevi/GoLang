package main

import "fmt"

type rango int

var x rango = 42
var y int

func main() {
	s := fmt.Sprintf("%v\n%T\n", x, x)
	fmt.Print(s)
	y = int(x)
	fmt.Println(y)
	j := fmt.Sprintf("%v\n%T\n", y, y)
	fmt.Println(j)
}
