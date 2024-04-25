package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < i; k++ {
				if k*k+j*j == i*i {
					if j < k {
						fmt.Printf("hipotenusa = %d, catetos %d e %d\n", i, j, k)
					}
				}

			}
		}
	}
}
