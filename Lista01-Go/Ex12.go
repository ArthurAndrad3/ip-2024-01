package main

import "fmt"

func main() {
	var T, t2 int
	fmt.Printf("Digite o tempo em horas: ")
	fmt.Scan(&T)
	t2 = T / 3
	v1 := t2 * 10
	v2 := T % 3 * 5

	fmt.Printf("o valor a pagar é %d\n", v1+v2)
}
