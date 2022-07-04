package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {
	pessoa1 := pessoa{"Claudio", "Da Silva", 85, "Comprador", 1236}
	pessoa2 := pessoa{"Carlos", "Dos Santos", 55, "Pedreiro", 5685}

	v1, err1 := json.Marshal(pessoa1)
	if err1 != nil {
		fmt.Println("Erro: ", err1)

	}

	v2, err2 := json.Marshal(pessoa2)
	if err2 != nil {
		fmt.Println("Erro: ", err2)

	}

	fmt.Println(string(v1))
	os.Stdout.Write(v2)
}
