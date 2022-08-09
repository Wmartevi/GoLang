package main

import "fmt"

func main() {

	fmt.Println("**************Horas e Minutos**************")
	for horas := 0; horas <= 12; horas++ {
		fmt.Printf("\nHora: %v\n", horas)
		for min := 0; min <= 60; min++ {
			fmt.Printf(" %v", min)
		}

	}

	fmt.Println("\n**************Dia e Mês********************")
	for mes := 1; mes <= 12; mes++ {
		fmt.Printf("\nMês: %v", mes)
		for dia := 1; dia <= 60; dia++ {
			fmt.Printf(" | Dia: %v", dia)
		}
		fmt.Println()
	}

	x := 0
	for x < 10 {
		fmt.Println(` "X" é menor que Dez `)
		x++
	}

	y := 0

	for {
		if y < 10 {
			y++
			fmt.Println(` "Y" é menor que Dez `)
		} else {
			fmt.Println(` "Y" é maior que Dez `)
			break
		}
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	j := 5

	switch true {
	case j < 5:
		fmt.Println(`"J" é menor que `, 5)
	case j == 5:
		fmt.Println(`"J" é igual a `, 5)
		fallthrough //permite testar a condição seguinte
	case j != 0:
		fmt.Println(`"J" é diferente de `, 0)
	case j > 0:
		fmt.Println(`"J" é maior que `, 5)
	default:
		fmt.Println("Opção inválida!!")

	}

	nome := "Carlos"

	switch nome {
	case "João":
		fmt.Println(`O nome é: `, nome)
	case "José":
		fmt.Println(`O nome é: `, nome)
	case "Carlos":
		fmt.Println(`O nome é: `, nome)
	case "Eduardo":
		fmt.Println(`O nome é: `, nome)

	}
}
