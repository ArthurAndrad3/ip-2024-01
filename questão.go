package main

import "fmt"

func main() {
	var n [3]float64
	var media float64
	for i := 0; i < 3; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&n[i])
		media += n[i] / 3
	}
	fmt.Printf("Media: %.2f\n", media)
}
