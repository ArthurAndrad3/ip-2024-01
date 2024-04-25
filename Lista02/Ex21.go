package main

import "fmt"

func main() {
	var tamanho int

	for {
		fmt.Scan(&tamanho)
		if tamanho == 0 {
			break
		}
		numeros := make([]float64, tamanho)
		for i := 0; i < tamanho; i++ {
			fmt.Scan(&numeros[i])
		}
		ordenada := true
		for i := 1; i < tamanho; i++ {
			if numeros[i] < numeros[i-1] {
				ordenada = false
				break
			}
		}

		if ordenada {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}
