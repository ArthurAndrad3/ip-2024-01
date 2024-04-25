package main

import "fmt"

func main() {
	var N, k, soma int
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		soma = 0
		fmt.Printf("%d*%d*%d = ", i, i, i)
		for {
			j := k
			j++
			if soma == i*i*i {
				fmt.Printf("\n")
				break
			}

			if j%2 != 0 {
				soma += j
				fmt.Printf("%d ", j)
				if soma != i*i*i {
					fmt.Printf("+ ")
				}
			}
			k = j
		}

	}
}
