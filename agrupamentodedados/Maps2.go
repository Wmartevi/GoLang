package main

import (
	"fmt"
)

func main() {
	departamento := map[int]string{
		100: "Comercial",
		105: "Tecnologia",
		200: "Compras",
		205: "RH",
	}

	for chave, valor := range departamento {

		fmt.Printf("Departamento: %v Nome: %v\n", chave, valor)
	}

	delete(departamento, 205)
	departamento[205] = "Recursos Humanos"

	fmt.Println(departamento)
}
