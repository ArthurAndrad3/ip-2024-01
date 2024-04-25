package main

import "fmt"

func main() {
	var a, b, t float64
	fmt.Printf("Digite o valor da População A: ")
	fmt.Scan(&a)
	fmt.Printf("Digite o valor da População B: ")
	fmt.Scan(&b)
	for a < b {
		t++
		a += a * 0.03
		b += b * 0.015
	}
	fmt.Printf("ANOS = %.0f\n", t)
}
