package main

import "fmt"

func calcularRendaTotal(testes int) {
	for i := 1; i <= testes; i++ {
		var ingressosTotal, popular, geral, arquibancada, cadeira float64

		fmt.Printf("Teste %d:\n", i)

		fmt.Println("Digite a quantidade total de ingressos:")
		fmt.Scanln(&ingressosTotal)

		fmt.Println("Digite a porcentagem de ingressos populares:")
		fmt.Scanln(&popular)

		fmt.Println("Digite a porcentagem de ingressos gerais:")
		fmt.Scanln(&geral)

		fmt.Println("Digite a porcentagem de ingressos de arquibancada:")
		fmt.Scanln(&arquibancada)

		fmt.Println("Digite a porcentagem de ingressos de cadeira:")
		fmt.Scanln(&cadeira)

		renda := (popular*0.1 + geral*0.5 + arquibancada + cadeira*2) * ingressosTotal

		fmt.Printf("Renda total do teste %d = %.2f\n\n", i, renda)
	}
}

func main() {
	var testes int

	fmt.Println("Digite o número de testes:")
	fmt.Scanln(&testes)

	calcularRendaTotal(testes)
}
