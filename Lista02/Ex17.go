package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, d, y, m, n, o, x, l, g int
	var b, c, plucro, somaC, somaV, lucro float64
	for {
		fmt.Scan(&a, &b, &c, &d)
		bufio.NewReader(os.Stdin).ReadByte()
		if a == 0 {
			break
		}
		plucro = (c - b) / b
		if plucro < 0.1 {
			m++
		}
		if plucro >= 0.1 && plucro <= 0.2 {
			n++
		}
		if plucro > 0.2 {
			o++
		}
		lucro = (float64(d) * c) - (float64(d) * b)
		if l < int(lucro) {
			l = int(lucro)
			x = a
		}
		if g < d {
			g = d
			y = a
		}
		somaC += (float64(d) * b)
		somaV += (float64(d) * c)

	}
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\nQuantidade de mercadorias que geraram lucro maior ou igual a 10%%, e menor ou igual a 20%%: %d\nQuantidade de mercadorias que geraram lucro maior do que 20%%: %d\nCodigo da mercadoria que gerou maior lucro: %d\nCodigo da mercadoria mais vendida: %d\nValor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%", m, n, o, int(x), y, somaC, somaV, (somaV-somaC)/somaC*100)
}
