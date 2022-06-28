package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type vencimentos struct {
	salario    int
	bonus      int
	beneficios int
}

type dentista struct {
	pessoa
	vencimentos
	dentesarrancados int
}

type arquiteto struct {
	pessoa
	vencimentos
	tipodeconstrucao string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é: ", x.nome, " Eu já arranquei: ", x.dentesarrancados, " Meu salario é: ", x.salario)
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é: ", x.nome, " Minha Construção é: ", x.tipodeconstrucao, " Meu salario é: ", x.salario)
}

type trabalhador interface {
	oibomdia()
}

func profissao(p trabalhador) {
	p.oibomdia()
}

func main() {

	dentista := dentista{pessoa{"Marlos", "Dos Santos", 45}, vencimentos{15000, 1000, 2000}, 3}
	arquiteto := arquiteto{pessoa{"Josué", "Silva", 35}, vencimentos{12000, 3000, 4000}, "Alvenaria"}
	//Chamada fora interface
	dentista.oibomdia()
	arquiteto.oibomdia()

	//Chamada pela interface
	profissao(dentista)
	profissao(arquiteto)

}
