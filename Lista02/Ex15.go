package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, x, y, z int
	var a, b float64

	fmt.Scan(&T)
	bufio.NewReader(os.Stdin).ReadByte()
	for i := 0; i < T; i++ {
		fmt.Scanf("%f %f", &a, &b)
		bufio.NewReader(os.Stdin).ReadByte()
		for j := 0; j <= int(a/100); j++ {
			if j == int(a/100) {
				x = j
			}

		}
		for j := 0; j <= int(a/10); j++ {
			if j == int(a/10) {
				y = j - x*10
			}
		}
		z = (int(a) - x*100 - y*10) * 100
		somaA := x + 10*y + z
		for j := 0; j <= int(b/100); j++ {
			if j == int(b/100) {
				x = j
			}
		}
		for j := 0; j <= int(b/10); j++ {
			if j == int(b/10) {
				y = j - x*10
			}
		}
		z = (int(b) - x*100 - y*10) * 100
		somaB := x + 10*y + z

		if somaA > somaB {
			fmt.Println(somaA)
		} else {
			fmt.Println(somaB)
		}
	}

}
