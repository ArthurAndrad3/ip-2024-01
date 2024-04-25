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
		for j := i * 2; j > 0; j-- {
			e *= float64(j)
		}
		f = math.Pow(x, float64(i*2))
		if i%2 == 0 {
			y += f / e
		} else {
			y -= f / e
		}
	}
	fmt.Printf("cos(%.2f) = %.6f\n", x, y)
}
