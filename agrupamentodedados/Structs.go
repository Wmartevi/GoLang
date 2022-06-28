package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Da Silva",
		fumante:   true,
	}

	cliente2 := cliente{"Joana", "Prado", false}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

	pessoa1 := pessoa{"Carlitp", 45}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Ana",
			idade: 45,
		},
		titulo:  "Compradora",
		salario: 1500}

	pessoa3 := profissional{pessoa{"Mauro", 55}, "Ajudante", 2000}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2.titulo)
	fmt.Println(pessoa3.salario)

}
