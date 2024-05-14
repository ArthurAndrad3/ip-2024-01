package main

import "fmt"

func main() {
	var testes int
	fmt.Scan(&testes)
	for i := 0; i < testes; i++ {
		cpf := make([]int, 11)
		for j := 0; j < 11; j++ {
			fmt.Scan(&cpf[j])
		}
		soma := 0
		for j := 0; j < 9; j++ {
			soma += cpf[j] * (j + 1)
		}
		if soma%11 == 10 {
			soma = 0
		} else {
			soma %= 11
		}
		soma2 := 0
		for j := 8; j >= 0; j-- {
			soma2 += cpf[j] * (9 - j)
		}
		if soma2%11 == 10 {
			soma2 = 0
		} else {
			soma2 %= 11
		}
		if soma == cpf[9] && soma2 == cpf[10] {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}
