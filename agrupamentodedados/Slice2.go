package main

import (
	"fmt"
)

func main() {
	sabores := []string{"Peperoni", "Mozzarela", "Quatro Queijos", "Calabresa"}

	fatia := sabores[1:3]

	fmt.Println(fatia)

	fatia2 := sabores[:]

	fmt.Println(fatia2)

	fatia2 = append(sabores[:2], sabores[3:]...) //Como excluir um item, para excluir um item tem que fatiar

	for i := 0; i < len(fatia2); i++ {
		fmt.Println(fatia2[i])
	}

}
