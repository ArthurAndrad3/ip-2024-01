package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {
			if j < i {
				fmt.Printf("(%d,%d)", i, j)
				if j < i-1 && j < n {
					fmt.Printf("-")
				}
				if j == i-1 || j >= n {
					fmt.Printf("\n")
				}

			}
		}
	}
}
