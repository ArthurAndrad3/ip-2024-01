package main

import "fmt"

func main() {
	var n, x float64
	x = 1
	var y int
	fmt.Scan(&n)
	for {
		n *= 10
		y = int(n)
		x *= 10
		if y%10 == 0 && n-float64(y) == 0 {
			break
		}
	}
	for i := 2; i <= int(x); i++ {
		for int(n)%i == 0 && int(x)%i == 0 {
			n /= float64(i)
			x /= float64(i)
		}
	}

	fmt.Printf("%.0f/%.0f\n", n, x)
}
