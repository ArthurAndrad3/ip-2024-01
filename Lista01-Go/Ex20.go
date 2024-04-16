package main

import "fmt"

func main() {
	var h, m, s int
	fmt.Printf("Digite o valor de horas: ")
	fmt.Scan(&h)
	fmt.Printf("Digite o valor de minutos: ")
	fmt.Scan(&m)
	fmt.Printf("Digite o valor de segundos: ")
	fmt.Scan(&s)
	tempo := h*3600 + m*60 + s
	fmt.Printf("O tempo em segundos é %d\n", tempo)
}
