package main

import "fmt"

func main() {
	var ValorIngresso, ValorMinimo, ValorMaximo, ingresso, despesa, x, y, z float64
	fmt.Scan(&ValorIngresso)
	fmt.Scan(&ValorMinimo)
	fmt.Scan(&ValorMaximo)
	for i := ValorMinimo; i <= ValorMaximo; i++ {

		if i < ValorIngresso {
			ingresso = 120 + (50 * (ValorIngresso - i))
		}
		if i == ValorIngresso {
			ingresso = 120
		}
		if i > ValorIngresso {
			ingresso = 120 - (60 * (i - ValorIngresso))
		}
		despesa = 200 + (0.05 * ingresso)
		lucro := ingresso*i - despesa
		fmt.Printf("V: %.2f, N: %.2f, L: %.2f \n", i, ingresso, lucro)
		if x < lucro {
			x = lucro
			y = i
			z = ingresso
		}
	}
	fmt.Printf("Melhor valor final: %.2f\nLucro: %.2f\nNumero de ingressos: %.2f\n", y, x, z)
}
