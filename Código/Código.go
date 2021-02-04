package main

import "fmt"

func main() {
	var Itens [30]string
	var Lista [30]int
	var Preço [30]int
	var Email [30]string
	var Individual [30]int
	s := 0
	n := 0
	i := 0
	m := 0
	g := 0
	k := 0
	Conta := 0
	Total := 0
	Totals := 0
	fmt.Println("Bem vindo\nDigite o nome do item, o valor em centavos e a quantidade desejada!")
	for g <= 1 {

		fmt.Printf("Item: ")
		fmt.Scanln(&Itens[m])
		fmt.Printf("Preço por unidade: ")
		fmt.Scanln(&Preço[m])
		fmt.Printf("Quantidade: ")
		fmt.Scanln(&Lista[m])
		fmt.Printf("Mais algum Item?\n Se sim digite 1\n Se não digite 2\n Digite agora por favor: ")
		fmt.Scanln(&g)
		m++
		if g == 1 {
			g = 0
		} else {
			g = 2
		}
	}
	for i <= 1 {

		fmt.Printf("Digite o seu e-mail: ")
		fmt.Scanln(&Email[s])
		fmt.Printf("Quer digitar mais algum e-mail?\n Digite 1 para sim\n Digite 2 para finalizar\n Digite agora por favor: ")
		fmt.Scanln(&n)
		s++
		if n == 1 {
			i = 0
		} else {
			i = 2
		}
	}
	for g = 0; g < 15; g++ {
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
