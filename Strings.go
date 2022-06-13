package main

import "fmt"

type hotdog int

var b int
var c hotdog = 10

func main() {
	//interpreted strings
	x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."
	fmt.Println(x)
	//raw strings
	y := `"oi bom dia\ncomo vai?\tespero \"que\" tudo bem."`
	fmt.Println(y)

	z := fmt.Sprint(x, " ", y, " OK!!!!!!!!!!!!!!!")
	fmt.Println(z)

	fmt.Printf("%T", b)
	fmt.Printf("\n%T", c)

	//type conversion
	b = int(c)
	fmt.Printf("\n%T", b)

}
