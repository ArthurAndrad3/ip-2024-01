package main

import "fmt"

func main() {
	var N int
	fmt.Printf("Digite o número de temperaturas em Fahrenheit à serem convertidas para °C: ")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var F float64
		fmt.Printf("Digite a temperatura em Fahrenheit: ")
		fmt.Scan(&F)
		C := (F - 32) * 5 / 9
		fmt.Printf("A temperatura %.2f Fahrenheit em Celsius é: %.2f\n", F, C)
	}
}
