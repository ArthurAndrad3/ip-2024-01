package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x, y, e, f float64
	fmt.Scan(&x)
	fmt.Scan(&N)
	for i := 0; i <= int(N); i++ {
		e = 1
		for j := i; j > 0; j-- {
			e *= float64(j)
		}
		f = math.Pow(x, float64(i))
		y += f / e
	}
	fmt.Printf("e^%.2f = %.6f\n", x, y)
}
