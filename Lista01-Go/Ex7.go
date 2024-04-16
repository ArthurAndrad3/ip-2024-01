package main

import "fmt"

func main() {
	var F, P float64
	fmt.Printf("Digite a temperatura em Fahrenheit: ")
	fmt.Scan(&F)
	C := (F - 32) * 5 / 9
	fmt.Printf("digite a precipitação em polegadas")
	fmt.Scan(&P)
	mm := P * 25.4
	fmt.Printf("A temperatura em Celsius é %.2f\n", C)
	fmt.Printf("A precipitação em milímetros é %.2f\n", mm)
}
