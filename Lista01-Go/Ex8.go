package main

import "fmt"

func main() {
	var R, H float64
	Pi := 3.14159
	fmt.Printf("Digite o raio da base do cilindro em metros: ")
	fmt.Scan(&R)
	fmt.Printf("Digite a altura do cilindro em metros: ")
	fmt.Scan(&H)
	Ac := Pi * R * R
	Al := 2 * Pi * R * H
	At := 2*Ac + Al
	custo := 100 * At
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)

}
