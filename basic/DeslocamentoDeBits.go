package main

import "fmt"

var a int
var b int

func main() {
	a = 42
	b = a << 2  //uma forma de multiplicar
	c := a >> 1 //uma forma de dividir

	fmt.Printf("%b\t%b\t%b\n%v\t%v,\t%v", a, b, c, a, b, c)
}
