package main

import (
	"fmt"
)

func main() {
	//argumentos("tarde")
	//resultado, tamanho := soma(10, 10, 10, 10, 10)
	resultado, tamanho := calculo("-", 100, 10, 10, 10, 10, 10)
	fmt.Println(resultado, tamanho)

}

func argumentos(s string) {
	if s == "manhã" {
		fmt.Println("Bom dia")
	} else if s == "tarde" {
		fmt.Println("Boa tarde")
	} else {
		fmt.Println("Boa noite")
	}
}

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x)
}

func calculo(operacao string, x ...int) (int, int) {
	result := 0

	if operacao == "+" {
		for _, v := range x {
			result += v
		}
	} else {
		for _, v := range x {
			result -= v
		}
	}

	return result, len(x)
}
