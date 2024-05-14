package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for {
		if n < 1 || n > 100000 {
			fmt.Println("Valor inválido")
			fmt.Scan(&n)
		} else {
			break
		}
	}
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	fmt.Scan(&m)
	for {
		if m < 1 || m > 1000 {
			fmt.Println("Valor inválido")
			fmt.Scan(&m)
		} else {
			break
		}
	}
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		for j := 0; j < n; j++ {
			if x == vetor[j] {
				fmt.Println("ACHEI")
				break
			} else if j == n-1 {
				fmt.Println("NAO ACHEI")
			}
		}
	}
}
