package main

import "fmt"

func main() {
	var n, x, s, k, y int
	y = 100000000
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if y < x {
			s++
		} else {
			if k < s {
				k = s
			}
			s = 0
		}
		y = x
	}

	if k < s {
		k = s
	}
	fmt.Println("O comprimento do segmento crescente maximo e: ", k)
}
