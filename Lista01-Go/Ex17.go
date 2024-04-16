package main

import "fmt"

func main() {
	var x, y int
	fmt.Printf("Digite o valor de x: ")
	fmt.Scan(&x)
	fmt.Printf("Digite o valor de y: ")
	fmt.Scan(&y)
	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x+i*2)
		}
	} else {
		fmt.Printf("O valor de x não é par\n")
	}
}
