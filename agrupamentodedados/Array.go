package main

import "fmt"

var x [5]int
var y [10]int

func main() {

	for i := 0; i < len(x); i++ {
		x[i] = i + 1
		fmt.Println("valor de x: ", x[i])

		for j := 0; j < len(y); j++ {
			y[j] = j + 1
			fmt.Println("valor de y: ", y[j])
		}
	}
}
