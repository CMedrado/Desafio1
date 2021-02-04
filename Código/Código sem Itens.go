package main

import "fmt"

func main() {
	var Itens [15]string
	var Lista [15]int
	var Preço [15]int
	var Email [30]string
	var Individual [30]int
	Itens[0], Itens[1], Itens[2], Itens[3], Itens[4], Itens[5], Itens[6], Itens[7], Itens[8], Itens[9], Itens[10], Itens[11], Itens[12], Itens[13], Itens[14] = "Arroz", "Feijão", "Banana", "Alface", "Morango", "Bolacha", "Tomate", "Macarrão", "Batata", "Refrigerante", "Maça", "Sal", "Farinha de Trigo", "Açucar", "Detergente"
	Preço[0], Preço[1], Preço[2], Preço[3], Preço[4], Preço[5], Preço[6], Preço[7], Preço[8], Preço[9], Preço[10], Preço[11], Preço[12], Preço[13], Preço[14] = 2002, 503, 156, 164, 603, 507, 546, 287, 382, 700, 242, 272, 361, 283, 193
	s := 0
	n := 0
	i := 0
	k := 0
	Conta := 0
	Total := 0
	Totals := 0
	fmt.Println("Bem vindo\nEscolha a quantidade de cada item abaixo e boas compras!")
	for m := 0; m < 15; m++ {
		fmt.Printf("Item: %v\nPreço por unidade: %v centavos\nQuantidade: ", Itens[m], Preço[m])
		fmt.Scanln(&Lista[m])
	}
	for i <= 1 {

		fmt.Println("Digite o seu e-mail:")
		fmt.Scanln(&Email[s])
		fmt.Printf("Quer digitar mais algum e-mail?\n Digite 1 para sim\n Digite 2 para continuar\n Digite agora por favor: ")
		fmt.Scanln(&n)
		s++
		if n == 1 {
			i = 0
		} else {
			i = 2
		}
	}
	for g := 0; g < 15; g++ {
		Conta = Preço[g] * Lista[g]
		Total = Total + Conta
	}
	fmt.Printf("Total: %v\n", Total)
	fmt.Printf("Divisão entre os participantes:\n")
	for z := 0; z < s; z++ {
		Individual[z] = Total / s
	}
	Totals = Individual[0] * s
	Resto := Total - Totals
	if Resto > 0 {
		for w := 0; w < Resto; w++ {
			Individual[k] = Individual[k] + 1
			k++
		}
	}
	for l := 0; l < s; l++ {
		fmt.Printf("E-mail: %v\nTotal a pagar: %v\n", Email[l], Individual[l])
	}
	fmt.Printf("Obrigado pela Compra")
}
