package main

import (
	"fmt"
	"math"
)

func main() {
	var H, aresta float64
	fmt.Printf("Digite a altura da piramide: ")
	fmt.Scan(&H)
	fmt.Printf("Digite a aresta da base da piramide: ")
	fmt.Scan(&aresta)
	Ab := (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2
	v := (Ab * H) / 3
	fmt.Printf("Volume da piramide: %.2f m³\n", v)
}
