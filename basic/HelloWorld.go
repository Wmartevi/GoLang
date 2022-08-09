package main

import "fmt"

func main() {
	numerodebytes, erros := fmt.Println("Hello,", "World, ", 100)
	fmt.Println(numerodebytes, erros)

	_, error := fmt.Println("Hello,", "World, ", 100)
	fmt.Println(error)
}
