package main

import "fmt"

const (
	_ = iota
	f = iota + 30
	x = 10
	y = 20
	z = "30"
	g = iota * 10
)

var a int
var b float64
var c string

func main() {
	a = x
	b = y
	c = z

	fmt.Printf("%v\t%v\t%v\t%v\t%v", a, b, c, f, g)
}
