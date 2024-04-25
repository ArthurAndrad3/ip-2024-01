package main

import "fmt"

func main() {
	var n, total int
	fmt.Scan(&n)
	for linha := 0; linha < 8; linha++ {
		for coluna := 0; coluna < 8; coluna++ {
			if (linha+coluna)%2 == 0 {
				total += n
			} else {
				total += 2 * n
			}

		}
	}
	fmt.Println(total - n)
}
