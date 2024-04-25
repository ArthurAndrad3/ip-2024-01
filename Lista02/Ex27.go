package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	for {
		if n < 2 {
			fmt.Printf("Fatoracao nao e possivel para o numero %d!", n)
			fmt.Scan(&n)
		} else {
			break
		}
	}
	fmt.Printf("%d = ", n)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if x == 0 {
				fmt.Printf("%d", i)
				x++
			} else {
				fmt.Printf(" * %d", i)
			}
			n /= i
			if n == 1 {
				fmt.Printf("\n")
				break
			}
		}
	}
}
